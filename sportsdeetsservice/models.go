package sportsdeetsservice


type Bot struct{
	Name string `json:"name"`
}

type BotRequest struct {
	UpdateId int `json:"update_id"`
	Message BotMessage `json:"message"`
}

type BotMessage struct {
	Text string `json:"text"`
	Chat Chat `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}