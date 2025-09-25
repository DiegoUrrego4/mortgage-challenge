package domain

import "time"

// Evaluation representa un registro completo de una decisión.
type Evaluation struct {
	ID            int       `json:"id"`
	MonthlyIncome float64   `json:"monthly_income"`
	MonthlyDebts  float64   `json:"monthly_debts"`
	LoanAmount    float64   `json:"loan_amount"`
	PropertyValue float64   `json:"property_value"`
	CreditScore   int       `json:"credit_score"`
	OccupancyType string    `json:"occupancy_type"`
	Decision      string    `json:"decision"`
	DTI           float64   `json:"dti"`
	LTV           float64   `json:"ltv"`
	Reasons       []string  `json:"reasons"`
	CreatedAt     time.Time `json:"created_at"`
}
