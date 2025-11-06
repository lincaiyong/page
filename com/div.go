package com

func Div() *DivComponent {
	ret := &DivComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret)
	ret.slotsAsChildren = true
	return ret
}

type DivComponent struct {
	*BaseComponent
}
