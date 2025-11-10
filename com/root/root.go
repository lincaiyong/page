package root

import (
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/js"
)

func Root(code string, comp com.Component) *Component {
	js.Set("Root", code)
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("div", ret, comp)
	return ret
}

type Component struct {
	*com.BaseComponent
}
