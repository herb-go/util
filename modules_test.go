package util

import "testing"

func TestModules(t *testing.T) {
	defer func() {
		CleanModules()
	}()
	CleanModules()
	output := ""
	unloadoutput := ""
	RegisterModule("2", func() {
		output = output + "module2"
		OnUnloadModules(func() {
			unloadoutput = unloadoutput + "unload2"
		})
	})
	RegisterModule("1", func() {
		output = output + "module1"
		OnUnloadModules(func() {
			unloadoutput = unloadoutput + "unload1"
		})
	})

	InitModulesOrderByName()
	if output != "module1module2" {
		t.Fatal(output)
	}
	UnloadModules()
	if unloadoutput != "unload2unload1" {
		t.Fatal(unloadoutput)
	}
}
