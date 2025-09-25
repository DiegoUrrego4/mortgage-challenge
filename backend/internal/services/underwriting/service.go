package underwriting

import (
	"fmt"
	"log"
	"time"

	"github.com/DiegoUrrego4/backend/internal/domain"
	"github.com/DiegoUrrego4/backend/internal/ports"
)

type Service struct {
	repository ports.EvaluationRepository
}

func NewService(repository ports.EvaluationRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Evaluate(input domain.Input) domain.Result {
	dti := input.MonthlyDebts / input.MonthlyIncome
	ltv := input.LoanAmount / input.PropertyValue

	var reasons []string
	var decision string

	switch {
	case dti <= 0.43 && ltv <= 0.80 && input.CreditScore >= 680:
		decision = "Approve"
		reasons = append(reasons, "Applicant meets all primary criteria.")
	case dti <= 0.50 && ltv <= 0.95 && input.CreditScore >= 660:
		decision = "Refer"
		reasons = append(reasons, fmt.Sprintf("Manual review required. DTI: %.2f, LTV: %.2f, FICO: %d", dti, ltv, input.CreditScore))
	default:
		decision = "Decline"
		reasons = append(reasons, "Applicant does not meet the minimum criteria for DTI, LTV, or FICO Score.")
	}

	result := domain.Result{Decision: decision, DTI: dti, LTV: ltv, Reasons: reasons}
	evaluationToSave := domain.Evaluation{
		MonthlyIncome: input.MonthlyIncome,
		MonthlyDebts:  input.MonthlyDebts,
		LoanAmount:    input.LoanAmount,
		PropertyValue: input.PropertyValue,
		CreditScore:   input.CreditScore,
		OccupancyType: input.OccupancyType,
		Decision:      result.Decision,
		DTI:           result.DTI,
		LTV:           result.LTV,
		Reasons:       result.Reasons,
		CreatedAt:     time.Now(),
	}
	err := s.repository.Save(evaluationToSave)
	if err != nil {
		log.Fatalln("Error saving in DB", err)
		return domain.Result{}
	}

	return result
}

func (s *Service) GetAllEvaluations() ([]domain.Evaluation, error) {
	return s.repository.GetAll()
}
