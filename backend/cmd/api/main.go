package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"negestat/internal/domain"
)

// API response types matching frontend expectations
type BrandResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Handle    string `json:"handle"`
	Followers int    `json:"followers"`
	AvatarURL string `json:"avatarUrl"`
}

// ProductResponse is the customer-facing product shape. It deliberately omits the
// internal Print Batch threshold and progress: per ADR 0001 the 10-unit threshold is a
// production-batching trigger, not a customer-facing pre-order condition, so it must not
// leak to the client.
type ProductResponse struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	Brand        BrandResponse `json:"brand"`
	PriceCents   int           `json:"priceCents"`
	RushFeeCents int           `json:"rushFeeCents"`
	Accent       string        `json:"accent"`
	ImageURL     string        `json:"imageUrl"`
	Sizes        []string      `json:"sizes"`
	Trending     bool          `json:"trending"`
}

type OrderRequest struct {
	ProductID string `json:"productId"`
	Size      string `json:"size"`
	Quantity  int    `json:"quantity"`
	IsRush    bool   `json:"isRush"`
}

type OrderResponse struct {
	OrderID string `json:"orderId"`
	Message string `json:"message"`
}

// Mock data store (in production this would be a database)
var (
	brands = []BrandResponse{
		{ID: "b1", Name: "Nova Collective", Handle: "@novacollective", Followers: 12340,
			AvatarURL: "https://images.unsplash.com/photo-1531123897727-8f129e1688ce?w=200&h=200&q=80&auto=format&fit=crop"},
		{ID: "b2", Name: "Axis Studios", Handle: "@axisstudios", Followers: 8670,
			AvatarURL: "https://images.unsplash.com/photo-1506277886164-e25aa3f4ef7f?w=200&h=200&q=80&auto=format&fit=crop"},
		{ID: "b3", Name: "Bloom Designs", Handle: "@bloomdesigns", Followers: 5420,
			AvatarURL: "https://images.unsplash.com/photo-1519699047748-de8e457a634e?w=200&h=200&q=80&auto=format&fit=crop"},
		{ID: "b4", Name: "Echo Merch", Handle: "@echomerch", Followers: 3210,
			AvatarURL: "https://images.unsplash.com/photo-1463453091185-61582044d556?w=200&h=200&q=80&auto=format&fit=crop"},
	}

	products = []ProductResponse{
		// Nova Collective (b1)
		{
			ID: "p1", Title: "Midnight Wave", Brand: brands[0],
			PriceCents: 3500, RushFeeCents: 1500, Accent: "#4A5FCF",
			ImageURL: "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"S", "M", "L", "XL"}, Trending: true,
		},
		{
			ID: "p2", Title: "Fractal Mind", Brand: brands[0],
			PriceCents: 3600, RushFeeCents: 1500, Accent: "#9B51E0",
			ImageURL: "https://images.unsplash.com/photo-1562157873-818bc0726f68?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"S", "M", "L", "XL"}, Trending: true,
		},
		{
			ID: "p3", Title: "Aurora", Brand: brands[0],
			PriceCents: 3400, RushFeeCents: 1500, Accent: "#6366F1",
			ImageURL: "https://images.unsplash.com/photo-1503342217505-b0a15ec3261c?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"S", "M", "L", "XL"}, Trending: true,
		},
		{
			ID: "p4", Title: "Nebula", Brand: brands[0],
			PriceCents: 3700, RushFeeCents: 1500, Accent: "#4338CA",
			ImageURL: "https://images.unsplash.com/photo-1554568218-0f1715e72254?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"M", "L", "XL"}, Trending: false,
		},

		// Axis Studios (b2)
		{
			ID: "p5", Title: "Velocity", Brand: brands[1],
			PriceCents: 3200, RushFeeCents: 1500, Accent: "#E94560",
			ImageURL: "https://images.unsplash.com/photo-1576566588028-4147f3842f27?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"M", "L", "XL"}, Trending: true,
		},
		{
			ID: "p6", Title: "Static", Brand: brands[1],
			PriceCents: 3100, RushFeeCents: 1500, Accent: "#6C757D",
			ImageURL: "https://images.unsplash.com/photo-1581655353564-df123a1eb820?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"M", "L"}, Trending: false,
		},
		{
			ID: "p7", Title: "Redline", Brand: brands[1],
			PriceCents: 3300, RushFeeCents: 1500, Accent: "#DC2626",
			ImageURL: "https://images.unsplash.com/photo-1583744946564-b52ac1c389c8?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"S", "M", "L", "XL"}, Trending: true,
		},
		{
			ID: "p8", Title: "Apex", Brand: brands[1],
			PriceCents: 3500, RushFeeCents: 1500, Accent: "#F97316",
			ImageURL: "https://images.unsplash.com/photo-1618517351616-38fb9c5210c6?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"M", "L", "XL"}, Trending: false,
		},

		// Bloom Designs (b3)
		{
			ID: "p9", Title: "Bloom", Brand: brands[2],
			PriceCents: 3800, RushFeeCents: 1500, Accent: "#F7B32B",
			ImageURL: "https://images.unsplash.com/photo-1583743814966-8936f5b7be1a?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"S", "M", "L"}, Trending: true,
		},
		{
			ID: "p10", Title: "Wildflower", Brand: brands[2],
			PriceCents: 3600, RushFeeCents: 1500, Accent: "#10B981",
			ImageURL: "https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"S", "M", "L", "XL"}, Trending: true,
		},
		{
			ID: "p11", Title: "Sunbeam", Brand: brands[2],
			PriceCents: 3400, RushFeeCents: 1500, Accent: "#FACC15",
			ImageURL: "https://images.unsplash.com/photo-1556821840-3a63f95609a7?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"M", "L"}, Trending: false,
		},
		{
			ID: "p12", Title: "Meadow", Brand: brands[2],
			PriceCents: 3500, RushFeeCents: 1500, Accent: "#22C55E",
			ImageURL: "https://images.unsplash.com/photo-1523275335684-37898b6baf30?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"S", "M", "L", "XL"}, Trending: false,
		},

		// Echo Merch (b4)
		{
			ID: "p13", Title: "Echo Chamber", Brand: brands[3],
			PriceCents: 3400, RushFeeCents: 1500, Accent: "#2EC4B6",
			ImageURL: "https://images.unsplash.com/photo-1618354691373-d851c5c3a990?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"M", "L", "XL"}, Trending: true,
		},
		{
			ID: "p14", Title: "Frequency", Brand: brands[3],
			PriceCents: 3300, RushFeeCents: 1500, Accent: "#06B6D4",
			ImageURL: "https://images.unsplash.com/photo-1620799140408-edc6dcb6d633?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"S", "M", "L"}, Trending: false,
		},
		{
			ID: "p15", Title: "Resonance", Brand: brands[3],
			PriceCents: 3600, RushFeeCents: 1500, Accent: "#0EA5E9",
			ImageURL: "https://images.unsplash.com/photo-1503341504253-dff4815485f1?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"S", "M", "L", "XL"}, Trending: true,
		},
		{
			ID: "p16", Title: "Soundwave", Brand: brands[3],
			PriceCents: 3200, RushFeeCents: 1500, Accent: "#14B8A6",
			ImageURL: "https://images.unsplash.com/photo-1571945153237-4929e783af4a?w=800&q=80&auto=format&fit=crop",
			Sizes:    []string{"M", "L", "XL"}, Trending: false,
		},
	}
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func handleBrands(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brands)
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use the real domain model to create an order
	var order domain.Order
	var err error

	if req.IsRush {
		order, err = domain.PlaceRushOrder(
			domain.OrderID("order-"+time.Now().Format("20060102150405")),
			domain.DesignID(req.ProductID),
			domain.CustomerID("customer-demo"),
			req.Quantity,
			time.Now(),
		)
	} else {
		order, err = domain.PlaceOrder(
			domain.OrderID("order-"+time.Now().Format("20060102150405")),
			domain.DesignID(req.ProductID),
			domain.CustomerID("customer-demo"),
			req.Quantity,
			time.Now(),
		)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kindStr := "standard"
	if order.IsRush() {
		kindStr = "rush"
	}

	resp := OrderResponse{
		OrderID: string(order.ID()),
		Message: "Order placed successfully! (" + kindStr + " delivery)",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/products", handleProducts)
	http.HandleFunc("/api/brands", handleBrands)
	http.HandleFunc("/api/orders", handleOrder)

	log.Println("🚀 API server running on http://localhost:8080")
	log.Println("   GET  /api/products")
	log.Println("   GET  /api/brands")
	log.Println("   POST /api/orders")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
