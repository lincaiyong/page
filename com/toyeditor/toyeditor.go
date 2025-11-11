package toyeditor

import "github.com/lincaiyong/page/com"

func ToyEditor() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
}
