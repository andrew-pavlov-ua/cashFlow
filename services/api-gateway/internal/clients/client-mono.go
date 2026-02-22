package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	utils "github.com/andrew-pavlov-ua/pkg"
	"github.com/andrew-pavlov-ua/pkg/logger"
	"github.com/andrew-pavlov-ua/pkg/models"
)

type MonoClient struct {
	client *http.Client
}

func NewMonoClient() *MonoClient {
	return &MonoClient{
		client: &http.Client{},
	}
}

func (s *MonoClient) SetUpMonoWebhook(token string) error {
	webhookURL := "https://" + utils.Getenv("GATEWAY_HOST", "http://localhost:8080") + "/mono-webhook"
	url := "https://api.monobank.ua/personal/webhook"

	logger.Infof("Setting up MonoBank webhook | url=%s", webhookURL)

	payload := map[string]string{
		"webHookUrl": webhookURL,
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		logger.Error("Failed to marshal webhook payload", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		logger.Error("Failed to create webhook setup request", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token", token)

	resp, err := s.client.Do(req)
	if err != nil {
		logger.Error("Failed to send webhook setup request", err)
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		logger.Errorf(
			"Monobank webhook setup failed | status=%d | body=%s",
			resp.StatusCode,
			string(respBody),
		)
		return fmt.Errorf("monobank webhook setup failed: %s", string(respBody))
	}

	logger.Infof("Monobank webhook successfully set | url=%s", webhookURL)

	return nil
}

func (s *MonoClient) FetchClient(token string) (*models.Client, error) {
	url := "https://api.monobank.ua/personal/client-info"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("Failed to create request for MonoBank API", err)
		return nil, err
	}
	req.Header.Set("X-Token", token)

	resp, err := s.client.Do(req)
	if err != nil {
		logger.Error("Failed to fetch client info from MonoBank", err)
		return nil, err
	}
	defer resp.Body.Close()

	var clientData models.Client
	err = json.NewDecoder(resp.Body).Decode(&clientData)
	if err != nil {
		logger.Error("Failed to decode MonoBank response", err)
		return nil, err
	}

	// delete this after testing
	logger.Infof("Fetched client data from MonoBank: %+v", clientData)

	return &clientData, nil
}
