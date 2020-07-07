package sportsdeetsservice


type BotRepository interface {
	BotCommands() []string
	SendBotResponse(message BotMessage)
	ParseBotRawInput(text string) string
}