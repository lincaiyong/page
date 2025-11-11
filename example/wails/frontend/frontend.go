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
	comp := Root(Editor().NameAs("editorEle")).OnCreated("Root.test").
		Code(`
function test() {
	setTimeout(function() {
		const editor = page.root.editorEle;
		editor.setValue('package main\n\nfunc main() {\n\n}');
		editor.setLanguage('go');
	});
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
