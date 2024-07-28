package config

import (
	"anna/osutils"
	"log"
	"strconv"
)

type serverConfig struct {
	Port   int
	IsProd bool
}

func newServerConfig() *serverConfig {
	serverPort, err := osutils.GetEnvVar("SERVER_PORT")
	if err != nil {
		log.Fatal(err)
	}
	portInt, err := strconv.ParseInt(serverPort, 0, 32)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	env, err := osutils.GetEnvVar("ENV")
	if err != nil {
		log.Fatal(err)
	}

	return &serverConfig{
		Port:   int(portInt),
		IsProd: env == "prod",
	}
}
