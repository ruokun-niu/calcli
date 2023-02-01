// ------------------------------------------------------------
// Copyright © 2022 HalfsugarDev halfsugardev7@gmail.com
// Licensed under the MIT License.
// ------------------------------------------------------------

package cmd

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

func VerifyCompleteExist() bool {
	directory := dir.CompleteDirectory
	if _, err := os.Stat(directory); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
