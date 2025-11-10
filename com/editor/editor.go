package editor

import "github.com/lincaiyong/page/com"

func Editor() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("div", ret)
	return ret
}

type Component struct {
	*com.BaseComponent
	onCreated com.Method
	_destroy  com.Method
	setValue  com.Method
}
