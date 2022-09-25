/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// AddCmdCmd represents the AddCmd command
var AddCmdCmd = &cobra.Command{
	Use:   "AddCmd",
	Short: "A Command to add student data to database",
	Long: `Adds student information to database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Blue+"Add Command:"+Reset)
		fmt.Println(`	It is used to add student information into database.
	When the command is invoked by running `+Yellow+`student add `+Reset+
	`It generates a form to fill in the student information.
	Like Name, Age and Programming Language the student prefers.`)
	},
}

func init() {
	helpCmd.AddCommand(AddCmdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
