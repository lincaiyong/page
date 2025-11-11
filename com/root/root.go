package root

import (
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/js"
)

func Root(code string, children ...com.Component) *Component {
	js.Set("Root", code)
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret, children...)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
}
