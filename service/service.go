package service

import (
	"context"
	"green-api-test-project/client"
	"green-api-test-project/models"
)

type Service struct {
	client *client.Client
}

func NewService(client *client.Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) GetInstanceSettings(ctx context.Context, instanceID, apiToken string) (*models.SettingsResponse, error) {
	result, err := s.client.GetInstanceSettings(ctx, instanceID, apiToken)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) GetInstanceState(ctx context.Context, instanceID, apiToken string) (*models.StateResponse, error) {
	result, err := s.client.GetInstanceState(ctx, instanceID, apiToken)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) GetAccountSettings(ctx context.Context, instanceID, apiToken string) (*models.AccountSettingsResponse, error) {
	result, err := s.client.GetAccountSettings(ctx, instanceID, apiToken)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) SendMessage(ctx context.Context, instanceID, apiToken string, body models.SendMessageJSONRequestBody) (*models.SendMessageResponse, error) {
	result, err := s.client.SendMessage(ctx, instanceID, apiToken, body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) SendFile(ctx context.Context, instanceID, apiToken string, body models.SendFileJSONRequestBody) (*models.SendFileResponse, error) {
	result, err := s.client.SendFile(ctx, instanceID, apiToken, body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
