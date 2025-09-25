package gin

import (
	"github.com/DiegoUrrego4/backend/internal/ports"
	"github.com/gin-gonic/gin"
)

// SetupRouter es el constructor de nuestro router. Recibe sus dependencias
// (en este caso, el servicio de underwriting) y devuelve el motor de Gin configurado.
func SetupRouter(underwritingService ports.UnderwritingService) *gin.Engine {
	// Creamos el motor de Gin con el logger y recovery por defecto.
	router := gin.Default()

	// Creamos una instancia de nuestro handler, inyectándole el servicio.
	handler := NewHandler(underwritingService)

	// Definimos la ruta y la asociamos con el método correspondiente del handler.
	router.GET("/ping", handler.Ping)
	router.POST("/evaluate", handler.Evaluate)
	router.GET("/evaluations", handler.GetAllEvaluations)

	// Aquí podrías añadir más rutas en el futuro:
	// router.GET("/evaluations/:id", handler.GetEvaluation)

	return router
}
