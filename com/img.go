package com

func Img() *ImgComponent {
	ret := &ImgComponent{}
	ret.BaseComponent = NewBaseComponent("img", ret)
	return ret
}

type ImgComponent struct {
	*BaseComponent
	src       Property `type:"string"`
	onUpdated Method   `bind:"img_onUpdated.js"`
}

func (b *ImgComponent) SVG() *ImgComponent {
	b.tag = "svg"
	return b
}

func (b *ImgComponent) Src(s string) *ImgComponent {
	b.props["src"] = s
	return b
}
