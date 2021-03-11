package main

import (
	"fmt"
	"strconv"

	"github.com/StephenDsouza90/StudentApp/core"
	"github.com/StephenDsouza90/StudentApp/database"
	"github.com/StephenDsouza90/StudentApp/utils"
)

// TODO: Update a student

// TODO: Add a course
// TODO: Update a course
// TODO: Delete a course
// TODO: Get a course
// TODO: Get all courses

// TODO: Assign a student to a course
// TODO: Update a student's course
// TODO: Delete a student's course
// TODO: Get all courses of a student
// TODO: Get all students in a course
// TODO: Get one student's courses
// TODO: Get one course's students

func main() {
	db := database.GetConnection()
	core.CreateTables(db)

	userSelection := getMessageAndUserInput()
	for userSelection != 0 {
		if userSelection == 1 {
			fmt.Printf("\nEnter student name to be added : \n")
			studentName := utils.UserInput()
			core.AddStudent(db, studentName)
		} else if userSelection == 2 {
			fmt.Printf("\nEnter student ID to get details : \n")
			studentID := utils.UserInput()
			intStudentID, _ := strconv.Atoi(studentID)
			core.GetStudent(db, intStudentID)
		} else if userSelection == 3 {
			core.GetAllStudents(db)
		} else if userSelection == 5 {
			fmt.Printf("\nEnter student ID to be deleted: \n")
			studentID := utils.UserInput()
			intStudentID, _ := strconv.Atoi(studentID)
			core.DeleteStudent(db, intStudentID)
		} else {
			fmt.Printf("\nInvalid user input! Please enter a number again.\n")
			fmt.Scanln(&userSelection)
		}
		userSelection = getMessageAndUserInput()
	}
}

func getMessageAndUserInput() int16 {
	var userSelection int16

	message := `

Welcome to the Student App!

1 : Add a student
2 : Get a student
3 : Get all students
4 : Update a student
5 : Delete a student
0 : Exit application

	`

	fmt.Println(message)
	fmt.Scanln(&userSelection)
	return userSelection
}
