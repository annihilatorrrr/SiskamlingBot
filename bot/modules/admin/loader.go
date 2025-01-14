package admin

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram/types"
)

// Module contains the state for an instance of this module.
type Module struct {
	App *app.MyApp
}

// Info returns basic information about this module.
func (*Module) Info() app.ModuleInfo {
	return app.ModuleInfo{
		Name: "Admin",
	}
}

// Commands returns a list of telegram provided by this module.
func (m *Module) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "user",
			Description: "get user info",
			Func:        m.getUser,
		},
		{
			Name:        "chat",
			Description: "get chat info",
			Func:        m.getChat,
		},
		{
			Name:        "dbg",
			Description: "debug",
			Func:        m.debug,
		},
		{
			Name:        "gban",
			Description: "gban",
			Func:        m.globalBan,
		},
		{
			Name:        "ungban",
			Description: "ungban",
			Func:        m.removeGlobalBan,
		},
	}
}

func (*Module) Messages() []types.Message {
	return nil
}

func (*Module) Callbacks() []types.Callback {
	return nil
}

// NewModule returns a new instance of this module.
func NewModule(bot *app.MyApp) (app.Module, error) {
	return &Module{
		App: bot,
	}, nil
}

func init() {
	app.RegisterModule("Admin", NewModule)
}
