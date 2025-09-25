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
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	repo, err := mysqlAdapter.NewRepository(dsn)
	if err != nil {
		log.Fatal("FATAL: No se pudo conectar a la base de datos: ", err)
	}

	underwritingService := underwriting.NewService(repo)

	router := ginAdapter.SetupRouter(underwritingService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado en http://localhost:%s", port)
	router.Run(":" + port)
}
