package cmd

import (
	"log"
	"os"

	"github.com/pixel365/rdnx-cli/helpers"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	createConfigFile()
}

func createConfigFile() {
	dirPath := helpers.GetConfigDirPath()
	if _, err := os.Stat(dirPath); err != nil {
		if e := os.Mkdir(dirPath, 0750); e != nil {
			log.Fatal(e)
		}
	}

	configFullPath := helpers.GetConfigFilePath()
	if _, err := os.Stat(configFullPath); err != nil {
		if _, err := os.Create(configFullPath); err != nil {
			log.Fatal("unable to create configuration file")
		} else {
			config := helpers.Config{}
			config.Accounts = make(map[string]helpers.Creds)
			helpers.SaveConfig(&config)
		}
	}
}
