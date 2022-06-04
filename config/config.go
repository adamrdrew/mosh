package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const UNINITIALIZED = "UNINITIALIZED"
const CONFIG_FILE_PART = "config.yaml"

var Version string
var defaultConfig = Config{
	Token:           UNINITIALIZED,
	Address:         UNINITIALIZED,
	Port:            UNINITIALIZED,
	Library:         UNINITIALIZED,
	ShowArt:         false,
	CacheMaxSizeMB:  4096,
	CacheMaxAgeDays: 14,
}

func GetPort() string {
	port := os.Getenv("MOSH_PORT")
	if port == "" {
		port = "9666"
	}
	return port
}

func GetCacheDir() string {
	dir := os.Getenv("MOSH_CACHE_DIR")
	if dir == "" {
		dir = "/tmp/mosh"
	}
	return dir + "/"
}

func GetPidDir() string {
	dir := os.Getenv("MOSH_PID_DIR")
	if dir == "" {
		dir = "/tmp"
	}
	return dir + "/"
}

func GetConfigDir() string {
	dir := os.Getenv("MOSH_CONFIG_DIR")
	if dir == "" {
		dir = "/etc/mosh/"
	}
	return dir + "/"
}

func GetLogir() string {
	dir := os.Getenv("MOSH_LOG_DIR")
	if dir == "" {
		dir = "/var/log/mosh"
	}
	return dir + "/"
}

func GetConfig() Config {
	conf := Config{}
	conf.Load()
	return conf
}

type Config struct {
	Token           string
	Address         string
	Port            string
	Library         string
	ShowArt         bool
	CacheMaxSizeMB  int
	CacheMaxAgeDays int
}

func (c *Config) Load() {
	c.createConfigFileIfNotThere()
	c.loadYAML()
}

func (c *Config) SetToken(token string) {
	c.Token = token
	c.Save()
}

func (c *Config) filePath() string {
	return GetConfigDir() + CONFIG_FILE_PART
}

//If we extend the config spec we'll have bogus values in whatever the new keys are
//So we have to do this on load
func (c *Config) setMissingValues() {
	if c.CacheMaxAgeDays == 0 {
		c.CacheMaxAgeDays = defaultConfig.CacheMaxAgeDays
	}
	if c.CacheMaxSizeMB == 0 {
		c.CacheMaxSizeMB = defaultConfig.CacheMaxSizeMB
	}
}

func (c *Config) loadYAML() {
	yfile, err := ioutil.ReadFile(c.filePath())
	if err != nil {
		panic(err)
	}

	errorUnmarshal := yaml.Unmarshal(yfile, &c)
	if errorUnmarshal != nil {
		panic(errorUnmarshal)
	}

	c.setMissingValues()

}

func (c *Config) Save() {
	data, err := yaml.Marshal(&c)
	if err != nil {
		panic(err)
	}

	err2 := ioutil.WriteFile(c.filePath(), data, 0)
	if err2 != nil {
		panic(err)
	}
}

func (c *Config) createConfigFileIfNotThere() {
	_, statErr := os.Stat(GetConfigDir())
	if os.IsNotExist(statErr) {
		os.Mkdir(GetConfigDir(), os.ModePerm)
	}

	_, statErr = os.Stat(c.filePath())
	if !os.IsNotExist(statErr) {
		return
	}

	yamlData, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(c.filePath(), yamlData, 0755)
	if err != nil {
		panic(err)
	}

}
