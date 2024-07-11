package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/pixel365/goreydenx"
	"github.com/pixel365/goreydenx/action"
	"github.com/pixel365/goreydenx/model"
)

const (
	ConfigDirName  = "/.reyden_x"
	ConfigFileName = "/config.json"
)

func GetConfigDirPath() (dirPath string) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir, err := os.UserHomeDir()
		if err == nil {
			dirPath = configDir + ConfigDirName
		}
	} else {
		dirPath = configDir + ConfigDirName
	}
	return
}

func GetConfigFilePath() string {
	return GetConfigDirPath() + ConfigFileName
}

func SaveConfig(config *Config) {
	data, err := json.Marshal(config)
	if err != nil {
		log.Fatal("unable to encode configuration file content")
	}

	if writeErr := os.WriteFile(GetConfigFilePath(), data, 0644); writeErr != nil {
		log.Fatal("unable to write configuration file content")
	}
}

func LoadConfig() *Config {
	content, err := os.ReadFile(GetConfigFilePath())
	if err != nil {
		log.Fatal("unable to read configuration file")
	}

	config := Config{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("unable to decode configuration file")
	}

	return &config
}

func AuthGuard() {
	config := LoadConfig()
	if !config.IsAuthenticated() {
		color.Red(Unauthorized)
		os.Exit(1)
	}
}

func Next[T any](r *model.Result[T]) bool {
	if !r.HasNext() {
		return false
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		color.Cyan("Load more? (y/n):")
		choice, _ := reader.ReadString('\n')
		choice = strings.Replace(choice, "\n", "", -1)
		switch strings.ToLower(choice) {
		case "y":
			return true
		case "n":
			return false
		}
	}
}

func AskMultipleIntValues() model.Identifiers {
	var identifiers model.Identifiers
	reader := bufio.NewReader(os.Stdin)
Loop:
	for {
		color.Cyan("Enter order numbers separated by commas (positive numbers only):")
		numbers, _ := reader.ReadString('\n')
		numbers = strings.Replace(numbers, "\n", "", -1)
		for _, n := range strings.Split(numbers, ",") {
			value, err := strconv.Atoi(strings.ReplaceAll(n, " ", ""))
			if err != nil {
				continue
			}

			if value < 1 {
				continue
			}

			identifiers.Identifiers = append(identifiers.Identifiers, uint32(value))
		}

		if len(identifiers.Identifiers) < 1 {
			continue Loop
		}

		return identifiers
	}
}

func AskIntegerValue(question, errorMessage string, allowZero bool) int {
	reader := bufio.NewReader(os.Stdin)
Loop:
	for {
		color.Cyan(question)
		number, _ := reader.ReadString('\n')
		number = strings.Replace(number, "\n", "", -1)
		value, err := strconv.Atoi(number)
		if err != nil {
			color.Red(errorMessage)
			continue Loop
		}

		if !allowZero && value == 0 {
			color.Red(errorMessage)
			continue Loop
		}

		if value < 0 {
			color.Red(errorMessage)
			continue Loop
		}

		return value
	}
}

func AskStringValue(question string, allowSpaces bool) string {
	reader := bufio.NewReader(os.Stdin)
Loop:
	for {
		color.Cyan(question)
		value, _ := reader.ReadString('\n')
		value = strings.Replace(value, "\n", "", -1)
		if !allowSpaces {
			value = strings.ReplaceAll(value, " ", "")
			if value == "" {
				color.Red("The value must not be empty")
				continue Loop
			}
		}

		return value
	}
}

func AskOrderId() int32 {
	reader := bufio.NewReader(os.Stdin)
Loop:
	for {
		color.Cyan(EnterOrderNumber)
		number, _ := reader.ReadString('\n')
		number = strings.Replace(number, "\n", "", -1)
		value, err := strconv.Atoi(number)
		if err != nil {
			color.Red(InvalidOrderNumber)
			continue Loop
		}

		if value < 1 {
			color.Red(InvalidOrderNumber)
			continue Loop
		}

		return int32(value)
	}
}

func Marshal(result any, e error) {
	if e != nil {
		color.Red(e.Error())
		os.Exit(1)
	}

	j, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println(string(j))
}

func WaitingTask(c *goreydenx.Client, result *model.ActionResult) {
	reader := bufio.NewReader(os.Stdin)
Loop:
	for {
		color.Cyan("Would you like to track the status of this operation? (y/n)")
		choice, _ := reader.ReadString('\n')
		choice = strings.Replace(choice, "\n", "", -1)
		switch strings.ToLower(choice) {
		case "y":
			var wg sync.WaitGroup
			wg.Add(1)
			go func(wg *sync.WaitGroup, result *model.ActionResult) {
				defer wg.Done()
				for {
					status, err := action.TaskStatus(c, result.OrderId, result.Task.Id)
					if err != nil {
						color.Red(err.Error())
						return
					}

					if status == "completed" {
						color.Green(fmt.Sprintf("Task %s is %s!", result.Task.Id, status))
						return
					}

					color.Yellow(fmt.Sprintf("Task %s: current status is '%s'", result.Task.Id, status))
					time.Sleep(time.Second * 5)
				}
			}(&wg, result)
			wg.Wait()
			break Loop
		case "n":
			break Loop
		default:
			continue Loop
		}
	}
}

func AskLaunchMode() (launchMode string, delayTime int) {
	delayTime = 0
	delayTimeFn := func() int {
		value := AskIntegerValue(
			"Delay start for (minutes):",
			"Invalid Value: The value must be >= 0",
			false,
		)
		return value
	}

	color.Green("STEP 5:")
	fmt.Println("1) Auto")
	fmt.Println("2) Manual")
	fmt.Println("3) Delay")

Loop:
	for {
		value := AskIntegerValue(
			"Select launch mode:",
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
			for delayTime < 5 || delayTime > 240 {
				color.Red("The number of minutes for delayed start should be from 5 to 240")
				delayTime = delayTimeFn()
			}
			break Loop
		}
		color.Red("Invalid number, please select another one")
	}
	return
}
