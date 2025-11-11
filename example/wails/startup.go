package main

import (
	"context"
)

func Startup(ctx context.Context) {
	app.ctx = ctx
	//runtime.EventsOn(ctx, "message", func(optionalData ...interface{}) {
	//	if len(optionalData) == 1 {
	//		data := optionalData[0].(string)
	//		logOutput(g.ctx, data)
	//	}
	//})
}
