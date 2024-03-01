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

func UpdateStudent(student Student, studentWithUpdatedFields Student) {
	db.Model(&student).Updates(&studentWithUpdatedFields)
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
