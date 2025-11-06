package com

func Divider() *DividerComponent {
	ret := &DividerComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret)
	return ret
}

type DividerComponent struct {
	*BaseComponent
}
