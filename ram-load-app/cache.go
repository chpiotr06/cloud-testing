package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func connectCache(cfg *Config) *memcache.Client {
	mc := memcache.New(fmt.Sprintf("%s:%d", cfg.CacheConfig.Host, cfg.CacheConfig.Port))

	return mc
}