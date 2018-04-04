package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//
type Config struct {
	Server struct {
		ServerURL string `json:"server_url,omitempty"`
	} `json:"server,omitempty"`
}

func readConfig() Config {

	f, err := os.Open("config.json")
	if err != nil {
		panic(fmt.Sprintf("error opening config file %v\n", err))
	}

	config := Config{}

	r := bufio.NewReader(f)
	json.NewDecoder(r).Decode(&config)

	return config
}
