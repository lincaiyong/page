package bar

import (
	_ "embed"
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/js"
)

//go:embed bar.js
var barJs string

func init() {
	js.Set("Bar", barJs)
}

func create() *Bar {
	ret := &Bar{}
	ret.BaseComponent = com.NewBaseComponent("div", ret)
	ret.OnMouseDown("Bar.handleMouseDown").ZIndex("1")
	return ret
}

func VBar() *Bar {
	ret := create()
	ret.Props()["leftRight"] = "[prev, next]"
	ret.Cursor("'col-resize'").W_(20)
	return ret
}

func HBar() *Bar {
	ret := create()
	ret.Props()["topBottom"] = "[prev, next]"
	ret.Cursor("'row-resize'").H_(20)
	return ret
}

type Bar struct {
	*com.BaseComponent
	leftRight       com.Property `default:"[undefined, undefined]"`
	topBottom       com.Property `default:"[undefined, undefined]"`
	onUpdated       com.Method
	handleMouseDown com.StaticMethod
}
