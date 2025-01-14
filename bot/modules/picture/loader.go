package picture

import (
	"SiskamlingBot/bot/core/app"
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/core/telegram/types"
)

// Module contains the state for an instance of this module.
type Module struct {
	App *app.MyApp
}

// Info returns basic information about this module.
func (Module) Info() app.ModuleInfo {
	return app.ModuleInfo{
		Name: "Picture",
	}
}

// Commands returns a list of telegram provided by this module.
func (Module) Commands() []types.Command {
	return []types.Command{}
}

func (m Module) Messages() []types.Message {
	return []types.Message{
		{
			Name:        "PictureScanner",
			Description: "Detect user without profile picture",
			Filter:      telegram.ProfileAndGroupFilter(m.App.Bot),
			Func:        m.pictureScan,
			Order:       1,
			Async:       false,
		},
	}
}

func (m Module) Callbacks() []types.Callback {
	return []types.Callback{
		{
			Name:        "PictureCallback",
			Description: "",
			Callback:    `picture\((.+?)\)`,
			Func:        m.pictureCallback,
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
	app.RegisterModule("Picture", NewModule)
}
