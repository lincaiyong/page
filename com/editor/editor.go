package editor

import "github.com/lincaiyong/page/com"

func Editor() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	onCreated com.Method
	_destroy  com.Method
	setValue  com.Method
}
