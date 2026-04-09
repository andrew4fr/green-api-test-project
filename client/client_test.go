package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"green-api-test-project/models"
)

func TestGetInstanceSettings(t *testing.T) {
	expected := models.SettingsResponse{
		Wid: ptr("79991234567@c.us"),
		TypeInstance: ptr("v3"),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/waInstance123/getSettings/token456" {
			t.Errorf("unexpected URL: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := New(server.URL, "")
	result, err := client.GetInstanceSettings(context.Background(), "123", "token456")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Wid == nil || *result.Wid != *expected.Wid {
		t.Errorf("expected wid %v, got %v", expected.Wid, result.Wid)
	}
}

func TestGetInstanceState(t *testing.T) {
	expected := models.StateResponse{
		StateInstance: func() *models.StateResponseStateInstance {
			v := models.Authorized
			return &v
		}(),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/waInstance123/getStateInstance/token456" {
			t.Errorf("unexpected URL: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := New(server.URL, "")
	result, err := client.GetInstanceState(context.Background(), "123", "token456")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.StateInstance == nil || *result.StateInstance != *expected.StateInstance {
		t.Errorf("expected state %v, got %v", *expected.StateInstance, *result.StateInstance)
	}
}

func TestGetAccountSettings(t *testing.T) {
	expected := models.AccountSettingsResponse{
		Phone: ptr("79991234567"),
		ChatId: ptr("10000000"),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/waInstance123/getAccountSettings/token456" {
			t.Errorf("unexpected URL: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := New(server.URL, "")
	result, err := client.GetAccountSettings(context.Background(), "123", "token456")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Phone == nil || *result.Phone != *expected.Phone {
		t.Errorf("expected phone %v, got %v", expected.Phone, result.Phone)
	}
}

func TestSendMessage(t *testing.T) {
	expected := models.SendMessageResponse{
		IdMessage: ptr("1763115112345"),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/waInstance123/sendMessage/token456" {
			t.Errorf("unexpected URL: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}

		var reqBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}

		if reqBody["chatId"] != "10000000" {
			t.Errorf("expected chatId 10000000, got %v", reqBody["chatId"])
		}
		if reqBody["message"] != "Hello" {
			t.Errorf("expected message Hello, got %v", reqBody["message"])
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := New(server.URL, "")
	body := models.SendMessageJSONRequestBody{
		ChatId:  "10000000",
		Message: "Hello",
	}
	result, err := client.SendMessage(context.Background(), "123", "token456", body)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IdMessage == nil || *result.IdMessage != *expected.IdMessage {
		t.Errorf("expected idMessage %v, got %v", expected.IdMessage, result.IdMessage)
	}
}

func TestSendFile(t *testing.T) {
	expected := models.SendFileResponse{
		IdMessage: ptr("1763115112345"),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/waInstance123/sendFileByUrl/token456" {
			t.Errorf("unexpected URL: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}

		var reqBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}

		if reqBody["chatId"] != "10000000" {
			t.Errorf("expected chatId 10000000, got %v", reqBody["chatId"])
		}
		if reqBody["urlFile"] != "https://example.com/file.pdf" {
			t.Errorf("expected urlFile https://example.com/file.pdf, got %v", reqBody["urlFile"])
		}
		if reqBody["fileName"] != "file.pdf" {
			t.Errorf("expected fileName file.pdf, got %v", reqBody["fileName"])
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := New(server.URL, "")
	body := models.SendFileJSONRequestBody{
		ChatId:   "10000000",
		UrlFile:  "https://example.com/file.pdf",
		FileName: "file.pdf",
	}
	result, err := client.SendFile(context.Background(), "123", "token456", body)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IdMessage == nil || *result.IdMessage != *expected.IdMessage {
		t.Errorf("expected idMessage %v, got %v", expected.IdMessage, result.IdMessage)
	}
}

func TestRequest_BadMethod(t *testing.T) {
	client := New("http://localhost:8080", "")
	err := client.request(context.Background(), "test", "PUT", nil, nil, "123", "token")
	if err == nil {
		t.Error("expected error for bad method")
	}
	if err.Error() != "bad http method: PUT" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRequest_NonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New(server.URL, "")
	err := client.request(context.Background(), "test", http.MethodGet, nil, nil, "123", "token")
	if err == nil {
		t.Error("expected error for non-OK status")
	}
}

func TestNew(t *testing.T) {
	c := New("api-url", "media-url")
	if c == nil {
		t.Fatal("expected non-nil client")
	}
	if c.api.APIURL != "api-url" {
		t.Errorf("expected APIURL api-url, got %s", c.api.APIURL)
	}
	if c.api.MediaURL != "media-url" {
		t.Errorf("expected MediaURL media-url, got %s", c.api.MediaURL)
	}
	if c.httpClient == nil {
		t.Error("expected non-nil httpClient")
	}
}

func ptr[T any](v T) *T {
	return &v
}
