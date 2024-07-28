package config

import (
	"log"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppContext struct {
	Logger       *zap.Logger
	Db           *gorm.DB
	ServerConfig *serverConfig
	JwtConfig    *jwtConfig
}

func NewAppContext(appConfig *AppConfig) *AppContext {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal()
	}

	db := newDb(appConfig.dbConfig)
	return &AppContext{
		Logger:       logger,
		Db:           db,
		ServerConfig: appConfig.serverConfig,
		JwtConfig:    appConfig.jwtConfig,
	}
}
