function onUpdated(k, v) {
    const onUpdatedFn = this.model.properties.onUpdatedFn[0]();
    onUpdatedFn(this, k, v);
}