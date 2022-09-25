/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Help Command",
	Long: `Use student command instead of student help
	student help is resersed to get information about single command 
	for example: student help Add`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// fmt.Println("help called")
	// },
}

func init() {
	rootCmd.AddCommand(helpCmd)
	// rootCmd.SetHelpCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
