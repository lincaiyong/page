package main

import (
	"github.com/lincaiyong/log"
	"github.com/lincaiyong/page"
	"github.com/lincaiyong/page/com"
	. "github.com/lincaiyong/page/com/all"
	"os"
)

func main() {
	com.BaseUrl = ""
	comp := Root(
		Div().NameAs("headerEle").H("34").BgColor(ColorBlack),
		Div().Y("prev.y2").H("parent.h-prev.h-next.h").SetSlots(
			Div().NameAs("leftSideEle").W("32").SetSlots(
				Button().OnClick("Root.handleClick"),
			),
			Div().X("prev.x2").W("parent.w-prev.w-next.w").BgColor("page.theme.grayPaneColor").SetSlots(
				Div().NameAs("navEle").W("next.x").SetSlots(
					Tree().NameAs("treeEle"),
				),
				VBar().X("parent.w/3").BgColor(ColorYellow).Opacity("0.1"),
				Div().NameAs("mainEle").X("prev.x2").W("parent.w-.x2").SetSlots(
					Editor().NameAs("editorEle"),
				),
			),
			Div().NameAs("rightSideEle").X("parent.w-.w").W("32").BgColor("page.theme.grayPaneColor"),
		),
		Div().NameAs("footerEle").Y("parent.h-.h").H("24").BgColor("page.theme.grayPaneColor"),
	).OnCreated("Root.test").
		Code(`
function test() {
	//setTimeout(function() {
	//	const editor = page.root.editorEle;
	//	editor.setValue('package main\n\nfunc main() {\n\n}');
	//	editor.setLanguage('go');
	//});
}
function handleClick() {
	go.main.App.SelectFolder().then(s => {
		const obj = JSON.parse(s)
		Root.log(obj);
		page.root.treeEle.items = obj.files;
	});
}
function log(v) {
	go.main.App.Log(v);
}
`)
	html, err := page.MakeHtml("CodeEdge", comp)
	if err != nil {
		log.ErrorLog("%v", err)
		return
	}
	err = os.WriteFile("./dist/index.html", []byte(html), 0644)
	if err != nil {
		log.ErrorLog("%v", err)
		return
	}
	log.InfoLog("done")
}
