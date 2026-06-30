package domain

import (
	"testing"
	"time"
)

func closedBatch(t *testing.T, design DesignID) Batch {
	t.Helper()
	b := OpenBatch("b1", design, time.Now())
	o, _ := PlaceOrder("o1", design, "c1", 10, time.Now())
	if err := b.Add(o); err != nil {
		t.Fatal(err)
	}
	if err := b.Close(time.Now()); err != nil {
		t.Fatal(err)
	}
	return b
}

func TestDesign_FirstBatchRequiresReview(t *testing.T) {
	// ADR 0007: first batch per design routes to manual review.
	d := NewDesign("d1")
	b := closedBatch(t, "d1")

	decision, err := d.SubmitForPrint(&b)
	if err != nil {
		t.Fatalf("SubmitForPrint: %v", err)
	}
	if !decision.RequiresReview || decision.Approved {
		t.Error("first batch should require review, not auto-approve")
	}
	if b.Approved() {
		t.Error("batch should not be approved before review")
	}
}

func TestDesign_SubsequentBatchesAutoApprove(t *testing.T) {
	// ADR 0007: after the first approval, auto-approval is on by default.
	d := NewDesign("d1")

	first := closedBatch(t, "d1")
	if _, err := d.SubmitForPrint(&first); err != nil {
		t.Fatal(err)
	}
	if err := d.ApproveReview(&first); err != nil {
		t.Fatalf("ApproveReview: %v", err)
	}
	if !d.AutoApproves() {
		t.Fatal("auto-approval should be on after first approval")
	}

	second := closedBatch(t, "d1")
	decision, err := d.SubmitForPrint(&second)
	if err != nil {
		t.Fatal(err)
	}
	if !decision.Approved || decision.RequiresReview {
		t.Error("second batch should auto-approve")
	}
	if !second.Approved() {
		t.Error("second batch should be marked approved")
	}
}

func TestDesign_SubmitOpenBatchFails(t *testing.T) {
	d := NewDesign("d1")
	b := OpenBatch("b1", "d1", time.Now())
	if _, err := d.SubmitForPrint(&b); err != ErrBatchNotClosed {
		t.Errorf("expected ErrBatchNotClosed, got %v", err)
	}
}

func TestDesign_ApproveReviewTwiceFails(t *testing.T) {
	d := NewDesign("d1")
	b := closedBatch(t, "d1")
	_ = d.ApproveReview(&b)
	if err := d.ApproveReview(&b); err != ErrNotAwaitingReview {
		t.Errorf("expected ErrNotAwaitingReview, got %v", err)
	}
}
