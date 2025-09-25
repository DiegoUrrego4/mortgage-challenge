package underwriting

import (
	"testing"

	"github.com/DiegoUrrego4/backend/internal/domain"
)

type mockRepository struct{}

func (m *mockRepository) Save(eval domain.Evaluation) error {
	return nil
}

func (m *mockRepository) GetAll() ([]domain.Evaluation, error) {
	return nil, nil
}

func TestEvaluate(t *testing.T) {
	testCases := []struct {
		name             string
		input            domain.Input
		expectedDecision string
	}{
		{
			name: "Caso de Aprobación Perfecto",
			input: domain.Input{
				MonthlyIncome: 10000,
				MonthlyDebts:  3000,
				LoanAmount:    200000,
				PropertyValue: 250000,
				CreditScore:   720,
			},
			expectedDecision: "Approve",
		},
		{
			name: "Caso de Referencia por DTI alto",
			input: domain.Input{
				MonthlyIncome: 10000,
				MonthlyDebts:  4500,
				LoanAmount:    200000,
				PropertyValue: 250000,
				CreditScore:   690,
			},
			expectedDecision: "Refer",
		},
		{
			name: "Caso de Referencia por Score bajo",
			input: domain.Input{
				MonthlyIncome: 10000,
				MonthlyDebts:  4000,
				LoanAmount:    200000,
				PropertyValue: 250000,
				CreditScore:   670,
			},
			expectedDecision: "Refer",
		},
		{
			name: "Caso de Rechazo por Score muy bajo",
			input: domain.Input{
				MonthlyIncome: 10000,
				MonthlyDebts:  3000,
				LoanAmount:    200000,
				PropertyValue: 250000,
				CreditScore:   650,
			},
			expectedDecision: "Decline",
		},
		{
			name: "Caso de Rechazo por LTV muy alto",
			input: domain.Input{
				MonthlyIncome: 10000,
				MonthlyDebts:  3000,
				LoanAmount:    490000,
				PropertyValue: 500000,
				CreditScore:   700,
			},
			expectedDecision: "Decline",
		},
	}

	mockRepo := &mockRepository{}
	service := NewService(mockRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := service.Evaluate(tc.input)

			if result.Decision != tc.expectedDecision {
				t.Errorf("Para el caso '%s', se esperaba la decisión '%s', pero se obtuvo '%s'",
					tc.name, tc.expectedDecision, result.Decision)
			}
		})
	}
}
