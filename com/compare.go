package com

func Compare() *CompareComponent {
	ret := &CompareComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret)
	return ret
}

type CompareComponent struct {
	*BaseComponent
	onCreated Method `bind:"compare_onCreated.js"`
	_destroy  Method `bind:"compare__destroy.js"`
	setValue  Method `bind:"compare_setValue.js"`
}
