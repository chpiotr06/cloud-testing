package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppPort int `json:"appPort"`	
	CacheConfig CacheConfig `json:"cache"`
}

type CacheConfig struct {
	Host string `json:"host"`
	Port int `json:"port"`
	ExpirationTime int `json:"expirationTime"`
}

func (c *Config) LoadConfig(path string) {
	data, err := os.ReadFile(path)
	Fail(err)

	jsonErr := json.Unmarshal(data, c)
	Fail(jsonErr)
}