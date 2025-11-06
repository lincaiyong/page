package com

func Textarea() *TextareaComponent {
	ret := &TextareaComponent{}
	ret.BaseComponent = NewBaseComponent("textarea", ret)
	return ret
}

type TextareaComponent struct {
	*BaseComponent
}
