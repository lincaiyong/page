function onCreated() {
    if (!this.list) {
        const child = page.createElement(this.model.slot[0], this);
        this.childWidth = child.w;
        this.childHeight = child.h;
        child.onUpdated((k, v) => {
            if (k === 'w') {
                this.childWidth = v;
            } else if (k === 'h') {
                this.childHeight = v;
            }
        });
    }

    if (this.scrollable) {
        this.hBarEle = new Scrollbar(this, 'h');
        this.vBarEle = new Scrollbar(this, 'v');
        const bars = [this.hBarEle, this.vBarEle];
        bars.forEach(bar => bar.initDraggable());
        this.onWheel = (_, ev) => {
            ev.preventDefault();
            bars.forEach(bar => bar.handleWheel(ev));
        };
    }
}