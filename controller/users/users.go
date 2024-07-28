package users

import (
	"anna/config"
	"anna/jwt"
	usersModel "anna/models/users"
	"anna/repo/userrepo"
	"anna/utils"
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	useerepo *userrepo.UserRepo
	appCtx   *config.AppContext
}

func NewUserController(appCtx *config.AppContext) *UserController {
	return &UserController{
		useerepo: userrepo.NewUserRepo(appCtx.Db),
		appCtx:   appCtx,
	}
}

func (userCtrl *UserController) RegisterUser(ctx *gin.Context, appCtx *config.AppContext) (*usersModel.User, *utils.AppError) {
	var newUser usersModel.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {

		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Error in Binding json user",
			Status: http.StatusBadRequest,
		}
	}

	if err := validateUser(newUser); err != nil {
		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Failed to validate user",
			Status: http.StatusBadRequest,
		}
	}

	passwordBytes := []byte(newUser.Password)
	hashPassword := md5.Sum(passwordBytes)
	hashPasswordStr := hex.EncodeToString(hashPassword[:])

	if err := userCtrl.useerepo.RegisterUser(newUser.Name, newUser.Email, hashPasswordStr); err != nil {
		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Failed to register user",
			Status: http.StatusInternalServerError,
		}
	}
	return &newUser, nil
}

func validateUser(user usersModel.User) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil
}

func (userCtrl *UserController) LoginUser(ctx *gin.Context, appCtx *config.AppContext) (*string, *utils.AppError) {
	var logUser usersModel.User

	if err := ctx.ShouldBindJSON(&logUser); err != nil {
		return nil, &utils.AppError{
			Error:  err,
			Msg:    "Error in binding json loguser",
			Status: http.StatusBadRequest,
		}
	}
	isValid, errr := userCtrl.useerepo.LoginUser(logUser.Name, logUser.Email, logUser.Password)
	if !isValid {
		return nil, &utils.AppError{
			Error:  errr,
			Msg:    "Invalid Credentials",
			Status: http.StatusUnauthorized,
		}
	}
	if errr != nil {
		return nil, &utils.AppError{
			Error:  errr,
			Msg:    "Error in Login user",
			Status: http.StatusBadRequest,
		}
	}
	token, err := jwt.GenerateJwt(logUser.Email, appCtx.JwtConfig.JwtSecretKey)

	if err != nil {
		return nil, &utils.AppError{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}
	return &token, nil
}
