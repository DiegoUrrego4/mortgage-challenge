package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DiegoUrrego4/backend/internal/domain"
	"github.com/DiegoUrrego4/backend/internal/ports"
)

type Handler struct {
	underwritingService ports.UnderwritingService
}

func NewHandler(underwritingService ports.UnderwritingService) *Handler {
	return &Handler{
		underwritingService: underwritingService,
	}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (h *Handler) Evaluate(c *gin.Context) {
	var input domain.Input

	if err := c.ShouldBindJSON(&input); err != nil {
		return
	}

	result := h.underwritingService.Evaluate(input)

	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetAllEvaluations(c *gin.Context) {
	evaluations, err := h.underwritingService.GetAllEvaluations()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Evaluations cannot be founded"})
		return
	}

	c.JSON(http.StatusOK, evaluations)
}
