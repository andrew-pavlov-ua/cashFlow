package clients

import (
	"encoding/json"
	"net/http"

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
