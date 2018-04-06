package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// Endpoint is the base URL to the Slack API.
	Endpoint = "https://slack.com/api"
)

// Client is a Slack API client.
type Client struct {
	token      string
	httpClient *http.Client
}

// NewClient creates a Client.
func NewClient(token string) *Client {
	return &Client{
		token: token,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// ChatPostMessage sends a message to a channel (chat.postMessage).
func (c *Client) ChatPostMessage(
	channel,
	text string,
) error {
	type Payload struct {
		Channel string `json:"channel"`
		Text    string `json:"text"`
	}
	payload := Payload{
		Channel: channel,
		Text:    text,
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %s", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/chat.postMessage", Endpoint),
		bytes.NewBuffer(buf),
	)
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// I don't know what the deal is with the token parameter that can be in the
	// request, but apparently we have to provide it in the Authorization header.
	// Oh, it appears to be usable when not sending JSON requests.
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error performing HTTP request: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_ = resp.Body.Close()
		return fmt.Errorf("error reading body: %s", err)
	}

	if err := resp.Body.Close(); err != nil {
		return fmt.Errorf("error closing body: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d from API", resp.StatusCode)
	}

	var p map[string]interface{}
	if err := json.Unmarshal(body, &p); err != nil {
		return fmt.Errorf("error unmarshaling body: %s", err)
	}

	ok, exists := p["ok"]
	if !exists {
		return fmt.Errorf("response did not include ok")
	}
	success, isBool := ok.(bool)
	if !isBool {
		return fmt.Errorf("response ok was not bool")
	}
	if !success {
		return fmt.Errorf("API said !ok: %+v (I sent %s)", p, buf)
	}

	return nil
}
