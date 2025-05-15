package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"scope3/go-interview/config"

	log "github.com/sirupsen/logrus"
)

const (
	BaseURL = "https://api.scope3.com/v2"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		apiKey:     config.Config.ApiKey,
	}
}

func (c *Client) Measure(inventoryIds []string, requestDate string) (*MeasureResponse, error) {
	rows := make([]*RequestRow, len(inventoryIds))
	for i, inventoryId := range inventoryIds {
		rows[i] = NewRequestRow(inventoryId, requestDate)
	}

	jsonData, err := json.Marshal(map[string]any{
		"rows": rows,
	})

	if err != nil {
    	return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	params := NewMeasureQueryParams()
	requestUrl := fmt.Sprintf("%s?%s", fmt.Sprintf("%s/measure", BaseURL), params.ToQueryString())

	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	log.WithFields(log.Fields{
		"URL": requestUrl,
		"Body": jsonData,
	}).Debug("Request Details")

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	log.WithFields(log.Fields{
		"Headers": request.Header,
	}).Debug("Request Headers")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer response.Body.Close()

    bodyBytes, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d, body: %s",
            response.StatusCode, string(bodyBytes))
    }

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error Response Body: %d\n", response.Body)
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var result MeasureResponse
    if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w, body: %s",
            err, string(bodyBytes))
    }

	return &result, nil
}
