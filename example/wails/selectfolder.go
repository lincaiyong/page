package main

import (
	"encoding/json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/fs"
	"os"
	"path/filepath"
)

func (a *App) SelectFolder() string {
	//logOutput(a.ctx, "select folder")
	opts := runtime.OpenDialogOptions{
		Title: "选择项目",
	}
	entry, err := runtime.OpenDirectoryDialog(a.ctx, opts)
	if err != nil {
		//logOutput(a.ctx, fmt.Sprintf("OpenDirectoryDialog error: %v", err))
		return ""
	}
	//logOutput(a.ctx, entry)
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
	//logOutput(a.ctx, string(ret))

	return string(ret)
}
