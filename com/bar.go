package com

func Bar() *BarComponent {
	ret := &BarComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret)
	ret.OnMouseDown("BarComponent.handleMouseDown").ZIndex("1")
	return ret
}

func VBar() *BarComponent {
	ret := Bar()
	ret.props["leftRight"] = "[prev, next]"
	ret.Cursor("'col-resize'").W_(20)
	return ret
}

func HBar() *BarComponent {
	ret := Bar()
	ret.props["topBottom"] = "[prev, next]"
	ret.Cursor("'row-resize'").H_(20)
	return ret
}

type BarComponent struct {
	*BaseComponent
	leftRight       Property     `default:"[undefined, undefined]"`
	topBottom       Property     `default:"[undefined, undefined]"`
	onUpdated       Method       `bind:"bar_onUpdated.js"`
	handleMouseDown StaticMethod `bind:"bar_handleMouseDown.js"`
}
