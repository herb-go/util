package util

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"
)

func Recover() {
	if r := recover(); r != nil {
		err := r.(error)
		LogError(err)
	}

}

func getStackLines(stack []byte, from int, to int) []string {
	lines := strings.Split(string(stack), "\n")
	if len(lines) < 2 {
		return []string{}
	}
	return lines[from:to]
}
func GetStackLines(from int, to int) []string {
	return getStackLines(debug.Stack(), from, to)
}
func LogError(err error) {
	if IsErrorIgnored(err) == false {
		lines := strings.Split(string(debug.Stack()), "\n")
		length := len(lines)
		maxLength := LoggerMaxLength*2 + 7
		if length > maxLength {
			length = maxLength
		}
		var output = make([]string, length-6)
		output[0] = fmt.Sprintf("Panic: %s", err.Error())
		output[0] += "\n" + lines[0]
		copy(output[1:], lines[7:])
		log.Println(strings.Join(output, "\n"))
	}
}
