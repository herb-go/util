package util

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestRecovery(t *testing.T) {
	var testerror = errors.New("testerror")
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	output := bytes.NewBuffer([]byte{})
	log.SetOutput(output)
	func() {
		defer Recover()
		panic(testerror)
	}()
	_, err := ioutil.ReadAll(output)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetStackLines(t *testing.T) {
	lines := getStackLines([]byte{}, 1, 10)
	if len(lines) != 0 {
		t.Fatal(lines)
	}
	lines = GetStackLines(0, 1)
	if len(lines) == 0 {
		t.Fatal(lines)
	}
}
