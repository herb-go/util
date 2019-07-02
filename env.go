package util

import (
	"os"
)

var systemEnvNamePrefix = "HerbGo."

//SystemEnvNamePrefix system env name prefix
func SystemEnvNamePrefix() string {
	return systemEnvNamePrefix
}

//SetSystemEnvNamePrefix set system env name prefix
func SetSystemEnvNamePrefix(p string) {
	systemEnvNamePrefix = p
}

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
func GetHerbEnv(envname string) string {
	return Getenv(EnvnameWithPrefix(envname))
}

func SetHerbEnv(envname string, val string) {
	Setenv(EnvnameWithPrefix(envname), val)
}

type AppEnv interface {
	Getenv(string) string
	Setenv(string, string) error
}

func Getenv(key string) string {
	return Env.Getenv(key)
}

func Setenv(key string, value string) error {
	return Env.Setenv(key, value)
}

type PrefixEnv struct {
	Prefix string
}

func (e *PrefixEnv) Getenv(key string) string {
	return os.Getenv(e.Prefix + key)
}

func (e *PrefixEnv) Setenv(key string, value string) error {
	return os.Setenv(key, value)
}
func NewPrefixEnv(prefix string) *PrefixEnv {
	return &PrefixEnv{
		Prefix: prefix,
	}
}

var OsEnv = NewPrefixEnv("")

var Env = OsEnv

type PresetEnv map[string]string

func (e PresetEnv) Getenv(key string) string {
	return e[key]
}

func (e PresetEnv) Setenv(key string, value string) error {
	e[key] = value
	return nil
}

func EnvnameWithPrefix(envname string) string {
	return SystemEnvNamePrefix() + envname
}
func initEnv() {
	ForceDebug = false
	Debug = false
	RootPath = GetHerbEnv(EnvRootPath)
	ResourcesPath = GetHerbEnv(EnvResourcesPath)
	AppDataPath = GetHerbEnv(EnvAppDataPath)
	ConfigPath = GetHerbEnv(EnvConfigPath)
	SystemPath = GetHerbEnv(EnvSystemPath)
	ConstantsPath = GetHerbEnv(EnvConstantsPath)
	if IgnoreEnv == false && GetHerbEnv(EnvForceDebugMode) != "" {
		ForceDebug = true
		Debug = true
	}
}
func init() {
	initEnv()
}
