package create

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/pixel365/goreydenx/model"

	p "github.com/pixel365/goreydenx/prices"

	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func NewYouTubeOrderCmd() *cobra.Command {
	return &cobra.Command{
		Use:     helpers.YouTube,
		Aliases: []string{"yt"},
		Short:   "Create new order for YouTube stream",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.AuthGuard()
			config := helpers.LoadConfig()
			client := config.RxClient()
			if client != nil {
				prices, balance := loadData(client, p.YouTube)

				channelUrl := youtubeStep1()
				numberOfViews := step2()
				numberOfViewers := step3()
				priceId := step4(prices)
				launchMode, delayTime := step5()
				smmothGainEnabled, smoothGainMinutes := step6()

				params := model.YouTubeParams{
					BaseOrderParams: model.BaseOrderParams{
						LaunchMode: launchMode,
						SmoothGain: model.SmoothGain{
							Enabled: smmothGainEnabled,
							Minutes: uint32(smoothGainMinutes),
						},
						PriceId:         uint32(priceId),
						NumberOfViews:   uint32(numberOfViews),
						NumberOfViewers: uint32(numberOfViewers),
						DelayTime:       uint32(delayTime),
					},
					ChannelUrl: channelUrl,
				}

				color.Yellow("=== SUMMARY ===")

				fmt.Println("Platform: YouTube")
				fmt.Printf("YouTube Channel URL: %s\n", params.ChannelUrl)

				summary(prices, balance, &params.BaseOrderParams)
				confirmation(client, &params)
			} else {
				color.Red(helpers.Unauthorized)
			}
		},
	}
}
