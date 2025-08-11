package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func connectCache(cfg *Config) *memcache.Client {
	address := fmt.Sprintf("%s:%d", cfg.CacheConfig.Host, cfg.CacheConfig.Port)
	var mc *memcache.Client

	for i := 0; i < 10; i++ {
			mc = memcache.New(address)
			err := mc.Set(&memcache.Item{Key: "ping", Value: []byte("pong")})
			if err == nil {
					fmt.Printf("Connected to Memcached at %s\n", address)
					return mc
			}
			fmt.Printf("Retrying Memcached connection (%d/10): %v\n", i+1, err)
			time.Sleep(2 * time.Second)
	}

	log.Fatalf("Could not connect to Memcached at %s after 10 retries", address)
	return nil
}