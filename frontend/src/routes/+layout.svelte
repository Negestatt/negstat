<script lang="ts">
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { base } from "$app/paths";
    import favicon from "$lib/assets/favicon.svg";

    let { children } = $props();

    let theme = $state<"dark" | "light">("dark");
    let menuOpen = $state(false);

    const path = $derived($page.url.pathname);

    function isActive(href: string): boolean {
        const current = path.replace(/\/$/, "") || "/";
        if (href === "/") return current === (base || "/");
        const target = base + href;
        return current === target || current.startsWith(target + "/");
    }

    onMount(() => {
        const current = document.documentElement.dataset.theme;
        theme = current === "light" ? "light" : "dark";
    });

    $effect(() => {
        if (typeof document === "undefined") return;
        document.documentElement.dataset.theme = theme;
        try {
            localStorage.setItem("theme", theme);
        } catch (e) {
            // ignore storage errors
        }
    });

    // Lock body scroll while the mobile menu is open.
    $effect(() => {
        if (typeof document === "undefined") return;
        document.body.style.overflow = menuOpen ? "hidden" : "";
    });

    function toggleTheme() {
        theme = theme === "dark" ? "light" : "dark";
    }
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
    <link
        href="https://fonts.googleapis.com/css2?family=Google+Sans:wght@400;500;700&display=swap"
        rel="stylesheet"
    />
</svelte:head>

