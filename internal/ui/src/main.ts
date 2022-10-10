import App from "./App.svelte";
import "./assets/global.css";
import { registerSW } from "virtual:pwa-register";
import { mountCloudflareAnalytics } from "./lib/analytics";

const IS_PROD = import.meta.env.PROD;

if ("serviceWorker" in navigator && IS_PROD) {
  // && !/localhost/.test(window.location) && !/lvh.me/.test(window.location)) {
  const updateSW = registerSW({
    onNeedRefresh() {
      const shouldUpdate = confirm("A new update is available");
      if (shouldUpdate) {
        updateSW(true);
      }
    },
  });
}


if (IS_PROD) {
  mountCloudflareAnalytics();
}

const target = document.getElementById("app") || document.body;
const app = new App({
  target,
});

export default app;
