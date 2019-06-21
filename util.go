package util

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

var QuitChan = make(chan int)
var SignalChan = make(chan os.Signal)
var LeaveMessage = "Bye."

var Debug = false
var DebugOutput = os.Stdout

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
		close(QuitChan)
	case <-QuitChan:
	}
	fmt.Println("Quiting ...")
}
func Bye() {
	if LeaveMessage != "" {
		fmt.Println(LeaveMessage)
	}
}
func Quit() {
	defer func() {
		recover()
	}()
	close(QuitChan)
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
