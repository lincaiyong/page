package textarea

import "github.com/lincaiyong/page/com"

func Textarea() *Component {
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent("textarea", ret)
	return ret
}

type Component struct {
	*com.BaseComponent
}
