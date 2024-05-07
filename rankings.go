package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var rankings map[string]int = make(map[string]int)

func initRankings() error {

	url := os.Getenv("URL")
	if url == "" {
		url = "https://tranco-list.eu/top-1m.csv.zip"
		//"https://s3.amazonaws.com/alexa-static/top-1m.csv.zip"
	}

	logger.Info("Downloading rankings", "url", url)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Failed to download rankings", "error", err, "url", url)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	zipData, err := zip.NewReader(bytes.NewReader(body), resp.ContentLength)
	if err != nil {
		logger.Error("Failed to read zip file", "error", err, "url", url)
		return err
	}

	zipFile, zipFileErr := zipData.File[0].Open()
	if zipFileErr != nil {
		logger.Error("Failed to open file from zip", "error", zipFileErr, "url", url)
		return zipFileErr
	}

	scanner := bufio.NewScanner(zipFile)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			logger.Error("Invalid ranking line", "line", line, "lineNumber", lineNo, "url", url)
			continue
		}
		rank, err := strconv.Atoi(parts[0])
		if err != nil {
			logger.Error("Invalid ranking rank", "line", line, "lineNumber", lineNo, "url", url)
			continue
		}
		rankings[parts[1]] = rank
	}
	if err := scanner.Err(); err != nil {
		logger.Error("Failed to scan zip file", "error", err, "url", url)
		return err
	}
	logger.Info("Loaded rankings", "url", url, "count", len(rankings))

	return nil
}
