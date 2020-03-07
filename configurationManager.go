package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	LastFmConfig LastFmConfig `json:"lastfm"`
}

type LastFmConfig struct {
	ApiKey    string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
}

func GetLastFmConfiguration() LastFmConfig {
	//open file
	configFile, err := os.Open("app.config")

	if err != nil {
		fmt.Println(err)
	}

	//read file
	byteValue, _ := ioutil.ReadAll(configFile)
	var parsedConfig Config

	//parse into struct
	json.Unmarshal(byteValue, &parsedConfig)

	return parsedConfig.LastFmConfig
}
