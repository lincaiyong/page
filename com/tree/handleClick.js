function handleClick(child, ev) {
    this.selectChild(child, true);
    // 通知发生点击事件
    if (this.onClickItem instanceof Function) {
        this.onClickItem(child, ev);
    }
    // 目录展开折叠
    if (!child.data.leaf) {
        const {key} = child.data;
        const node = this.nodeMap[key];
        node.collapsed = !node.collapsed;
        this.containerEle.items = this.nodeToItems(this.nodeMap, '', 0, 0);
    }
    // 处理blur
    this.onClickOutside = (_, event) => {
        if (ev !== event) {
            this.focus = false;
            return true;
        }
        return false;
    };
}