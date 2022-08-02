/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// todoaddCmd represents the todoadd command
var todoaddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an item to your todo list ",
	Long:  `The cli todo list behaves like a queue (FIFO)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todoadd called")
	},
}

func init() {
	rootCmd.AddCommand(todoaddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todoaddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todoaddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
