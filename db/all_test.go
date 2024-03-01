package db

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	log.Println("setup suite")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("There's no .env file in directory! %v", err)
	}
	host := os.Getenv("TEST_DB_HOST")
	user := os.Getenv("TEST_DB_USER")
	password := os.Getenv("TEST_DB_PASSWORD")
	dbname := os.Getenv("TEST_DB_NAME")
	port := os.Getenv("TEST_DB_PORT")
	sslmode := os.Getenv("TEST_SSL_MODE")
	timeZone := os.Getenv("TEST_TIME_ZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timeZone=%s", host, user, password, dbname, port, sslmode, timeZone)

	Connect(dsn)

	fmt.Println(dsn)

	MigrateAllTables()

	return func(tb testing.TB) {
		log.Println("teardown suite")

		var tableNames []string
		rows, err := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Rows()
		if err != nil {
			log.Fatalf("Error retrieving table names: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var tableName string
			if err := rows.Scan(&tableName); err != nil {
				log.Fatalf("Error scanning table name: %v", err)
			}
			tableNames = append(tableNames, tableName)
		}

		for _, tableName := range tableNames {
			if err := db.Exec(fmt.Sprintf("DROP TABLE %s CASCADE", tableName)).Error; err != nil {
				log.Printf("Error dropping table %s: %v", tableName, err)
			} else {
				log.Printf("Table %s dropped successfully", tableName)
			}
		}
	}
}

var departmentsInput = []Department{
	{
		Name: "Engineering and Natural Sciences",
	},
	{
		Name: "Pedagogical and humanitarian sciences",
	},
	{
		Name: "Business school",
	},
	{
		Name: "Law and social sciences and humanities",
	},
}

func createDepartments() {
	for _, department := range departmentsInput {
		CreateDepartment(department)
	}
}

func TestDepartmentCreate(t *testing.T) {
	tearDownSuite := setupSuite(t)
	defer tearDownSuite(t)

	createDepartments()

	expected := len(departmentsInput)

	actualDepartments := FindAllDepartments()

	if actual := len(actualDepartments); expected != actual {
		log.Fatalf("After insert expected %d departments. Found %d!", expected, actual)
	}
}

func TestDepartmentUpdate(t *testing.T) {
	tearDownSuite := setupSuite(t)
	defer tearDownSuite(t)

	createDepartments()

	departmentIdToUpdate := 1
	updatedDepartmentName := "Pedagogy"
	department := FindDepartmentById(departmentIdToUpdate)
	departmentUpdatedFields := Department{
		Name: updatedDepartmentName,
	}
	UpdateDepartment(department, departmentUpdatedFields)

	department = FindDepartmentById(departmentIdToUpdate)

	if department.Name != departmentUpdatedFields.Name {
		log.Fatalf("After updating department name expected: %s, but actual: %s", updatedDepartmentName, department.Name)
	}
}

func TestDepartmentDelete(t *testing.T) {
	tearDownSuite := setupSuite(t)
	defer tearDownSuite(t)

	createDepartments()

	departmentIdToDelete := 1
	department := FindDepartmentById(departmentIdToDelete)

	expected := 3
	DeleteDepartment(department)
	actualDepartments := FindAllDepartments()

	if actual := len(actualDepartments); expected != actual {
		log.Fatalf("After deletion one department expected: %d departments, but found: %d", expected, actual)
	}
}

var studentsInput = []Student{
	{
		FullName: "Askar Bekbergen", Age: 20, City: "Almaty", DepartmentId: 1,
	},
	{
		FullName: "Ramazan Mamyrbek", Age: 20, City: "Turkistan", DepartmentId: 1,
	},
	{
		FullName: "Nurdaulet Agabek", Age: 19, City: "Kaskelen", DepartmentId: 2,
	},
	{
		FullName: "Asset Tagvay", Age: 19, City: "New York", DepartmentId: 2,
	},
}

func createStudents() {
	createDepartments()

	for _, student := range studentsInput {
		CreateStudent(student)
	}
}

func TestStudentCreate(t *testing.T) {
	tearDownSuite := setupSuite(t)
	defer tearDownSuite(t)

	createStudents()

	expectedStudentsSize := len(studentsInput)
	actualStudents := FindAllStudents()

	if actualStudentsSize := len(actualStudents); actualStudentsSize != expectedStudentsSize {
		log.Fatalf("After insert expected %d students. Found %d!", expectedStudentsSize, actualStudentsSize)
	}
}

