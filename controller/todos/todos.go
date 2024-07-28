package todos

import (
	"anna/config"
	todosModel "anna/models/todos"
	"anna/repo/todorepo"
	"anna/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todorepo *todorepo.TodoRepo
	appCtx   *config.AppContext
}

func NewTodoController(appCtx *config.AppContext) *TodoController {
	todoRepo := todorepo.NewTodoRepo(appCtx.Db)
	return &TodoController{
		todorepo: todoRepo,
		appCtx:   appCtx,
	}
}

func (todoCtrl *TodoController) CreateTodo(ctx *gin.Context, appCtx *config.AppContext) (*todorepo.Todo, *utils.AppError) {
	userEmail := ctx.GetString("email")

	var newTodo todosModel.Todo
	if err := ctx.ShouldBindJSON(&newTodo); err != nil {

		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Error in binding json",
			Status: http.StatusInternalServerError,
		}
	}
	todo, err := todoCtrl.todorepo.CreateTodo(&newTodo, userEmail)
	if err != nil {
		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Error in create todo",
			Status: http.StatusInternalServerError,
		}

	}

	return todo, nil
}

func (todoCtrl *TodoController) GetTodos(ctx *gin.Context, appCtx *config.AppContext) (*[]todorepo.Todo, *utils.AppError) {
	userEmail := ctx.GetString("email")

	todosList, err := todoCtrl.todorepo.GetTodos(userEmail)
	if err != nil {
		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Error in getting todos",
			Status: http.StatusInternalServerError,
		}
	}

	return &todosList, nil
}

func (todoCtrl *TodoController) UpdateTodo(ctx *gin.Context, appCtx *config.AppContext) (*todorepo.Todo, *utils.AppError) {
	todoIdStr := ctx.Param("id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {

		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Invalid todo id",
			Status: http.StatusBadRequest,
		}
	}
	var updateTodo todorepo.Todo
	if err := ctx.ShouldBindJSON(&updateTodo); err != nil {

		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Error in getting todos",
			Status: http.StatusBadRequest,
		}
	}
	todo, err := todoCtrl.todorepo.UpdateTodo(todoId, updateTodo)
	if err != nil {
		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Error in updating Todo",
			Status: http.StatusInternalServerError,
		}
	}
	return todo, nil
}

func (todoCtrl *TodoController) DeleteTodo(ctx *gin.Context, appCtx *config.AppContext) (*todorepo.Todo, *utils.AppError) {
	todoIdStr := ctx.Param("id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Invalid todo id",
			Status: http.StatusBadRequest,
		}
	}
	if err := todoCtrl.todorepo.DeleteTodo(todoId); err != nil {

		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Error in deleting todo",
			Status: http.StatusInternalServerError,
		}
	}
	return nil, nil
}
