// ------------------------------------------------------------
// Copyright © 2022 HalfsugarDev halfsugardev7@gmail.com
// Licensed under the MIT License.
// ------------------------------------------------------------
package cmd

import (
	"fmt"

	initializer "github.com/ruokun-niu/calcli/pkg/init"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var todoInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the notebook txt file",
	Long: `This command will create a txt file called 'todo.txt' on your computer
The exact directory (MacOS) where the file lives is as follows:
'/Users/{username}/calcli/todo.txt'
Party on with calcli!
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		initializer.InitializeText()
	},
}

func init() {
	todoCmd.AddCommand(todoInitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
