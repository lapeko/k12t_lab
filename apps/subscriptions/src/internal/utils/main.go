package utils

import (
	"fmt"
	"os"
)

type APP_ENV string

const (
	APP_ENV_PROD  APP_ENV = "prod"
	APP_ENV_STAGE APP_ENV = "stage"
	APP_ENV_DEV   APP_ENV = "dev"
)

func ParseAppEnv() (APP_ENV, error) {
	env := APP_ENV(os.Getenv("APP_ENV"))
	switch env {
	case APP_ENV_PROD, APP_ENV_STAGE, APP_ENV_DEV:
		return env, nil
	default:
		return "", fmt.Errorf("unknown environment: %q", env)
	}
}
