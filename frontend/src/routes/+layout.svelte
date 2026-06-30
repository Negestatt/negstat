<script lang="ts">
    import favicon from "$lib/assets/favicon.svg";

    let { children } = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
    <link
        href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700&display=swap"
        rel="stylesheet"
    />
</svelte:head>

<div class="app">
    <header>
        <nav>
            <a href="/" class="logo">Negestat</a>
            <div class="links">
                <a href="/">Browse</a>
                <a href="/designers">Designers</a>
            </div>
        </nav>
    </header>
    <main>
        {@render children()}
    </main>
    <footer>
        <p>Negestat — print-on-demand marketplace</p>
    </footer>
</div>

<style>
    :global(body) {
        margin: 0;
        font-family:
            "Inter",
            system-ui,
            -apple-system,
            sans-serif;
        background: #f9fafb;
        color: #111827;
        line-height: 1.5;
    }
    :global(*) {
        box-sizing: border-box;
    }
    .app {
        display: flex;
        flex-direction: column;
        min-height: 100vh;
    }
    header {
        background: white;
        border-bottom: 1px solid #e5e7eb;
        position: sticky;
        top: 0;
        z-index: 10;
    }
    nav {
        max-width: 1280px;
        margin: 0 auto;
        padding: 1rem 1.5rem;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .logo {
        font-size: 1.5rem;
        font-weight: 700;
        text-decoration: none;
        color: #111827;
    }
    .links {
        display: flex;
        gap: 1.5rem;
    }
    .links a {
        text-decoration: none;
        color: #6b7280;
        font-weight: 600;
        transition: color 0.2s ease;
    }
    .links a:hover {
        color: #111827;
    }
    main {
        flex: 1;
        max-width: 1280px;
        width: 100%;
        margin: 0 auto;
        padding: 2rem 1.5rem;
    }
    footer {
        background: white;
        border-top: 1px solid #e5e7eb;
        padding: 2rem 1.5rem;
        text-align: center;
        color: #6b7280;
    }
    footer p {
        margin: 0;
    }
</style>
