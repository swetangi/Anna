package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
	Email  string `json:"email"`
}

var todos []Todo

func CreateTodo(ctx *gin.Context) {
	var newTodo Todo
	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo.ID = len(todos) + 1
	email, ok := ctx.Get("email")

	if ok {
		newTodo.Email = email.(string)
	}
	todos = append(todos, newTodo)

	ctx.JSON(http.StatusCreated, gin.H{"todos": lo.Filter(todos, func(v Todo, i int) bool {
		return v.Email == email
	})})
}
