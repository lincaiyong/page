package root

import (
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/js"
)

func Root(code string) *Component {
	js.Set("RootComponent", code)
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("div", ret)
	ret.BaseComponent.SetSlotsAsChildren()
	return ret
}

type Component struct {
	*com.BaseComponent
}
