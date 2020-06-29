package util

import (
	"os"
)

var DefaultFileMode = os.FileMode(0660)

var SecretFileMode = os.FileMode(0600)
