package tree

import (
	"fmt"
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/container"
	"github.com/lincaiyong/page/com/div"
)

func Tree() *TreeComponent {
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
	ret := &TreeComponent{}
	ret.BaseComponent = com.NewBaseComponent("div", ret,
		div.Div().X("10").Y("this.selectedChildTop-next.scrollTop").W("parent.w-20").H("this.itemHeight").BorderRadius("4").
			BackgroundColor("this.focus ? page.theme.treeFocusSelectedBgColor : page.theme.treeSelectedBgColor"),
		container.Container().List(true).Virtual(true).Align("'fill'").X("10").W("parent.w - .x").Contains(
			treeItem(),
		),
	)
	return ret
}

type TreeComponent struct {
	*com.BaseComponent
	focus            com.Property `type:"bool"`
	items            com.Property `type:"array"`
	nodeMap          com.Property `type:"object"`
	onClickItem      com.Property `type:"element"`
	selectedChildTop com.Property `type:"number"`
	itemHeight       com.Property `type:"number"`
	onCreated        com.Method
	onUpdated        com.Method
	makeNodeMap      com.Method
	nodeToItems      com.Method
	selectChild      com.Method
	handleClick      com.Method
	sortChildren     com.StaticMethod
}

func (b *TreeComponent) Focus(v bool) *TreeComponent {
	b.Props()["focus"] = fmt.Sprintf("%v", v)
	return b
}

func (b *TreeComponent) Items(s string) *TreeComponent {
	b.Props()["items"] = s
	return b
}

func (b *TreeComponent) NodeMap(s string) *TreeComponent {
	b.Props()["nodeMap"] = s
	return b
}
