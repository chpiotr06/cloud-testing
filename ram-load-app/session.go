package main

import (
	"encoding/json"

	"github.com/bradfitz/gomemcache/memcache"
)

type SessionDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Session struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Exp      string `json:"exp"`
	Iat      string `json:"iat"`
}

func (s *Session) insertToCache(memcacheClient *memcache.Client, exp int32) (err error) {
	session, err := json.Marshal(s)
	if err != nil {
		return WarnAndReturn(err, "json Marshall err")
	}

	err = memcacheClient.Set(&memcache.Item{Key: s.Uuid, Value: session, Expiration: exp})
	if err != nil {
		return WarnAndReturn(err, "memcache Set error")
	}

	return nil
}

func (s *Session) getFromCache(memecacheClient *memcache.Client, key string) (err error) {
	item, err := memecacheClient.Get(key)
	if err != nil {
		return WarnAndReturn(err, "error while getting item from memcache")
	}

	err = json.Unmarshal(item.Value, &s)
	if err != nil {
		return WarnAndReturn(err, "unmarshal failed")
	}

	return nil
}
