package ports

import "github.com/DiegoUrrego4/backend/internal/domain"

type EvaluationRepository interface {
	Save(evaluation domain.Evaluation) error
	GetAll() ([]domain.Evaluation, error)
}
