package util

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

var ApplieationLock sync.Mutex

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

var initQuitChan = make(chan int)
var quitChan = unsafe.Pointer(&initQuitChan)

func resetQuitChan() chan int {
	cnew := make(chan int)
	cold := atomic.SwapPointer(&quitChan, unsafe.Pointer(&cnew))
	return *(*chan int)(cold)
}

func QuitChan() chan int {
	c := atomic.LoadPointer(&quitChan)
	return *(*chan int)(c)
}

var SignalChan = make(chan os.Signal)
var LeaveMessage = "Bye."

var Debug = false
var DebugOutput = os.Stdout
var Stdout io.Writer = os.Stdout

func Println(args ...interface{}) (n int, err error) {
	return fmt.Fprintln(Stdout, args...)
}

func Printf(format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(Stdout, format, args...)
}
func Print(args ...interface{}) (n int, err error) {
	return fmt.Fprint(Stdout, args...)
}
func DebugPrintln(args ...interface{}) {
	if Debug || ForceDebug {
		fmt.Fprintln(DebugOutput, args...)
	}
}

func DebugPrintf(format string, args ...interface{}) {
	if Debug || ForceDebug {
		fmt.Fprintf(DebugOutput, format, args...)
	}
}
func DebugPrint(args ...interface{}) {
	if Debug || ForceDebug {
		fmt.Fprint(DebugOutput, args...)
	}
}

var QuitDelayDuration = 500 * time.Microsecond

func DelayAndQuit() {
	Println("Delay for quit.")
	time.Sleep(QuitDelayDuration)
	Println("Quiting ...")
}
func WaitingQuit() {
	signal.Notify(SignalChan, os.Interrupt, os.Kill)
	select {
	case <-SignalChan:
		Quit()
	case <-QuitChan():
	}
}

func OnQuit(handlers ...func()) {
	for k := range handlers {
		handler := handlers[k]
		go func() {
			<-QuitChan()
			handler()
		}()
	}
}

func OnQuitAndLogError(handlers ...func() error) {
	for k := range handlers {
		handler := handlers[k]
		go func() {
			<-QuitChan()
			err := handler()
			if err != nil {
				LogError(err)
			}
		}()
	}
}

func OnQuitAndIgnoreError(handlers ...func() error) {
	for k := range handlers {
		handler := handlers[k]
		go func() {
			<-QuitChan()
			handler()
		}()
	}
}
func Bye() {
	if LeaveMessage != "" {
		Println(LeaveMessage)
	}
}

func Quit() {
	defer func() {
		recover()
	}()
	c := resetQuitChan()
	if c != nil {
		close(c)
	}
}

var LoggerMaxLength = 5
var LoggerIgnoredErrorsChecker = []func(error) bool{
	func(err error) bool {
		if err == http.ErrHandlerTimeout || err == http.ErrAbortHandler {
			return true
		}
		if err == context.Canceled {
			return true
		}
		oe, ok := err.(*net.OpError)
		if ok {
			if oe.Err == syscall.EPIPE || oe.Err == syscall.ECONNRESET {
				return true
			}
			se, ok := oe.Err.(*os.SyscallError)
			if ok && (se.Err == syscall.EPIPE || se.Err == syscall.ECONNRESET) {
				return true
			}
		}
		return false
	},
}

var IsErrorIgnored = func(err error) bool {
	for k := range LoggerIgnoredErrorsChecker {
		if LoggerIgnoredErrorsChecker[k](err) {
			return true
		}
	}
	return false
}
var RegisterLoggerIgnoredErrorsChecker = func(f func(error) bool) {
	LoggerIgnoredErrorsChecker = append(LoggerIgnoredErrorsChecker, f)
}

func init() {
	resetQuitChan()
}
