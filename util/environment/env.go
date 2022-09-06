package environment

import (
	"os"
	"strings"
)

type ZeeEnvironment string

const (
	EnvDev  = ZeeEnvironment("DEV")
	EnvProd = ZeeEnvironment("PROD")
)

var ZeeEnv ZeeEnvironment

func Init() {
	envStr := os.Getenv("ZEE_ENV")
	if envStr == "" {
		envStr = string(EnvProd)
	}
	ZeeEnv = ZeeEnvironment(strings.ToUpper(envStr))
}
