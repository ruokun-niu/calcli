// ------------------------------------------------------------
// Copyright © 2022 HalfsugarDev halfsugardev7@gmail.com
// Licensed under the MIT License.
// ------------------------------------------------------------
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	dir "github.com/ruokun-niu/calcli/constants"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an item from your todo list",
	Long: `Delete an item from your todo list.
The deleted item will not be added to the completed list
`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO add a check to see if an argument is passed
		// maybe a fatal check?
		toDelIndex, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("an error has occurred when trying to delete an item, err: %d", err)
		}
		err = deleteItem(toDelIndex)
		if err != nil {
			log.Fatalf("an error has occurred when trying to delete an item, err: %d", err)
		}
		fmt.Println("delete called")
	},
}

func init() {
	todoCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteItem(index int) error {
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

func editIndex(currItem string) string {
	originalStrIndex := strings.Split(currItem, " ")[0]
	index, _ := strconv.Atoi(originalStrIndex)
	index--
	strIndex := strconv.Itoa(index)
	result := strings.Replace(currItem, originalStrIndex, strIndex, 1)
	return result
}
