package com

func Iframe() *IframeComponent {
	ret := &IframeComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret)
	return ret
}

type IframeComponent struct {
	*BaseComponent
	setHtml Method `bind:"iframe_setHtml.js"`
}
