package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

var config configuration

func init() {
	err := envconfig.Process("CUEWATCH", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
}

type configuration struct {
	PlantUMLServerAddress string `envconfig:"PLANTML_SERVER" default:"http://localhost:8080"`
	ListenAddress         string `envconfig:"ADDR" default:":9090"`
	PollingDir            string `envconfig:"POLLING_DIR" default:"./"`
	RecursivePoll         bool   `envconfig:"RECURSIVE" default:"true"`
}

func main() {
	mux, err := processConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(config.ListenAddress, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
