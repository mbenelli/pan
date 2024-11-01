package jenkins

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	defaultConfPath = "jenkins-sample.json"
	jobPath         = "api/json"
	manifestPath    = "artifact/reports/manifest.xml"
)

type Config struct {
	Url   string
	Jobs  []string
	User  string
	Token string
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
	err = json.Unmarshal(cfg, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

func GetConfig() Config {
	if config == nil {
		c, err := ReadConf(DefaultConfPath())
		if err != nil {
			panic(fmt.Sprintf("Error in reading config file: %v", err.Error()))
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
	c := GetConfig()
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	}
	req.SetBasicAuth(c.User, c.Token)
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	if resp.StatusCode != 200 {
		return fmt.Sprintf("Wrong status code: %d", resp.StatusCode)
	}
	return "ok"
}
