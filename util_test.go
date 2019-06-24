package util

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"syscall"
	"testing"
	"time"
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

func TestPrint(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})
	Output = buffer
	defer func() {
		Output = os.Stdout
	}()
	_, err := Print("123", "456")
	if err != nil {
		t.Fatal(err)
	}
	bs, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Fatal(err)
	}
	if string(bs) != fmt.Sprint("123", "456") {
		t.Fatal(string(bs))
	}
	_, err = Println("123", "456")
	if err != nil {
		t.Fatal(err)
	}
	bs, err = ioutil.ReadAll(buffer)
	if err != nil {
		t.Fatal(err)
	}
	if string(bs) != fmt.Sprintln("123", "456") {
		t.Fatal(string(bs))
	}
	_, err = Printf("%s %s", "123", "456")
	if err != nil {
		t.Fatal(err)
	}
	bs, err = ioutil.ReadAll(buffer)
	if err != nil {
		t.Fatal(err)
	}
	if string(bs) != fmt.Sprintf("%s %s", "123", "456") {
		t.Fatal(string(bs))
	}
	Bye()
	bs, err = ioutil.ReadAll(buffer)
	if err != nil {
		t.Fatal(err)
	}
	if string(bs) != fmt.Sprintln(LeaveMessage) {
		t.Fatal(string(bs))
	}
}

func TestQuit(t *testing.T) {
	go func() {
		time.Sleep(time.Millisecond)
		Quit()
	}()
	WaitingQuit()
	go func() {
		time.Sleep(time.Millisecond)
		SignalChan <- nil
	}()
	WaitingQuit()
	Quit()

}
