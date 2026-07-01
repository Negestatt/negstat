import type { Brand, Product } from "./types";
import productsData from "./mock/products.json";
import brandsData from "./mock/brands.json";

// Data is bundled as static JSON so the app deploys as a fully static site
// (GitHub Pages) with no backend. The Go service in ../backend owns the
// authoritative domain model and is exercised by its own tests.

export async function fetchProducts(): Promise<Product[]> {
  return productsData as Product[];
}

export async function fetchBrands(): Promise<Brand[]> {
  return brandsData as Brand[];
}

export interface PlaceOrderRequest {
  productId: string;
  size: string;
  quantity: number;
  isRush: boolean;
}

export interface PlaceOrderResponse {
  orderId: string;
  message: string;
}

// Client-side order placement mirroring the backend's domain validation
// (quantity must be positive). On a static deploy there is no server to call.
export async function placeOrder(req: PlaceOrderRequest): Promise<PlaceOrderResponse> {
  if (req.quantity <= 0) {
    throw new Error("order quantity must be positive");
  }
  const kind = req.isRush ? "rush" : "standard";
  const orderId = `order-${Date.now()}`;
  return {
    orderId,
    message: `Order placed successfully! (${kind} delivery)`,
  };
}
