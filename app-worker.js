const cacheName = "app-" + "a0418fc0c8b03c6a6e8449066a5fa628149afa51";

self.addEventListener("install", event => {
  console.log("installing app worker a0418fc0c8b03c6a6e8449066a5fa628149afa51");

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
  console.log("app worker a0418fc0c8b03c6a6e8449066a5fa628149afa51 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
