const cacheName = "app-" + "58da5a1078cb88eb153465dd23b810339f593ccd";

self.addEventListener("install", event => {
  console.log("installing app worker 58da5a1078cb88eb153465dd23b810339f593ccd");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/keygaen",
          "/keygaen/app.css",
          "/keygaen/app.js",
          "/keygaen/manifest.webmanifest",
          "/keygaen/wasm_exec.js",
          "/keygaen/web/app.wasm",
          "/keygaen/web/default.png",
          "/keygaen/web/index.css",
          "/keygaen/web/large.png",
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
  console.log("app worker 58da5a1078cb88eb153465dd23b810339f593ccd is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
