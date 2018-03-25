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
		DBName string `json:"db_name,omitempty"`
	} `json:"database,omitempty"`
	Server struct {
		CertPath string `json:"cert_path,omitempty"`
		KeyPath  string `json:"key_path,omitempty"`
		Port     string `json:"port,omitempty"`
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
