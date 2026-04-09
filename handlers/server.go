package handlers

import (
	"encoding/json"
	"net/http"

	"green-api-test-project/models"
	"green-api-test-project/service"
)

type Server struct {
	svc *service.Service
}

func NewServer(svc *service.Service) *Server {
	return &Server{
		svc: svc,
	}
}

func (s *Server) GetInstanceSettings(w http.ResponseWriter, r *http.Request, params GetInstanceSettingsParams) {
	instanceId := params.InstanceId
	apiToken := params.ApiToken

	result, err := s.svc.GetInstanceSettings(r.Context(), instanceId, apiToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{
			Code:    "internal_error",
			Details: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (s *Server) GetInstanceState(w http.ResponseWriter, r *http.Request, params GetInstanceStateParams) {
	instanceId := params.InstanceId
	apiToken := params.ApiToken

	result, err := s.svc.GetInstanceState(r.Context(), instanceId, apiToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{
			Code:    "internal_error",
			Details: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (s *Server) SendFile(w http.ResponseWriter, r *http.Request, params SendFileParams) {
	instanceId := params.InstanceId
	apiToken := params.ApiToken

	var body models.SendFileJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{
			Code:    "invalid_request",
			Details: err.Error(),
		})
		return
	}

	result, err := s.svc.SendFile(r.Context(), instanceId, apiToken, body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{
			Code:    "internal_error",
			Details: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (s *Server) SendMessage(w http.ResponseWriter, r *http.Request, params SendMessageParams) {
	instanceId := params.InstanceId
	apiToken := params.ApiToken

	var body models.SendMessageJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{
			Code:    "invalid_request",
			Details: err.Error(),
		})
		return
	}

	result, err := s.svc.SendMessage(r.Context(), instanceId, apiToken, body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{
			Code:    "internal_error",
			Details: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (s *Server) GetAccountSettings(w http.ResponseWriter, r *http.Request, params GetAccountSettingsParams) {
	instanceId := params.InstanceId
	apiToken := params.ApiToken

	result, err := s.svc.GetAccountSettings(r.Context(), instanceId, apiToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{
			Code:    "internal_error",
			Details: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
