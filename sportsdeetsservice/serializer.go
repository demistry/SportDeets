package sportsdeetsservice

import "net/http"

type BotSerializer interface {
	ParseRepositoryRequest(r *http.Request) (*BotRequest,error)
}