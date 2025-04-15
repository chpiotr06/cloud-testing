package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	configPath := flag.String("config", "./config.json", "Path to the config json file")
	flag.Parse()

	config := new(Config)
	config.LoadConfig(*configPath)

	s := newServer(config)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /session", s.getSession)
	mux.HandleFunc("POST /session", s.postSession)

	err := http.ListenAndServe(fmt.Sprintf(":%d",config.AppPort), mux)
	Fail(err)
}