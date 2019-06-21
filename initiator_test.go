package util

import "testing"

func TestInitiator(t *testing.T) {
	defer func() {
		CleanInitiators()
		Debug = false
	}()
	Debug = true
	CleanInitiators()
	var output = ""
	RegisterInitiator("test", "100", func() {
		output = output + "100"
	})
	RegisterInitiator("test", "200", func() {
		output = output + "200"
	})
	InitOrderByName("test")
	if output != "100200" {
		t.Fatal(output)
	}
	InitOrderByName("Notexist")
}
