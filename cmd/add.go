/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	db "github.com/UtkarshM-hub/student/DB"
	"github.com/spf13/cobra"
)


func GenerateForm() (string,int,string) {
	var firstName, LastName,programming string
	var age int

	fmt.Println(Blue+`Student Info:`+Reset)
	fmt.Println("\nEnter First Name of the student:")
	fmt.Scanln(&firstName)
	fmt.Println("\nEnter Last Name of the student:")
	fmt.Scanln(&LastName)

	fmt.Println("\nEnter Age of the student:")
	fmt.Scanln(&age)

	fmt.Println("\nEnter the Programming Language Student Prefers:")
	fmt.Scan(&programming)
	
	name:=firstName+" "+LastName
	return name,age,programming
}

// AddCmd represents the Add command
var AddCmd = &cobra.Command{
	Use:   "Add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name,age,programming:=GenerateForm()
		// fmt.Println(name,age,programming)
		err:=db.AddStudent(name,age,programming)
		if err!=nil{
			fmt.Println("❗"+Red+"Error Adding Student"+Reset)
		} else {
			fmt.Println("✔"+Green+" Successfully Added Student"+Reset)
		}
	},
}

func init() {
	rootCmd.AddCommand(AddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
