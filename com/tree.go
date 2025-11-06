package com

import "fmt"

func Tree() *TreeComponent {
	ret := &TreeComponent{}
	ret.BaseComponent = NewBaseComponent("div", ret)
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
	onUpdated        Method         `bind:"tree_onUpdated.js"`
	makeNodeMap      Method         `bind:"tree_makeNodeMap.js"`
	nodeToItems      Method         `bind:"tree_nodeToItems.js"`
	selectChild      Method         `bind:"tree_selectChild.js"`
	handleClick      Method         `bind:"tree_handleClick.js"`
	sortChildren     StaticMethod   `bind:"tree_sortChildren.js"`
	compute          StaticMethod   `bind:"tree_compute.js"`
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
