package main

import (
	"encoding/json"
	"os"
)

type AppConfiguration struct {
	DbConfig DbConfiguration
	Token    string
}

type DbConfiguration struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	SslMode  string
}

var configuration AppConfiguration

func GetConfiguration() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = AppConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
}
