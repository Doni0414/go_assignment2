package db

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	Id           uint `gorm:"primaryKey"`
	FullName     string
	Age          uint
	City         string
	Courses      []Course `gorm:"many2many:enrollments;constraint:OnDelete:CASCADE;"`
	DepartmentId uint
	CreatedAt    time.Time
}

type Course struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	Students     []Student `gorm:"many2many:enrollments;constraint:OnDelete:CASCADE;"`
	DepartmentId uint
	InstructorId uint
	DeletedAt    gorm.DeletedAt
}

type Department struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Students    []Student    `gorm:"foreignKey:DepartmentId"`
	Courses     []Course     `gorm:"foreignKey:DepartmentId"`
	Instructors []Instructor `gorm:"foreignKey:DepartmentId"`
}

type Instructor struct {
	Id           uint `gorm:"primaryKey"`
	FullName     string
	Age          uint
	DepartmentId uint
	Courses      []Course `gorm:"foreignKey:InstructorId;constraint:OnDelete:SET NULL;"`
	UpdatedAt    time.Time
}

type Enrollment struct {
	StudentId uint
	CourseId  uint
}

func (student *Student) BeforeCreate(tx *gorm.DB) error {
	currentTime := time.Now()
	student.CreatedAt = currentTime
	return nil
}

func (instructor *Instructor) BeforeUpdate(tx *gorm.DB) error {
	instructor.UpdatedAt = time.Now()
	return nil
}
