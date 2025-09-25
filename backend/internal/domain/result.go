package domain

type Result struct {
	Decision string   `json:"decision"`
	DTI      float64  `json:"dti"`
	LTV      float64  `json:"ltv"`
	Reasons  []string `json:"reasons"`
}
