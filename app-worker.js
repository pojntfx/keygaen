const cacheName = "app-" + "97ed3c4a56843fe40c30b34f6599eebbf0e10c09";

self.addEventListener("install", event => {
  console.log("installing app worker 97ed3c4a56843fe40c30b34f6599eebbf0e10c09");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/gridge",
          "/gridge/app.css",
          "/gridge/app.js",
          "/gridge/manifest.webmanifest",
          "/gridge/wasm_exec.js",
          "/gridge/web/app.wasm",
          "/gridge/web/default.png",
          "/gridge/web/index.css",
          "/gridge/web/large.png",
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
  console.log("app worker 97ed3c4a56843fe40c30b34f6599eebbf0e10c09 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
