package main

import (
	"encoding/base64"
	"io"
	"net/http"
	"strings"
	"time"
)

func fetchSubscription(link string) ([]string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	decoded, err := decodeBase64(strings.TrimSpace(string(body)))
	if err != nil {
		decoded = string(body)
	}

	lines := strings.Split(strings.ReplaceAll(decoded, "\r\n", "\n"), "\n")
	return lines, nil
}

func decodeBase64(input string) (string, error) {
	decStd := base64.StdEncoding.WithPadding(base64.NoPadding)

	data, err := decStd.DecodeString(input)
	if err != nil {
		decUrl := base64.URLEncoding.WithPadding(base64.NoPadding)

		data, err = decUrl.DecodeString(input)
		if err != nil {
			return "", err
		}
	}
	return string(data), nil
}
