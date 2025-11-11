package containeritem

import (
	"github.com/lincaiyong/page/com"
)

func ContainerItem(compute, update string, children ...com.Component) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret, children...)
	ret.Y("0").X("0")
	ret.Props()["compute"] = compute
	ret.Props()["update"] = update
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	data      com.Property `type:"object"`
	compute   com.Property `default:"null"`
	update    com.Property `default:"null"`
	onUpdated com.Method
}
