import {
  cleanupOutdatedCaches,
  createHandlerBoundToURL,
  precacheAndRoute,
} from "workbox-precaching";
import { NavigationRoute, registerRoute } from "workbox-routing";

const ctx: ServiceWorkerGlobalScope = self as any;

ctx.addEventListener("message", (event) => {
  if (event.data && event.data.type === "SKIP_WAITING") ctx.skipWaiting();
});

// self.__WB_MANIFEST is default injection point
precacheAndRoute(ctx.__WB_MANIFEST);

// clean old assets
cleanupOutdatedCaches();

// to allow work offline
registerRoute(new NavigationRoute(createHandlerBoundToURL("index.html")));
