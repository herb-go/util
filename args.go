package util

import (
	"os"
)

var Args []string

func init() {
	Args = os.Args
}
