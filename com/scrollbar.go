package com

import "fmt"

func Scrollbar() *ScrollbarComponent {
	ret := &ScrollbarComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret)
	ret.ZIndex("1").
		BackgroundColor("page.theme.scrollbarBgColor").
		Opacity("0.5").
		BorderRadius(".w / 2").
		Cursor("default").
		X(".vertical ? parent.cw - parent.scrollBarMargin - parent.scrollBarWidth : 0").
		Y(".vertical ? 0 : parent.ch - parent.scrollBarMargin - parent.scrollBarWidth").
		W(".vertical ? parent.scrollBarWidth : 0").
		H(".vertical ? 0 : parent.scrollBarWidth").
		V("0")
	return ret
}

func VScrollbar() *ScrollbarComponent {
	ret := Scrollbar()
	ret.props["vertical"] = "true"
	return ret
}

func HScrollbar() *ScrollbarComponent {
	ret := Scrollbar()
	ret.props["vertical"] = "false"
	return ret
}

type ScrollbarComponent struct {
	*BaseComponent
	showLeft Property `type:"bool"`
	showTop  Property `type:"bool"`
	vertical Property `type:"bool"`
}

func (b *ScrollbarComponent) ShowLeft(v bool) *ScrollbarComponent {
	b.props["showLeft"] = fmt.Sprintf("%v", v)
	return b
}

func (b *ScrollbarComponent) ShowTop(v bool) *ScrollbarComponent {
	b.props["showTop"] = fmt.Sprintf("%v", v)
	return b
}
