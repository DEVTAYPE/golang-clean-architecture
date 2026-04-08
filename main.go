package main

import (
	"api-basico-dev/config"
	"api-basico-dev/database"
	"api-basico-dev/handlers"
	"api-basico-dev/middleware"
	"api-basico-dev/repositories"
	"api-basico-dev/server"
	"api-basico-dev/services"
	"fmt"
	"log"
)

func main() {
	config := config.LoadConfig(".env")

	// establecer conexion
	defer database.Close()
	if err := database.Connect(config.DATABASE_URL); err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		return
	}

	// iniciar repos
	userRepo := repositories.NewUserRepository(database.DB)

	// iniciar servicios
	userService := services.NewUserService(userRepo)

	// inicializar el handler
	handler := handlers.NewNewUserHandler(userService)

	// iniciar el servidor
	app := server.NewApp()
	app.Get("/health", func(ctx *server.Context) {
		ctx.RWriter.Header().Set("Content-Type", "application/json")
		ctx.Send(`{"status": "ok"}`)
	})

	app.Post("/signup", handler.SignUpHandler)
	app.Post("/login", handler.LoginHandler)
	app.Get("/me", middleware.AuthMiddleware(handler.MeHandler))

	if err := app.RunServer(config.PORT); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
