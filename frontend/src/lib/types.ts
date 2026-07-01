// Frontend domain types, mirroring the glossary in CONTEXT.md.
// These are the customer-facing shapes; product/brand data lives in src/lib/mock/*.json.

/** A size/colour a design can be printed on. */
export type Size = "S" | "M" | "L" | "XL";

/** A designer or brand whose designs customers can browse and follow. */
export interface Brand {
  id: string;
  name: string;
  handle: string;
  followers: number;
  avatarUrl: string;
}

/**
 * A Product is a design available to order on merch (t-shirts to start).
 * Prices are in Ethiopian Birr (ETB), whole units. The Print Batch threshold is an
 * internal production trigger (ADR 0001) and is deliberately NOT part of this shape.
 */
export interface Product {
  id: string;
  title: string;
  brand: Brand;
  priceBirr: number;
  rushFeeBirr: number;
  accent: string; // CSS colour used as the artwork backdrop while the image loads
  imageUrl: string; // t-shirt mockup photo
  sizes: Size[];
  trending: boolean;
}

/** How an order is fulfilled: batched (standard) or printed immediately (rush, ADR 0003). */
export type OrderKind = "standard" | "rush";
