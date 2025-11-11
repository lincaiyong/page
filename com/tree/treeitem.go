package tree

import (
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/img"
	"github.com/lincaiyong/page/com/text"
)

func treeItem() *ItemComponent {
	ret := &ItemComponent{}
	ret.BaseComponent = com.NewBaseComponent[ItemComponent]("div", ret,
		img.Img().Src("parent.data.collapsed ? 'svg/el/arrow-right.svg' : 'svg/el/arrow-down.svg'").X("parent.data.depth * 20 + 4").Y("5").W("11").H(".w").V("parent.data.leaf ? 0 : 1"),
		img.Img().Src("parent.data.leaf ? 'svg/mantis/file.svg' : 'svg/mantis/folder.svg'").X("next.x-18").Y("4").W("14").H(".w"),
		text.Text("parent.data.text").X("parent.data.depth * 20 + 40").Y("2").H("this.itemHeight - 2 * .y"),
	)
	return ret
}

type ItemComponent struct {
	*com.BaseComponent[ItemComponent]
	data com.Property `type:"object"`
	//compute com.StaticMethod
}
