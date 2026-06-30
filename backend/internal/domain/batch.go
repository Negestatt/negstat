package domain

import (
	"errors"
	"time"
)

const (
	// BatchThreshold is the number of ordered units that closes a batch immediately (ADR 0001).
	BatchThreshold = 10

	// BatchTimeBox bounds how long a batch can stay open before it closes regardless of
	// unit count, keeping the fulfillment guarantee honest (ADR 0001).
	BatchTimeBox = 14 * 24 * time.Hour
)

// CloseReason records why (and whether) a Batch closed. It directly encodes whether the
// batch was under-filled: a time-box close always means the threshold was not reached.
type CloseReason int

const (
	NotClosed CloseReason = iota
	ClosedByThreshold
	ClosedByTimeBox
)

var (
	ErrBatchClosed    = errors.New("batch is already closed")
	ErrRushNotBatched = errors.New("rush orders bypass batching")
	ErrDesignMismatch = errors.New("order belongs to a different design")
	ErrNotReady       = errors.New("batch has not met its closing condition")
)

// Batch groups standard Orders for a single design that print together. It closes once
// it reaches BatchThreshold units or its time-box elapses, whichever comes first.
// Rush Orders never belong to a Batch (ADR 0003).
type Batch struct {
	id       BatchID
	design   DesignID
	orders   []Order
	openedAt time.Time
	reason   CloseReason
	closedAt time.Time
	approved bool
}

// OpenBatch starts a new, empty Batch for a design.
func OpenBatch(id BatchID, design DesignID, at time.Time) Batch {
	return Batch{id: id, design: design, openedAt: at, reason: NotClosed}
}

// Add appends a standard Order to the batch. The time-box is measured from the batch's
// open time, so adding orders never extends the deadline.
func (b *Batch) Add(o Order) error {
	if b.reason != NotClosed {
		return ErrBatchClosed
	}
	if o.IsRush() {
		return ErrRushNotBatched
	}
	if o.Design() != b.design {
		return ErrDesignMismatch
	}
	b.orders = append(b.orders, o)
	return nil
}

// Units returns the total ordered units in the batch, derived from its orders rather
// than tracked separately — there is a single source of truth.
func (b *Batch) Units() int {
	total := 0
	for _, o := range b.orders {
		total += o.Units()
	}
	return total
}

// reasonAt computes why the batch would close as of now, or NotClosed. The threshold is
// checked first, so ClosedByTimeBox always implies the threshold was not met.
func (b *Batch) reasonAt(now time.Time) CloseReason {
	if b.Units() >= BatchThreshold {
		return ClosedByThreshold
	}
	if now.Sub(b.openedAt) >= BatchTimeBox {
		return ClosedByTimeBox
	}
	return NotClosed
}

// ReadyToClose reports whether the batch has met a closing condition as of now.
func (b *Batch) ReadyToClose(now time.Time) bool {
	return b.reasonAt(now) != NotClosed
}

// Close finalizes the batch, recording why it closed. It is an error to close a batch
// that has not met a closing condition, or one that is already closed.
func (b *Batch) Close(now time.Time) error {
	if b.reason != NotClosed {
		return ErrBatchClosed
	}
	reason := b.reasonAt(now)
	if reason == NotClosed {
		return ErrNotReady
	}
	b.reason = reason
	b.closedAt = now
	return nil
}

// Underfilled reports whether the batch closed via its time-box without reaching the
// threshold. The platform absorbs the extra per-unit cost for these batches (ADR 0002).
func (b *Batch) Underfilled() bool {
	return b.reason == ClosedByTimeBox
}

func (b *Batch) ID() BatchID         { return b.id }
func (b *Batch) Design() DesignID    { return b.design }
func (b *Batch) IsClosed() bool      { return b.reason != NotClosed }
func (b *Batch) Reason() CloseReason { return b.reason }
func (b *Batch) Approved() bool      { return b.approved }
func (b *Batch) ClosedAt() time.Time { return b.closedAt }
