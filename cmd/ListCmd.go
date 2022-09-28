/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ListCmdCmd represents the ListCmd command
var ListCmdCmd = &cobra.Command{
	Use:   "ListCmd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Blue+"List Command:"+Reset)
		fmt.Println(`	It is used to show student information in a tabular form.
	When the command is invoked by running `+Yellow+`student List `+Reset+
	`It shows the student information in the tabular form.`)
	},
}

func init() {
	helpCmd.AddCommand(ListCmdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ListCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ListCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
