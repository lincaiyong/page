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
	html, err := page.MakeHtml("demo", Root(Text("'hello'")))
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
