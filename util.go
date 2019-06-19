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
	QuitChan <- 1
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
