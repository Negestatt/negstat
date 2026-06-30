import type { Brand, Product } from './types';

// Mock data reflecting the domain: designers, products with batch progress, trending.
// In production these would come from a backend API.

export const brands: Brand[] = [
	{ id: 'b1', name: 'Nova Collective', handle: '@novacollective', followers: 12340 },
	{ id: 'b2', name: 'Axis Studios', handle: '@axisstudios', followers: 8670 },
	{ id: 'b3', name: 'Bloom Designs', handle: '@bloomdesigns', followers: 5420 },
	{ id: 'b4', name: 'Echo Merch', handle: '@echomerch', followers: 3210 },
];

export const products: Product[] = [
	{
		id: 'p1',
		title: 'Midnight Wave',
		brand: brands[0],
		priceCents: 3500,
		rushFeeCents: 1500,
		accent: '#4A5FCF',
		emoji: '🌊',
		sizes: ['S', 'M', 'L', 'XL'],
		unitsOrdered: 7,
		batchThreshold: 10,
		batchClosesInDays: 9,
		trending: true,
	},
	{
		id: 'p2',
		title: 'Velocity',
		brand: brands[1],
		priceCents: 3200,
		rushFeeCents: 1500,
		accent: '#E94560',
		emoji: '⚡️',
		sizes: ['M', 'L', 'XL'],
		unitsOrdered: 4,
		batchThreshold: 10,
		batchClosesInDays: 12,
		trending: true,
	},
	{
		id: 'p3',
		title: 'Bloom',
		brand: brands[2],
		priceCents: 3800,
		rushFeeCents: 1500,
		accent: '#F7B32B',
		emoji: '🌻',
		sizes: ['S', 'M', 'L'],
		unitsOrdered: 9,
		batchThreshold: 10,
		batchClosesInDays: 2,
		trending: true,
	},
	{
		id: 'p4',
		title: 'Echo Chamber',
		brand: brands[3],
		priceCents: 3400,
		rushFeeCents: 1500,
		accent: '#2EC4B6',
		emoji: '🔊',
		sizes: ['M', 'L', 'XL'],
		unitsOrdered: 2,
		batchThreshold: 10,
		batchClosesInDays: 14,
		trending: false,
	},
	{
		id: 'p5',
		title: 'Fractal Mind',
		brand: brands[0],
		priceCents: 3600,
		rushFeeCents: 1500,
		accent: '#9B51E0',
		emoji: '🧠',
		sizes: ['S', 'M', 'L', 'XL'],
		unitsOrdered: 5,
		batchThreshold: 10,
		batchClosesInDays: 7,
		trending: true,
	},
	{
		id: 'p6',
		title: 'Static',
		brand: brands[1],
		priceCents: 3100,
		rushFeeCents: 1500,
		accent: '#6C757D',
		emoji: '📺',
		sizes: ['M', 'L'],
		unitsOrdered: 1,
		batchThreshold: 10,
		batchClosesInDays: 13,
		trending: false,
	},
];

export function formatCents(cents: number): string {
	return `$${(cents / 100).toFixed(2)}`;
}

export function getProduct(id: string): Product | undefined {
	return products.find((p) => p.id === id);
}
