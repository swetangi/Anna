package routes

import (
	"anna/config"
	"anna/controller/todos"
	"anna/controller/users"
	"anna/middleware"
	"anna/utils"

	"github.com/gin-gonic/gin"
)

func NewRoutes(appCtx *config.AppContext) *gin.Engine {
	userController := users.NewUserController(appCtx)
	todoController := todos.NewTodoController(appCtx)

	router := gin.Default()
	userRoutes := router.Group("/users")
	userRoutes.POST("/register", utils.RequestHandler(appCtx, userController.RegisterUser))
	userRoutes.POST("/login", utils.RequestHandler(appCtx, userController.LoginUser))

	router.Use(middleware.AuthMiddleware(appCtx))
	router.POST("todos/create", utils.RequestHandler(appCtx, todoController.CreateTodo))
	router.PATCH("todos/update/:id",utils.RequestHandler(appCtx,todoController.UpdateTodo))
	router.GET("todos", utils.RequestHandler(appCtx, todoController.GetTodos))
	return router
}
