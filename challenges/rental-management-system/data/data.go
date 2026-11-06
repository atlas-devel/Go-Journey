package data

import "github.com/atlas-devel/Go-Journey/challenges/rental-management-system/models"

var Properties = []models.Property{{

	ID:          1,
	Title:       "Beautiful Apartment",
	Location:    "New York",
	MonthlyRent: 1500.00,
	IsAvailable: true,
},
	{
		ID:          2,
		Title:       "Cozy House",
		Location:    "Los Angeles",
		MonthlyRent: 1200.00,
		IsAvailable: true,
	},

	{
		ID:          3,
		Title:       "Modern Loft",
		Location:    "Chicago",
		MonthlyRent: 1800.00,
		IsAvailable: true,
	}}

var Tenants = []models.Tenant{{
	ID:         1,
	Name:       "John Doe",
	Email:      "john.doe@example.com",
	PropertyID: 1,
},

	{ID: 2,
		Name:       "Jane Smith",
		Email:      "jane.smith@example.com",
		PropertyID: 2,
	}}
	
var Payments = []models.Payment{{
	ID:       1,
	TenantID: 1,
	Amount:   1500.00,
	Month:    "January",
	IsPaid:   true,
},
	{
		ID:       2,
		TenantID: 2,
		Amount:   1200.00,
		Month:    "January",
		IsPaid:   false,
	},
	{
		ID:       3,
		TenantID: 1,
		Amount:   1500.00,
		Month:    "February",
		IsPaid:   false,
	},
}
