package user

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram/types"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

// Module contains the state for an instance of this module.
type Module struct {
	App *app.MyApp
}

// Info returns basic information about this module.
func (Module) Info() app.ModuleInfo {
	return app.ModuleInfo{
		Name: "Misc",
	}
}

// Commands returns a list of telegram provided by this module.
func (m Module) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "ping",
			Description: "ping the bot.",
			Func:        m.ping,
		},
		{
			Name:        "about",
			Description: "about the bot.",
			Func:        m.about,
		},
		{
			Name:        "start",
			Description: "start the bot.",
			Func:        m.start,
		},
	}
}

func (m Module) Messages() []types.Message {
	return []types.Message{
		{
			Name:        "antispam",
			Description: "Detect user without username",
			Filter:      message.All,
			Func:        m.antispam,
			Order:       0,
			Async:       false,
		},
	}
}

func (m Module) Callbacks() []types.Callback {
	return []types.Callback{
		{
			Name:        "HelpCallback",
			Description: "",
			Callback:    `help\((.+?)\)`,
			Func:        m.helpCallback,
		},
	}
}

// NewModule returns a new instance of this module.
func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{
		App: bot,
	}, nil
}

func init() {
	app.RegisterModule("Misc", NewModule)
}
