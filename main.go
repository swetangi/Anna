package main

import (
	"fmt"
	"net/http"
	"practice/middleware"
	"practice/todos"
	"practice/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	publicRoutes := router.Group("/users")
	publicRoutes.POST("/register", users.RegisterUser)
	publicRoutes.POST("/login", users.LoginUser)

	router.Use(middleware.AuthMiddleware())
	router.POST("todos/create", todos.CreateTodo)

	fmt.Println("Server Listening on port no 8080")
	http.ListenAndServe("localhost:8080", router)
}
