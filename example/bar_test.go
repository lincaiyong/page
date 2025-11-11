package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/page"
	"github.com/lincaiyong/page/com/bar"
	"github.com/lincaiyong/page/com/button"
	"github.com/lincaiyong/page/com/compare"
	"github.com/lincaiyong/page/com/div"
	"github.com/lincaiyong/page/com/editor"
	"github.com/lincaiyong/page/com/root"
	"github.com/lincaiyong/page/com/text"
	"testing"
)

func TestBar(t *testing.T) {
	common.StartServer("page", "v1.0.1", "",
		func(_ []string, r *gin.RouterGroup) error {
			baseUrl := "http://127.0.0.1:9123"
			r.GET("/res/*filepath", page.HandleRes(baseUrl))
			r.GET("/", func(c *gin.Context) {
				comp := root.Root("", div.Div().Contains(
					div.Div().W("next.x").Contains(
						editor.Editor().X("20").Y("0").W("800").H("next.y - .y").BackgroundColor("'blue'"),
						bar.HBar().BackgroundColor("'blue'").Opacity("0.1").Y("parent.h/2").W("parent.w"),
						div.Div().X("20").Y("prev.y2").W("800").H("parent.h-prev.y2").BackgroundColor("'yellow'").Contains(
							text.Text("'hello world!'").H("200"),
						),
					),
					bar.VBar().X("parent.w/2").BackgroundColor("'blue'").Opacity("0.1"),
					div.Div().X("prev.x2").W("parent.w-prev.x2").Contains(
						compare.Compare().Y("0").H("next.y").BackgroundColor("'red'"),
						bar.HBar().BackgroundColor("'blue'").Opacity("0.1").Y("parent.h/2").W("parent.w"),
						div.Div().Y("prev.y2").W("40").H("40").BackgroundColor("'green'"),
						button.Button().Icon("'svg/el/folder.svg'").X("prev.x2").Y("prev.y2 + 100").W("40").H("40"),
					),
				))
				page.MakePage(c, "debug", comp, baseUrl)
			})
			return nil
		},
	)
}
