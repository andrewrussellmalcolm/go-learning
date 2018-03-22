package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//
type Config struct {
	Database struct {
		DBName   string `json:"db_name,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}
	Server struct {
		CertPath string
		KeyPath  string
		Port     string
	}
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
