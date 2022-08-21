// ------------------------------------------------------------
// Copyright © 2022 HalfsugarDev halfsugardev7@gmail.com
// Licensed under the MIT License.
// ------------------------------------------------------------
package cmd

import (
	"fmt"
	"log"
	"strconv"

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
		toDelIndex, err := strconv.Atoi(args[0])
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
