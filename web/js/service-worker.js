
var cacheName = 'DIEGO';
var dataCacheName = 'DIEGO';
var filesToCache = [
    "/favicon.ico",
    "/img/logo.png",
    "/lib.wasm",
    "/manifest.json",
    "/service-worker.js",
];

self.addEventListener('install', function(e) {
    console.log('[INFO ] Installing Service Worker ...');
    e.waitUntil(
        caches.open(cacheName).then(function(cache) {
            console.log('[INFO ] Service Worker cache created.');
            return cache.addAll(filesToCache);
        })
    );
});

self.addEventListener('activate', function(e) {
    console.log('[INFO ] Activating Service Worker ...');
    e.waitUntil(
        caches.keys().then(function(keyList) {
            return Promise.all(keyList.map(function(key) {
                if (key !== cacheName && key !== dataCacheName) {
                    console.log('[INFO ] Removing old cache', key);
                    return caches.delete(key);
                }
            }));
        })
    );
    return self.clients.claim();
});

self.addEventListener('fetch', function(e) {
    console.log('[INFO ] Fetching Service Worker ...', e.request.url);
    e.respondWith(
        caches.match(e.request).then(function(response) {
            console.log('[INFO ] Service Worker fetched.');
            return response || fetch(e.request);
        })
    );
});
