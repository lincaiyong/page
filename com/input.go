package com

func Input() *InputComponent {
	ret := &InputComponent{}
	ret.BaseComponent = NewBaseComponent("input", ret)
	ret.LineHeight(".h").FontSize("Math.floor(.h * 2 / 3)")
	return ret
}

type InputComponent struct {
	*BaseComponent
	placeholder Property `type:"string"`
	onUpdated   Method   `bind:"input_onUpdated.js"`
}

func (b *InputComponent) Placeholder(s string) *InputComponent {
	b.props["placeholder"] = s
	return b
}
