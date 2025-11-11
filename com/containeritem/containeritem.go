package containeritem

import (
	"github.com/lincaiyong/page/com"
)

func ContainerItem(compute string) *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret)
	ret.BaseComponent.SetSlotsAsChildren()
	ret.Y("0").X("0")
	ret.Props()["computeFn"] = compute
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	data      com.Property `type:"object"`
	onUpdated com.Method
}

func (c *Component) OnUpdated(v string) *Component {
	c.SetMethod("onUpdated", v)
	return c
}
