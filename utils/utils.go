package utils

import (
	"anna/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppError struct {
	Error  error
	Msg    string
	Status int
}

func RequestHandler[T any](appCtx *config.AppContext, controllerFunc func(c *gin.Context, appCtx *config.AppContext) (*T, *AppError)) func(c *gin.Context) {
	return func(c *gin.Context) {
		logger, err := zap.NewProduction()

		defer logger.Sync()
		if err != nil {
			log.Println("failed to initialize zapp ", zap.Error(err))
		}

		data, errr := controllerFunc(c, appCtx)
		if errr != nil {
			logger.Error(errr.Msg, zap.Error(errr.Error))
			c.JSON(errr.Status, gin.H{"error": errr})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
