package actions

import (
	"net/http"

	"github.com/SmurfsAtWork/azrael/cfmt"
	"github.com/SmurfsAtWork/azrael/config"
)

type CreateSmurfParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateSmurfPayload struct {
	Id     uint   `json:"id"`
	NanoId string `json:"nano_id"`
}

func (a *Actions) CreateSmurf(params CreateSmurfParams) error {
	payload, err := makeRequest[CreateSmurfParams, CreateSmurfPayload](makeRequestConfig[CreateSmurfParams]{
		method:   http.MethodPost,
		endpoint: "/v1/smurf",
		headers: map[string]string{
			"Authorization": config.SessionToken(),
		},
		body: params,
	})
	if err != nil {
		return err
	}

	cfmt.White().Bold().Print("Created a new Smurf with the id ")
	cfmt.Blue().Print(payload.Id)
	cfmt.Reset().Print(" and nano id ")
	cfmt.Green().Println(payload.NanoId)

	return nil
}
