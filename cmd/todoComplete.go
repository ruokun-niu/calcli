/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path"

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

func completeItem() error {

}

func initCompleteList() error {
	directory := dir.TodoFolderDirectory
	filePath := path.Join(directory, "complete.txt")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("The completed list has been successfully created!")
}

func verifyCompleteExists() (bool, error) {

}
