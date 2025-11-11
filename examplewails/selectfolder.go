package main

import (
	"encoding/json"
	"github.com/lincaiyong/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/fs"
	"path/filepath"
	"strings"
)

func (a *App) SelectFolder() string {
	log.DebugLog("select folder...")
	opts := runtime.OpenDialogOptions{
		Title: "选择项目",
	}
	folder, err := runtime.OpenDirectoryDialog(a.ctx, opts)
	if err != nil {
		log.ErrorLog("fail to OpenDirectoryDialog: %v", err)
		return ""
	}
	log.DebugLog("folder: %s", folder)
	var files []string
	err = filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
			return fs.SkipDir
		}
		if !d.IsDir() {
			relPath, _ := filepath.Rel(folder, path)
			files = append(files, relPath)
		}
		return nil
	})
	data := map[string]any{
		"folder": folder,
		"files":  files,
	}
	ret, _ := json.Marshal(data)
	log.DebugLog("result: %s", string(ret))
	return string(ret)
}
