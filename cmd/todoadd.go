// ------------------------------------------------------------
// Copyright © 2022 HalfsugarDev halfsugardev7@gmail.com
// Licensed under the MIT License.
// ------------------------------------------------------------
package cmd

import (
	"fmt"
	"log"
	"os"

	dir "github.com/ruokun-niu/calcli/constants"
	"github.com/spf13/cobra"
)

// todoaddCmd represents the todoadd command
var todoaddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an item to your todo list ",
	Long: `The cli todo list behaves like a queue (FIFO)
This command will add a new item to the end of the list
Type -h to see other ways of adding an item`,
	Run: func(cmd *cobra.Command, args []string) {
		fileExists := VerifyFileExist()
		if !fileExists {
			// File does not exist; asks the user to run init
			fmt.Println(`Hmmm seems like the todo list is not found
Have you run the command 'calcli todo init'?`)
		} else {
			item := args[0]
			err := writeFile(item)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	todoCmd.AddCommand(todoaddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todoaddCmd.PersistentFlags().String("foo", "", "A help for foo")
	todoaddCmd.PersistentFlags().String("temp", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todoaddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func writeFile(item string) error {
	toWrite := item + "\n"
	directory := dir.TodoDirectory
	file, err := os.OpenFile(directory, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to open the file, err: %d", err)
	}
	defer file.Close()
	if _, err = file.Write([]byte(toWrite)); err != nil {
		return fmt.Errorf("failed to write into the todo list, err: %d", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("Error closing the file")
	}

	return nil
}
