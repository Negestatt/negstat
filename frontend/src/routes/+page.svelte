<script lang="ts">
    import { onMount } from "svelte";
    import { base } from "$app/paths";
    import { fetchProducts, fetchBrands } from "$lib/api";
    import type { Product, Brand } from "$lib/types";
    import ProductCard from "$lib/components/ProductCard.svelte";

    let products = $state<Product[]>([]);
    let brands = $state<Brand[]>([]);
    let loading = $state(true);
    let error = $state<string | null>(null);

    let sliderEl = $state<HTMLDivElement | null>(null);

    const trendingProducts = $derived(products.filter((p) => p.trending));
    const trendingBrands = $derived(brands.slice(0, 4));

    onMount(async () => {
        try {
            [products, brands] = await Promise.all([fetchProducts(), fetchBrands()]);
        } catch (err) {
            error = err instanceof Error ? err.message : "Failed to load data";
        } finally {
            loading = false;
        }
    });

    function slide(dir: number) {
        if (!sliderEl) return;
        sliderEl.scrollBy({ left: dir * 320, behavior: "smooth" });
    }
</script>

<svelte:head>
    <title>Negestat — Print-on-Demand Marketplace</title>
</svelte:head>

{#if loading}
    <div class="state">
        <span class="spinner"></span>
        <p>Loading marketplace…</p>
    </div>
{:else if error}
    <div class="state error">
        <p>❌ {error}</p>
        <p class="hint">Make sure the backend API is running on http://localhost:8080</p>
    </div>
{:else}
    <section class="hero">
        <span class="badge">✨ Fresh drops every week</span>
        <h1>Wear something <span class="grad">nobody else has.</span></h1>
        <p>
            Discover bold designs from independent creators. Order custom merch — printed on demand,
            shipped to your door.
        </p>
        <div class="hero-actions">
            <a href="#trending" class="btn-primary">Browse drops</a>
            <a href="{base}/designers" class="btn-ghost">Meet the designers</a>
        </div>
    </section>

    <section id="trending">
        <div class="section-head">
            <div>
                <h2>🔥 Trending Products</h2>
                <p>What everyone's ordering right now</p>
            </div>
            <div class="slider-controls">
                <button onclick={() => slide(-1)} aria-label="Scroll left">‹</button>
                <button onclick={() => slide(1)} aria-label="Scroll right">›</button>
            </div>
        </div>

        <div class="slider" bind:this={sliderEl}>
            {#each trendingProducts as product, i (product.id)}
                <div class="slide" style="animation-delay: {i * 70}ms">
                    <ProductCard {product} />
                </div>
            {/each}
        </div>
    </section>

    <section>
        <div class="section-head">
            <div>
                <h2>✨ Trending Designers</h2>
                <p>Creators worth following</p>
            </div>
        </div>
        <div class="brand-list">
            {#each trendingBrands as brand, i (brand.id)}
                <a
                    href="{base}/designer/{brand.id}"
                    class="brand-card reveal"
                    style="animation-delay: {i * 80}ms"
                >
                    <img class="avatar" src={brand.avatarUrl} alt={brand.name} loading="lazy" />
                    <div class="brand-meta">
                        <h3>{brand.name}</h3>
                        <p>{brand.handle}</p>
                        <p class="followers">{brand.followers.toLocaleString()} followers</p>
                    </div>
                </a>
            {/each}
        </div>
    </section>
{/if}

<style>
    .state {
        text-align: center;
        padding: 5rem 0;
        color: var(--text-dim);
    }
    .state.error {
        color: #f87171;
    }
    .hint {
        margin-top: 0.5rem;
        font-size: 0.875rem;
        color: var(--text-dim);
    }
    .spinner {
        display: inline-block;
        width: 38px;
        height: 38px;
        border: 3px solid var(--border);
        border-top-color: var(--brand-2);
        border-radius: 50%;
        animation: spin 0.8s linear infinite;
        margin-bottom: 1rem;
    }
    @keyframes spin {
        to {
            transform: rotate(360deg);
        }
    }

    .hero {
        text-align: center;
        padding: 4rem 0 5rem;
        max-width: 760px;
        margin: 0 auto;
        animation: reveal 0.7s ease both;
    }
    .badge {
        display: inline-block;
        padding: 0.4rem 1rem;
        border-radius: 999px;
        background: var(--surface);
        border: 1px solid var(--border);
        color: var(--text-dim);
        font-size: 0.85rem;
        font-weight: 500;
        margin-bottom: 1.5rem;
    }
    h1 {
        margin: 0 0 1.25rem;
        font-size: clamp(2.5rem, 6vw, 4.25rem);
        font-weight: 700;
        letter-spacing: -0.03em;
        line-height: 1.05;
    }
    .grad {
        background: var(--gradient);
        -webkit-background-clip: text;
        background-clip: text;
        color: transparent;
    }
    .hero p {
        margin: 0 auto;
        max-width: 540px;
        font-size: 1.2rem;
        color: var(--text-dim);
    }
    .hero-actions {
        display: flex;
        gap: 0.75rem;
        justify-content: center;
        margin-top: 2rem;
        flex-wrap: wrap;
    }
    .btn-primary {
        padding: 0.85rem 1.75rem;
        border-radius: 12px;
        background: var(--gradient);
        color: white;
        font-weight: 700;
        text-decoration: none;
        box-shadow: 0 10px 30px rgba(168, 85, 247, 0.35);
        transition:
            transform 0.2s ease,
            box-shadow 0.2s ease;
    }
    .btn-primary:hover {
        transform: translateY(-2px);
        box-shadow: 0 14px 36px rgba(168, 85, 247, 0.5);
    }
    .btn-ghost {
        padding: 0.85rem 1.75rem;
        border-radius: 12px;
        background: var(--surface);
        border: 1px solid var(--border);
        color: var(--text);
        font-weight: 700;
        text-decoration: none;
        transition: background 0.2s ease;
    }
    .btn-ghost:hover {
        background: var(--surface-hover);
    }

    section {
        margin-bottom: 4.5rem;
    }
    .section-head {
        display: flex;
        align-items: flex-end;
        justify-content: space-between;
        gap: 1rem;
        margin-bottom: 1.75rem;
    }
    h2 {
        margin: 0;
        font-size: 1.85rem;
        font-weight: 700;
        letter-spacing: -0.02em;
    }
    .section-head p {
        margin: 0.35rem 0 0;
        color: var(--text-dim);
    }
    .slider-controls {
        display: flex;
        gap: 0.5rem;
        flex-shrink: 0;
    }
    .slider-controls button {
        width: 42px;
        height: 42px;
        border-radius: 12px;
        border: 1px solid var(--border);
        background: var(--surface);
        color: var(--text);
        font-size: 1.5rem;
        line-height: 1;
        cursor: pointer;
        transition:
            background 0.2s ease,
            transform 0.2s ease;
    }
    .slider-controls button:hover {
        background: var(--surface-hover);
        transform: scale(1.05);
    }

    .slider {
        display: flex;
        gap: 1.5rem;
        overflow-x: auto;
        scroll-snap-type: x mandatory;
        padding: 0.5rem 0.25rem 1.25rem;
        margin: 0 -0.25rem;
        scrollbar-width: none;
    }
    .slider::-webkit-scrollbar {
        display: none;
    }
    .slide {
        flex: 0 0 270px;
        scroll-snap-align: start;
        animation: reveal 0.55s ease both;
    }

    .brand-list {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(290px, 1fr));
        gap: 1rem;
    }
    .brand-card {
        display: flex;
        align-items: center;
        gap: 1rem;
        padding: 1.25rem;
        background: var(--surface);
        border: 1px solid var(--border);
        border-radius: var(--radius);
        text-decoration: none;
        color: inherit;
        transition:
            transform 0.2s ease,
            border-color 0.2s ease,
            background 0.2s ease;
    }
    .brand-card:hover {
        transform: translateY(-3px);
        border-color: rgba(168, 85, 247, 0.5);
        background: var(--surface-hover);
    }
    .reveal {
        animation: reveal 0.55s ease both;
    }
    .avatar {
        width: 60px;
        height: 60px;
        border-radius: 50%;
        object-fit: cover;
        background: var(--surface);
        flex-shrink: 0;
    }
    .brand-meta h3 {
        margin: 0;
        font-size: 1.05rem;
        font-weight: 700;
    }
    .brand-meta p {
        margin: 0.2rem 0 0;
        font-size: 0.85rem;
        color: var(--text-dim);
    }
    .followers {
        font-weight: 500;
    }

    @keyframes reveal {
        from {
            opacity: 0;
            transform: translateY(18px);
        }
        to {
            opacity: 1;
            transform: none;
        }
    }
</style>
