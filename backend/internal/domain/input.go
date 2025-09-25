package domain

type Input struct {
	MonthlyIncome float64 `json:"monthly_income"`
	MonthlyDebts  float64 `json:"monthly_debts"`
	LoanAmount    float64 `json:"loan_amount"`
	PropertyValue float64 `json:"property_value"`
	CreditScore   int     `json:"credit_score"`
	OccupancyType string  `json:"occupancy_type"`
}
