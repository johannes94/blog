package main

import (
	"clean-config/src/config"
	"fmt"
)

func main() {
	appConfig := &config.AppConfig{}
	// Load default config file
	appConfig.ReadYaml("config.yaml")
	// Load specific configuration with env variables
	appConfig.ReadEnv()

	if appConfig.ConfigFile != "" {
		appConfig.ReadYaml(appConfig.ConfigFile)
	}

	fmt.Printf("%+v\n", *appConfig)
}
