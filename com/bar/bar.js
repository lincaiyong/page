function handleMouseDown(ele, mouseDownEvent) {
    const prevEle = ele._('prev');
    const nextEle = ele._('next');
    const [left, right] = [prevEle, nextEle];
    const [top, bottom] = [prevEle, nextEle];
    const state = {prevX: mouseDownEvent.clientX, prevY: mouseDownEvent.clientY};
    const cancelMouseMoveListener = page.event.addListener(window, 'mousemove', ev => {
        const safeDist = 80;
        if (ele.cursor === 'col-resize') {
            const newX = ele.x + ev.clientX - state.prevX;
            state.prevX = ev.clientX;
            if (newX < left.x + safeDist) {
                ele.x = left.x + safeDist;
            } else if (newX > right.x + right.w - safeDist) {
                ele.x = right.x + right.w - safeDist;
            } else {
                ele.x = newX;
            }
        } else {
            const newY = ele.y + ev.clientY - state.prevY;
            state.prevY = ev.clientY;
            if (newY < top.y + safeDist) {
                ele.y = top.y + safeDist;
            } else if (newY > bottom.y + bottom.h - safeDist) {
                ele.y = bottom.y + bottom.h - safeDist;
            } else {
                ele.y = newY;
            }
        }
    });
    page.event.onceListener(window, 'mouseup', () => {
        cancelMouseMoveListener();
    });
}

function onUpdated(k, v) {
    switch (k) {
    }
}