package main

import "net/http"

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Input   string `json:"input"`
	Rank    int    `json:"rank"`
	Notice  string `json:"notice"`
	AsOf    string `json:"as_of"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	retVal := ApiResponse{}
	retVal.Notice = "Free for light, non-commericial use. For heavy or commercial use, please contact us."

	input := r.URL.Query().Get("domain")
	if input == "" {
		retVal.Success = false
		retVal.Message = "domain query parameter is required"
		handleJson(w, r, retVal)
		return
	}
	retVal.Input = input

	domain, pureErr := purifyDomain(input)
	if pureErr != nil {
		retVal.Success = false
		retVal.Message = pureErr.Error()
		handleJson(w, r, retVal)
		return
	}

	retVal.Domain = domain
	rank, ok := rankings[domain]
	if !ok {
		retVal.Success = false
		retVal.Message = "Domain not ranked"
		handleJson(w, r, retVal)
		return
	}

	retVal.AsOf = rankDate
	retVal.Success = true
	retVal.Message = "OK"
	retVal.Rank = rank

	handleJson(w, r, retVal)
}
