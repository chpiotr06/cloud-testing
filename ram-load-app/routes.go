package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (s *server) health(w http.ResponseWriter, _ *http.Request) {
	jsonResponse(w, &s.cfg, 200)
}

func (s *server) getSession(w http.ResponseWriter, req *http.Request) {
	param := req.PathValue("uuid")

	sess := Session{}
	err := sess.getFromCache(s.cache, param)

	if err != nil {
		WarnAndRespond(w, err, "Unable to get session", http.StatusBadRequest)
		return
	}

	jsonResponse(w, &sess, 200)
}

func (s *server) postSession(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	userSession := new(SessionDto)
	err := decoder.Decode(&userSession)
	if err != nil {
		WarnAndRespond(w, err, "Unable to decode session", http.StatusBadRequest)
		return 
	}

	now := time.Now()
	serverSession := Session{
		Uuid:     uuid.New().String(),
		Username: userSession.Username,
		Email:    userSession.Email,
		Iat:      now.Format(time.RFC3339),
		Exp:      now.Add(time.Hour * 24).Format(time.RFC3339),
		Payload: strings.Repeat("A", 1<<12),
	}

	err = serverSession.insertToCache(s.cache, int32(s.cfg.CacheConfig.ExpirationTime))
	if err != nil {
		WarnAndRespond(w, err, "Unable to set session in cache", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, &serverSession, 201)
}
