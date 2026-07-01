<script lang="ts">
    import { onMount } from "svelte";
    import { base } from "$app/paths";
    import { fetchBrands } from "$lib/api";
    import type { Brand } from "$lib/types";

    let brands = $state<Brand[]>([]);
    let loading = $state(true);
    let error = $state<string | null>(null);

    onMount(async () => {
        try {
            brands = await fetchBrands();
        } catch (err) {
            error = err instanceof Error ? err.message : "Failed to load designers";
        } finally {
            loading = false;
        }
    });
</script>

<svelte:head>
    <title>Designers — Negestat</title>
</svelte:head>

<div class="head">
    <h1>Meet the <span class="grad">designers</span></h1>
    <p>Independent creators behind every drop</p>
</div>

{#if loading}
    <div class="state">
        <span class="spinner"></span>
        <p>Loading designers…</p>
    </div>
{:else if error}
    <div class="state error">❌ {error}</div>
{:else}
    <div class="grid">
        {#each brands as brand}
            <a href="{base}/designer/{brand.id}" class="card">
                <img class="avatar" src={brand.avatarUrl} alt={brand.name} loading="lazy" />
                <div class="meta">
                    <h3>{brand.name}</h3>
                    <p>{brand.handle}</p>
                    <p class="followers">{brand.followers.toLocaleString()} followers</p>
                </div>
            </a>
        {/each}
    </div>
{/if}

<style>
    .head {
        margin-bottom: 2.5rem;
    }
    h1 {
        margin: 0;
        font-size: clamp(2.25rem, 5vw, 3.25rem);
        font-weight: 900;
        letter-spacing: -0.03em;
    }
    .grad {
        background: var(--gradient);
        -webkit-background-clip: text;
        background-clip: text;
        color: transparent;
    }
    .head p {
        margin: 0.5rem 0 0;
        color: var(--text-dim);
        font-size: 1.1rem;
    }
    .state {
        text-align: center;
        padding: 4rem 0;
        color: var(--text-dim);
    }
    .state.error {
        color: #f87171;
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
    .grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(290px, 1fr));
        gap: 1rem;
    }
    .card {
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
    .card:hover {
        transform: translateY(-3px);
        border-color: rgba(168, 85, 247, 0.5);
        background: rgba(255, 255, 255, 0.06);
    }
    .avatar {
        width: 60px;
        height: 60px;
        border-radius: 50%;
        object-fit: cover;
        background: var(--surface);
        flex-shrink: 0;
    }
    .meta h3 {
        margin: 0;
        font-size: 1.05rem;
        font-weight: 700;
    }
    .meta p {
        margin: 0.2rem 0 0;
        font-size: 0.85rem;
        color: var(--text-dim);
    }
    .followers {
        font-weight: 600;
    }
</style>
