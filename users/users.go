package users

import (
	"log"
	"net/http"
	"practice/jwt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var users []User

func RegisterUser(context *gin.Context) {
	var newUser User

	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newUser)
	// users := model.User{
	// 	Name:     newUser.Name,
	// 	Password: newUser.Password,
	// 	Email:    newUser.Email,
	// }
	// savedUser, err := users.Save()

	context.JSON(http.StatusCreated, gin.H{"user": users})
}

func LoginUser(context *gin.Context) {
	var logUser User
	if err := context.ShouldBindJSON(&logUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal("Failed to decode body", err)
		return
	}
	email := GetUser(logUser.Email)
	token, err := jwt.GenerateJwt(email)
	if err != nil {
		log.Println("Failed to generate JWT token ", err)
		return
	}
	context.Writer.Header().Set("Content-Type", "application/json")
	context.Writer.WriteHeader(200)
	context.Writer.Write([]byte(token))
}

func GetUser(email string) string {
	for _, v := range users {
		if email == v.Email {
			return email
		}
	}
	return ""

}
