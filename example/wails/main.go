package main

import (
	"context"
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

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

func (a *App) SelectFolder() string {
	logOutput(a.ctx, "select folder")
	opts := runtime.OpenDialogOptions{
		Title: "选择项目",
	}
	entry, err := runtime.OpenDirectoryDialog(a.ctx, opts)
	if err != nil {
		logOutput(a.ctx, fmt.Sprintf("OpenDirectoryDialog error: %v", err))
		return ""
	}
	logOutput(a.ctx, entry)
	files := map[string]string{}
	err = filepath.WalkDir(entry, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			var b []byte
			b, err = os.ReadFile(path)
			if err != nil {
				return err
			}
			files[path] = string(b)
		}
		return nil
	})
	data := map[string]any{
		"folder": entry,
		"files":  files,
	}
	ret, _ := json.Marshal(data)
	logOutput(a.ctx, string(ret))

	return string(ret)
}

//go:embed build/appicon.png
var icon []byte

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	err := wails.Run(&options.App{
		Title:  "CodeEdge",
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
				Title:   "CodeEdge",
				Message: "©2025 lincaiyong@codeedge.cc",
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
