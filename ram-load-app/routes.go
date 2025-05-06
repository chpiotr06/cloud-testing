package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (s *server) health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
}

func (s *server) getSession(w http.ResponseWriter, req *http.Request) {
	param := req.PathValue("uuid")

	sess := Session{}
	err := sess.getFromCache(s.cache, param)
	if err != nil {
		Warn(err)
		jsonResponse(w, resp{Msg: "Unable to get session"}, http.StatusBadRequest)
	}

	jsonResponse(w, &sess, 200)
}

func (s *server) postSession(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	userSession := new(SessionDto)
	err := decoder.Decode(&userSession)
	if err != nil {
		Warn(err)
		jsonResponse(w, resp{Msg: "Unable to decode session"}, http.StatusBadRequest)
	}

	now := time.Now()
	serverSession := Session{
		Uuid:     uuid.New().String(),
		Username: userSession.Username,
		Email:    userSession.Email,
		Iat:      now.Format(time.RFC3339),
		Exp:      now.Add(time.Hour * 24).Format(time.RFC3339),
	}

	err = serverSession.insertToCache(s.cache, int32(s.cfg.CacheConfig.ExpirationTime))
	if err != nil {
		Warn(err)
		jsonResponse(w, resp{Msg: "Unable to set session in cache"}, http.StatusBadRequest)
	}

	jsonResponse(w, &serverSession, 201)
}
