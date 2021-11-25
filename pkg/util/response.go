package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Msg string `json:"msg"`
}

type ResponseSuccess struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewApiResponse(w http.ResponseWriter, status int, r interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(r)
}
