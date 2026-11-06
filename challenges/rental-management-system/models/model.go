package models

type Property struct {
	ID          int
	Title       string
	Location    string
	MonthlyRent float64
	IsAvailable bool
}

type Tenant struct {
	ID         int
	Name       string
	Email      string
	PropertyID int
}

type Payment struct {
	ID       int
	TenantID int
	Amount   float64
	Month    string
	IsPaid   bool
}