package translateconfig

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/herb-go/herb/translate"
)

func getPath() string {
	_, p, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(p), "testdata")
}
func TestConfig(t *testing.T) {
	p := getPath()
	c := &Config{
		Language:  "testlang",
		Avaliable: []string{"test", "testlang", "zh"},
	}
	err := c.Apply(p)
	if err != nil {
		t.Fatal(err)
	}
	if translate.Lang != "testlang" {
		t.Fatal(translate.Lang)
	}
	message := translate.GetIn("zh", "testmodule", "test")
	if message != "translated test" {
		t.Fatal(message)
	}
}
