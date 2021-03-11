package models

// Students : Students table
type Students struct {
	ID          int
	StudentName string
}

// Courses : Courses table
type Courses struct {
	ID         int
	CourseName string
}

// Assignments : Assignments table, One student can have many courses
type Assignments struct {
	ID         int
	StudentsID int
	CourseID   int
}
