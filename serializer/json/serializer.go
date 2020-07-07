package json

import (
	"SportDeets/sportsdeetsservice"
	"encoding/json"
	"net/http"
)

type bot struct {}

func NewJsonSerializer() sportsdeetsservice.BotSerializer{
	return &bot{}
}

func (b bot) ParseRepositoryRequest(r *http.Request) (*sportsdeetsservice.BotRequest, error) {
	var request sportsdeetsservice.BotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	return &request,nil
}
