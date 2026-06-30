<script lang="ts">
	import type { Product } from '../types';

	let { product, compact = false }: { product: Product; compact?: boolean } = $props();

	const percentage = $derived((product.unitsOrdered / product.batchThreshold) * 100);
	const remaining = $derived(product.batchThreshold - product.unitsOrdered);
</script>

<div class="batch-progress" class:compact>
	<div class="bar">
		<div class="fill" style="width: {percentage}%"></div>
	</div>
	<p class="label">
		{#if remaining === 0}
			<span class="ready">✓ Batch ready to print!</span>
		{:else if compact}
			{product.unitsOrdered} / {product.batchThreshold} ordered
		{:else}
			<strong>{remaining}</strong>
			{remaining === 1 ? 'order' : 'orders'} until batch prints, or ships in
			<strong>{product.batchClosesInDays} days</strong>
		{/if}
	</p>
</div>

<style>
	.batch-progress {
		margin: 0.5rem 0;
	}
	.batch-progress.compact {
		margin: 0;
	}
	.bar {
		height: 6px;
		background: #E5E7EB;
		border-radius: 4px;
		overflow: hidden;
	}
	.fill {
		height: 100%;
		background: linear-gradient(90deg, #10B981 0%, #059669 100%);
		transition: width 0.3s ease;
	}
	.label {
		margin: 0.25rem 0 0;
		font-size: 0.75rem;
		color: #666;
	}
	.compact .label {
		font-size: 0.7rem;
	}
	.ready {
		color: #10B981;
		font-weight: 600;
	}
</style>
