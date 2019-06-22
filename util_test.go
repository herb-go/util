package util

import (
	"errors"
	"fmt"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestUtil(t *testing.T) {
	func() {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal(r)
			}
			err := r.(error)
			if err == nil {
				t.Fatal(err)
			}
		}()
		Must(fmt.Errorf("%s", "testerror"))
	}()
	err := &net.OpError{
		Op:     "test",
		Net:    "tcp",
		Source: nil,
		Addr:   nil,
		Err:    syscall.EPIPE,
	}
	if !IsErrorIgnored(err) {
		t.Fatal(err)
	}
	oserr := &os.SyscallError{
		Err: syscall.EPIPE,
	}
	err = &net.OpError{
		Err: oserr,
	}
	if !IsErrorIgnored(err) {
		t.Fatal(err)
	}
	testerror := errors.New("testIgnoredError")
	RegisterLoggerIgnoredErrorsChecker(func(err error) bool {
		return err == testerror
	})
	if !IsErrorIgnored(testerror) {
		t.Fatal(err)
	}
	e := errors.New("newerr")
	if IsErrorIgnored(e) {
		t.Fatal(e)
	}
}
