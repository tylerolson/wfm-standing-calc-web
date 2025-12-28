import type { HandleFetch } from "@sveltejs/kit";

import { env } from "$env/dynamic/private";
const API_URL = env.API_URL ?? "http://localhost:8080";

// We require external proxying of our /api route.
// In the browser this works fine by going to our proxied route.
// But for SSR in a container, the server hostname will be different so we must change it manually.
export const handleFetch: HandleFetch = async ({ request, fetch }) => {
  const url = new URL(request.url);

  if (url.pathname.startsWith("/api") || url.pathname.startsWith("/docs")) {
    return fetch(`${API_URL}${url.pathname}`);
  }

  return fetch(request);
};
