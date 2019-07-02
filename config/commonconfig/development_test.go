package commonconfig

import (
	"strings"
	"testing"

	"github.com/herb-go/util"
)

func TestDevelopment(t *testing.T) {
	config := &DevelopmentConfig{}
	func() {
		defer func() {
			r := recover()
			if r != nil {
				t.Fatal(r)
			}
		}()
		config.InitializeAndPanicIfNeeded()
	}()
	func() {
		defer func() {
			defer config.CleanWatehedEnvs()
			defer util.SetHerbEnv("test", "")
			r := recover()
			if r == nil {
				t.Fatal(r)
			}
			err := r.(error)
			if !strings.Contains(err.Error(), util.EnvnameWithPrefix("test")) {
				t.Fatal(err)
			}
		}()
		util.SetHerbEnv("test", "testval")
		config.OnEnv("test").ThenInitalize(func() bool {
			return true
		})
		config.InitializeAndPanicIfNeeded()
	}()
	config.Initializing = true
	func() {
		defer func() {
			defer config.CleanWatehedEnvs()
			util.SetHerbEnv("test", "")
			r := recover()
			if r == nil {
				t.Fatal(r)
			}
			err := r.(error)
			if err != ErrAppInitialized {
				t.Fatal(err)
			}
		}()
		config.OnEnv("test").ThenInitalize(func() bool {
			val := config.GetInitializeEnv("test")
			if val != "testval" {
				t.Fatal(val)
			}
			return true
		})
		util.SetHerbEnv("test", "testval")
		config.InitializeAndPanicIfNeeded()
	}()

	func() {
		defer func() {
			defer config.CleanWatehedEnvs()
			util.SetHerbEnv("test", "")
			r := recover()
			if r == nil {
				t.Fatal(r)
			}
			err := r.(error)
			if err != ErrAppIsInInitializingMode {
				t.Fatal(err)
			}
		}()
		config.OnEnv("test").ThenInitalize(func() bool {
			return false
		})
		util.SetHerbEnv("test", "testval")
		config.InitializeAndPanicIfNeeded()
	}()

	func() {
		defer func() {
			defer config.CleanWatehedEnvs()
			util.SetHerbEnv("test", "")
			r := recover()
			if r == nil {
				t.Fatal(r)
			}
			err := r.(error)
			if err != ErrAppIsInInitializingMode {
				t.Fatal(err)
			}
		}()
		config.OnEnv("test").ThenInitalize(func() bool {
			return true
		})
		util.SetHerbEnv("test", "")
		config.InitializeAndPanicIfNeeded()
	}()
}
