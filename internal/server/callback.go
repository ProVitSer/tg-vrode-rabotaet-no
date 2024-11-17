package server

import (
	"net/http"
)

type VerificationRequest struct {
	Code string `json:"code"`
}

type CallbackRequest struct {
	Event string `json:"event"`
	Data  struct {
		ID      string `json:"id"`
		Keyword string `json:"keyword"`
		Content string `json:"content"`
	} `json:"data"`
}

func SetleCallback(w http.ResponseWriter, r *http.Request) {

}
