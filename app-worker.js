const cacheName = "app-" + "727d3be3173effae88720553b8636813bfed72e0";

self.addEventListener("install", event => {
  console.log("installing app worker 727d3be3173effae88720553b8636813bfed72e0");

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
  console.log("app worker 727d3be3173effae88720553b8636813bfed72e0 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
