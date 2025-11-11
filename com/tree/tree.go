package tree

import (
	"fmt"
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/container"
	"github.com/lincaiyong/page/com/div"
	"github.com/lincaiyong/page/com/img"
	"github.com/lincaiyong/page/com/text"
)

func Tree() *Component {
	/*
	   selectedEle:div(x=10, y=this.selectedChildTop-next.scrollTop, w=parent.w-20, h=this.itemHeight, borderRadius=4,
	       backgroundColor=this.focus ? webapp.theme.treeFocusSelectedBgColor : webapp.theme.treeSelectedBgColor,
	   )
	   containerEle:container(x=10, w=parent.w - .x, list=true, virtual=true, align=fill) {
	       container_item<TreeItemData>(compute=compute, onClick=handleClick, onContextMenu=handleContextMenu) {
	           img(x=parent.data.depth * 20 + 4, y=5, w=11, h=.w, src=parent.data.collapsed ? 'svg/el/arrow-right.svg' : 'svg/el/arrow-down.svg', v=parent.data.leaf ? 0 : 1)
	           img(x=next.x - 18, y=4, w=14, h=.w, src=parent.data.leaf ? 'svg/mantis/file.svg' : 'svg/mantis/folder.svg')
	           text(x=parent.data.depth * 20 + 40, y=2, h=this.itemHeight-2*.y, text=parent.data.text)
	       }
	   }
	*/
	ret := &Component{}
	ret.BaseComponent = com.NewBaseComponent[Component]("div", ret,
		div.Div().X("10").Y("this.selectedChildTop-next.scrollTop").W("parent.w-20").H("this.itemHeight").
			BorderRadius("4").BackgroundColor("this.focus ? page.theme.treeFocusSelectedBgColor : page.theme.treeSelectedBgColor").
			NameAs("selectedEle"),
		container.VListContainer("", "",
			img.Img("parent.data.collapsed ? 'svg/el/arrow-right.svg' : 'svg/el/arrow-down.svg'").X("parent.data.depth * 20 + 4").Y("5").W("11").H(".w").V("parent.data.leaf ? 0 : 1"),
			img.Img("parent.data.leaf ? 'svg/mantis/file.svg' : 'svg/mantis/folder.svg'").X("next.x-18").Y("4").W("14").H(".w"),
			text.Text("parent.data.text").X("parent.data.depth * 20 + 40").Y("2").H("this.itemHeight - 2 * .y"),
		).Align("'fill'").X("10").W("parent.w - .x").
			NameAs("containerEle"),
	)
	return ret
}

type Component struct {
	*com.BaseComponent[Component]
	focus            com.Property `default:"false"`
	items            com.Property `default:"[]"`
	nodeMap          com.Property `default:"undefined"`
	onClickItem      com.Property `default:"undefined"`
	selectedChildTop com.Property `default:"0"`
	itemHeight       com.Property `default:"0"`
	onCreated        com.Method
	onUpdated        com.Method
	makeNodeMap      com.Method
	nodeToItems      com.Method
	selectChild      com.Method
	handleClick      com.Method
}

func (b *Component) Focus(v bool) *Component {
	b.SetProp("focus", fmt.Sprintf("%v", v))
	return b
}

func (b *Component) Items(s string) *Component {
	b.SetProp("items", s)
	return b
}

func (b *Component) NodeMap(s string) *Component {
	b.SetProp("nodeMap", s)
	return b
}
