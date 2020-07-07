package webhook

import (
	"SportDeets/serializer/json"
	"SportDeets/sportsdeetsservice"
)


type HookHandler interface {
	CallLyricsAPI(searchText string) (string, error)
}

type webHookHandle struct {
	botService sportsdeetsservice.BotService
}


func NewWebHookHandler(botService sportsdeetsservice.BotService) HookHandler{
	return &webHookHandle{botService: botService}
}

func (w webHookHandle) CallLyricsAPI(searchText string) (string, error) {
	//make call to external lyrics api to respond to bot
}


func (w webHookHandle) serializer() sportsdeetsservice.BotSerializer {
	return json.NewJsonSerializer()
}
