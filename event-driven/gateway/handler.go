package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type MessageDTO struct {
	Message string `json:"message"`
}

func (app *APP) Version(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseNotFound(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(VERSION))
}

func (app *APP) Message(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		ResponseNotFound(w)
		return
	}

	var msg MessageDTO

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		ResponseServerError(w)
		return
	}

	// validation
	msg.Message = strings.TrimSpace(msg.Message)

	if msg.Message == "" {
		ResponseBadRequest(w)
		LogWarn(map[string]interface{}{
			"IpAddress": r.RemoteAddr,
			"Message":   "bad request",
		})
		return
	}
	// end validation

	if err := app.RabbitMQ.Publish([]byte(msg.Message)); err != nil {
		ResponseServerError(w)
		LogError(err, map[string]interface{}{
			"Message": "cannot publish data",
		})
		return
	}

	json.NewEncoder(w).Encode(struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		"Success",
		"notification has been sent",
	})
}
