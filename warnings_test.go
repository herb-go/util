package util

import "testing"

func TestWarnings(t *testing.T) {
	defer func() {
		CleanWarnings()
	}()
	CleanWarnings()
	if HasWarning() {
		t.Fatal(Warnings())
	}
	SetWarning("test", "test info", "test info3")
	if !HasWarning() {
		t.Fatal(Warnings())
	}
	w := Warnings()
	if len(w) != 1 || len(w["test"]) != 2 || w["test"][0] != "test info" || w["test"][1] != "test info3" {
		t.Fatal(w)
	}
	DelWarning("notexists")
	if !HasWarning() {
		t.Fatal(Warnings())
	}
	DelWarning("test")
	if HasWarning() {
		t.Fatal(Warnings())
	}
	SetWarning("test2", "test2")
	SetWarning("test3", "test3")
	if !HasWarning() {
		t.Fatal(Warnings())
	}
	CleanWarnings()
	if HasWarning() {
		t.Fatal(Warnings())
	}

}
