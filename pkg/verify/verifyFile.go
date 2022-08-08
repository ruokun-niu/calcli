package verify

import (
	"errors"
	"os"

	dir "github.com/ruokun-niu/calcli/constants"
)

func VerifyFileExist() bool {
	directory := dir.TodoDirectory
	if _, err := os.Stat(directory); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
