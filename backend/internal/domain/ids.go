package domain

// Distinct identifier types prevent accidentally passing, say, a CustomerID where a
// DesignID is expected — the compiler catches the mistake instead of a runtime bug.
type (
	OrderID    string
	DesignID   string
	CustomerID string
	DesignerID string
	BatchID    string
	PayoutID   string
)
