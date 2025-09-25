package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DiegoUrrego4/backend/internal/domain"
	"github.com/DiegoUrrego4/backend/internal/ports"
)

// Handler es una struct que contiene las dependencias para nuestros handlers,
// en este caso, el servicio de underwriting.
type Handler struct {
	underwritingService ports.UnderwritingService
}

// NewHandler es el constructor para nuestro Handler.
func NewHandler(underwritingService ports.UnderwritingService) *Handler {
	return &Handler{
		underwritingService: underwritingService,
	}
}

// Ping maneja la petición a la ruta GET /evaluations.
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Evaluate es el método que maneja la petición a la ruta /evaluate.
// Es la misma lógica que teníamos antes, pero ahora es un método del Handler.
func (h *Handler) Evaluate(c *gin.Context) {
	var input domain.Input

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("--- ❌ ERROR: Falló el ShouldBindJSON:", err)
		return
	}

	fmt.Println("--- ▶️ DATOS RECIBIDOS DEL FRONTEND:", input)
	// Llama a la lógica de negocio a través de la interfaz.
	result := h.underwritingService.Evaluate(input)
	fmt.Println("--- ◀️ RESULTADO CALCULADO POR EL SERVICIO:", result)

	c.JSON(http.StatusOK, result)
	fmt.Println("--- ✅ FINALIZANDO HANDLER EVALUATE (JSON Enviado) ---")
}

// GetAllEvaluations maneja la petición a la ruta GET /evaluations.
func (h *Handler) GetAllEvaluations(c *gin.Context) {
	// 1. Llama al servicio para obtener la lista de evaluaciones.
	evaluations, err := h.underwritingService.GetAllEvaluations()

	// 2. Si hay un error (ej. la base de datos se cayó), devuelve un error 500.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las evaluaciones"})
		return
	}

	// 3. Si todo sale bien, devuelve la lista con un estado 200 OK.
	c.JSON(http.StatusOK, evaluations)
}
