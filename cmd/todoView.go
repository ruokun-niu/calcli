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

	dir "github.com/ruokun-niu/calcli/constants"
	"github.com/spf13/cobra"
)

// todoViewCmd represents the todoView command
var todoViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the todo list",
	Long: `This command will display all of the items in your todo list
If the todo list has not been initialized yet, it will prompt the user to do so`,
	Run: func(cmd *cobra.Command, args []string) {
		fileExists := VerifyFileExist()
		if !fileExists {
			// File does not exist; asks the user to run init
			fmt.Println(`Hmmm seems like the todo list is not found
Have you run the command 'calcli todo init'?`)
		} else {
			index, err := ViewIndex(dir.TodoDirectory)
			if err != nil {
				log.Fatalf("An error has occurred, err: %d", err)
			}
			_ = index
			err = viewItems()
			if err != nil {
				log.Fatalf("An error has occurred, err: %d", err)
			}
		}
	},
}

func init() {
	todoCmd.AddCommand(todoViewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todoViewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todoViewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func viewItems() error {
	directory := dir.TodoDirectory
	file, err := os.Open(directory)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
