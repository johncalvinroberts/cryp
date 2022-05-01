import {
  cleanupOutdatedCaches,
  createHandlerBoundToURL,
  precacheAndRoute,
} from "workbox-precaching";
import { clientsClaim } from "workbox-core";
import { NavigationRoute, registerRoute } from "workbox-routing";

const ctx: ServiceWorkerGlobalScope = self as any;

// self.__WB_MANIFEST is default injection point
precacheAndRoute(ctx.__WB_MANIFEST);

// clean old assets
cleanupOutdatedCaches();

// to allow work offline
registerRoute(new NavigationRoute(createHandlerBoundToURL("index.html")));

ctx.skipWaiting();
clientsClaim();
