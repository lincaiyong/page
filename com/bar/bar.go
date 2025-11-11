package bar

import (
	"github.com/lincaiyong/page/com"
)

func create() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.OnMouseDown("e.handleMouseDown").ZIndex("1")
	return ret
}

func VBar() *Component {
	ret := create()
	ret.SetProp("leftRight", "[prev, next]")
	ret.Cursor("'col-resize'").W("20")
	return ret
}

func HBar() *Component {
	ret := create()
	ret.SetProp("topBottom", "[prev, next]")
	ret.Cursor("'row-resize'").H("20")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	leftRight       com.Property `default:"[undefined, undefined]"`
	topBottom       com.Property `default:"[undefined, undefined]"`
	onUpdated       com.Method
	handleMouseDown com.Method
}
