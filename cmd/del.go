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

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "Del",
	Short: "Delets the student information of given index",
	Long: `Delets the student information of given index from the database`,
	Run: func(cmd *cobra.Command, args []string) {
		intVal,err:=strconv.Atoi(args[0])
		if err!=nil{
			log.Fatal("The Value provided by the user is not valid!❗")
		}
		err=db.DeleteStudent(intVal)
		if err!=nil{
			fmt.Printf("❗"+Red+"Error Deleting Student from index %v"+Reset,intVal)
		} else {
			fmt.Println("✔"+Green+" Successfully Deleted the Student Information"+Reset)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
