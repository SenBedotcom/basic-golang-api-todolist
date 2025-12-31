package main

import (
	"fmt"
	"log"

	"github.com/SenBedotcom/todo-api/config"
	httpHandler "github.com/SenBedotcom/todo-api/internaldeliveryhttp"
	internalrepository "github.com/SenBedotcom/todo-api/internalrepository"
	internalusecase "github.com/SenBedotcom/todo-api/internalusecase"
	pkgdatabase "github.com/SenBedotcom/todo-api/pkgdatabase"
)

func main() {
	// Load configuration
	cfg := config.Load()
	log.Println("Configuration loaded")

	// Initialize database connection
	dbConfig := pkgdatabase.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	}

	db, err := pkgdatabase.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Database connected successfully")

	// Initialize database schema
	if err := pkgdatabase.InitSchema(db); err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}
	log.Println("Database schema initialized")

	// Initialize repository
	todoRepo := internalrepository.NewPostgresTodoRepository(db)

	// Initialize use case
	todoUseCase := internalusecase.NewTodoUseCase(todoRepo)

	// Initialize HTTP handler
	todoHandler := httpHandler.NewTodoHandler(todoUseCase)

	// Setup router
	router := httpHandler.SetupRouter(todoHandler)

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Printf("Health check: http://localhost:%s/health", cfg.Server.Port)
	log.Printf("API endpoint: http://localhost:%s/api/v1/todos", cfg.Server.Port)

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