func TestFindAllStudentsByDepartmentId(t *testing.T) {
	tearDownSuite := setupSuite(t)
	defer tearDownSuite(t)

	createStudents()

	departmentId := 1
	expected := 2
	actual := len(FindAllStudentsByDepartmentId(uint(departmentId)))

	if actual != expected {
		log.Fatalf("The number of students studying at departmentId: %d is expected to be %d, but found %d!", departmentId, expected, actual)
	}
}

func TestFindStudentsByAge(t *testing.T) {
	tearDownSuite := setupSuite(t)
	defer tearDownSuite(t)

	createStudents()

	age := 19
	students := FindStudentsByAge(age)
	expected := 2

	if actual := len(students); actual != expected {
		log.Fatalf("The number of students with age %d is expected to be %d, but found %d!", age, expected, actual)
	}
}

func TestStudentDelete(t *testing.T) {
	tearDownSuite := setupSuite(t)
	defer tearDownSuite(t)

	createStudents()

	studentId := 1
	student := FindStudentById(studentId)
	DeleteStudent(student)

	expected := 3

	students := FindAllStudents()

	if actual := len(students); actual != expected {
		log.Fatalf("The number of students after deletion studentId %d is expected to be %d, but found %d!", studentId, expected, actual)
	}
}

var instructorsInput = []Instructor{
	{
		FullName: "Azamat Serek", Age: 28, DepartmentId: 1,
	},
	{
		FullName: "Nurbol Sabitov", Age: 25, DepartmentId: 1,
	},
	{
		FullName: "Alisher Duzmagambetov", Age: 24, DepartmentId: 2,
	},
	{
		FullName: "Sufyan Mustafa", Age: 30, DepartmentId: 1,
	},
}

func createInstructors() {
	createDepartments()

	for _, instructor := range instructorsInput {
		CreateInstructor(instructor)
	}
}

func TestCreateInstructor(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createInstructors()

	expected := 4
	actualInstructors := FindAllInstructors()

	if actual := len(actualInstructors); actual != expected {
		log.Fatalf("After insert expected %d instructors. Found %d!", expected, actual)
	}
}

func TestFindInstructorById(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createInstructors()

	id := 4
	expected := "Sufyan Mustafa"
	instructor := FindInstructorById(id)

	if actual := instructor.FullName; expected != actual {
		log.Fatalf("The instructors name with id %d is expected to be %s, but found %s!", id, expected, actual)
	}
}

func TestUpdateInstructor(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createInstructors()

	id := 3
	instructor := FindInstructorById(id)

	expectedFullName := "Alisher"
	expectedAge := 26
	instructorWithUpdatedFields := Instructor{
		FullName: expectedFullName, Age: uint(expectedAge),
	}

	UpdateInstructor(instructor, instructorWithUpdatedFields)

	instructor = FindInstructorById(id)
	if instructor.FullName != expectedFullName || instructor.Age != uint(expectedAge) {
		log.Fatalf("The instructor after update is expected to be with name: %s and age: %d, but found name: %s, age: %d", expectedFullName, expectedAge, instructor.FullName, instructor.Age)
	}
}

func TestDeleteInstructor(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createInstructors()

	id := 2
	instructor := FindInstructorById(id)

	DeleteInstructor(instructor)

	expected := 3
	instructors := FindAllInstructors()

	if actual := len(instructors); actual != expected {
		log.Fatalf("After deletion one instructor instructor's size expected to be %d, but found %d", expected, actual)
	}
}

var coursesInput = []Course{
	{
		Name: "The Go programming language", DepartmentId: 1, InstructorId: 1,
	},
	{
		Name: "Frontend Development", DepartmentId: 1, InstructorId: 2,
	},
	{
		Name: "The Virtualization", DepartmentId: 1, InstructorId: 1,
	},
	{
		Name: "The Fundamentals of Programming", DepartmentId: 1, InstructorId: 3,
	},
	{
		Name: "Culturology", DepartmentId: 4, InstructorId: 4,
	},
}

func createCourses() {
	createInstructors()
	for _, course := range coursesInput {
		CreateCourse(course)
	}
}

