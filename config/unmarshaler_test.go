package config

import (
	"strings"
	"testing"
)

func TestImarshaler(t *testing.T) {
	err := Unmarshal("dirvernotexsit", nil, nil)
	if err == nil || !strings.Contains(err.Error(), "dirvernotexsit") {
		t.Fatal(err)
	}
}
