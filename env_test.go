package util

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	defer func() {
		RootPath = ""
		ResourcesPath = ""
		AppDataPath = ""
		ConfigPath = ""
		SystemPath = ""
		ConstantsPath = ""
		Setenv(EnvRootPath, "")
		Setenv(EnvResourcesPath, "")
		Setenv(EnvAppDataPath, "")
		Setenv(EnvConfigPath, "")
		Setenv(EnvSystemPath, "")
		Setenv(EnvConstantsPath, "")
		Setenv(EnvForceDebugMode, "")
		Debug = false
		ForceDebug = false
	}()
	SetSystemEnvNamePrefix("HerbGo.")
	if SystemEnvNamePrefix() != "HerbGo." {
		t.Fatal(SystemEnvNamePrefix())
	}
	Setenv("testenv", "testenvvalue")
	env := os.Getenv("HerbGo.testenv")
	if env != "testenvvalue" {
		t.Fatal(env)
	}
	if RootPath != "" {
		t.Fatal(RootPath)
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	Setenv(EnvRootPath, cwd)
	Setenv(EnvResourcesPath, cwd+"/Resoures")
	Setenv(EnvAppDataPath, cwd+"/AppData")
	Setenv(EnvConfigPath, cwd+"/Config")
	Setenv(EnvSystemPath, cwd+"/System")
	Setenv(EnvConstantsPath, cwd+"/Constants")
	initEnv()
	if RootPath != cwd {
		t.Fatal(RootPath)
	}
	if ResourcesPath != cwd+"/Resoures" {
		t.Fatal(RootPath)
	}
	if AppDataPath != cwd+"/AppData" {
		t.Fatal(RootPath)
	}
	if ConfigPath != cwd+"/Config" {
		t.Fatal(RootPath)
	}
	if SystemPath != cwd+"/System" {
		t.Fatal(RootPath)
	}
	if ConstantsPath != cwd+"/Constants" {
		t.Fatal(RootPath)
	}
	if ForceDebug != false || Debug != false {
		t.Fatal(ForceDebug, Debug)
	}
	IgnoreEnv = true
	Setenv(EnvForceDebugMode, "true")
	initEnv()
	if ForceDebug != false || Debug != false {
		t.Fatal(ForceDebug, Debug)
	}

	IgnoreEnv = false
	Setenv("Debug", "true")
	initEnv()
	if ForceDebug != true || Debug != true {
		t.Fatal(ForceDebug, Debug)
	}
}
