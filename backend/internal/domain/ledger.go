package domain

import (
	"errors"
	"time"
)

var ErrNothingToPayOut = errors.New("no available balance to pay out")

// entryState marks where an earning sits in its lifecycle.
type entryState int

const (
	available      entryState = iota // spendable; counts toward the next Payout
	pendingConsent                   // held until copyright consent is verified (ADR 0005)
	paidOut                          // already disbursed via a Payout
)

type ledgerEntry struct {
	design DesignID
	amount Money
	state  entryState
	at     time.Time
}

// Ledger is a designer's append-only record of earnings. Earnings accrue the moment a
// Batch closes or a Rush Order ships (ADR 0004). Earnings for a design whose copyright
// consent is not yet verified are held as pending until released (ADR 0005). The
// designer's spendable Balance is derived from the ledger rather than stored separately,
// which keeps it auditable: every cent traces back to an entry.
type Ledger struct {
	designer DesignerID
	entries  []ledgerEntry
}

// NewLedger creates an empty ledger for a designer.
func NewLedger(designer DesignerID) Ledger {
	return Ledger{designer: designer}
}

// Accrue records an earning for a design. When consent is not yet verified the amount is
// held pending; otherwise it is immediately available.
func (l *Ledger) Accrue(design DesignID, amount Money, consentVerified bool, at time.Time) error {
	if amount < 0 {
		return ErrNegativeMoney
	}
	state := available
	if !consentVerified {
		state = pendingConsent
	}
	l.entries = append(l.entries, ledgerEntry{design: design, amount: amount, state: state, at: at})
	return nil
}

// ReleaseConsent moves every pending earning for a design into the available balance,
// once consent from the copyright holder has been verified (ADR 0005).
func (l *Ledger) ReleaseConsent(design DesignID, at time.Time) {
	for i := range l.entries {
		if l.entries[i].design == design && l.entries[i].state == pendingConsent {
			l.entries[i].state = available
			l.entries[i].at = at
		}
	}
}

func (l *Ledger) sum(state entryState) Money {
	var total Money
	for _, e := range l.entries {
		if e.state == state {
			total = total.Add(e.amount)
		}
	}
	return total
}

// Available is the designer's spendable Balance — earnings eligible for the next Payout.
func (l *Ledger) Available() Money { return l.sum(available) }

// Pending is the total held awaiting copyright consent.
func (l *Ledger) Pending() Money { return l.sum(pendingConsent) }

// PaidOut is the total already disbursed.
func (l *Ledger) PaidOut() Money { return l.sum(paidOut) }

// DrawPayout disburses the designer's entire available Balance as one Payout, marking
// those earnings paid out. This models the biweekly cadence: each cycle pays out whatever
// has accrued and is available (ADR 0004). It returns an error if nothing is available.
func (l *Ledger) DrawPayout(id PayoutID, at time.Time) (Payout, error) {
	amount := l.Available()
	if !amount.IsPositive() {
		return Payout{}, ErrNothingToPayOut
	}
	for i := range l.entries {
		if l.entries[i].state == available {
			l.entries[i].state = paidOut
			l.entries[i].at = at
		}
	}
	return Payout{id: id, designer: l.designer, amount: amount, createdAt: at}, nil
}

// Payout is a biweekly disbursement of a designer's available Balance to their payment
// method (CONTEXT.md, ADR 0004).
type Payout struct {
	id        PayoutID
	designer  DesignerID
	amount    Money
	createdAt time.Time
}

func (p Payout) ID() PayoutID         { return p.id }
func (p Payout) Designer() DesignerID { return p.designer }
func (p Payout) Amount() Money        { return p.amount }
func (p Payout) CreatedAt() time.Time { return p.createdAt }
