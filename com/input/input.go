package input

import "github.com/lincaiyong/page/com"

func Input() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("input", ret)
	ret.LineHeight(".h").FontSize("Math.floor(.h * 2 / 3)")
	return ret
}

type Component struct {
	*com.BaseComponent
	placeholder com.Property
	onUpdated   com.Method
}

func (b *Component) Placeholder(s string) *Component {
	b.Props()["placeholder"] = s
	return b
}
