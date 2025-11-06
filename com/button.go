package com

import "fmt"

func Button() *ButtonComponent {
	ret := &ButtonComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret,
		Img().SVG().Src("parent.icon").X_(4).Y(".x").W("parent.w - 2 * .x").H(".w").Color("parent.color"),
		Div().X("prev.x2 - .w + 1").Y("prev.y - 1").W_(6).H(".w").V("0").BorderRadius("3"),
	)
	ret.W_(24).H(".w").BorderRadius("6").
		BackgroundColor(".selected ? page.theme.buttonSelectedBgColor : page.theme.buttonBgColor").
		Color(".selected ? page.theme.buttonSelectedColor : page.theme.buttonColor").
		OnHover("e.handleHover").
		OnActive("e.handleActive")
	return ret
}

type ButtonComponent struct {
	*BaseComponent
	flag         Property `type:"string"`
	icon         Property `default:"'svg/el/folder.svg'"`
	selected     Property `type:"bool"`
	handleHover  Method   `bind:"button_handleHover.js"`
	handleActive Method   `bind:"button_handleActive.js"`
}

func (b *ButtonComponent) Flag(s string) *ButtonComponent {
	b.props["flag"] = s
	return b
}

func (b *ButtonComponent) Icon(s string) *ButtonComponent {
	b.props["icon"] = s
	return b
}

func (b *ButtonComponent) Selected(v bool) *ButtonComponent {
	b.props["selected"] = fmt.Sprintf("%v", v)
	return b
}
