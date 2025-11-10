package iframe

import "github.com/lincaiyong/page/com"

func Iframe() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("div", ret)
	return ret
}

type Component struct {
	*com.BaseComponent
	setHtml com.Method
}
