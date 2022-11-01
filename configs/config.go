package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	SecretKey string `json:"secret_key"`
	AppPort   string `json:"app_port"`
}

const FILENAME = "configs/config.json"

var App Config

func GetConfig() {
	configFile, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		msg := fmt.Sprintf("Failed to open config file: %s", err)
		panic(msg)
	}
	err = json.Unmarshal(configFile, &App)
	if err != nil {
		msg := fmt.Sprintf("Failed to read config file: %s", err)
		panic(msg)
	}
}
