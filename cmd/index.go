package cmd

import (
	"bufio"
	"os"
	"strconv"

	dir "github.com/ruokun-niu/calcli/constants"
)

func modifyIndex() error {
	index, err := ViewIndex()
	if err != nil {
		return err
	}
	index++
	return nil
}

func ViewIndex() (int, error) {
	directory := dir.TodoDirectory
	file, err := os.Open(directory)
	if err != nil {
		return -1, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	result, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return -1, err
	}
	return result, nil
}
