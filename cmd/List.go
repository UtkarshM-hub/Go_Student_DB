/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	db "github.com/UtkarshM-hub/student/DB"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// ListCmd represents the List command
var ListCmd = &cobra.Command{
	Use:   "List",
	Short: "shows all student data in tabular form",
	Long: `Shows all data in tabular form`,
	Run: func(cmd *cobra.Command, args []string) {
		data,err:=db.GetStudents()
		if err!=nil{
			fmt.Printf("❗"+Red+"Error getting data from the database"+Reset)
		}
		output:=[][]string{}

		for index,student:=range data{
			output=append(output,[]string{fmt.Sprint(index+1),student.Name,fmt.Sprint(student.Age),student.ProgrammingLanguage})
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "Name", "Age","Programming Language"})

		for _, v := range output {
			table.Append(v)
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(ListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
