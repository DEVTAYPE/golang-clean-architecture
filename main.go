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
	postRepo := repositories.NewPostRepository(database.DB)

	// iniciar servicios
	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo)

	// inicializar el handler
	userHandler := handlers.NewNewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)

	// iniciar el servidor
	app := server.NewApp()
	app.Get("/health", func(ctx *server.Context) {
		ctx.RWriter.Header().Set("Content-Type", "application/json")
		ctx.Send(`{"status": "ok"}`)
	})

	app.Post("/signup", userHandler.SignUpHandler)
	app.Post("/login", userHandler.LoginHandler)
	app.Get("/me", middleware.AuthMiddleware(userHandler.MeHandler))

	// posts
	app.Post("/posts", middleware.AuthMiddleware(postHandler.CreatePostHandler))
	app.Get("/posts", postHandler.GetAllPostsHandler)
	app.Get("/posts/{id}", postHandler.GetPostByIDHandler)
	app.Get("/users/posts", middleware.AuthMiddleware(postHandler.GetPostsByUserIDHandler))
	app.Put("/posts/{id}", middleware.AuthMiddleware(postHandler.UpdatePostHandler))
	app.Delete("/posts/{id}", middleware.AuthMiddleware(postHandler.DeletePostHandler))

	if err := app.RunServer(config.PORT); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
