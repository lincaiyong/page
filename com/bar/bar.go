package bar

import (
	"github.com/lincaiyong/page/com"
)

func create() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("div", ret)
	ret.OnMouseDown("BarComponent.handleMouseDown").ZIndex("1")
	return ret
}

func VBar() *Component {
	ret := create()
	ret.Props()["leftRight"] = "[prev, next]"
	ret.Cursor("'col-resize'").W("20")
	return ret
}

func HBar() *Component {
	ret := create()
	ret.Props()["topBottom"] = "[prev, next]"
	ret.Cursor("'row-resize'").H("20")
	return ret
}

type Component struct {
	*com.BaseComponent
	leftRight       com.Property `default:"[undefined, undefined]"`
	topBottom       com.Property `default:"[undefined, undefined]"`
	onUpdated       com.Method
	handleMouseDown com.StaticMethod
}
