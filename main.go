package main

import (
	"anna/config"
	"anna/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("..........")
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Failed to load env")
	}
	appConfig := config.NewAppConfig()
	fmt.Println("....wsa")
	appCtx := config.NewAppContext(appConfig)

	fmt.Println("Server Listening on port no 8080")
	http.ListenAndServe("localhost:8080", routes.NewRoutes(appCtx))
}
