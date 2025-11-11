package containeritem

import (
	"github.com/lincaiyong/page/com"
)

func ContainerItem(compute string) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.BaseComponent.SetSlotsAsChildren()
	ret.Y("0").X("0")
	ret.Props()["compute"] = compute
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	data    com.Property `type:"object"`
	compute com.Property `default:"null"`
}
