package config

import (
	"anna/osutils"
	"log"
)

type jwtConfig struct {
	JwtSecretKey string
}

func newJwtConfig() *jwtConfig {
	jwtSecretKey, err := osutils.GetEnvVar("JWT_SECRETKEY")
	if err != nil {
		log.Fatal(err)
	}
	return &jwtConfig{
		JwtSecretKey: jwtSecretKey,
	}
}
