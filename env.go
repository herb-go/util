package util

import (
	"os"
)

//SystemEnvNamePrefix system env name prefix
var SystemEnvNamePrefix = "HerbGo."

//EnvForceDebugMode env field to set force demog mode
const EnvForceDebugMode = "Debug"

//EnvRootPath env field to set root path
const EnvRootPath = "Path.Root"

const EnvResourcesPath = "Path.Resources"

const EnvConfigPath = "Path.Config"

const EnvAppDataPath = "Path.AppData"

const EnvSystemPath = "Path.System"

const EnvConstantsPath = "Path.Constants"

//ForceDebug force useing debug mode
var ForceDebug bool

//IgnoreEnv ignore os env settings.
var IgnoreEnv = false

//Getenv get env value with SystemEnvNamePrefix
func Getenv(envname string) string {
	return os.Getenv(SystemEnvNamePrefix + envname)
}

func Setenv(envname string, val string) {
	os.Setenv(SystemEnvNamePrefix+envname, val)
}
func initEnv() {
	ForceDebug = false
	Debug = false
	RootPath = Getenv(EnvRootPath)
	ResourcesPath = Getenv(EnvResourcesPath)
	AppDataPath = Getenv(EnvAppDataPath)
	ConfigPath = Getenv(EnvConfigPath)
	SystemPath = Getenv(EnvSystemPath)
	ConstantsPath = Getenv(EnvConstantsPath)
	if IgnoreEnv == false && Getenv(EnvForceDebugMode) != "" {
		ForceDebug = true
		Debug = true
	}
}
func init() {
	initEnv()
}
