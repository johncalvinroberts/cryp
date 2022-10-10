export const mountCloudflareAnalytics = () => {
  const script = document.createElement("script");
  script.src = "https://static.cloudflareinsights.com/beacon.min.js";
  script.defer = true;
  script.setAttribute(
    "data-cf-beacon",
    '{"token": "d75931f6ae3f4c1c918fce41593cbb98"}'
  );
  document.body.appendChild(script);
};
