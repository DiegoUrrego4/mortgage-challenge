package ports

import "github.com/DiegoUrrego4/backend/internal/domain"

// EvaluationRepository es el "enchufe" para cualquier tipo de almacenamiento.
type EvaluationRepository interface {
	Save(evaluation domain.Evaluation) error
	GetAll() ([]domain.Evaluation, error)
}
