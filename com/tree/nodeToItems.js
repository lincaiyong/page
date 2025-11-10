function nodeToItems(nodeMap, key, index, depth){
    const node = nodeMap[key];
    if (!node || !node.children || node.collapsed) {
        return [];
    }
    let ret = [];
    const children = node.children;
    for (let i = 0; i < children.length; i++) {
        const childNode = children[i];
        const item = {
            index: index,
            key: childNode.key,
            depth: depth,
            leaf: childNode.children.length === 0,
            collapsed: childNode.collapsed,
            text: childNode.text,
        };
        ret.push(item);
        const tmp = this.nodeToItems(
            nodeMap,
            childNode.key,
            index + 1,
            depth + 1,
        );
        ret = ret.concat(tmp);
        index = index + tmp.length + 1;
    }
    return ret;
}