// ------------------------------------------------------------
// Copyright © 2022 HalfsugarDev halfsugardev7@gmail.com
// Licensed under the MIT License.
// ------------------------------------------------------------
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	dir "github.com/ruokun-niu/calcli/constants"
	"github.com/spf13/cobra"
)

const (
	runHelpComplete = "Type 'calcli delete -h' for more details on using this command."
	ContactRepo     = "Please submit an issue or email halfsugardev7@gmail.com"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove an item from your todo list",
	Long: `Remove an item from your todo list.
The deleted item will not be added to the completed list.
To delete an item, simply input the index of the item:
e.g. 'calcli delete 1' will delete the item with index #1.
`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO add a check to see if an argument is passed
		// maybe a fatal check?
		if len(args) == 0 {
			fmt.Println("Please input an index")
			fmt.Println("e.g. calcli delete 1")
			fmt.Println(runHelpComplete)
			os.Exit(0)
		}
		toDelIndex, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("An error has occurred when trying to remove the item")
			fmt.Println(ContactRepo)
			fmt.Println(runHelpComplete)
			os.Exit(0)
		}
		err = deleteItem(toDelIndex)
		if err != nil {
			fmt.Println("An error has occurred when trying to remove the item")
			fmt.Println(ContactRepo)
			fmt.Println(runHelpComplete)
			os.Exit(0)
		}
		err = decrementIndex()
		if err != nil {
			fmt.Println("An error has occurred when trying to decrement the index")
			fmt.Println(ContactRepo)
			fmt.Println(runHelpComplete)
			os.Exit(0)
		}
		fmt.Println("Item deleted")
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

	folderDir := dir.TodoRenameDirectory

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
		newText := EditIndex(scanner.Text())
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
