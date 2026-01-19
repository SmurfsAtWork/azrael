package actions

import (
	"net/http"

	"github.com/SmurfsAtWork/azrael/config"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginPayload struct {
	SessionToken string `json:"session_token"`
}

func (a *Actions) Login(params LoginParams) error {
	payload, err := makeRequest[LoginParams, LoginPayload](makeRequestConfig[LoginParams]{
		method:   http.MethodPost,
		endpoint: "/v1/login",
		body:     params,
	})
	if err != nil {
		return err
	}
	config.SetSessionToken(payload.SessionToken)

	return nil
}
