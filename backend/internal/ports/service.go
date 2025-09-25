package ports

import "github.com/DiegoUrrego4/backend/internal/domain"

// UnderwritingService es la INTERFAZ. Es el contrato.
// Define que CUALQUIER COSA que quiera ser un servicio de underwriting
// DEBE tener un método llamado Evaluate que reciba un Input y devuelva un Result.
type UnderwritingService interface {
	Evaluate(input domain.Input) domain.Result
	GetAllEvaluations() ([]domain.Evaluation, error)
}
