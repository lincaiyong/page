function onCreated() {
    const queryString = window.location.search;
    page.util.fetch('<base_url>/code/get' + queryString).then(v => {
        const leftModel = monaco.editor.createModel('原始文本', 'text/plain');
        const rightModel = monaco.editor.createModel('修改后的文本', 'text/plain');

        this._editor = monaco.editor.createDiffEditor(this.ref, {
            automaticLayout: true,
        });
        this._editor.setModel({
            original: leftModel,
            modified: rightModel
        });
    }).catch(err => page.log.error(err));
}