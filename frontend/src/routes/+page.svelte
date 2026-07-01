<script lang="ts">
    import { products, brands } from "$lib/data";
    import ProductCard from "$lib/components/ProductCard.svelte";

    const trendingProducts = $derived(products.filter((p) => p.trending));
    const trendingBrands = $derived(brands.slice(0, 4));
</script>

<svelte:head>
    <title>Negestat — Print-on-Demand Marketplace</title>
</svelte:head>

<section class="hero">
    <h1>Discover unique designs</h1>
    <p>Support creators. Get custom merch.</p>
</section>

<section class="trending">
    <h2>🔥 Trending Products</h2>
    <div class="grid">
        {#each trendingProducts as product}
            <ProductCard {product} />
        {/each}
    </div>
</section>

<section class="brands">
    <h2>✨ Trending Designers</h2>
    <div class="brand-list">
        {#each trendingBrands as brand}
            <a href="/designer/{brand.id}" class="brand-card">
                <div class="avatar">{brand.name.slice(0, 2).toUpperCase()}</div>
                <div>
                    <h3>{brand.name}</h3>
                    <p>{brand.handle}</p>
                    <p class="followers">{brand.followers.toLocaleString()} followers</p>
                </div>
            </a>
        {/each}
    </div>
</section>

<style>
    .hero {
        text-align: center;
        margin-bottom: 3rem;
        padding: 2rem 0;
    }
    h1 {
        margin: 0 0 0.5rem;
        font-size: 3rem;
        font-weight: 700;
    }
    .hero p {
        margin: 0;
        font-size: 1.25rem;
        color: #6b7280;
    }
    section {
        margin-bottom: 4rem;
    }
    h2 {
        margin: 0 0 1.5rem;
        font-size: 1.75rem;
        font-weight: 700;
    }
    .grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
        gap: 1.5rem;
    }
    .brand-list {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
        gap: 1rem;
    }
    .brand-card {
        display: flex;
        align-items: center;
        gap: 1rem;
        padding: 1.25rem;
        background: white;
        border-radius: 12px;
        text-decoration: none;
        color: inherit;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
        transition:
            transform 0.2s ease,
            box-shadow 0.2s ease;
    }
    .brand-card:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    }
    .avatar {
        width: 64px;
        height: 64px;
        border-radius: 50%;
        background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
        color: white;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 1.25rem;
        font-weight: 700;
        flex-shrink: 0;
    }
    .brand-card h3 {
        margin: 0;
        font-size: 1.125rem;
        font-weight: 600;
    }
    .brand-card p {
        margin: 0.25rem 0 0;
        font-size: 0.875rem;
        color: #6b7280;
    }
    .followers {
        font-weight: 600;
    }
</style>
