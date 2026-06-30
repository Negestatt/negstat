<script lang="ts">
	import { page } from '$app/stores';
	import { getProduct, formatCents } from '$lib/data';
	import type { Size, OrderKind } from '$lib/types';
	import BatchProgress from '$lib/components/BatchProgress.svelte';

	const product = $derived(getProduct($page.params.id));

	// Order state using Svelte 5 runes
	let selectedSize = $state<Size | null>(null);
	let quantity = $state(1);
	let orderKind = $state<OrderKind>('standard');

	const totalCents = $derived(
		product ? product.priceCents * quantity + (orderKind === 'rush' ? product.rushFeeCents : 0) : 0
	);

	function placeOrder() {
		if (!product || !selectedSize) return;
		alert(
			`[DEMO] Placing ${orderKind} order:\n` +
				`${quantity}x ${product.title} (${selectedSize})\n` +
				`Total: ${formatCents(totalCents)}`
		);
	}
</script>

<svelte:head>
	<title>{product ? product.title : 'Product'} — Negestat</title>
</svelte:head>

{#if !product}
	<div class="not-found">
		<h1>Product not found</h1>
		<a href="/">← Back to browse</a>
	</div>
{:else}
	<div class="product-detail">
		<div class="visual">
			<div
				class="artwork"
				style="background: linear-gradient(135deg, {product.accent} 0%, {product.accent}CC 100%);"
			>
				<span class="emoji">{product.emoji}</span>
			</div>
		</div>

		<div class="info">
			<a href="/" class="back">← Browse</a>
			<h1>{product.title}</h1>
			<p class="brand">{product.brand.name}</p>

			<BatchProgress {product} />

			<div class="options">
				<div class="option-group">
					<label>Size</label>
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
					<label>Quantity</label>
					<input type="number" min="1" max="10" bind:value={quantity} />
				</div>

				<div class="option-group">
					<label class="rush-toggle">
						<input type="checkbox" checked={orderKind === 'rush'} onchange={(e) => (orderKind = e.currentTarget.checked ? 'rush' : 'standard')} />
						<span>Rush delivery — skip the batch, get it now (+{formatCents(product.rushFeeCents)})</span>
					</label>
					<p class="rush-note">
						{#if orderKind === 'rush'}
							Your order prints individually and ships within 3 days. (ADR 0003)
						{:else}
							Standard delivery: ships when the batch fills or in {product.batchClosesInDays} days. (ADR 0001)
						{/if}
					</p>
				</div>
			</div>

			<div class="purchase">
				<p class="price">
					{formatCents(product.priceCents)}
					{#if orderKind === 'rush'}
						<span class="rush-fee">+ {formatCents(product.rushFeeCents)} rush</span>
					{/if}
				</p>
				<button class="cta" disabled={!selectedSize} onclick={placeOrder}>
					{#if !selectedSize}
						Select a size
					{:else}
						Place Order — {formatCents(totalCents)}
					{/if}
				</button>
			</div>
		</div>
	</div>
{/if}

<style>
	.not-found {
		text-align: center;
		padding: 4rem 0;
	}
	.not-found a {
		color: #6366F1;
		text-decoration: none;
	}
	.product-detail {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 3rem;
		align-items: start;
	}
	@media (max-width: 768px) {
		.product-detail {
			grid-template-columns: 1fr;
		}
	}
	.visual {
		position: sticky;
		top: 100px;
	}
	.artwork {
		aspect-ratio: 1;
		border-radius: 16px;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 8rem;
	}
	.emoji {
		filter: drop-shadow(0 4px 16px rgba(0, 0, 0, 0.15));
	}
	.info {
		padding: 1rem 0;
	}
	.back {
		display: inline-block;
		color: #6B7280;
		text-decoration: none;
		margin-bottom: 1rem;
		font-weight: 600;
	}
	.back:hover {
		color: #111827;
	}
	h1 {
		margin: 0 0 0.5rem;
		font-size: 2.5rem;
		font-weight: 700;
	}
	.brand {
		margin: 0 0 1.5rem;
		font-size: 1.125rem;
		color: #6B7280;
	}
	.options {
		margin: 2rem 0;
		display: flex;
		flex-direction: column;
		gap: 1.5rem;
	}
	.option-group label {
		display: block;
		font-weight: 600;
		margin-bottom: 0.5rem;
	}
	.size-picker {
		display: flex;
		gap: 0.5rem;
	}
	.size-btn {
		width: 4rem;
		height: 3rem;
		border: 2px solid #E5E7EB;
		background: white;
		border-radius: 8px;
		font-weight: 600;
		cursor: pointer;
		transition: all 0.2s ease;
	}
	.size-btn:hover {
		border-color: #9CA3AF;
	}
	.size-btn.selected {
		border-color: #6366F1;
		background: #6366F1;
		color: white;
	}
	input[type='number'] {
		width: 5rem;
		padding: 0.5rem;
		border: 2px solid #E5E7EB;
		border-radius: 8px;
		font-size: 1rem;
	}
	.rush-toggle {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		cursor: pointer;
	}
	.rush-toggle input {
		width: 1.25rem;
		height: 1.25rem;
		cursor: pointer;
	}
	.rush-note {
		margin: 0.5rem 0 0;
		font-size: 0.875rem;
		color: #6B7280;
		font-style: italic;
	}
	.purchase {
		padding: 1.5rem;
		background: white;
		border-radius: 12px;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
	}
	.price {
		margin: 0 0 1rem;
		font-size: 2rem;
		font-weight: 700;
	}
	.rush-fee {
		font-size: 1rem;
		color: #6B7280;
		font-weight: 400;
	}
	.cta {
		width: 100%;
		padding: 1rem 2rem;
		background: #6366F1;
		color: white;
		border: none;
		border-radius: 8px;
		font-size: 1.125rem;
		font-weight: 700;
		cursor: pointer;
		transition: background 0.2s ease;
	}
	.cta:hover:not(:disabled) {
		background: #4F46E5;
	}
	.cta:disabled {
		background: #9CA3AF;
		cursor: not-allowed;
	}
</style>
