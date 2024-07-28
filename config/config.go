package config

type AppConfig struct {
	dbConfig     *dbConfig
	jwtConfig    *jwtConfig
	serverConfig *serverConfig
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		dbConfig:     newDbConfig(),
		jwtConfig:    newJwtConfig(),
		serverConfig: newServerConfig(),
	}
}
