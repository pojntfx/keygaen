const cacheName = "app-" + "fd7f88ebd94a32079523a8cf43590a3ad9ea7c48";

self.addEventListener("install", event => {
  console.log("installing app worker fd7f88ebd94a32079523a8cf43590a3ad9ea7c48");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/keygean",
          "/keygean/app.css",
          "/keygean/app.js",
          "/keygean/manifest.webmanifest",
          "/keygean/wasm_exec.js",
          "/keygean/web/app.wasm",
          "/keygean/web/default.png",
          "/keygean/web/index.css",
          "/keygean/web/large.png",
          "https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly-addons.css",
          "https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly.css",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
      })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker fd7f88ebd94a32079523a8cf43590a3ad9ea7c48 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
