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
			name: "Perfect approving case",
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
			name: "Reference case for high DTI",
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
			name: "Reference Case for Low Score",
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
			name: "Case of rejection due to very low score",
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
			name: "Case of rejection due to very high LTV",
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
				t.Errorf("For the case ‘%s’, the decision ‘%s’ was expected, but ‘%s’ was obtained.",
					tc.name, tc.expectedDecision, result.Decision)
			}
		})
	}
}
