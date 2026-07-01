<script lang="ts">
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { base } from "$app/paths";
    import { fetchProducts, placeOrder } from "$lib/api";
    import { formatBirr } from "$lib/data";
    import type { Size, OrderKind, Product } from "$lib/types";

    let products = $state<Product[]>([]);
    let loading = $state(true);

    const product = $derived(products.find((p) => p.id === $page.params.id));

    // Order state using Svelte 5 runes
    let selectedSize = $state<Size | null>(null);
    let quantity = $state(1);
    let orderKind = $state<OrderKind>("standard");
    let orderPending = $state(false);
    let orderMessage = $state<string | null>(null);

    const totalBirr = $derived(
        product
            ? product.priceBirr * quantity + (orderKind === "rush" ? product.rushFeeBirr : 0)
            : 0,
    );

    onMount(async () => {
        try {
            products = await fetchProducts();
        } finally {
            loading = false;
        }
    });

    async function handlePlaceOrder() {
        if (!product || !selectedSize || orderPending) return;

        orderPending = true;
        orderMessage = null;

        try {
            const resp = await placeOrder({
                productId: product.id,
                size: selectedSize,
                quantity,
                isRush: orderKind === "rush",
            });
            orderMessage = `✅ ${resp.message} (Order ID: ${resp.orderId})`;
        } catch (err) {
            orderMessage = `❌ ${err instanceof Error ? err.message : "Order failed"}`;
        } finally {
            orderPending = false;
        }
    }
</script>

<svelte:head>
    <title>{product ? product.title : "Product"} — Negestat</title>
</svelte:head>

