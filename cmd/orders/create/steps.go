package create

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
	rx "github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/model"
	o "github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/goreydenx/user"
	"github.com/pixel365/rdnx-cli/helpers"
)

func twitchStep1() (twitchId int32) {
	color.Green("STEP 1:")
	twitchId = int32(helpers.AskIntegerValue(
		"Enter Twitch Channel Id:",
		"Invalid Value: Channel Id must be greater than zero",
		false,
	))
	return
}

func youtubeStep1() (channelUrl string) {
	channelUrl = ""
	color.Green("STEP 1:")
Loop:
	for {
		channelUrl = helpers.AskStringValue(
			"Enter YouTube Channel URL:",
			false,
		)
		if !strings.Contains(channelUrl, "http") || !strings.Contains(channelUrl, "://") {
			color.Red("Only full URL format supports. Example: https://www.youtube.com/@ThePrimeTimeagen")
			continue Loop
		}

		if !(strings.Contains(channelUrl, "youtu.be") || strings.Contains(channelUrl, "youtube.com")) {
			color.Red("Invalid URL. Example: https://www.youtube.com/@ThePrimeTimeagen")
			continue Loop
		}

		break Loop
	}
	return
}

func step2() (numberOfViews int32) {
	color.Green("STEP 2:")
	numberOfViews = int32(helpers.AskIntegerValue(
		"Enter the number of views:",
		"Invalid Value: Value must be greater than zero",
		false,
	))
	return
}

func step3() (numberOfViewers int32) {
	color.Green("STEP 3:")
	numberOfViewers = int32(helpers.AskIntegerValue(
		"Enter the number of viewers:",
		"Invalid Value: Value must be greater than zero",
		false,
	))
	return
}

func step4(prices *model.Result[[]model.Price]) (priceId int32) {
	color.Green("STEP 4:")
	for i, price := range prices.Result {
		fmt.Printf("%d) %s (Cost: %.3f)\n", i+1, price.Name, price.Price)
	}

	priceId = 0
Loop:
	for {
		value := helpers.AskIntegerValue(
			"Enter the tariff number from the list:",
			"Invalid Value: Value must be greater than zero",
			false,
		)
		for i, price := range prices.Result {
			if i+1 == value {
				priceId = int32(price.Id)
				break Loop
			}
		}
		color.Red("Invalid tariff number, please select another one")
	}
	return
}

func step5() (launchMode string, delayTime int) {
	delayTime = 0
	delayTimeFn := func() int {
		value := helpers.AskIntegerValue(
			"Delay the launch for (minutes):",
			"Invalid Value: The value must be >= 0",
			true,
		)
		return value
	}

	color.Green("STEP 5:")
	fmt.Println("1) Auto")
	fmt.Println("2) Manual")
	fmt.Println("3) Delay")

Loop:
	for {
		value := helpers.AskIntegerValue(
			"Select startup mode:",
			"Invalid Value: The value must be from 1 to 3",
			false,
		)
		switch value {
		case 1:
			launchMode = "auto"
			break Loop
		case 2:
			launchMode = "manual"
			break Loop
		case 3:
			launchMode = "delay"
			delayTime = delayTimeFn()
			break Loop
		}
		color.Red("Invalid number, please select another one")
	}
	return
}

func step6() (enabled bool, minutes int) {
	enabled = false
	minutes = 0
	minutesFn := func() int {
		value := helpers.AskIntegerValue(
			"Period of gradual increase in viewers in minutes:",
			"Invalid Value: The value must be >= 0",
			true,
		)
		return value
	}

	color.Green("STEP 6:")
Loop:
	for {
		enableSmoothGain := helpers.AskStringValue(
			"Enable a smooth increase in viewers after the start? (y/n):",
			false,
		)

		switch strings.ToLower(enableSmoothGain) {
		case "y":
			minutes = minutesFn()
			break Loop
		case "n":
			break Loop
		default:
			color.Red("Invalid value: enter 'y' or 'n'")
		}
	}
	return
}

func confirmation[T model.OrderParams](client *rx.Client, params T) {
	allow := false
Loop:
	for {
		choice := helpers.AskStringValue(
			"Would you like to create an order with these parameters? (y/n):",
			false,
		)

		switch strings.ToLower(choice) {
		case "y":
			allow = true
			break Loop
		case "n":
			allow = false
			break Loop
		default:
			color.Red("Invalid value: enter 'y' or 'n'")
		}
	}

	if allow {
		result, err := o.CreateStream(client, params)
		helpers.Marshal(result, err)
		helpers.WaitingTask(client, result)
	} else {
		color.White("You have canceled the order creation")
	}
}

func loadData(client *rx.Client, pricesFn func(*rx.Client) (*model.Result[[]model.Price], error)) (*model.Result[[]model.Price], *model.Balance) {
	var prices *model.Result[[]model.Price]
	var balance *model.Balance
	var wg sync.WaitGroup
	var errors []error

	wg.Add(2)

	go func(wg *sync.WaitGroup, client *rx.Client) {
		defer wg.Done()
		res, err := pricesFn(client)
		if err != nil {
			errors = append(errors, err)
		} else {
			prices = res
		}
	}(&wg, client)

	go func(wg *sync.WaitGroup, client *rx.Client) {
		defer wg.Done()
		res, err := user.Balance(client)
		if err != nil {
			errors = append(errors, err)
		} else {
			balance = res
		}
	}(&wg, client)

	wg.Wait()

	if len(errors) > 0 {
		for _, e := range errors {
			color.Red(e.Error())
		}
		os.Exit(1)
	}

	return prices, balance
}

func summary(prices *model.Result[[]model.Price], balance *model.Balance, params *model.BaseOrderParams) {
	fmt.Printf("Number of Views: %d\n", params.NumberOfViews)
	fmt.Printf("Number of Viewers: %d\n", params.NumberOfViewers)

	if params.LaunchMode == "delay" && params.DelayTime > 0 {
		fmt.Printf("Launch Mode: %s (%d min.)\n", params.LaunchMode, params.DelayTime)
	} else {
		fmt.Printf("Launch Mode: %s\n", params.LaunchMode)
	}

	if params.SmoothGain.Enabled && params.SmoothGain.Minutes > 0 {
		fmt.Printf("Smooth Gain: %d\n", params.SmoothGain.Minutes)
	}

	for _, price := range prices.Result {
		if price.Id == params.PriceId {
			fmt.Printf("Tariff (id: %d): %s\n", price.Id, price.Name)
			fmt.Printf("Price: %.3f %s\n", price.Price, balance.Currency)
			fmt.Printf("Total Cost (%.3f * %d): %.3f %s\n", price.Price, params.NumberOfViews, price.Price*float64(params.NumberOfViews), balance.Currency)

			if (price.Price * float64(params.NumberOfViews)) > balance.Amount {
				color.Red("There are not enough funds in your balance to create this order")
				os.Exit(1)
			}
		}
	}
}