func TestCreateCourse(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createCourses()

	expected := 5
	actualCourses := FindAllCourses()

	if actual := len(actualCourses); actual != expected {
		log.Fatalf("After insert expected %d courses. Found %d!", expected, actual)
	}
}

func TestFindAllCoursesByInstructorId(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createCourses()

	expected := 2

	instructorId := 1
	instructor := FindInstructorById(instructorId)
	courses := FindAllCoursesByInstructorId(uint(instructorId))

	if actual := len(courses); expected != actual {
		log.Fatalf("Expected that %s teaches %d courses, but found %d!", instructor.FullName, expected, actual)
	}
}

func TestFindCourseById(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createCourses()

	expected := "The Virtualization"
	id := 3
	course := FindCourseById(id)

	if actual := course.Name; expected != actual {
		log.Fatalf("The expected name of course is %s, but it's %s", expected, actual)
	}
}

func TestUpdateCourse(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createCourses()

	id := 4
	course := FindCourseById(id)

	expectedCourse := Course{
		Name: "Server Administration", InstructorId: 2,
	}

	UpdateCourse(course, expectedCourse)

	course = FindCourseById(id)

	if course.Name != expectedCourse.Name || course.InstructorId != expectedCourse.InstructorId {
		log.Fatalf("The expected course name is %s and instructor id is %d, but it's %s and %d", expectedCourse.Name, expectedCourse.InstructorId, course.Name, course.InstructorId)
	}
}

func TestDeleteCourse(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createCourses()

	id := 3
	course := FindCourseById(id)

	expected := len(coursesInput) - 1

	DeleteCourse(course)

	courses := FindAllCourses()

	if actual := len(courses); actual != expected {
		log.Fatalf("After deletion expected %d courses, but found %d!", expected, actual)
	}
}

var enrollments = []Enrollment{
	{
		StudentId: 1, CourseId: 2,
	},
	{
		StudentId: 1, CourseId: 3,
	},
	{
		StudentId: 3, CourseId: 2,
	},
	{
		StudentId: 3, CourseId: 4,
	},
	{
		StudentId: 3, CourseId: 5,
	},
	{
		StudentId: 2, CourseId: 1,
	},
	{
		StudentId: 2, CourseId: 2,
	},
	{
		StudentId: 2, CourseId: 3,
	},
}

func createEnrollments() {
	createCourses()
	for _, student := range studentsInput {
		CreateStudent(student)
	}

	for _, enrollment := range enrollments {
		EnrollStudentForCourse(enrollment.StudentId, enrollment.CourseId)
	}
}
func TestEnrollStudentForCourse(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createEnrollments()

	id := 3
	expected := 3

	courses := GetStudentEnrolledCoursesByStudentId(uint(id))

	if actual := len(courses); actual != expected {
		log.Fatalf("The student with id %d should be enrolled for %d courses, but found %d!", id, expected, actual)
	}
}

func TestStudentCountForEachDepartment(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createStudents()

	expectedCounts := make([]uint, len(departmentsInput))

	for _, student := range studentsInput {
		expectedCounts[student.DepartmentId-1] += 1
	}

	actualCounts := GetStudentCountForEachDepartment()

	for i := 0; i < len(actualCounts); i++ {
		idx := actualCounts[i].Id - 1
		if expectedCounts[idx] != actualCounts[i].StudentCount {
			log.Fatalf("The %s department should contain %d students, but found %d!", actualCounts[i].Name, expectedCounts[idx], actualCounts[i].StudentCount)
		}
	}
}

func TestGetStudentsOfInstructor(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	createEnrollments()

	instructorId := 1

	courses := FindAllCoursesByInstructorId(uint(instructorId))

	expectedStudentName := map[string]struct{}{}

	for _, course := range courses {
		students := GetCourseEnrolledStudentsByCourseId(course.Id)
		for _, student := range students {
			expectedStudentName[student.FullName] = struct{}{}
		}
	}

	actualStudents := GetStudentsOfInstructor(uint(instructorId))

	for _, student := range actualStudents {
		_, exists := expectedStudentName[student.FullName]
		if !exists {
			log.Fatalf("Instructor with id %d doesn't teach for %s!", instructorId, student.FullName)
		}
	}
}
