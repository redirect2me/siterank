package main

import (
	"net/http"
	"os"
	"strconv"
)

func main() {

	rankErr := initRankings()
	if rankErr != nil {
		logger.Error("unable to load rankings", "error", rankErr)
		return
	}

	var listenPort, portErr = strconv.Atoi(os.Getenv("PORT"))
	if portErr != nil {
		listenPort = 4000
	}
	var listenAddress = os.Getenv("ADDRESS")

	http.HandleFunc("/status.json", statusHandler)
	http.HandleFunc("/{$}", staticHandler.ServeHTTP)
	http.HandleFunc("/robots.txt", staticHandler.ServeHTTP)
	http.HandleFunc("/favicon.ico", staticHandler.ServeHTTP)
	http.HandleFunc("/favicon.svg", staticHandler.ServeHTTP)
	http.HandleFunc("/images/", staticHandler.ServeHTTP)
	http.HandleFunc("/api/rank.json", apiHandler)
	http.HandleFunc("/api/multiple.json", multiHandler)
	http.HandleFunc("/api/random.json", randomHandler)

	err := http.ListenAndServe(listenAddress+":"+strconv.Itoa(listenPort), nil)
	if err != nil {
		logger.Error("unable to listen", "address", listenAddress, "port", listenPort, "error", err)
	}
}
