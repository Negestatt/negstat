package domain

import (
	"errors"
	"time"
)

// OrderKind distinguishes a standard, batched order from a Rush Order that bypasses
// batching entirely (ADR 0003).
type OrderKind int

const (
	Standard OrderKind = iota
	Rush
)

var (
	ErrInvalidQuantity = errors.New("order quantity must be positive")
	ErrMissingIdentity = errors.New("order requires design and customer identifiers")
)

// Order is a customer's commitment to purchase a printed product, which the platform
// guarantees to fulfill regardless of how many other customers order the same design
// (CONTEXT.md, ADR 0001). An Order's kind is fixed at creation: a Rush Order can never
// become standard, nor vice versa.
type Order struct {
	id       OrderID
	design   DesignID
	customer CustomerID
	units    int
	kind     OrderKind
	placedAt time.Time
}

// PlaceOrder creates a standard (batched) Order.
func PlaceOrder(id OrderID, design DesignID, customer CustomerID, units int, at time.Time) (Order, error) {
	return newOrder(id, design, customer, units, Standard, at)
}

// PlaceRushOrder creates a Rush Order, which prints individually and immediately,
// bypassing the Print Batch threshold and time-box (ADR 0003).
func PlaceRushOrder(id OrderID, design DesignID, customer CustomerID, units int, at time.Time) (Order, error) {
	return newOrder(id, design, customer, units, Rush, at)
}

func newOrder(id OrderID, design DesignID, customer CustomerID, units int, kind OrderKind, at time.Time) (Order, error) {
	if units <= 0 {
		return Order{}, ErrInvalidQuantity
	}
	if design == "" || customer == "" {
		return Order{}, ErrMissingIdentity
	}
	return Order{
		id:       id,
		design:   design,
		customer: customer,
		units:    units,
		kind:     kind,
		placedAt: at,
	}, nil
}

func (o Order) ID() OrderID          { return o.id }
func (o Order) Design() DesignID     { return o.design }
func (o Order) Customer() CustomerID { return o.customer }
func (o Order) Units() int           { return o.units }
func (o Order) Kind() OrderKind      { return o.kind }
func (o Order) IsRush() bool         { return o.kind == Rush }
func (o Order) PlacedAt() time.Time  { return o.placedAt }
