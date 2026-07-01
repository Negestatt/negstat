<script lang="ts">
    import { base } from "$app/paths";
    import type { Product } from "../types";
    import { formatBirr } from "../data";

    let { product }: { product: Product } = $props();
</script>

<a href="{base}/product/{product.id}" class="card">
    <div class="artwork" style="background-color: {product.accent}1A;">
        <img src={product.imageUrl} alt={product.title} loading="lazy" />
        <span class="ship-tag">🚚 ~2 weeks</span>
    </div>
    <div class="details">
        <h3>{product.title}</h3>
        <p class="brand">{product.brand.name}</p>
        <p class="price">{formatBirr(product.priceBirr)}</p>
    </div>
</a>

<style>
    .card {
        display: block;
        border-radius: var(--radius);
        overflow: hidden;
        background: var(--surface);
        border: 1px solid var(--border);
        text-decoration: none;
        color: inherit;
        transition:
            transform 0.25s ease,
            border-color 0.25s ease,
            box-shadow 0.25s ease;
    }
    .card:hover {
        transform: translateY(-6px);
        border-color: rgba(168, 85, 247, 0.5);
        box-shadow: var(--shadow-hover);
    }
    .artwork {
        position: relative;
        aspect-ratio: 1;
        overflow: hidden;
    }
    .artwork img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        transition: transform 0.4s ease;
    }
    .card:hover .artwork img {
        transform: scale(1.06);
    }
    .ship-tag {
        position: absolute;
        bottom: 0.6rem;
        left: 0.6rem;
        padding: 0.3rem 0.7rem;
        border-radius: 999px;
        background: var(--tag-bg);
        border: 1px solid var(--border);
        backdrop-filter: blur(8px);
        font-size: 0.7rem;
        font-weight: 600;
        color: var(--text);
    }
    .details {
        padding: 1rem 1.1rem 1.25rem;
    }
    h3 {
        margin: 0 0 0.2rem;
        font-size: 1.1rem;
        font-weight: 700;
        letter-spacing: -0.01em;
    }
    .brand {
        margin: 0 0 0.75rem;
        font-size: 0.85rem;
        color: var(--text-dim);
    }
    .price {
        margin: 0;
        font-size: 1.3rem;
        font-weight: 800;
        background: var(--gradient);
        -webkit-background-clip: text;
        background-clip: text;
        color: transparent;
        width: fit-content;
    }
</style>
