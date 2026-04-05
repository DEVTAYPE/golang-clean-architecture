package main

import (
	"api-basico-dev/config"
	"api-basico-dev/database"
	"fmt"
)

func main() {
	config := config.LoadConfig(".env")

	// establecer conexion
	defer database.Close()
	if err := database.Connect(config.DATABASE_URL); err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		return
	}

	fmt.Println("Hello, World!")
}
