<script lang="ts">
	import type { Product } from '../types';
	import { formatCents } from '../data';
	import BatchProgress from './BatchProgress.svelte';

	let { product }: { product: Product } = $props();
</script>

<a href="/product/{product.id}" class="card">
	<div class="artwork" style="background: linear-gradient(135deg, {product.accent} 0%, {product.accent}CC 100%);">
		<span class="emoji">{product.emoji}</span>
	</div>
	<div class="details">
		<h3>{product.title}</h3>
		<p class="brand">{product.brand.name}</p>
		<BatchProgress {product} compact />
		<p class="price">{formatCents(product.priceCents)}</p>
	</div>
</a>

<style>
	.card {
		display: block;
		border-radius: 12px;
		overflow: hidden;
		background: white;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
		transition: transform 0.2s ease, box-shadow 0.2s ease;
		text-decoration: none;
		color: inherit;
	}
	.card:hover {
		transform: translateY(-4px);
		box-shadow: 0 4px 16px rgba(0, 0, 0, 0.12);
	}
	.artwork {
		aspect-ratio: 1;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 5rem;
	}
	.emoji {
		filter: drop-shadow(0 2px 8px rgba(0, 0, 0, 0.1));
	}
	.details {
		padding: 1rem;
	}
	h3 {
		margin: 0 0 0.25rem;
		font-size: 1.1rem;
		font-weight: 600;
	}
	.brand {
		margin: 0 0 0.75rem;
		font-size: 0.875rem;
		color: #666;
	}
	.price {
		margin: 0.75rem 0 0;
		font-size: 1.25rem;
		font-weight: 700;
		color: #111;
	}
</style>
