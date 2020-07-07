package sportsdeetsservice



type botService struct {
	botRepository BotRepository
}

func NewBotService(repository BotRepository) BotService{
	return &botService{botRepository: repository}
}

func (b botService) SendBotResponse(message BotMessage) {
	b.SendBotResponse(message)
}

func (b botService) ParseBotRawInput(text string) string {
	return b.ParseBotRawInput(text)
}
