package com

import (
	"fmt"
	"strconv"
)

func ContainerItem() *ContainerItemComponent {
	ret := &ContainerItemComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret)
	return ret
}

type ContainerItemComponent struct {
	*BaseComponent
	data    Property       `type:"object"`
	compute StaticProperty `type:"function"`
}

func (b *ContainerItemComponent) Data(s string) *ContainerItemComponent {
	b.props["data"] = s
	return b
}

func (b *ContainerItemComponent) Compute(s string) *ContainerItemComponent {
	b.props["compute"] = s
	return b
}

func Container() *ContainerComponent {
	ret := &ContainerComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret,
		HScrollbar().NameAs("hBarEle"),
		VScrollbar().NameAs("vBarEle"),
	)
	ret.ScrollLeft("0").ScrollTop("0")
	return ret
}

type ContainerComponent struct {
	*BaseComponent
	align              Property `default:"'none'"`
	childHeight        Property `type:"number"`
	childWidth         Property `type:"number"`
	items              Property `type:"array"`
	list               Property `type:"bool"`
	minWidth           Property `type:"number"`
	reuseItem          Property `type:"bool"`
	scrollBarFadeTime  Property `default:"500"`
	scrollBarMinLength Property `default:"20"`
	scrollBarWidth     Property `default:"6"`
	scrollable         Property `default:"true"`
	virtual            Property `type:"bool"`
	onCreated          Method   `bind:"container_onCreated.js"`
	_updateList        Method   `bind:"container__updateList.js"`
	onUpdated          Method   `bind:"container_onUpdated.js"`
	hBarEle            Property `type:"element"`
	vBarEle            Property `type:"element"`
	scrollBarMargin    Property `type:"number"`
}

func (b *ContainerComponent) Align(s string) *ContainerComponent {
	b.props["align"] = s
	return b
}

func (b *ContainerComponent) ChildHeight(v int) *ContainerComponent {
	b.props["childHeight"] = strconv.Itoa(v)
	return b
}

func (b *ContainerComponent) ChildWidth(v int) *ContainerComponent {
	b.props["childWidth"] = strconv.Itoa(v)
	return b
}

func (b *ContainerComponent) Items(s string) *ContainerComponent {
	b.props["items"] = s
	return b
}

func (b *ContainerComponent) List(v bool) *ContainerComponent {
	b.props["list"] = fmt.Sprintf("%v", v)
	return b
}

func (b *ContainerComponent) MinWidth(v int) *ContainerComponent {
	b.props["minWidth"] = strconv.Itoa(v)
	return b
}

func (b *ContainerComponent) ReuseItem(v bool) *ContainerComponent {
	b.props["reuseItem"] = fmt.Sprintf("%v", v)
	return b
}

func (b *ContainerComponent) ScrollBarFadeTime(v int) *ContainerComponent {
	b.props["scrollBarFadeTime"] = strconv.Itoa(v)
	return b
}

func (b *ContainerComponent) ScrollBarMinLength(v int) *ContainerComponent {
	b.props["scrollBarMinLength"] = strconv.Itoa(v)
	return b
}

func (b *ContainerComponent) ScrollBarWidth(v int) *ContainerComponent {
	b.props["scrollBarWidth"] = strconv.Itoa(v)
	return b
}

func (b *ContainerComponent) Scrollable(v bool) *ContainerComponent {
	b.props["scrollable"] = fmt.Sprintf("%v", v)
	return b
}

func (b *ContainerComponent) Virtual(v bool) *ContainerComponent {
	b.props["virtual"] = fmt.Sprintf("%v", v)
	return b
}
