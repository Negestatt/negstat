package domain

import (
	"testing"
	"time"
)

func standardOrder(t *testing.T, id OrderID, design DesignID, units int) Order {
	t.Helper()
	o, err := PlaceOrder(id, design, "customer-1", units, time.Now())
	if err != nil {
		t.Fatalf("PlaceOrder(%s): %v", id, err)
	}
	return o
}

func TestPlaceOrder_RejectsNonPositiveQuantity(t *testing.T) {
	if _, err := PlaceOrder("o1", "d1", "c1", 0, time.Now()); err != ErrInvalidQuantity {
		t.Errorf("expected ErrInvalidQuantity, got %v", err)
	}
}

func TestPlaceRushOrder_IsRush(t *testing.T) {
	o, err := PlaceRushOrder("o1", "d1", "c1", 1, time.Now())
	if err != nil {
		t.Fatalf("PlaceRushOrder: %v", err)
	}
	if !o.IsRush() || o.Kind() != Rush {
		t.Error("expected order to be a rush order")
	}
}

func TestBatch_Add_RejectsRushOrders(t *testing.T) {
	// ADR 0003: rush orders bypass batching.
	b := OpenBatch("b1", "d1", time.Now())
	rush, _ := PlaceRushOrder("o1", "d1", "c1", 1, time.Now())
	if err := b.Add(rush); err != ErrRushNotBatched {
		t.Errorf("expected ErrRushNotBatched, got %v", err)
	}
}

func TestBatch_Add_RejectsForeignDesign(t *testing.T) {
	b := OpenBatch("b1", "d1", time.Now())
	if err := b.Add(standardOrder(t, "o1", "d2", 1)); err != ErrDesignMismatch {
		t.Errorf("expected ErrDesignMismatch, got %v", err)
	}
}

func TestBatch_ClosesByThreshold(t *testing.T) {
	// ADR 0001: closes at 10 units.
	b := OpenBatch("b1", "d1", time.Now())
	if err := b.Add(standardOrder(t, "o1", "d1", 10)); err != nil {
		t.Fatal(err)
	}
	if !b.ReadyToClose(time.Now()) {
		t.Fatal("batch should be ready to close at 10 units")
	}
	if err := b.Close(time.Now()); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if b.Reason() != ClosedByThreshold {
		t.Errorf("expected ClosedByThreshold, got %v", b.Reason())
	}
	if b.Underfilled() {
		t.Error("a threshold close is never under-filled")
	}
}

func TestBatch_ClosesByTimeBox_AndIsUnderfilled(t *testing.T) {
	// ADR 0001 + 0002: closes at 14 days; under-filled cost absorbed by platform.
	opened := time.Now().Add(-15 * 24 * time.Hour)
	b := OpenBatch("b1", "d1", opened)
	if err := b.Add(standardOrder(t, "o1", "d1", 4)); err != nil {
		t.Fatal(err)
	}
	if !b.ReadyToClose(time.Now()) {
		t.Fatal("batch should be ready to close after 14 days")
	}
	if err := b.Close(time.Now()); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if b.Reason() != ClosedByTimeBox {
		t.Errorf("expected ClosedByTimeBox, got %v", b.Reason())
	}
	if !b.Underfilled() {
		t.Error("a time-box close below threshold should be under-filled")
	}
}

func TestBatch_DoesNotCloseEarly(t *testing.T) {
	b := OpenBatch("b1", "d1", time.Now())
	if err := b.Add(standardOrder(t, "o1", "d1", 5)); err != nil {
		t.Fatal(err)
	}
	if b.ReadyToClose(time.Now()) {
		t.Error("batch with 5 units inside the time-box should stay open")
	}
	if err := b.Close(time.Now()); err != ErrNotReady {
		t.Errorf("expected ErrNotReady, got %v", err)
	}
}

func TestBatch_UnitsCountByQuantity(t *testing.T) {
	// Threshold counts units, not distinct orders.
	b := OpenBatch("b1", "d1", time.Now())
	_ = b.Add(standardOrder(t, "o1", "d1", 7))
	_ = b.Add(standardOrder(t, "o2", "d1", 3))
	if b.Units() != 10 {
		t.Errorf("expected 10 units, got %d", b.Units())
	}
	if !b.ReadyToClose(time.Now()) {
		t.Error("10 units across two orders should close the batch")
	}
}

func TestBatch_AddAfterCloseFails(t *testing.T) {
	b := OpenBatch("b1", "d1", time.Now())
	_ = b.Add(standardOrder(t, "o1", "d1", 10))
	_ = b.Close(time.Now())
	if err := b.Add(standardOrder(t, "o2", "d1", 1)); err != ErrBatchClosed {
		t.Errorf("expected ErrBatchClosed, got %v", err)
	}
}
