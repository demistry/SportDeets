package sportsdeetsservice



type BotService interface{
	SendBotResponse(message BotMessage)
	ParseBotRawInput(text string) string
}