package scrollbar

import (
	"fmt"
	"github.com/lincaiyong/page/com"
)

func create() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("div", ret)
	ret.ZIndex("1").
		BackgroundColor("page.theme.scrollbarBgColor").
		Opacity("0.5").
		BorderRadius(".w / 2").
		Cursor("'default'").
		X(".vertical ? parent.cw - parent.scrollBarMargin - parent.scrollBarWidth : 0").
		Y(".vertical ? 0 : parent.ch - parent.scrollBarMargin - parent.scrollBarWidth").
		W(".vertical ? parent.scrollBarWidth : 0").
		H(".vertical ? 0 : parent.scrollBarWidth").
		V("0")
	return ret
}

func VScrollbar() *Component {
	ret := create()
	ret.Props()["vertical"] = "true"
	return ret
}

func HScrollbar() *Component {
	ret := create()
	ret.Props()["vertical"] = "false"
	return ret
}

type Component struct {
	*com.BaseComponent
	showLeft com.Property `type:"bool"`
	showTop  com.Property `type:"bool"`
	vertical com.Property `type:"bool"`
}

func (b *Component) ShowLeft(v bool) *Component {
	b.Props()["showLeft"] = fmt.Sprintf("%v", v)
	return b
}

func (b *Component) ShowTop(v bool) *Component {
	b.Props()["showTop"] = fmt.Sprintf("%v", v)
	return b
}
