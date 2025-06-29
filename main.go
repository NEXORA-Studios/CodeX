package main

import (
	"embed"

	"CodeX/backend/bind"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "CodeX",
		Width:     1280,
		Height:    720,
		MinWidth:  1280,
		MinHeight: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.onStartUpHanlder,
		CSSDragProperty:  "--drag-region",
		CSSDragValue:     "y",
		Bind: []any{
			app,
			&bind.ProfileBind{},
			&bind.ConfigBind{},
			&bind.SystemBind{},
			&bind.EnvBind{},
		},
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: "00c00d10-8e8c-4522-b425-abc00369c076",
			OnSecondInstanceLaunch: func(secondInstanceData options.SecondInstanceData) {
				runtime.WindowShow(app.ctx)
			},
		},
		Frameless: true,
		Windows: &windows.Options{
			WindowIsTranslucent: true,
			BackdropType:        windows.Mica,
		},
		Mac: &mac.Options{
			TitleBar:            mac.TitleBarHiddenInset(),
			WindowIsTranslucent: true,
			About: &mac.AboutInfo{
				Title:   "Code X",
				Message: "Code X is a tutorial application that is used to learn how to code.\n\nCopyright © 2025~present NEXORA Studios",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
			ProgramName:         "wails",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
