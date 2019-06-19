package util

import (
	"fmt"
	"os"
)

//SystemEnvNamePrefix system env name prefix
var SystemEnvNamePrefix = "HerbGo."

//EnvForceDebugMode env field to set force demog mode
const EnvForceDebugMode = "HerbDebug"

//EnvRootPath env field to set root path
const EnvRootPath = "HerbRoot"

//ForceDebug force useing debug mode
var ForceDebug bool

//IgnoreEnv ignore os env settings.
var IgnoreEnv = false

//Getenv get env value with SystemEnvNamePrefix
func Getenv(envname string) string {
	return os.Getenv(SystemEnvNamePrefix + EnvForceDebugMode)
}

func init() {
	RootPath = os.Getenv(EnvRootPath)
	fmt.Println(RootPath)
	if IgnoreEnv == false && Getenv(EnvForceDebugMode) != "" {
		ForceDebug = true
		Debug = true
	}
}
