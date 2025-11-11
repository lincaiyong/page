package container

import (
	"fmt"
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/containeritem"
	"github.com/lincaiyong/page/com/scrollbar"
	"strconv"
)

func Container(compute, update string, children ...com.Component) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		scrollbar.HScrollbar().NameAs("hBarEle"),
		scrollbar.VScrollbar().NameAs("vBarEle"),
	)
	ret.ScrollLeft("0").ScrollTop("0")
	ret.Contains(containeritem.ContainerItem(compute, update, children...))
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	align              com.Property `default:"'none'"`
	childHeight        com.Property `type:"number"`
	childWidth         com.Property `type:"number"`
	items              com.Property `type:"array"`
	list               com.Property `type:"bool"`
	minWidth           com.Property `type:"number"`
	reuseItem          com.Property `type:"bool"`
	scrollBarFadeTime  com.Property `default:"500"`
	scrollBarMinLength com.Property `default:"20"`
	scrollBarWidth     com.Property `default:"6"`
	scrollBarMargin    com.Property `type:"number"`
	scrollable         com.Property `default:"true"`
	virtual            com.Property `type:"bool"`
	onCreated          com.Method
	_updateList        com.Method
	onUpdated          com.Method
}

func (b *Component) Align(s string) *Component {
	b.Props()["align"] = s
	return b
}

func (b *Component) ChildHeight(v int) *Component {
	b.Props()["childHeight"] = strconv.Itoa(v)
	return b
}

func (b *Component) ChildWidth(v int) *Component {
	b.Props()["childWidth"] = strconv.Itoa(v)
	return b
}

func (b *Component) List(v bool) *Component {
	b.Props()["list"] = fmt.Sprintf("%v", v)
	return b
}

func (b *Component) MinWidth(v int) *Component {
	b.Props()["minWidth"] = strconv.Itoa(v)
	return b
}

func (b *Component) ReuseItem(v bool) *Component {
	b.Props()["reuseItem"] = fmt.Sprintf("%v", v)
	return b
}

func (b *Component) ScrollBarFadeTime(v int) *Component {
	b.Props()["scrollBarFadeTime"] = strconv.Itoa(v)
	return b
}

func (b *Component) ScrollBarMinLength(v int) *Component {
	b.Props()["scrollBarMinLength"] = strconv.Itoa(v)
	return b
}

func (b *Component) ScrollBarWidth(v int) *Component {
	b.Props()["scrollBarWidth"] = strconv.Itoa(v)
	return b
}

func (b *Component) Scrollable(v bool) *Component {
	b.Props()["scrollable"] = fmt.Sprintf("%v", v)
	return b
}

func (b *Component) Virtual(v bool) *Component {
	b.Props()["virtual"] = fmt.Sprintf("%v", v)
	return b
}