{#if loading}
    <div class="state">
        <span class="spinner"></span>
        <p>Loading product…</p>
    </div>
{:else if !product}
    <div class="state">
        <h1>Product not found</h1>
        <a href="{base}/">← Back to browse</a>
    </div>
{:else}
    <a href="{base}/" class="back">← Browse</a>

    <div class="product-detail">
        <div class="visual">
            <div class="artwork" style="background-color: {product.accent}1A;">
                <img src={product.imageUrl} alt={product.title} />
            </div>
        </div>

        <div class="info">
            <p class="brand-line">{product.brand.name}</p>
            <h1>{product.title}</h1>

            <p class="price">
                {formatBirr(product.priceBirr)}
                {#if orderKind === "rush"}
                    <span class="rush-fee">+ {formatBirr(product.rushFeeBirr)} rush</span>
                {/if}
            </p>

            <p class="delivery-promise">🚚 Ships within 2 weeks, guaranteed</p>

            {#if orderMessage}
                <div class="order-message" class:success={orderMessage.startsWith("✅")}>
                    {orderMessage}
                </div>
            {/if}

            <div class="options">
                <div class="option-group">
                    <span class="field-label">Size</span>
                    <div class="size-picker">
                        {#each product.sizes as size}
                            <button
                                class="size-btn"
                                class:selected={selectedSize === size}
                                onclick={() => (selectedSize = size)}
                            >
                                {size}
                            </button>
                        {/each}
                    </div>
                </div>

                <div class="option-group">
                    <label class="field-label" for="quantity">Quantity</label>
                    <input id="quantity" type="number" min="1" max="10" bind:value={quantity} />
                </div>

                <button
                    class="rush-card"
                    class:active={orderKind === "rush"}
                    onclick={() => (orderKind = orderKind === "rush" ? "standard" : "rush")}
                >
                    <div class="rush-check" class:on={orderKind === "rush"}>
                        {orderKind === "rush" ? "✓" : ""}
                    </div>
                    <div class="rush-text">
                        <strong>⚡ Rush delivery — skip the wait</strong>
                        <span>
                            {#if orderKind === "rush"}
                                Printed just for you and shipped within 3 days.
                            {:else}
                                Get it in 3 days for +{formatBirr(product.rushFeeBirr)}.
                            {/if}
                        </span>
                    </div>
                </button>
            </div>

            <button class="cta" disabled={!selectedSize || orderPending} onclick={handlePlaceOrder}>
                {#if orderPending}
                    Placing order…
                {:else if !selectedSize}
                    Select a size
                {:else}
                    Place Order — {formatBirr(totalBirr)}
                {/if}
            </button>
        </div>
    </div>
{/if}

<style>
    .state {
        text-align: center;
        padding: 5rem 0;
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
    .product-detail {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 3rem;
        align-items: start;
    }
    @media (max-width: 820px) {
        .product-detail {
            grid-template-columns: 1fr;
            gap: 2rem;
        }
    }
    .visual {
        position: sticky;
        top: 100px;
    }
    .artwork {
        aspect-ratio: 1;
        border-radius: var(--radius);
        overflow: hidden;
        border: 1px solid var(--border);
    }
    .artwork img {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }
    .brand-line {
        margin: 0 0 0.4rem;
        color: var(--brand-2);
        font-weight: 700;
        font-size: 0.95rem;
    }
    h1 {
        margin: 0 0 1rem;
        font-size: 2.5rem;
        font-weight: 900;
        letter-spacing: -0.03em;
        line-height: 1.05;
    }
    .price {
        margin: 0 0 0.75rem;
        font-size: 2rem;
        font-weight: 800;
    }
    .rush-fee {
        font-size: 1rem;
        color: var(--text-dim);
        font-weight: 500;
    }
    .delivery-promise {
        margin: 0 0 1.5rem;
        font-size: 0.95rem;
        font-weight: 600;
        color: #34d399;
    }
    .order-message {
        padding: 0.9rem 1rem;
        border-radius: 12px;
        margin: 0 0 1.5rem;
        font-size: 0.9rem;
        background: rgba(248, 113, 113, 0.12);
        border: 1px solid rgba(248, 113, 113, 0.3);
        color: #fca5a5;
    }
    .order-message.success {
        background: rgba(52, 211, 153, 0.12);
        border-color: rgba(52, 211, 153, 0.3);
        color: #6ee7b7;
    }
    .options {
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
        margin-bottom: 2rem;
    }
    .field-label {
        display: block;
        font-weight: 700;
        margin-bottom: 0.6rem;
        font-size: 0.9rem;
    }
    .size-picker {
        display: flex;
        gap: 0.5rem;
    }
    .size-btn {
        min-width: 3.5rem;
        height: 3rem;
        padding: 0 0.5rem;
        border: 1px solid var(--border);
        background: var(--surface);
        color: var(--text);
        border-radius: 12px;
        font-weight: 700;
        cursor: pointer;
        transition: all 0.2s ease;
    }
    .size-btn:hover {
        border-color: rgba(168, 85, 247, 0.6);
    }
    .size-btn.selected {
        border-color: transparent;
        background: var(--gradient);
        color: white;
        box-shadow: 0 8px 20px rgba(168, 85, 247, 0.35);
    }
    input[type="number"] {
        width: 5.5rem;
        padding: 0.7rem;
        border: 1px solid var(--border);
        background: var(--surface);
        color: var(--text);
        border-radius: 12px;
        font-size: 1rem;
        font-weight: 600;
    }
    .rush-card {
        display: flex;
        align-items: center;
        gap: 0.9rem;
        text-align: left;
        padding: 1rem 1.1rem;
        border-radius: 14px;
        border: 1px solid var(--border);
        background: var(--surface);
        cursor: pointer;
        transition: all 0.2s ease;
        width: 100%;
        color: var(--text);
    }
    .rush-card:hover {
        border-color: rgba(168, 85, 247, 0.5);
    }
    .rush-card.active {
        border-color: rgba(168, 85, 247, 0.7);
        background: rgba(168, 85, 247, 0.1);
    }
    .rush-check {
        width: 26px;
        height: 26px;
        border-radius: 8px;
        border: 2px solid var(--border);
        display: grid;
        place-items: center;
        flex-shrink: 0;
        font-size: 0.85rem;
        font-weight: 800;
        color: white;
        transition: all 0.2s ease;
    }
    .rush-check.on {
        background: var(--gradient);
        border-color: transparent;
    }
    .rush-text {
        display: flex;
        flex-direction: column;
        gap: 0.15rem;
    }
    .rush-text strong {
        font-size: 0.95rem;
    }
    .rush-text span {
        font-size: 0.85rem;
        color: var(--text-dim);
    }
    .cta {
        width: 100%;
        padding: 1.1rem 2rem;
        background: var(--gradient);
        color: white;
        border: none;
        border-radius: 14px;
        font-size: 1.1rem;
        font-weight: 800;
        cursor: pointer;
        box-shadow: 0 12px 30px rgba(168, 85, 247, 0.35);
        transition:
            transform 0.2s ease,
            box-shadow 0.2s ease;
    }
    .cta:hover:not(:disabled) {
        transform: translateY(-2px);
        box-shadow: 0 16px 38px rgba(168, 85, 247, 0.5);
    }
    .cta:disabled {
        background: var(--surface);
        color: var(--text-dim);
        box-shadow: none;
        cursor: not-allowed;
        border: 1px solid var(--border);
    }
</style>
