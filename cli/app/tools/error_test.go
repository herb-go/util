package tools

import (
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	if err := ErrorIfStringFieldNotInList("testfield", "1", "1", "2", "3"); err != nil {
		t.Fatal(err)
	}
	if err := ErrorIfStringFieldNotInList("testfield", "a", "1", "2", "3"); err == nil || !strings.Contains(err.Error(), "testfield") {
		t.Fatal(err)
	}

	if err := ErrorIfIntFieldNotInRange("testfield", 1, 0, 100); err != nil {
		t.Fatal(err)
	}
	if err := ErrorIfIntFieldNotInRange("testfield", 0, 0, 100); err != nil {
		t.Fatal(err)
	}
	if err := ErrorIfIntFieldNotInRange("testfield", 100, 0, 100); err != nil {
		t.Fatal(err)
	}
	if err := ErrorIfIntFieldNotInRange("testfield", -1, 0, 100); err == nil || !strings.Contains(err.Error(), "testfield") {
		t.Fatal(err)
	}
	if err := ErrorIfIntFieldNotInRange("testfield", 101, 0, 100); err == nil || !strings.Contains(err.Error(), "testfield") {
		t.Fatal(err)
	}

	if err := ErrorIfInt64FieldNotInRange("testfield", 1, 0, 100); err != nil {
		t.Fatal(err)
	}
	if err := ErrorIfInt64FieldNotInRange("testfield", 0, 0, 100); err != nil {
		t.Fatal(err)
	}
	if err := ErrorIfInt64FieldNotInRange("testfield", 100, 0, 100); err != nil {
		t.Fatal(err)
	}
	if err := ErrorIfInt64FieldNotInRange("testfield", -1, 0, 100); err == nil || !strings.Contains(err.Error(), "testfield") {
		t.Fatal(err)
	}
	if err := ErrorIfInt64FieldNotInRange("testfield", 101, 0, 100); err == nil || !strings.Contains(err.Error(), "testfield") {
		t.Fatal(err)
	}
}
