function onCreated() {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    if (urlParams.has('code')) {
        const codeValue = urlParams.get('code');
        const decodedValue = decodeURIComponent(codeValue);
        const options = {
            value: decodedValue,
            language: 'go',
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
    } else {
        page.util.fetch('<base_url>/code/get' + queryString).then(v => {
            const options = {
                value: v,
                language: 'go',
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
}