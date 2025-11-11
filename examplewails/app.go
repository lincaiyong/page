package main

import (
	"context"
	"github.com/lincaiyong/log"
)

var app App

type App struct {
	ctx context.Context
}

//func logOutput(ctx context.Context, msg string) {
//	runtime.EventsEmit(ctx, "output", msg)
//}

func (a *App) Log(v any) {
	log.DebugLog("%v", v)
}
