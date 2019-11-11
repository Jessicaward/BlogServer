package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	LastFmConfig LastFmConfig `json:"lastfm"`
}

type LastFmConfig struct {
    ApiKey      string `json:"apiKey"`
    ApiSecret   string `json:"apiSecret"`
}

func GetLastFmConfiguration() LastFmConfig {
	//open file
	configFile, err := os.Open("app.config")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened config file")

	//read file
	byteValue, _ := ioutil.ReadAll(configFile)
	var parsedConfig Config

	fmt.Println(parsedConfig)

	//parse into struct
	json.Unmarshal(byteValue, &parsedConfig)

	//log contents to console
	fmt.Println("key: " + parsedConfig.LastFmConfig.ApiKey)
	fmt.Println("secret: " + (parsedConfig.LastFmConfig.ApiSecret))

	return parsedConfig.LastFmConfig
}