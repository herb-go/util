package util

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
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
var Output io.Writer = os.Stdout

func Println(args ...interface{}) (n int, err error) {
	return fmt.Fprintln(Output, args...)
}

func Printf(format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(Output, format, args...)
}
func Print(args ...interface{}) (n int, err error) {
	return fmt.Fprint(Output, args...)
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

func WaitingQuit() {
	signal.Notify(SignalChan, os.Interrupt, os.Kill)
	select {
	case <-SignalChan:
		Quit()
	case <-QuitChan():
	}
	Println("Quiting ...")
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
