package com

func Text(text string) *TextComponent {
	ret := &TextComponent{}
	ret.BaseComponent = NewBaseComponent("span", ret)
	ret.props["text"] = text
	ret.FontSize("Math.floor(.h * 2 / 3)").
		LineHeight(".h").
		W("page.util.textWidth(.text, .fontFamily, .fontSize)")
	return ret
}

type TextComponent struct {
	*BaseComponent
	align     Property `default:"'left'"`
	text      Property `type:"string"`
	onUpdated Method   `bind:"text_onUpdated.js"`
}

func (b *TextComponent) Align(s string) *TextComponent {
	b.props["align"] = s
	return b
}
