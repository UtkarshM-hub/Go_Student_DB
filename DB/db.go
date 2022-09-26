package db

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Student struct {
	Name                string
	Age                 int
	ProgrammingLanguage string
}

// Get the student data from db.json file
func GetStudents() ([]Student, error) {
	byte, err := os.ReadFile("DB/db.json")
	if err!=nil{
		return nil,errors.New("error occured while reading from the file")
	}
	var students []Student
	json.Unmarshal(byte,&students)
	return students,nil
}

// Add the Student to database
func AddStudent(name string,age int,language string){
	// get students from db to append data
	data,err:=GetStudents()
	if err!=nil{
		log.Fatal(err)
	}
	var newStudent Student
	newStudent.Name=name
	newStudent.Age=age
	newStudent.ProgrammingLanguage=language

	data = append(data, newStudent)

	UpdateDB(data)
}

// Update database
func UpdateDB(students []Student){
	// convert byte data to json
	input,err:=json.Marshal(students)
	if err!=nil{
		log.Fatal(err)
	}
	os.WriteFile("DB/db.json",input,0644)
}