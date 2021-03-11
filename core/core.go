package core

import (
	"fmt"

	"github.com/StephenDsouza90/StudentApp/models"

	"github.com/jinzhu/gorm"
)

// CreateTables : ...
func CreateTables(db *gorm.DB) {
	db.Exec("PRAGMA foreign_keys = ON;")
	db.AutoMigrate(models.Students{}, models.Courses{}, models.Assignments{})
}

// AddStudent : ...
func AddStudent(db *gorm.DB, studentName string) {
	user := models.Students{StudentName: studentName}
	db.Create(&user)
	fmt.Printf("\nUser added to the database. User ID : %v \n", user.ID)
}

// GetStudent : ...
func GetStudent(db *gorm.DB, studentID int) {
	var student models.Students
	db.First(&student, studentID)
	fmt.Printf("\nStudent ID : %v\n", student.ID)
	fmt.Printf("Student Name : %s\n", student.StudentName)
}

// GetAllStudents : ...
func GetAllStudents(db *gorm.DB) {
	var student models.Students
	rows, _ := db.Model(&student).Rows()
	for rows.Next() {
		db.ScanRows(rows, &student)
		fmt.Printf("Student ID: %v \n", student.ID)
		fmt.Printf("Student Name : %s\n\n", student.StudentName)
	}
}

// DeleteStudent : ...
func DeleteStudent(db *gorm.DB, studentID int) {
	// FIXME: student.ID is returning 0
	var student models.Students
	db.Delete(&student, studentID)
	fmt.Printf("\nStudent ID deleted : %v\n", student.ID)
}
