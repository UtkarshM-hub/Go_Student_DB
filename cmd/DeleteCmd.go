/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// DeleteCmdCmd represents the DeleteCmd command
var DeleteCmdCmd = &cobra.Command{
	Use:   "DeleteCmd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Blue+"Delete Command:"+Reset)
		fmt.Println(`	It is used to Delete student information from the database.
	When the command is invoked by running `+Yellow+`student Del <INDEX_NUMBER> `+Reset+
	`Removes the student information from the database.`)
	},
}

func init() {
	helpCmd.AddCommand(DeleteCmdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DeleteCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DeleteCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
