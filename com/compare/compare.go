package compare

import (
	"github.com/lincaiyong/page/com"
)

func Compare() *Component {
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
