function handleActive(ele) {
    if (ele.selected) {
        return;
    }
    const oldBgColor = ele.backgroundColor;
    ele.backgroundColor = page.theme.buttonActiveBgColor;
    return () => {
        ele.backgroundColor = oldBgColor;
    };
}