<div class="app">
    <div class="bg-orbs" aria-hidden="true">
        <span class="orb orb-1"></span>
        <span class="orb orb-2"></span>
        <span class="orb orb-3"></span>
    </div>

    <header>
        <nav>
            <a href="{base}/" class="logo">
                <span class="logo-mark">N</span>
                <span class="logo-text">Negestat</span>
            </a>
            <div class="links desktop-only">
                <a href="{base}/" class:active={isActive("/")}>Browse</a>
                <a href="{base}/designers" class:active={isActive("/designers")}>Designers</a>
            </div>
            <button
                class="theme-toggle"
                onclick={toggleTheme}
                aria-label="Toggle light and dark theme"
                title="Toggle theme"
            >
                {theme === "dark" ? "☀️" : "🌙"}
            </button>
        </nav>
    </header>

    <main>
        {@render children()}
    </main>

    <footer>
        <p>© {new Date().getFullYear()} | Negestat. All rights reserved.</p>
    </footer>

    <!-- Mobile: full-screen menu + bottom trigger -->
    <button
        class="menu-trigger mobile-only"
        onclick={() => (menuOpen = true)}
        aria-label="Open menu"
    >
        <span class="menu-icon">☰</span>
        Menu
    </button>

    {#if menuOpen}
        <div class="mobile-menu">
            <button class="menu-close" onclick={() => (menuOpen = false)} aria-label="Close menu"
                >✕</button
            >
            <nav class="mobile-links">
                <a href="{base}/" class:active={isActive("/")} onclick={() => (menuOpen = false)}
                    >Browse</a
                >
                <a
                    href="{base}/designers"
                    class:active={isActive("/designers")}
                    onclick={() => (menuOpen = false)}
                >
                    Designers
                </a>
            </nav>
            <button class="menu-theme" onclick={toggleTheme}>
                {theme === "dark" ? "☀️ Light mode" : "🌙 Dark mode"}
            </button>
        </div>
    {/if}
</div>

<style>
    :global(:root),
    :global([data-theme="dark"]) {
        --bg: #0b0b14;
        --surface: rgba(255, 255, 255, 0.04);
        --surface-hover: rgba(255, 255, 255, 0.07);
        --border: rgba(255, 255, 255, 0.08);
        --active-border: rgba(168, 85, 247, 0.5);
        --text: #f4f4f8;
        --text-dim: #a0a0b8;
        --header-bg: rgba(11, 11, 20, 0.6);
        --tag-bg: rgba(11, 11, 20, 0.7);
        --menu-bg: rgba(11, 11, 20, 0.97);
        --orb-opacity: 0.5;
        --shadow-hover: 0 20px 40px rgba(0, 0, 0, 0.45);
        --brand-1: #6366f1;
        --brand-2: #a855f7;
        --brand-3: #ec4899;
        --gradient: linear-gradient(120deg, #6366f1 0%, #a855f7 50%, #ec4899 100%);
        --radius: 18px;
    }
    :global([data-theme="light"]) {
        --bg: #f5f5fb;
        --surface: #ffffff;
        --surface-hover: #ffffff;
        --border: rgba(17, 17, 30, 0.1);
        --active-border: rgba(168, 85, 247, 0.6);
        --text: #15151f;
        --text-dim: #61617a;
        --header-bg: rgba(245, 245, 251, 0.72);
        --tag-bg: rgba(255, 255, 255, 0.85);
        --menu-bg: rgba(245, 245, 251, 0.98);
        --orb-opacity: 0.28;
        --shadow-hover: 0 20px 40px rgba(80, 60, 160, 0.18);
    }

    :global(body) {
        margin: 0;
        font-family:
            "Google Sans",
            system-ui,
            -apple-system,
            sans-serif;
        background: var(--bg);
        color: var(--text);
        line-height: 1.5;
        -webkit-font-smoothing: antialiased;
        transition:
            background 0.3s ease,
            color 0.3s ease;
    }
    :global(*) {
        box-sizing: border-box;
    }
    :global(::selection) {
        background: rgba(168, 85, 247, 0.35);
    }

    .app {
        position: relative;
        display: flex;
        flex-direction: column;
        min-height: 100vh;
        overflow: hidden;
    }

    .bg-orbs {
        position: fixed;
        inset: 0;
        z-index: 0;
        pointer-events: none;
    }
    .orb {
        position: absolute;
        border-radius: 50%;
        filter: blur(90px);
        opacity: var(--orb-opacity);
        transition: opacity 0.3s ease;
    }
    .orb-1 {
        width: 480px;
        height: 480px;
        background: #6366f1;
        top: -120px;
        left: -100px;
        animation: float 18s ease-in-out infinite;
    }
    .orb-2 {
        width: 420px;
        height: 420px;
        background: #ec4899;
        top: 30%;
        right: -120px;
        animation: float 22s ease-in-out infinite reverse;
    }
    .orb-3 {
        width: 380px;
        height: 380px;
        background: #a855f7;
        bottom: -140px;
        left: 30%;
        animation: float 20s ease-in-out infinite;
    }
    @keyframes float {
        0%,
        100% {
            transform: translate(0, 0) scale(1);
        }
        50% {
            transform: translate(40px, -30px) scale(1.1);
        }
    }

    header {
        position: sticky;
        top: 0;
        z-index: 20;
        background: var(--header-bg);
        backdrop-filter: blur(16px);
        border-bottom: 1px solid var(--border);
    }
    nav {
        max-width: 1240px;
        margin: 0 auto;
        padding: 1rem 1.5rem;
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 1rem;
    }
    .logo {
        display: flex;
        align-items: center;
        gap: 0.6rem;
        text-decoration: none;
        margin-right: auto;
    }
    .logo-mark {
        display: grid;
        place-items: center;
        width: 36px;
        height: 36px;
        border-radius: 10px;
        background: var(--gradient);
        color: white;
        font-weight: 700;
        font-size: 1.1rem;
        box-shadow: 0 6px 20px rgba(168, 85, 247, 0.4);
    }
    .logo-text {
        font-size: 1.4rem;
        font-weight: 700;
        color: var(--text);
        letter-spacing: -0.02em;
    }
    .links {
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }
    .links a {
        text-decoration: none;
        color: var(--text-dim);
        font-weight: 500;
        font-size: 0.95rem;
        padding: 0.5rem 0.9rem;
        border-radius: 10px;
        border: 1px solid transparent;
        transition: all 0.2s ease;
    }
    .links a:hover {
        color: var(--text);
        background: var(--surface);
    }
    .links a.active {
        color: var(--text);
        border-color: var(--active-border);
    }
    .theme-toggle {
        display: grid;
        place-items: center;
        width: 38px;
        height: 38px;
        border-radius: 10px;
        border: 1px solid var(--border);
        background: var(--surface);
        font-size: 1rem;
        cursor: pointer;
        transition:
            background 0.2s ease,
            transform 0.2s ease;
        flex-shrink: 0;
    }
    .theme-toggle:hover {
        background: var(--surface-hover);
        transform: rotate(-15deg);
    }

    main {
        position: relative;
        z-index: 1;
        flex: 1;
        max-width: 1240px;
        width: 100%;
        margin: 0 auto;
        padding: 2.5rem 1.5rem 4rem;
    }

    footer {
        position: relative;
        z-index: 1;
        border-top: 1px solid var(--border);
        padding: 2rem 1.5rem;
        text-align: center;
        color: var(--text-dim);
        font-size: 0.9rem;
    }
    footer p {
        margin: 0;
    }

    /* Mobile menu */
    .mobile-only {
        display: none;
    }
    .menu-trigger {
        position: fixed;
        bottom: 1.25rem;
        left: 50%;
        transform: translateX(-50%);
        z-index: 30;
        align-items: center;
        gap: 0.5rem;
        padding: 0.85rem 1.6rem;
        border: none;
        border-radius: 999px;
        background: var(--gradient);
        color: white;
        font-family: inherit;
        font-size: 1rem;
        font-weight: 700;
        cursor: pointer;
        box-shadow: 0 10px 30px rgba(168, 85, 247, 0.45);
    }
    .menu-icon {
        font-size: 1.1rem;
    }
    .mobile-menu {
        position: fixed;
        inset: 0;
        z-index: 40;
        background: var(--menu-bg);
        backdrop-filter: blur(20px);
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        gap: 1.5rem;
        animation: fadein 0.2s ease both;
    }
    .menu-close {
        position: absolute;
        top: 1.25rem;
        right: 1.5rem;
        width: 44px;
        height: 44px;
        border-radius: 12px;
        border: 1px solid var(--border);
        background: var(--surface);
        color: var(--text);
        font-size: 1.1rem;
        cursor: pointer;
    }
    .mobile-links {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 1rem;
    }
    .mobile-links a {
        text-decoration: none;
        color: var(--text-dim);
        font-size: 2rem;
        font-weight: 700;
        letter-spacing: -0.02em;
    }
    .mobile-links a.active {
        background: var(--gradient);
        -webkit-background-clip: text;
        background-clip: text;
        color: transparent;
    }
    .menu-theme {
        margin-top: 1rem;
        padding: 0.75rem 1.5rem;
        border-radius: 12px;
        border: 1px solid var(--border);
        background: var(--surface);
        color: var(--text);
        font-family: inherit;
        font-size: 1rem;
        font-weight: 600;
        cursor: pointer;
    }
    @keyframes fadein {
        from {
            opacity: 0;
        }
        to {
            opacity: 1;
        }
    }

    @media (max-width: 768px) {
        .desktop-only {
            display: none;
        }
        .mobile-only {
            display: flex;
        }
        main {
            padding-bottom: 6rem;
        }
    }
</style>
