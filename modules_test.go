package util

import "testing"

func TestModules(t *testing.T) {
	defer func() {
		CleanModules()
		Debug = false
	}()
	CleanModules()
	Debug = true
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
	CleanModules()
	output = ""
	unloadoutput = ""

	RegisterModule("3", func() {
		output = output + "module3"
		OnUnloadModules(func() {
			unloadoutput = unloadoutput + "unload3"
		})
	})
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
	InitModulesOrderByName("3", "2")
	if output != "module2module3" {
		t.Fatal(output)
	}
	UnloadModules()
	if unloadoutput != "unload3unload2" {
		t.Fatal(unloadoutput)
	}

	CleanModules()
	output = ""
	unloadoutput = ""

	StageInit.RegisterModule("3", func() {
		output = output + "module3"
		OnUnloadModules(func() {
			unloadoutput = unloadoutput + "unload3"
		})
	})
	RegisterModule("2", func() {
		output = output + "module2"
		OnUnloadModules(func() {
			unloadoutput = unloadoutput + "unload2"
		})
	})
	StageFinish.RegisterModule("1", func() {
		output = output + "module1"
		OnUnloadModules(func() {
			unloadoutput = unloadoutput + "unload1"
		})
	})
	InitModulesOrderByName()
	if output != "module3module2module1" {
		t.Fatal(output)
	}
	UnloadModules()
	if unloadoutput != "unload1unload2unload3" {
		t.Fatal(unloadoutput)
	}
}
