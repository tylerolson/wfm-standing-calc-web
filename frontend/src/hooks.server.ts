import type { Handle } from "@sveltejs/kit";

const API_URL = process.env.API_URL || "http://localhost:8080";

export const handle: Handle = async ({ event, resolve }) => {
  const pathname = event.url.pathname;

  if (pathname.startsWith("/api") || pathname.startsWith("/docs")) {
    const apiUrl = `${API_URL}${pathname}${event.url.search}`;

    try {
      return await fetch(apiUrl, {
        method: event.request.method,
        headers: event.request.headers,
        body: event.request.body,
      });
    } catch (error) {
      console.error("API proxy error:", error);
      return new Response("API server unavailable", { status: 502 });
    }
  }

  return resolve(event);
};
