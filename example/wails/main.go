package main

import (
	"context"
	"embed"
	_ "embed"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed build/appicon.png
var icon []byte

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	err := wails.Run(&options.App{
		Title:  "demo",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarDefault(),
			Appearance:           mac.DefaultAppearance,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "demo",
				Message: "©2025 lincaiyong <lincaiyong@codeedge.cc>",
				Icon:    icon,
			},
		},
		OnStartup: app.startup,
		Bind:      []any{app},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func logOutput(ctx context.Context, msg string) {
	runtime.EventsEmit(ctx, "output", msg)
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsOn(ctx, "message", func(optionalData ...interface{}) {
		if len(optionalData) == 1 {
			data := optionalData[0].(string)
			logOutput(a.ctx, data)
		}
	})
}
