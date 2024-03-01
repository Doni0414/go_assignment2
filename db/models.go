package db

type Student struct {
	Id           uint `gorm:"primaryKey"`
	FullName     string
	Age          uint
	City         string
	Courses      []Course
	DepartmentId uint
}

type Course struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	Students     []Student
	DepartmentId uint
	InstructorId uint
}

type Department struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Students    []Student
	Courses     []Course
	Instructors []Instructor
}

type Instructor struct {
	Id           uint `gorm:"primaryKey"`
	FullName     string
	Age          uint
	DepartmentId uint
	Courses      []Course
}

type Enrollment struct {
	StudentId uint
	CourseId  uint
}
