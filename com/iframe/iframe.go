package iframe

import "github.com/lincaiyong/page/com"

func Iframe() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	setHtml com.Method
}
