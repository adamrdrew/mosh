package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const CONFIG_FILE = "config/config.yaml"

type Config struct {
	Token string
}

func (c *Config) Load() {
	c.createConfigFileIfNotThere()
	c.loadYAML()
}

func (c *Config) loadYAML() {
	yfile, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		panic(err)
	}

	errorUnmarshal := yaml.Unmarshal(yfile, &c)
	if errorUnmarshal != nil {
		panic(errorUnmarshal)
	}
}

func (c *Config) Save() {
	data, err := yaml.Marshal(&c)
	if err != nil {
		panic(err)
	}

	err2 := ioutil.WriteFile(CONFIG_FILE, data, 0)
	if err2 != nil {
		panic(err)
	}
}

func (c *Config) createConfigFileIfNotThere() {
	_, statErr := os.Stat(CONFIG_FILE)
	if !os.IsNotExist(statErr) {
		return
	}

	defaultConfig := Config{
		Token: "NOTYETCREATED",
	}

	yamlData, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(CONFIG_FILE, yamlData, 0755)
	if err != nil {
		panic(err)
	}

}
