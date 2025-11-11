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
					editor.Editor().X("20").Y("20").W("800").H("next.y - .y").BackgroundColor("'blue'"),
					bar.HBar().Y("200").W("parent.w").H("20").Opacity("0.1"),
					div.Div().X("20").Y("prev.y2").W("800").H("400").BackgroundColor("'yellow'").Contains(
						text.Text("'hello world!'").H("200"),
					),
					compare.Compare().X("800").Y("20").W("400").H("400").BackgroundColor("'red'"),
					div.Div().X("prev.x2").Y("prev.y2").W("40").H("40").BackgroundColor("'green'"),
					button.Button().Icon("'svg/el/folder.svg'").X("prev.x2").Y("prev.y2 + 100").W("40").H("40"),
				))
				page.MakePage(c, "debug", comp, baseUrl)
			})
			return nil
		},
	)
}
