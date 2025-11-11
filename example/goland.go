package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/page"
	. "github.com/lincaiyong/page/com/all"
)

func goland(c *gin.Context) {
	comp := Root(
		Text("'hello'"),
	)
	page.MakePage(c, "goland", comp)
}
