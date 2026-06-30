# Negestat Backend (opus design)

Go backend for the Negestat print-on-demand marketplace. This is the `opus` branch's
take on the domain model — a second, independent implementation of the same ADR rules
as the `sonnet` branch, for comparison.

## Structure

```
backend/
└── internal/domain/
    ├── ids.go        # Distinct typed identifiers (OrderID, DesignID, ...)
    ├── money.go      # Money value type (cents)
    ├── order.go      # Order with fixed Kind (Standard / Rush)
    ├── batch.go      # Batch with an explicit CloseReason state machine
    ├── design.go     # Design: first-batch review, then auto-approve (ADR 0007)
    ├── ledger.go     # Append-only earnings Ledger + Payout (ADR 0004/0005)
    └── *_test.go
```

## Design choices specific to this branch

- **Strongly-typed IDs** (`OrderID`, `DesignID`, ...) so the compiler catches mixed-up
  arguments.
- **`Money` value type** instead of bare `int` cents.
- **`CloseReason` state machine** on `Batch`: a time-box close *is* the under-fill signal
  (ADR 0002), so under-fill isn't recomputed from counts.
- **Derived `Units()`** from the batch's orders — a single source of truth, no
  denormalized counter to keep in sync.
- **Append-only `Ledger`** instead of mutable balance fields: available / pending /
  paid-out are all derived, and copyright-consent gating (ADR 0005) is just an entry
  state transition. Every cent traces back to an entry.

## ADR rules encoded

| Rule | ADR | Where |
| --- | --- | --- |
| Batch closes at 10 units or 14 days | 0001 | `batch.go` `reasonAt` |
| Platform absorbs under-fill cost | 0002 | `batch.go` `Underfilled` |
| Rush orders bypass batching | 0003 | `order.go`, `batch.go` `Add` |
| Biweekly payout of accrued balance | 0004 | `ledger.go` `DrawPayout` |
| Copyright consent gates payout | 0005 | `ledger.go` `Accrue`/`ReleaseConsent` |
| First batch reviewed, then auto-approve | 0007 | `design.go` |

## Running tests

```bash
cd backend
go test ./internal/domain -v
```
