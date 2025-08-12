import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vite.dev/config/
export default defineConfig({
  plugins: [tailwindcss(), svelte()],
  server: {
    //when using vite server, instead look for our go server
    host: "0.0.0.0",
    proxy: { "/api": "http://localhost:8080" },
  },
});
