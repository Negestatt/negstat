package domain

import (
	"testing"
	"time"
)

func TestLedger_AccrueWithConsentIsAvailable(t *testing.T) {
	l := NewLedger("designer-1")
	if err := l.Accrue("d1", Cents(5000), true, time.Now()); err != nil {
		t.Fatal(err)
	}
	if l.Available() != Cents(5000) {
		t.Errorf("expected 5000 available, got %d", l.Available().Cents())
	}
	if !l.Pending().IsZero() {
		t.Error("expected nothing pending")
	}
}

func TestLedger_AccrueWithoutConsentIsPending(t *testing.T) {
	// ADR 0005: copyright consent gates payout, not listing.
	l := NewLedger("designer-1")
	if err := l.Accrue("d1", Cents(5000), false, time.Now()); err != nil {
		t.Fatal(err)
	}
	if !l.Available().IsZero() {
		t.Error("unconsented earnings must not be available")
	}
	if l.Pending() != Cents(5000) {
		t.Errorf("expected 5000 pending, got %d", l.Pending().Cents())
	}
}

func TestLedger_ReleaseConsentMakesAvailable(t *testing.T) {
	l := NewLedger("designer-1")
	_ = l.Accrue("d1", Cents(5000), false, time.Now())
	l.ReleaseConsent("d1", time.Now())
	if l.Available() != Cents(5000) {
		t.Errorf("expected 5000 available after release, got %d", l.Available().Cents())
	}
	if !l.Pending().IsZero() {
		t.Error("expected nothing pending after release")
	}
}

func TestLedger_ReleaseConsentIsPerDesign(t *testing.T) {
	l := NewLedger("designer-1")
	_ = l.Accrue("d1", Cents(3000), false, time.Now())
	_ = l.Accrue("d2", Cents(2000), false, time.Now())
	l.ReleaseConsent("d1", time.Now())

	if l.Available() != Cents(3000) {
		t.Errorf("expected 3000 available, got %d", l.Available().Cents())
	}
	if l.Pending() != Cents(2000) {
		t.Errorf("expected 2000 still pending, got %d", l.Pending().Cents())
	}
}

func TestLedger_DrawPayoutDisbursesAvailable(t *testing.T) {
	// ADR 0004: biweekly payout disburses the accrued, available balance.
	l := NewLedger("designer-1")
	_ = l.Accrue("d1", Cents(4000), true, time.Now())
	_ = l.Accrue("d2", Cents(6000), true, time.Now())

	p, err := l.DrawPayout("payout-1", time.Now())
	if err != nil {
		t.Fatalf("DrawPayout: %v", err)
	}
	if p.Amount() != Cents(10000) {
		t.Errorf("expected payout of 10000, got %d", p.Amount().Cents())
	}
	if !l.Available().IsZero() {
		t.Error("available balance should be zero after payout")
	}
	if l.PaidOut() != Cents(10000) {
		t.Errorf("expected 10000 paid out, got %d", l.PaidOut().Cents())
	}
}

func TestLedger_DrawPayoutExcludesPending(t *testing.T) {
	l := NewLedger("designer-1")
	_ = l.Accrue("d1", Cents(4000), true, time.Now())
	_ = l.Accrue("d2", Cents(6000), false, time.Now()) // pending consent

	p, err := l.DrawPayout("payout-1", time.Now())
	if err != nil {
		t.Fatalf("DrawPayout: %v", err)
	}
	if p.Amount() != Cents(4000) {
		t.Errorf("expected payout of 4000 (pending excluded), got %d", p.Amount().Cents())
	}
	if l.Pending() != Cents(6000) {
		t.Errorf("expected 6000 still pending, got %d", l.Pending().Cents())
	}
}

func TestLedger_DrawPayoutWithNothingAvailableFails(t *testing.T) {
	l := NewLedger("designer-1")
	_ = l.Accrue("d1", Cents(6000), false, time.Now()) // all pending
	if _, err := l.DrawPayout("payout-1", time.Now()); err != ErrNothingToPayOut {
		t.Errorf("expected ErrNothingToPayOut, got %v", err)
	}
}
