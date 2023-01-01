package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	dir "github.com/ruokun-niu/calcli/constants"
)

func incrementIndex() error {
	index, err := ViewIndex(dir.TodoDirectory)
	if err != nil {
		return err
	}
	index++
	folderDir := "/Users/ruokunniu/calcli/foo.txt"

	newFile, err := os.Create(folderDir)

	if err != nil {
		return fmt.Errorf("encountered an error when trying to create a dummy txt, err: %d", err)
	}
	defer newFile.Close()

	// Adding the index to the top of the text file
	strIndex := strconv.Itoa(index) + "\n"
	_, err = newFile.WriteString(strIndex)
	if err != nil {
		return fmt.Errorf("encountered an error when trying to append the new index, err: %d", err)
	}
	originalFile, err := os.Open(dir.TodoDirectory)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(originalFile)

	//Skipping first line
	scanner.Scan()
	for scanner.Scan() {
		// Copy the content from the original txt to the new one
		_, err = newFile.WriteString(scanner.Text())
		_, err = newFile.WriteString("\n")
	}
	newFile.Sync()

	err = os.Rename(folderDir, dir.TodoDirectory)
	if err != nil {
		return err
	}
	return nil
}

func decrementIndex() error {
	index, err := ViewIndex(dir.TodoDirectory)
	if err != nil {
		return err
	}
	index--
	folderDir := "/Users/ruokunniu/calcli/foo.txt"

	newFile, err := os.Create(folderDir)

	if err != nil {
		return fmt.Errorf("encountered an error when trying to create a dummy txt, err: %d", err)
	}
	defer newFile.Close()

	// Adding the index to the top of the text file
	strIndex := strconv.Itoa(index) + "\n"
	_, err = newFile.WriteString(strIndex)
	if err != nil {
		return fmt.Errorf("encountered an error when trying to append the new index, err: %d", err)
	}
	originalFile, err := os.Open(dir.TodoDirectory)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(originalFile)

	//Skipping first line
	scanner.Scan()
	for scanner.Scan() {
		// Copy the content from the original txt to the new one
		_, err = newFile.WriteString(scanner.Text())
		_, err = newFile.WriteString("\n")
	}
	newFile.Sync()

	err = os.Rename(folderDir, dir.TodoDirectory)
	if err != nil {
		return err
	}
	return nil
}

func ViewIndex(directory string) (int, error) {
	file, err := os.Open(directory)
	if err != nil {
		return -1, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if scanner.Text() == "" {
		return 0, nil
	}
	result, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return -1, err
	}
	return result, nil
}

func EditIndex(currItem string) string {
	originalStrIndex := strings.Split(currItem, " ")[0]
	index, _ := strconv.Atoi(originalStrIndex)
	index--
	strIndex := strconv.Itoa(index)
	result := strings.Replace(currItem, originalStrIndex, strIndex, 1)
	return result
}

func IncrementIndexForComplete() error {
	index, err := ViewIndex(dir.CompleteDirectory)
	if index == -1 {
		index = 0
	} else if err != nil {
		return err
	}
	index++
	folderDir := "/Users/ruokunniu/calcli/foo.txt"

	newFile, err := os.Create(folderDir)

	if err != nil {
		return fmt.Errorf("encountered an error when trying to create a dummy txt, err: %d", err)
	}
	defer newFile.Close()

	// Adding the index to the top of the text file
	strIndex := strconv.Itoa(index) + "\n"
	_, err = newFile.WriteString(strIndex)
	if err != nil {
		return fmt.Errorf("encountered an error when trying to append the new index, err: %d", err)
	}
	originalFile, err := os.Open(dir.CompleteDirectory)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(originalFile)

	//Skipping first line
	scanner.Scan()
	for scanner.Scan() {
		// Copy the content from the original txt to the new one
		_, err = newFile.WriteString(scanner.Text())
		_, err = newFile.WriteString("\n")
	}
	newFile.Sync()

	err = os.Rename(folderDir, dir.CompleteDirectory)
	if err != nil {
		return err
	}
	return nil
}
