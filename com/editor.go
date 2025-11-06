package com

func Editor() *EditorComponent {
	ret := &EditorComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret)
	return ret
}

type EditorComponent struct {
	*BaseComponent
	onCreated Method `bind:"editor_onCreated.js"`
	_destroy  Method `bind:"editor__destroy.js"`
	setValue  Method `bind:"editor_setValue.js"`
}
