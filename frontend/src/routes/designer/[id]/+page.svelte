<script lang="ts">
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { base } from "$app/paths";
    import { fetchBrands, fetchProducts } from "$lib/api";
    import type { Brand, Product } from "$lib/types";
    import ProductCard from "$lib/components/ProductCard.svelte";

    let brands = $state<Brand[]>([]);
    let products = $state<Product[]>([]);
    let loading = $state(true);
    let following = $state(false);

    const brand = $derived(brands.find((b) => b.id === $page.params.id));
    const designerProducts = $derived(products.filter((p) => p.brand.id === $page.params.id));
    const followerCount = $derived(brand ? brand.followers + (following ? 1 : 0) : 0);

    onMount(async () => {
        try {
            [brands, products] = await Promise.all([fetchBrands(), fetchProducts()]);
        } finally {
            loading = false;
        }
    });
</script>

<svelte:head>
    <title>{brand ? brand.name : "Designer"} — Negestat</title>
</svelte:head>

{#if loading}
    <div class="state">
        <span class="spinner"></span>
        <p>Loading designer…</p>
    </div>
{:else if !brand}
    <div class="state">
        <h1>Designer not found</h1>
        <a href="{base}/designers">← All designers</a>
    </div>
{:else}
    <a href="{base}/designers" class="back">← All designers</a>

    <header class="profile">
        <img class="avatar" src={brand.avatarUrl} alt={brand.name} />
        <div class="meta">
            <h1>{brand.name}</h1>
            <p class="handle">{brand.handle}</p>
            <p class="followers">{followerCount.toLocaleString()} followers</p>
        </div>
        <button class="follow" class:following onclick={() => (following = !following)}>
            {following ? "✓ Following" : "+ Follow"}
        </button>
    </header>

    <section>
        <h2>Designs <span class="count">{designerProducts.length}</span></h2>
        {#if designerProducts.length === 0}
            <p class="empty">No designs yet.</p>
        {:else}
            <div class="grid">
                {#each designerProducts as product}
                    <ProductCard {product} />
                {/each}
            </div>
        {/if}
    </section>
{/if}

<style>
    .state {
        text-align: center;
        padding: 4rem 0;
        color: var(--text-dim);
    }
    .state a {
        color: var(--brand-2);
        text-decoration: none;
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
    .back {
        display: inline-block;
        color: var(--text-dim);
        text-decoration: none;
        margin-bottom: 1.5rem;
        font-weight: 600;
    }
    .back:hover {
        color: var(--text);
    }
    .profile {
        display: flex;
        align-items: center;
        gap: 1.5rem;
        padding: 2rem;
        background: var(--surface);
        border: 1px solid var(--border);
        border-radius: var(--radius);
        margin-bottom: 3rem;
    }
    @media (max-width: 560px) {
        .profile {
            flex-direction: column;
            text-align: center;
        }
    }
    .avatar {
        width: 96px;
        height: 96px;
        border-radius: 50%;
        object-fit: cover;
        background: var(--surface);
        flex-shrink: 0;
        box-shadow: 0 10px 30px rgba(168, 85, 247, 0.35);
    }
    .meta {
        flex: 1;
    }
    h1 {
        margin: 0;
        font-size: 2rem;
        font-weight: 900;
        letter-spacing: -0.02em;
    }
    .handle {
        margin: 0.25rem 0 0;
        color: var(--brand-2);
        font-weight: 600;
    }
    .followers {
        margin: 0.5rem 0 0;
        font-weight: 600;
        color: var(--text-dim);
    }
    .follow {
        padding: 0.8rem 1.6rem;
        background: var(--gradient);
        color: white;
        border: none;
        border-radius: 12px;
        font-size: 1rem;
        font-weight: 800;
        cursor: pointer;
        box-shadow: 0 10px 26px rgba(168, 85, 247, 0.35);
        transition: transform 0.2s ease;
    }
    .follow:hover {
        transform: translateY(-2px);
    }
    .follow.following {
        background: var(--surface);
        border: 1px solid rgba(52, 211, 153, 0.5);
        color: #6ee7b7;
        box-shadow: none;
    }
    h2 {
        margin: 0 0 1.5rem;
        font-size: 1.6rem;
        font-weight: 800;
        display: flex;
        align-items: center;
        gap: 0.6rem;
    }
    .count {
        font-size: 0.95rem;
        padding: 0.15rem 0.6rem;
        border-radius: 999px;
        background: var(--surface);
        border: 1px solid var(--border);
        color: var(--text-dim);
        font-weight: 700;
    }
    .empty {
        color: var(--text-dim);
    }
    .grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(255px, 1fr));
        gap: 1.5rem;
    }
</style>
