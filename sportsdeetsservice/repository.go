package sportsdeetsservice


type BotRepository interface {
	BotCommands() []string
}