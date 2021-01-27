package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port   int    `json:"port"`
	DBPath string `json:"dbPath"`
	Token  string `json:"token"`
}

var Conf Config

func GetConfig(inputConfPath string) error {
	configPath := "./data.json"
	if inputConfPath != "" {
		configPath = inputConfPath
	}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	var data Config
	json.Unmarshal(file, &data)
	Conf = data
	return nil
}
