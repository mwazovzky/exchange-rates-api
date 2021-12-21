package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Rate struct {
	Currency string
	Value    string
}

type ApiHandlers struct {
	//
}

func NewApiHandlers() *ApiHandlers {
	return &ApiHandlers{}
}

func (*ApiHandlers) Index(rw http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	date := params.Get("date")

	log.Println("Api Rates Index Request", date)

	rates := []Rate{
		{Currency: "USD", Value: "73,5000"},
		{Currency: "EUR", Value: "83,4000"},
		{Currency: "GBP", Value: "90,0000"},
		{Currency: "RUB", Value: "1,0000"},
	}

	rw.Header().Add("Content-Type", "application/json")

	e := json.NewEncoder(rw)
	err := e.Encode(rates)
	if err != nil {
		log.Println("Api Rates Index Request: Unable to marshall json")
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
	}
}
