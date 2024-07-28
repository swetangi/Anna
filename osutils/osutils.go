package osutils

import (
	"fmt"
	"syscall"
)

func GetEnvVar(varName string) (string, error) {
	val, found := syscall.Getenv(varName)
	if !found || val == "" {
		err := fmt.Errorf("env var %s not found", varName)
		return "", err
	} else {
		return val, nil
	}
}
