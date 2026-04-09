package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"green-api-test-project/models"
	"io"
	"net/http"
	"time"
)

type GreenAPI struct {
	APIURL           string
	MediaURL         string
}

type Client struct {
	api            GreenAPI
	httpClient     *http.Client
}

func New(apiUrl, mediaUrl string) *Client {
	api := GreenAPI{
		APIURL:           apiUrl,
		MediaURL:         mediaUrl,
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		api:        api,
		httpClient: httpClient,
	}
}

func (c *Client) GetInstanceSettings(ctx context.Context, instanceID, apiToken string) (*models.SettingsResponse, error) {
	var settings models.SettingsResponse
	if err := c.request(ctx, "getSettings", http.MethodGet, nil, &settings, instanceID, apiToken); err != nil {
		return nil, fmt.Errorf("error on get instance settings: %w", err)
	}

	return &settings, nil
}

func (c *Client) GetInstanceState(ctx context.Context, instanceID, apiToken string) (*models.StateResponse, error) {
	var state models.StateResponse
	if err := c.request(ctx, "getStateInstance", http.MethodGet, nil, &state, instanceID, apiToken); err != nil {
		return nil, fmt.Errorf("error on get instance state: %w", err)
	}

	return &state, nil
}

func (c *Client) GetAccountSettings(ctx context.Context, instanceID, apiToken string) (*models.AccountSettingsResponse, error) {
	var settings models.AccountSettingsResponse
	if err := c.request(ctx, "getAccountSettings", http.MethodGet, nil, &settings, instanceID, apiToken); err != nil {
		return nil, fmt.Errorf("error on get account settings: %w", err)
	}

	return &settings, nil
}

func (c *Client) SendMessage(ctx context.Context, instanceID, apiToken string, body models.SendMessageJSONRequestBody) (*models.SendMessageResponse, error) {
	params := map[string]interface{}{
		"chatId": body.ChatId,
		"message": body.Message,
	}

	var resp models.SendMessageResponse
	if err := c.request(ctx, "sendMessage", http.MethodPost, params, &resp, instanceID, apiToken); err != nil {
		return nil, fmt.Errorf("error on send message: %w", err)
	}

	return &resp, nil
}

func (c *Client) SendFile(ctx context.Context, instanceID, apiToken string, body models.SendFileJSONRequestBody) (*models.SendFileResponse, error) {
	params := map[string]interface{}{
		"chatId": body.ChatId,
		"urlFile": body.UrlFile,
		"fileName": body.FileName,
		"caption": body.Caption,
	}

	var resp models.SendFileResponse
	if err := c.request(ctx, "sendFileByUrl", http.MethodPost, params, &resp, instanceID, apiToken); err != nil {
		return nil, fmt.Errorf("error on send file: %w", err)
	}

	return &resp, nil
}

func (c *Client) request(ctx context.Context, endpoint, method string, params map[string]interface{}, data interface{}, instanceID, apiToken string) error {
	url := fmt.Sprintf("%s/waInstance%s/%s/%s", c.api.APIURL, instanceID, endpoint, apiToken)
	var req *http.Request
	var err error

	switch method {
	case http.MethodPost:
		body, err := json.Marshal(params)
		if err != nil {
			return fmt.Errorf("error on marshal body params. endpoint: %s, error: %w", endpoint, err)
		}

		req, err = http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			url,
			bytes.NewBuffer(body),
		)
		if err != nil {
			return fmt.Errorf("error on create post request. endpoint: %s, error: %w", endpoint, err)
		}
	case http.MethodGet:
		req, err = http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			url,
			nil,
		)
		if err != nil {
			return fmt.Errorf("error on create get request. endpoint: %s, error: %w", endpoint, err)
		}
	default:
		return fmt.Errorf("bad http method: %s", method)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error on do request. endpoint: %s, error: %w", endpoint, err)
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error on request. endpoint: %s, status: %s", endpoint, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response. endpoint: %s, method: %s, error: %w", endpoint, method, err)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("error unmarshaling response. endpoint: %s, method: %s, error: %w", endpoint, method, err)
	}

	return nil
}
