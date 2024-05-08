package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var rankings map[string]int = make(map[string]int)
var rankDate string

func getLastModified(resp *http.Response) time.Time {
	lastModified, lmErr := http.ParseTime(resp.Header.Get("Last-Modified"))
	if lmErr == nil {
		return lastModified
	}
	dateFromHeader, dateErr := time.Parse("2006-01-02 15:04:05", resp.Header.Get("Date"))
	if dateErr == nil {
		return dateFromHeader
	}

	return time.Now()
}

func initRankings() error {

	urlStr := os.Getenv("URL")
	if urlStr == "" {
		urlStr = "file:top-1m.csv.zip"
	}

	logger.Info("Loading rankings", "url", urlStr)
	url, urlErr := url.Parse(urlStr)
	if urlErr != nil {
		logger.Error("Invalid URL", "error", urlErr, "url", urlStr)
		return urlErr
	}
	logger.Debug("Parsed rankings URL", "url", urlStr, "parsed", url)

	var body []byte
	var readErr error
	if url.Scheme == "file" {
		var filePath string
		if url.Path == "" {
			filePath = url.Opaque
		} else {
			filePath = url.Path
		}
		stats, statErr := os.Stat(filePath)
		if statErr != nil {
			logger.Error("Failed to stat rankings file", "error", statErr, "url", urlStr, "path", filePath)
			return statErr
		}
		rankDate = stats.ModTime().Format("2006-01-02 15:04:05")
		body, readErr = os.ReadFile(filePath)
	} else if url.Scheme == "http" || url.Scheme == "https" {
		resp, err := http.Get(urlStr)
		if err != nil {
			logger.Error("Failed to download rankings", "error", err, "url", urlStr)
			return err
		}
		defer resp.Body.Close()

		rankDate = getLastModified(resp).Format("2006-01-02 15:04:05")
		body, readErr = io.ReadAll(resp.Body)

		logger.Info("Downloaded rankings", "url", urlStr, "size", len(body), "contentLength", resp.ContentLength)
	} else {
		logger.Error("Unsupported URL scheme", "scheme", url.Scheme, "url", urlStr)
		return fmt.Errorf("unsupported URL scheme %s", url.Scheme)
	}
	if readErr != nil {
		logger.Error("Failed to read rankings", "error", readErr, "url", urlStr)
		return readErr
	}

	zipData, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		logger.Error("Failed to read zip file", "error", err, "url", urlStr)
		return err
	}

	zipFile, zipFileErr := zipData.File[0].Open()
	if zipFileErr != nil {
		logger.Error("Failed to open file from zip", "error", zipFileErr, "url", urlStr)
		return zipFileErr
	}

	scanner := bufio.NewScanner(zipFile)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			logger.Error("Invalid ranking line", "line", line, "lineNumber", lineNo, "url", urlStr)
			continue
		}
		rank, err := strconv.Atoi(parts[0])
		if err != nil {
			logger.Error("Invalid ranking rank", "line", line, "lineNumber", lineNo, "url", urlStr)
			continue
		}
		rankings[parts[1]] = rank
	}
	if err := scanner.Err(); err != nil {
		logger.Error("Failed to scan zip file", "error", err, "url", urlStr)
		return err
	}
	logger.Info("Loaded rankings", "url", urlStr, "count", len(rankings))

	return nil
}
