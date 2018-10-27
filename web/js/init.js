
if ('serviceWorker' in navigator) {
    window.addEventListener('load', function() {
        navigator.serviceWorker.register('/service-worker.js').then(function(registration) {
            // Registration was successful
            console.log('[INFO ] ServiceWorker registration successful with scope: ', registration.scope);
        }, function(err) {
            // registration failed :(
            console.log('[ERROR] ServiceWorker registration failed: ', err);
        });
    });
}

if (!WebAssembly.instantiateStreaming) {
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();
let mod, inst;
console.log('[INFO ] Initiating WebAssembly ...');

WebAssembly.instantiateStreaming(fetch("/lib.wasm"), go.importObject).then(async (result) => {
    mod = result.module;
    inst = result.instance;
    console.log('[INFO ] Running Go ...');
    await go.run(inst);
});
