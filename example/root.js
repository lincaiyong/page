function compute(container, idx, prev) {
    return {
        key: ''+idx,
        x: 0,
        y: 20 * idx,
        w: 200,
        h: 20,
        text: 'hello world!' + idx,
    }
}

function onUpdated(k, v) {
    if (k === 'data') {
        this.children[0].children[0].text = v.text;
    }
}

function onCreated() {
    const container = this.containerEle;
    container.items = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16];
}