package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func randomTemperature(w http.ResponseWriter, req *http.Request) {
	type Resp struct {
		Value float64 `json:"value"`
	}

	resp, err := json.Marshal(&Resp{Value: float64(rand.Intn(100)) + rand.Float64()})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
