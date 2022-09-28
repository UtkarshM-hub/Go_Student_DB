/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// UpdateCmdCmd represents the UpdateCmd command
var UpdateCmdCmd = &cobra.Command{
	Use:   "UpdateCmd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Blue+"Update Command:"+Reset)
		fmt.Println(`	It is used to update student information into database.
	When the command is invoked by running `+Yellow+`student Update <INDEX_NUMBER>`+Reset+
	`It generates a form to fill in the student information.
	Like Name, Age and Programming Language the student prefers, to update the student information.`)
	},
}

func init() {
	helpCmd.AddCommand(UpdateCmdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UpdateCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UpdateCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
