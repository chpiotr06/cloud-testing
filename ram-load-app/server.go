package main

import (
	"encoding/json"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
)

type server struct {
	cache *memcache.Client
	cfg *Config
}

type resp struct {
	Msg string `json:"msg"`
}

func newServer(cfg *Config) *server {
	s := server{cfg: cfg}
	s.cache = connectCache(cfg)

	return &s
}

func jsonResponse(w http.ResponseWriter, value any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(value)
}