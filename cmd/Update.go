/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	db "github.com/UtkarshM-hub/student/DB"
	"github.com/spf13/cobra"
)

// UpdateCmd represents the Update command
var UpdateCmd = &cobra.Command{
	Use:   "Update",
	Short: "Updates the student information at given index",
	Long: `Updates the student information at given index in database`,
	Run: func(cmd *cobra.Command, args []string) {
		intVal,err:=strconv.Atoi(args[0])
		if err!=nil{
			log.Fatal("The Value provided by the user is not valid!❗")
		}
		name,age,language:=GenerateForm()
		err=db.UpdateUserInfo(name,age,language,intVal)
		if err!=nil{
			fmt.Printf("❗"+Red+"Error Updating Student from index %v"+Reset,intVal)
		} else {
			fmt.Println("✔"+Green+" Successfully Updated the Student Information"+Reset)
		}
	},
}

func init() {
	rootCmd.AddCommand(UpdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
