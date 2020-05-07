package config

import (
	"fmt"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
	// split word splits the env var name on uppercase letters with "_"
	// the env var name is APP_CONFIG_FILE
	ConfigFile string `split_words:"true"`
	Db         struct {
		User     string `yaml:"user" envconfig:"DB_USER"`
		Password string `yaml:"password" envconfig:"DB_PASSWORD"`
		Host     string `yaml:"host" envconfig:"DB_HOST"`
		Port     string `yaml:"port" envconfig:"DB_PORT"`
	}
}

func (config *AppConfig) ReadEnv() {
	envconfig.Process("app", config)
}

func (config *AppConfig) ReadYaml(filepath string) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error Reading Config file with path: %v\n", filepath)
	}

	yaml.Unmarshal(yamlFile, config)
}
