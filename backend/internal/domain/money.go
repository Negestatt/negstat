package domain

import "errors"

// Money is an amount in the platform's minor currency unit (cents). Using a dedicated
// type — rather than a bare int — keeps monetary values from being mixed up with counts,
// quantities, or other integers.
type Money int64

// ErrNegativeMoney is returned when a negative monetary amount is supplied where only
// non-negative amounts make sense (earnings, payouts).
var ErrNegativeMoney = errors.New("money amount cannot be negative")

// Cents constructs a Money value from a whole number of cents.
func Cents(n int64) Money { return Money(n) }

// Cents returns the underlying amount in cents.
func (m Money) Cents() int64 { return int64(m) }

// Add returns the sum of two amounts.
func (m Money) Add(other Money) Money { return m + other }

// IsZero reports whether the amount is exactly zero.
func (m Money) IsZero() bool { return m == 0 }

// IsPositive reports whether the amount is greater than zero.
func (m Money) IsPositive() bool { return m > 0 }
