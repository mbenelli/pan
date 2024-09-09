package jenkins

import (
	"encoding/json"
	"net/http"
	"os"
)

const (
	defaultConfPath = "jenkins-sample.json"
	jobPath         = "api/json"
	manifestPath    = "artifact/reports/manifest.xml"
)

type Config struct {
	Url  string
	Jobs []string
}

var config *Config

func DefaultConfPath() string {
	home, err := os.UserConfigDir()
	if err != nil {
		return defaultConfPath
	}
	return home + "/pan/" + defaultConfPath
}

func ReadConf(path string) (Config, error) {
	cfg, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = json.Unmarshal([]byte(cfg), &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

func GetConfig() Config {
	if config == nil {
		c, err := ReadConf(DefaultConfPath())
		if err != nil {
			os.Exit(-1)
		} else {
			config = &c
		}
	}
	return *config
}

func JobsList() []string {
	var urls []string
	c := GetConfig()
	for _, j := range c.Jobs {
		urls = append(urls, c.Url+j)
	}
	return urls
}

func Job(url string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		return "error in creating request"
	}
	resp, err := client.Do(req)
	if err == nil {
		return "error"
	}
}
