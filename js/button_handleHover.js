function handleHover(ele, hover) {
    if (hover) {
        this.backgroundColor = this.selected ? page.theme.buttonSelectedBgColor : page.theme.buttonHoverBgColor;
    } else {
        this.backgroundColor = this.selected ? page.theme.buttonSelectedBgColor : page.theme.buttonBgColor;
    }
}