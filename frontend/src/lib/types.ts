// Frontend domain types, mirroring the glossary in CONTEXT.md.
// These are the customer-facing shapes; the Go backend owns the authoritative model.

/** A size/colour a design can be printed on. */
export type Size = 'S' | 'M' | 'L' | 'XL';

/** A designer or brand whose designs customers can browse and follow. */
export interface Brand {
	id: string;
	name: string;
	handle: string;
	followers: number;
}

/**
 * A Product is a design available to order on merch (t-shirts to start).
 * `unitsOrdered` / `batchThreshold` express the Print Batch progress (ADR 0001):
 * the design prints once it reaches the threshold or its time-box expires.
 */
export interface Product {
	id: string;
	title: string;
	brand: Brand;
	priceCents: number;
	rushFeeCents: number;
	accent: string; // CSS colour used for the placeholder artwork
	emoji: string;
	sizes: Size[];
	unitsOrdered: number;
	batchThreshold: number;
	batchClosesInDays: number; // remaining days in the 14-day time-box
	trending: boolean;
}

/** How an order is fulfilled: batched (standard) or printed immediately (rush, ADR 0003). */
export type OrderKind = 'standard' | 'rush';
