package main

import (
	"fmt"
	"github.com/lincaiyong/page"
	"github.com/lincaiyong/page/com"
	. "github.com/lincaiyong/page/com/all"
	"os"
	"testing"
)

func TestPage(t *testing.T) {
	com.BaseUrl = ""
	html, err := page.MakeHtml("demo", Root(Text("'hello'")))
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("frontend/dist/index.html", []byte(html), 0644)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("done")
}
