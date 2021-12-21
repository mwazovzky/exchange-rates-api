package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"exchange-rates/pkg/rater"
)

type ApiHandlers struct {
	r *rater.Rater
}

func NewApiHandlers(r *rater.Rater) *ApiHandlers {
	return &ApiHandlers{r}
}

func (ah *ApiHandlers) Index(rw http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	date := params.Get("date")

	log.Println("Api Rates Index Request", date)

	var rates rater.Rates
	err := ah.r.Load(date, &rates)
	if err != nil {
		log.Println("Api Rates Index Request: Unable load exchange rates")
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
	}

	rw.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(rw)
	err = e.Encode(rates.Currencies)
	if err != nil {
		log.Println("Api Rates Index Request: Unable to marshall json")
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
	}
}
