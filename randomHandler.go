package main

import "net/http"

type RandomResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Notice  string `json:"notice"`
}

func randomHandler(w http.ResponseWriter, r *http.Request) {

	retVal := RandomResponse{}
	retVal.Notice = "Free for light, non-commericial use. For heavy or commercial use, please contact us."

	// pick a random key from the rankings map: since go tries to be non-deterministic, we'll just pick the first one
	for k := range rankings {
		retVal.Domain = k
		break
	}

	retVal.Success = true
	retVal.Message = "OK"

	handleJson(w, r, retVal)
}
