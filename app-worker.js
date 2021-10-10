const cacheName = "app-" + "d8bd5654891b4383c231baa45f1901b2ef68dd5d";

self.addEventListener("install", event => {
  console.log("installing app worker d8bd5654891b4383c231baa45f1901b2ef68dd5d");

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
  console.log("app worker d8bd5654891b4383c231baa45f1901b2ef68dd5d is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
