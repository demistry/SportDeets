package telegram

import "SportDeets/sportsdeetsservice"


var botCommands = []string{"/start","/latest_fixtures"}

const botName = "@sportdeets_bot"

type TelegramRepository struct {
	BaseURL string
}

func newTelegramRepository() sportsdeetsservice.BotRepository{
	return &TelegramRepository{}
}


func (t TelegramRepository) BotCommands() []string {
	return botCommands
}

func (t TelegramRepository) SendBotResponse(message sportsdeetsservice.BotMessage) {

}

//Clean up raw bot text to remove the commands
//TODO:- Use regex to make this cleaner
func (t TelegramRepository) ParseBotRawInput(text string) string {
	for _,command := range botCommands{
		lengthOfCmd := len(command)
		if len(text) >= lengthOfCmd{
			if text[:lengthOfCmd] == command{
				text = text[lengthOfCmd:]
			}
		}
	}

	if len(text) >= len(botName){
		text = text[len(botName):]
	}

	return text
}

