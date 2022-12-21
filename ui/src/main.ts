import App from "./App.svelte";
import "./assets/global.css";
import { registerSW } from "virtual:pwa-register";

const IS_PROD = import.meta.env.PROD;

const mountCloudflareAnalytics = () => {
  const script = document.createElement("script");
  script.src = "https://static.cloudflareinsights.com/beacon.min.js";
  script.defer = true;
  script.setAttribute(
    "data-cf-beacon",
    '{"token": "d75931f6ae3f4c1c918fce41593cbb98"}'
  );
  document.body.appendChild(script);
};

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
