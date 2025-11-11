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
		//Editor().NameAs("editorEle"),
		Button().OnClick("Root.handleClick"),
		Tree().Y("prev.y2").H("parent.h-.y").NameAs("treeEle"),
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
