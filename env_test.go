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
		SetHerbEnv(EnvRootPath, "")
		SetHerbEnv(EnvResourcesPath, "")
		SetHerbEnv(EnvAppDataPath, "")
		SetHerbEnv(EnvConfigPath, "")
		SetHerbEnv(EnvSystemPath, "")
		SetHerbEnv(EnvConstantsPath, "")
		SetHerbEnv(EnvForceDebugMode, "")
		Debug = false
		ForceDebug = false
	}()
	SetSystemEnvNamePrefix("HerbGo.")
	if SystemEnvNamePrefix() != "HerbGo." {
		t.Fatal(SystemEnvNamePrefix())
	}
	SetHerbEnv("testenv", "testenvvalue")
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
	SetHerbEnv(EnvRootPath, cwd)
	SetHerbEnv(EnvResourcesPath, cwd+"/Resoures")
	SetHerbEnv(EnvAppDataPath, cwd+"/AppData")
	SetHerbEnv(EnvConfigPath, cwd+"/Config")
	SetHerbEnv(EnvSystemPath, cwd+"/System")
	SetHerbEnv(EnvConstantsPath, cwd+"/Constants")
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
	SetHerbEnv(EnvForceDebugMode, "true")
	initEnv()
	if ForceDebug != false || Debug != false {
		t.Fatal(ForceDebug, Debug)
	}

	IgnoreEnv = false
	SetHerbEnv("Debug", "true")
	initEnv()
	if ForceDebug != true || Debug != true {
		t.Fatal(ForceDebug, Debug)
	}
}

func TestOsEnv(t *testing.T) {
	envvalue := os.Getenv("testenvname")
	if envvalue != "" {
		t.Fatal(envvalue)
	}
	value := OsEnv.Getenv("testenvname")
	if value != "" {
		t.Fatal(value)
	}
	err := OsEnv.Setenv("testenvname", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		OsEnv.Setenv("testenvname", "")
	}()
	value = OsEnv.Getenv("testenvname")
	if value != "test" {
		t.Fatal(value)
	}
	envvalue = os.Getenv("testenvname")
	if envvalue != "test" {
		t.Fatal(envvalue)
	}
}

func TestPresetEnv(t *testing.T) {
	var TestEnv = PresetEnv(map[string]string{
		"test": "testvalue",
	})
	value := TestEnv.Getenv("test")
	if value != "testvalue" {
		t.Fatal(TestEnv)
	}
	err := TestEnv.Setenv("test", "testvalue2")
	if err != nil {
		t.Fatal(err)
	}
	value = TestEnv.Getenv("test")
	if value != "testvalue2" {
		t.Fatal(TestEnv)
	}
}
