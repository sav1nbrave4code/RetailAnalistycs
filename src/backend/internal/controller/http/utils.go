package http

import (
	"encoding/json"
	"github.com/gocarina/gocsv"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type Error struct {
	Error string `json:"Error"`
}

func WriteError(w http.ResponseWriter, error error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	e := &Error{
		Error: error.Error(),
	}

	byteMessage, _ := json.Marshal(e)

	if _, err := w.Write(byteMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteResponseJson(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := jsoniter.Marshal(body)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(resp); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func WriteResponseCsv(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "text/csv; charset=UTF-8")

	resp, err := gocsv.MarshalBytes(body)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(resp); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
