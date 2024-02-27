package db

type Student struct {
	Id       uint `gorm:"primaryKey"`
	FullName string
	Age      uint
	City     string
}

type Course struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}

type Department struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}

type Enrollment struct {
	Id        uint `gorm:"primaryKey"`
	StudentId uint
	CourseId  uint
}

type Instructor struct {
	Id       uint `gorm:"primaryKey"`
	FullName string
	Age      uint
}
