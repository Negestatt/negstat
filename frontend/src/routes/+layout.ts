// Render as a pure client-side SPA: no server-side rendering or prerendering.
// Data comes from bundled JSON at runtime, and the static adapter's fallback shell
// (plus a copied 404.html in the deploy workflow) makes deep links work on GitHub Pages.
export const ssr = false;
export const prerender = false;
