/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	dir "github.com/ruokun-niu/calcli/constants"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark an item as completed",
	Long: `Marking an item as completed will remove it from the current list.
The item will be added to the complete list, which holds a maximum
of 10 items.
You can view the complete list by running the command 'calcli view --complete`,
	Run: func(cmd *cobra.Command, args []string) {
		if !VerifyCompleteExist() {
			initCompleteList()
		}
		toCompIndex, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("an error has occurred when trying to complete an item, err: %d", err)
		}
		err = completeItem(toCompIndex)
		if err != nil {
			log.Fatalf("an error has occurred when trying to complete an item, err: %d", err)
		}
		err = decrementIndex()
		if err != nil {
			log.Fatalf("an error has occurred when trying to decrement the index, err: %d", err)
		}
		fmt.Println("complete called")
	},
}

func init() {
	todoCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func completeItem(index int) error {
	directory := dir.TodoDirectory
	folderDir := "/Users/ruokunniu/calcli/foo.txt"

	newFile, err := os.Create(folderDir)
	if err != nil {
		return err
	}
	defer newFile.Close()

	originalFile, err := os.Open(directory)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(originalFile)
	scanner.Scan()
	_, err = newFile.WriteString(scanner.Text())
	_, err = newFile.WriteString("\n")
	for i := 0; i < index-1; i++ {
		// Prior to reaching the item
		scanner.Scan()
		_, err = newFile.WriteString(scanner.Text())
		_, err = newFile.WriteString("\n")
	}
	//Skipping one line
	scanner.Scan()
	toComplete := scanner.Text()

	err = writeToComplete(toComplete)
	if err != nil {
		return err
	}

	for scanner.Scan() {
		// Scan to the end of the file
		// changing the indices along the way
		newText := editIndex(scanner.Text())
		_, err = newFile.WriteString(newText)
		_, err = newFile.WriteString("\n")
	}
	newFile.Sync()

	//rename foo
	err = os.Rename(folderDir, directory)
	if err != nil {
		return err
	}
	return nil
}

func initCompleteList() error {
	directory := dir.TodoFolderDirectory
	filePath := path.Join(directory, "complete.txt")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("The completed list has been successfully created!")
	return nil
}

func writeToComplete(item string) error {
	completeDir := dir.CompleteDirectory
	completeFile, err := os.OpenFile(completeDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer completeFile.Close()

	_, err = completeFile.WriteString(item)
	_, err = completeFile.WriteString("\n")

	return nil
}
