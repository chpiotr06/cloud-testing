package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func WarnAndRespond(w http.ResponseWriter, err error, message string, status int) {
	Warn(err)
	jsonResponse(w, resp{Msg: message}, status)
}

func WarnAndReturn(err error, message string) error {
		Warn(err)
		if err != nil {
			return fmt.Errorf("%s : %w", message, err)
		}
		return nil
}

func Warn(err error){
	if err != nil {
		slog.Warn(fmt.Sprintf("%s", err))
	}
}


func Fail(err error){
	if err != nil {
		slog.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
}