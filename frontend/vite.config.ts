import process from "node:process";
import adapter from "@sveltejs/adapter-static";
import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [
    sveltekit({
      compilerOptions: {
        // Force runes mode for the project, except for libraries. Can be removed in svelte 6.
        runes: ({ filename }) =>
          filename.split(/[/\\]/).includes("node_modules") ? undefined : true,
      },

      // Static adapter so the app deploys as a fully static SPA (GitHub Pages).
      // `fallback` produces an app shell served for any client-routed path.
      adapter: adapter({ fallback: "index.html" }),

      // On GitHub Pages the project site is served from /<repo>/, so the base path
      // is provided at build time via BASE_PATH (empty for local dev).
      paths: {
        base: (process.env.BASE_PATH || "") as "" | `/${string}`,
      },
    }),
  ],
});
