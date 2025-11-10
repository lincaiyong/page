function onCreated() {
    const queryString = window.location.search; // ?language=javascript&...
    const urlParams = new URLSearchParams(queryString);
    page.util.fetch('<base_url>/code/get' + queryString).then(v => {
        const options = {
            value: v,
            language: urlParams.get('language') || 'go',
            theme: 'vs',
            automaticLayout: true,
            lineNumbers: 'on', // 'off'
            minimap: {
                enabled: false,
            },
            readOnly: true,
            // fontFamily: '',
            // glyphMargin: false,
            // suggestOnTriggerCharacters: false,
        };
        this._editor = monaco.editor.create(this.ref, options);
    }).catch(err => page.log.error(err));
}