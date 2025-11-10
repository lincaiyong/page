package div

import "github.com/lincaiyong/page/com"

func Div() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("div", ret)
	ret.BaseComponent.SlotsAsChildren(true)
	return ret
}

type Component struct {
	*com.BaseComponent
}
