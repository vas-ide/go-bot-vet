package care

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vas-atc/go-bot-vet/internal/app/path"
	"github.com/vas-atc/go-bot-vet/internal/service/veterinary/care"
)

type VeterinaryCareCommander struct {
	bot         *tgbotapi.BotAPI
	careService *care.Service
}

func NewVeterinaryCareCommander(
	bot *tgbotapi.BotAPI,
) *VeterinaryCareCommander {
	careService := care.NewService()

	return &VeterinaryCareCommander{
		bot:         bot,
		careService: careService,
	}
}

func (c *VeterinaryCareCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("VeterinaryCareCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *VeterinaryCareCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
