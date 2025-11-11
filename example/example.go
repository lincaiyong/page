package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/page"
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/bar"
	"github.com/lincaiyong/page/com/button"
	"github.com/lincaiyong/page/com/compare"
	"github.com/lincaiyong/page/com/container"
	"github.com/lincaiyong/page/com/div"
	"github.com/lincaiyong/page/com/editor"
	"github.com/lincaiyong/page/com/root"
	"github.com/lincaiyong/page/com/text"
)

//go:embed example.js
var exampleJs string

func main() {
	common.StartServer("page", "v1.0.1", "",
		func(_ []string, r *gin.RouterGroup) error {
			r.GET("/res/*filepath", page.HandleRes(com.BaseUrl))
			r.GET("/hello", func(c *gin.Context) {
				comp := root.Root(
					"",
					text.Text("'hello world'").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2"),
				)
				page.MakePage(c, "debug", comp)
			})
			r.GET("/click", func(c *gin.Context) {
				comp := root.Root(
					`
function test(msg) {
	console.log("test: " + msg);
}
function handleClick() {
	console.log(...arguments);
}`,
					text.Text("'hello world'").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2").
						OnCreated("() => Root.test('onCreated')").
						OnUpdated("() => Root.test('onUpdated')").
						OnClick("Root.handleClick"),
				)
				page.MakePage(c, "debug", comp)
			})
			r.GET("/bar", func(c *gin.Context) {
				comp := root.Root("", div.Div().SetSlots(
					div.Div().W("next.x").SetSlots(
						editor.Editor().X("20").Y("0").W("800").H("next.y - .y").BackgroundColor(com.ColorBlue),
						bar.HBar().BackgroundColor(com.ColorBlue).Opacity("0.1").Y("parent.h/2").W("parent.w"),
						div.Div().X("20").Y("prev.y2").W("800").H("parent.h-prev.y2").BackgroundColor(com.ColorYellow).SetSlots(
							text.Text("'hello world!'").H("200"),
						),
					),
					bar.VBar().X("parent.w/2").BackgroundColor(com.ColorBlue).Opacity("0.1"),
					div.Div().X("prev.x2").W("parent.w-prev.x2").SetSlots(
						compare.Compare().Y("0").H("next.y").BackgroundColor(com.ColorRed),
						bar.HBar().BackgroundColor(com.ColorBlue).Opacity("0.1").Y("parent.h/2").W("parent.w"),
						div.Div().Y("prev.y2").W("40").H("40").BackgroundColor(com.ColorGreen),
						button.Button().Icon(com.SvgElAddLocation).X("prev.x2").Y("prev.y2 + 100").W("40").H("40"),
					),
				))
				page.MakePage(c, "debug", comp)
			})
			r.GET("/vlist", func(c *gin.Context) {
				comp := root.Root(exampleJs,
					div.Div().BackgroundColor("'#eee'").W("200").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2").SetSlots(
						container.VListContainer("Root.computeItem", "Root.updateItem",
							div.Div().OnHover("Root.hoverItem").SetSlots(
								text.Text("''").NameAs("textEle"),
							),
						).NameAs("containerEle"),
					),
				).OnCreated("Root.onCreated")
				page.MakePage(c, "debug5", comp)
			})
			r.GET("/container", func(c *gin.Context) {
				comp := root.Root("", container.Container(text.Text("'hello world!'").H("400")).Scrollable(true).BackgroundColor("'#eee'").W("200").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2"))
				page.MakePage(c, "debug4", comp)
			})
			r.GET("/editor", func(c *gin.Context) {
				comp := root.Root(`
function test() {
	setTimeout(function() {
		const editor = page.root.editorEle;
		editor.setValue('package main\n\nfunc main() {\n\n}');
		editor.setLanguage('go');
	}, 1000);
}
`, editor.Editor().NameAs("editorEle")).OnCreated("Root.test")
				page.MakePage(c, "editor", comp)
			})
			return nil
		},
	)
}
