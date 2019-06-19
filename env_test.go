package util

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("HerbGo.testenv", "testenvvalue")
	env := Getenv("testenv")
	if env != "testenvvalue" {
		t.Fatal(env)
	}
	if ForceDebug != false || Debug != false {
		t.Fatal(ForceDebug, Debug)
	}
	IgnoreEnv = true
	os.Setenv("HerbGo.HerbDebug", "true")
	initEnv()
	if ForceDebug != false || Debug != false {
		t.Fatal(ForceDebug, Debug)
	}

	IgnoreEnv = false
	os.Setenv("HerbGo.HerbDebug", "true")
	initEnv()
	if ForceDebug != true || Debug != true {
		t.Fatal(ForceDebug, Debug)
	}
}
