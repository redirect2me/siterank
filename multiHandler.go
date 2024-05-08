package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type MultiResponse struct {
	Success  bool           `json:"success"`
	Message  string         `json:"message"`
	Results  map[string]int `json:"results"`
	Messages []string       `json:"messages"`
	Notice   string         `json:"notice"`
}

var re = regexp.MustCompile("[ ,\t\n\r]+")

func multiHandler(w http.ResponseWriter, r *http.Request) {

	retVal := MultiResponse{}
	retVal.Notice = "Free for light, non-commericial use. For heavy or commercial use, please contact us."
	retVal.Results = make(map[string]int)

	input := r.URL.Query().Get("domains")
	if input == "" {
		retVal.Success = false
		retVal.Message = "domains query parameter is required"
		handleJson(w, r, retVal)
		return
	}

	domains := re.Split(input, -1)

	for _, rawDomain := range domains {
		domain, pureErr := purifyDomain(rawDomain)
		if pureErr != nil {
			retVal.Messages = append(retVal.Messages, fmt.Sprintf("Error: %s does not appear to be valid (%s)", rawDomain, pureErr.Error()))
			continue
		}
		if rawDomain != domain {
			retVal.Messages = append(retVal.Messages, fmt.Sprintf("Warning: %s was purified to %s", rawDomain, domain))
		}

		rank, ok := rankings[domain]
		if ok {
			retVal.Results[domain] = rank
		} else {
			retVal.Results[domain] = -1
		}
	}

	retVal.Success = true
	retVal.Message = "OK"

	handleJson(w, r, retVal)
}
