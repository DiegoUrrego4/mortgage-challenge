package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	ginAdapter "github.com/DiegoUrrego4/backend/internal/adapters/gin"
	mysqlAdapter "github.com/DiegoUrrego4/backend/internal/adapters/mysql"
	"github.com/DiegoUrrego4/backend/internal/services/underwriting"
)

func main() {
	// 1. Cargar la configuración desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env")
	}

	// 2. Construir la cadena de conexión (DSN) para MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// 3. Crear la primera pieza: el adaptador de la base de datos
	repo, err := mysqlAdapter.NewRepository(dsn)
	if err != nil {
		log.Fatal("FATAL: No se pudo conectar a la base de datos: ", err)
	}

	// 4. Crear la segunda pieza: el servicio, inyectándole el repositorio
	underwritingService := underwriting.NewService(repo)

	// 5. Crear la tercera pieza: el router, inyectándole el servicio
	router := ginAdapter.SetupRouter(underwritingService)

	// 6. Leer el puerto y arrancar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado en http://localhost:%s", port)
	router.Run(":" + port)
}
