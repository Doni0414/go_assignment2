package db

func CreateStudent(student Student) {
	db.Create(&student)
}

func FindAllStudents() []Student {
	var students []Student
	db.Find(&students)
	return students
}

func FindStudentById(id int) Student {
	student := Student{}
	db.First(&student, id)
	return student
}

func FindAllStudentsByDepartmentId(departmentId uint) []Student {
	var students []Student
	db.Where("department_id = ?", departmentId).Find(&students)
	return students
}

func FindStudentsByAge(age int) []Student {
	var students []Student
	db.Where("age = ?", age).Find(&students)
	return students
}

func GetStudentEnrolledCoursesByStudentId(studentId uint) []Course {
	var student Student
	db.Model(&Student{}).Where("id = ?", studentId).Preload("Courses").First(&student)
	return student.Courses
}

func UpdateStudentAge(student Student, age int) {
	db.Model(&student).Update("Age", age)
}

func DeleteStudent(student Student) {
	db.Delete(&student)
}

// COURSES
func CreateCourse(course Course) {
	db.Create(&course)
}

func FindAllCourses() []Course {
	var courses []Course
	db.Find(&courses)
	return courses
}

func FindAllCoursesByInstructorId(instructorId uint) []Course {
	var courses []Course
	db.Where("instructor_id = ?", instructorId).Find(&courses)
	return courses
}

func FindCourseById(id int) Course {
	course := Course{}
	db.First(&course, id)
	return course
}

func GetCourseEnrolledStudentsByCourseId(courseId uint) []Student {
	var course Course
	db.Model(&Course{}).Where("id = ?", courseId).Preload("Students").First(&course)
	return course.Students
}

func UpdateCourse(course Course, courseWithUpdatedFields Course) {
	db.Model(&course).Updates(&courseWithUpdatedFields)
}

func DeleteCourse(course Course) {
	db.Delete(&course)
}

// DEPARTMENT
func CreateDepartment(department Department) {
	db.Create(&department)
}

func FindAllDepartments() []Department {
	var departments []Department
	db.Find(&departments)
	return departments
}

func FindDepartmentById(id int) Department {
	department := Department{}
	db.First(&department, id)
	return department
}

func UpdateDepartment(department Department, departmentWithUpdatedFields Department) {
	db.Model(&department).Updates(&departmentWithUpdatedFields)
}

func DeleteDepartment(department Department) {
	db.Delete(&department)
}

//Enrollment

func EnrollStudentForCourse(studentId, courseId uint) error {
	tx := db.Begin()

	var student Student
	if err := tx.First(&student, studentId).Error; err != nil {
		tx.Rollback()
		return err
	}

	var course Course
	if err := tx.First(&course, courseId).Error; err != nil {
		tx.Rollback()
		return err
	}

	student.Courses = append(student.Courses, course)

	if err := tx.Save(&student).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Instructor
func CreateInstructor(instructor Instructor) {
	db.Create(&instructor)
}

func FindAllInstructors() []Instructor {
	var instructors []Instructor
	db.Find(&instructors)
	return instructors
}

func FindInstructorById(id int) Instructor {
	instructor := Instructor{}
	db.First(&instructor, id)
	return instructor
}

func UpdateInstructor(instructor Instructor, instructorWithUpdatedFields Instructor) {
	db.Model(&instructor).Updates(&instructorWithUpdatedFields)
}

func DeleteInstructor(instructor Instructor) {
	db.Delete(&instructor)
}

// CUSTOM QUERIES

type APIDepartment struct {
	Id           uint
	Name         string
	StudentCount uint
}

type APICourse struct {
	ID           uint
	Name         string
	StudentCount uint
}

func GetStudentCountForEachDepartment() []APIDepartment {
	var apiDepartments []APIDepartment
	db.Model(&Department{}).Select("departments.id, departments.name, COUNT(*) as student_count").Joins("inner join students on departments.id = students.department_id").Group("departments.id, departments.name").Find(&apiDepartments)
	return apiDepartments
}

func GetStudentsOfInstructor(instructorId uint) []Student {
	var students []Student
	db.Model(&Instructor{}).Select("students.id, students.full_name, students.age, students.city, students.department_id, students.created_at").Where("instructor_id = ?", instructorId).Joins("inner join courses on instructors.id = courses.instructor_id inner join enrollments on courses.id = enrollments.course_id inner join students on enrollments.student_id = students.id").Group("students.id, students.full_name, students.age, students.city, students.department_id, students.created_at").Find(&students)
	return students
}
