package com

import "fmt"

func TreeItem() *TreeItemComponent {
	ret := &TreeItemComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret,
		Img().Src("parent.data.collapsed ? 'svg/el/arrow-right.svg' : 'svg/el/arrow-down.svg'").X("parent.data.depth * 20 + 4").Y("5").W("11").H(".w").V("parent.data.leaf ? 0 : 1"),
		Img().Src("parent.data.leaf ? 'svg/mantis/file.svg' : 'svg/mantis/folder.svg'").X("next.x-18").Y("4").W("14").H(".w"),
		Text("parent.data.text").X("parent.data.depth * 20 + 40").Y("2").H("this.itemHeight - 2 * .y"),
	)
	return ret
}

type TreeItemComponent struct {
	*BaseComponent
	data    Property     `type:"object"`
	compute StaticMethod `bind:"treeItem_compute.js"`
}

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
	ret.BaseComponent = NewBaseComponent("div", ret,
		Div().X("10").Y("this.selectedChildTop-next.scrollTop").W("parent.w-20").H("this.itemHeight").BorderRadius("4").
			BackgroundColor("this.focus ? page.theme.treeFocusSelectedBgColor : page.theme.treeSelectedBgColor"),
		Container().List(true).Virtual(true).Align("'fill'").X("10").W("parent.w - .x").Contains(
			TreeItem(),
		),
	)
	return ret
}

type TreeComponent struct {
	*BaseComponent
	focus            Property       `type:"bool"`
	items            Property       `type:"array"`
	nodeMap          Property       `type:"object"`
	onClickItem      Property       `type:"element"`
	selectedChildTop Property       `type:"number"`
	itemHeight       StaticProperty `type:"number"`
	onCreated        Method         `bind:"tree_onCreated.js"`
	onUpdated        Method         `bind:"tree_onUpdated.js"`
	makeNodeMap      Method         `bind:"tree_makeNodeMap.js"`
	nodeToItems      Method         `bind:"tree_nodeToItems.js"`
	selectChild      Method         `bind:"tree_selectChild.js"`
	handleClick      Method         `bind:"tree_handleClick.js"`
	sortChildren     StaticMethod   `bind:"tree_sortChildren.js"`
}

func (b *TreeComponent) Focus(v bool) *TreeComponent {
	b.props["focus"] = fmt.Sprintf("%v", v)
	return b
}

func (b *TreeComponent) Items(s string) *TreeComponent {
	b.props["items"] = s
	return b
}

func (b *TreeComponent) NodeMap(s string) *TreeComponent {
	b.props["nodeMap"] = s
	return b
}
