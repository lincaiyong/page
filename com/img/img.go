package img

import "github.com/lincaiyong/page/com"

func Img() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("img", ret)
	return ret
}

func Svg() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("svg", ret)
	return ret
}

type Component struct {
	*com.BaseComponent
	src       com.Property `type:"string"`
	onUpdated com.Method
}

func (b *Component) Src(s string) *Component {
	b.Props()["src"] = s
	return b
}
