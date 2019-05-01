package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Database struct {
	Jdbcurl string `json:"jdbcurl"`
	Driver  string `json:"driver"`
	Uname   string `json:"uname"`
	Pwd     string `json:"pwd"`
}

type Stream struct {
	Topic       string `json:"topic"`
	Registryurl string `json:"registryurl"`
}

// Configuration structure includes
// 1. server - app server
// 2. database - db server
// 3. stream - Kafka server
type Configuration struct {
	Server Server   `json:"server"`
	Db     Database `json:"database"`
	Stream Stream   `json:"stream"`
}

func getJsonAsString() ([]byte, error) {

	props := Configuration{
		Server{"localhost", "8080"},
		Database{"jdbcurl", "driver", "uname", "pwd"},
		Stream{"topic", "registryurl"},
	}
	js, err := json.MarshalIndent(props, "", "\t")
	return js, err
}

func main2() {
	json, err := getJsonAsString()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("json %s : ", json)
}

func main() {

	fmt.Println("Initializing Rest server.....")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json, err := getJsonAsString()
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(json))
	})

	http.ListenAndServe(":8080", nil)
}
