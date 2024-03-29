// ------------------------------------------------------------
// Copyright © 2022 HalfsugarDev halfsugardev7@gmail.com
// Licensed under the MIT License.
// ------------------------------------------------------------
package cmd

import (
	"fmt"
	"os"
	"strconv"

	dir "github.com/ruokun-niu/calcli/constants"
	"github.com/spf13/cobra"
)

const (
	runHelpAdd = "Type 'calcli add -h' for more details on using this command."
)

// todoaddCmd represents the todoadd command
var todoaddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an item to your todo list ",
	Long: `The cli todo list behaves like a queue (FIFO)
This command will add a new item to the end of the list
e.g. 'calcli add coffee'`,
	Run: func(cmd *cobra.Command, args []string) {
		fileExists := VerifyFileExist()
		if !fileExists {
			// File does not exist; asks the user to run init
			fmt.Println(`Hmmm seems like the todo list is not found
Have you run the command 'calcli todo init'?`)
			os.Exit(0)
		} else {
			item := args[0]
			err := writeFile(item)
			if err != nil {
				fmt.Println(err)
				fmt.Println(ContactRepo)
				fmt.Println(runHelpAdd)
				os.Exit(0)
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
	currIndex, err := ViewIndex(dir.TodoDirectory)
	if err != nil {
		return fmt.Errorf("encountered an error when trying to retrieve the index; err: %d", err)
	}
	currIndex++

	toWrite := strconv.Itoa(currIndex) + " " + item + "\n"
	err = incrementIndex()
	if err != nil {
		return fmt.Errorf("encountered an error when trying to increment the index; err: %d", err)
	}
	directory := dir.TodoDirectory
	file, err := os.OpenFile(directory, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to open the todo list, err: %d", err)
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
