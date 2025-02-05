package veterinary

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-ide/go-bot-vet/internal/app/commands/veterinary/care"
	"github.com/vas-ide/go-bot-vet/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type VeterinaryCommander struct {
	bot           *tgbotapi.BotAPI
	careCommander Commander
}

func NewVeterinaryCommander(
	bot *tgbotapi.BotAPI,
) *VeterinaryCommander {
	return &VeterinaryCommander{
		bot: bot,
		// careCommander
		careCommander: care.NewVeterinaryCareCommander(bot),
	}
}

func (c *VeterinaryCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Care {
	case "care":
		c.careCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("VeterinaryCommander.HandleCallback: unknown care - %s", callbackPath.Care)
	}
}

func (c *VeterinaryCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Care {
	case "care":
		c.careCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("VeterinaryCommander.HandleCommand: unknown care - %s", commandPath.Care)
	}
}
