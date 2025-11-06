function onUpdated(k, v) {
    switch (k) {
        case 'src':
            if (this.tag === 'svg') {
                page.util.fetch('<base_url>/res/' + this.src).then(v => this.ref.innerHTML = v).catch(err => page.log.error(err));
            } else if (this.tag === 'img') {
                this.ref.setAttribute(k, '<base_url>/res/' + this.src);
            }
            break;
        default:
            break;
    }
}
