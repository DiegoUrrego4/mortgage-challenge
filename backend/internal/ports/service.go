package ports

import "github.com/DiegoUrrego4/backend/internal/domain"

type UnderwritingService interface {
	Evaluate(input domain.Input) domain.Result
	GetAllEvaluations() ([]domain.Evaluation, error)
}
