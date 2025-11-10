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