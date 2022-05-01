import App from "./App.svelte";
import "./assets/global.css";
import { registerSW } from "virtual:pwa-register";

if ("serviceWorker" in navigator) {
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
const app = new App({
  target: document.getElementById("app"),
});

export default app;
