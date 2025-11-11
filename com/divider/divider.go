package divider

import "github.com/lincaiyong/page/com"

func VDivider() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.BackgroundColor(com.ColorBlack).W("2")
	return ret
}

func HDivider() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.BackgroundColor(com.ColorBlack).H("2")
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
}
