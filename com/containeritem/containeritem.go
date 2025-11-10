package containeritem

import (
	"github.com/lincaiyong/page/com"
)

func ContainerItem(compute, onUpdated string) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("div", ret)
	ret.BaseComponent.SetSlotsAsChildren()
	ret.Y("0").X("0")
	ret.Props()["computeFn"] = compute
	ret.Props()["onUpdatedFn"] = onUpdated
	return ret
}

type Component struct {
	*com.BaseComponent
	data      com.Property `type:"object"`
	onUpdated com.Method
}
