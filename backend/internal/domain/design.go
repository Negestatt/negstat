package domain

import "errors"

var (
	ErrBatchNotClosed    = errors.New("batch must be closed before printing")
	ErrNotAwaitingReview = errors.New("batch is not awaiting manual review")
)

// Design carries the per-design moderation state that governs batch approval. A design's
// first Print Batch requires manual review; once approved, auto-approval turns on by
// default for all of its future batches (ADR 0007).
type Design struct {
	id          DesignID
	autoApprove bool
}

// NewDesign creates a design with auto-approval off, so its first batch routes to review.
func NewDesign(id DesignID) Design {
	return Design{id: id}
}

// PrintDecision is the outcome of submitting a closed batch toward production.
type PrintDecision struct {
	Approved       bool
	RequiresReview bool
}

// SubmitForPrint routes a closed batch toward production. If auto-approval is on it is
// approved immediately; otherwise it is held for manual review.
func (d *Design) SubmitForPrint(b *Batch) (PrintDecision, error) {
	if !b.IsClosed() {
		return PrintDecision{}, ErrBatchNotClosed
	}
	if d.autoApprove {
		b.approved = true
		return PrintDecision{Approved: true}, nil
	}
	return PrintDecision{RequiresReview: true}, nil
}

// ApproveReview approves a batch that was held for manual review and turns on
// auto-approval for this design's future batches (ADR 0007).
func (d *Design) ApproveReview(b *Batch) error {
	if !b.IsClosed() {
		return ErrBatchNotClosed
	}
	if d.autoApprove {
		return ErrNotAwaitingReview
	}
	b.approved = true
	d.autoApprove = true
	return nil
}

// AutoApproves reports whether future batches for this design skip manual review.
func (d *Design) AutoApproves() bool { return d.autoApprove }
