package gin

import (
	"github.com/DiegoUrrego4/backend/internal/ports"
	"github.com/gin-gonic/gin"
)

func SetupRouter(underwritingService ports.UnderwritingService) *gin.Engine {
	router := gin.Default()

	handler := NewHandler(underwritingService)

	// Routes
	router.GET("/ping", handler.Ping)
	router.POST("/evaluate", handler.Evaluate)
	router.GET("/evaluations", handler.GetAllEvaluations)

	return router
}
