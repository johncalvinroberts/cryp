import { defineConfig, loadEnv } from "vite";
import { VitePWA } from "vite-plugin-pwa";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, "env"); // reuse vite's env parser to inject into our index.html
  return {
    plugins: [
      svelte(),
      VitePWA({
        includeAssets: [
          "gou.svg",
          "favicon.ico",
          "robots.txt",
          "apple-touch-icon.png",
        ],
        manifest: {
          name: "cryp.",
          short_name: "cryp.",
          description: "Encrypt & decrypt any files",
          theme_color: "#282828",
          icons: [
            {
              src: "pwa-192x192.png",
              sizes: "192x192",
              type: "image/png",
            },
            {
              src: "pwa-512x512.png",
              sizes: "512x512",
              type: "image/png",
            },
            {
              src: "pwa-512x512.png",
              sizes: "512x512",
              type: "image/png",
              purpose: "any maskable",
            },
          ],
        },
      }),
      htmlPlugin(loadEnv(mode, ".")),
    ],
    server: {
      port: 1234,
    },
  };
});

/**
 * Replace env variables in index.html
 * @see https://github.com/vitejs/vite/issues/3105#issuecomment-939703781
 * @see https://vitejs.dev/guide/api-plugin.html#transformindexhtml
 */
function htmlPlugin(env: ReturnType<typeof loadEnv>) {
  return {
    name: "html-transform",
    transformIndexHtml: {
      enforce: "pre" as const,
      transform: (html: string): string =>
        html.replace(/%(.*?)%/g, (match, p1) => env[p1] ?? match),
    },
  };
}
