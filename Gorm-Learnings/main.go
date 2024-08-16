package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// --- students table ---

// type Student struct {
// 	StudentId int
// 	Name      string
// 	Major     string
// }

// --- employees table ---

// type Employee struct {
// 	value        int       `gorm:"primaryKey"`
// 	FirstName    string    `gorm:"size:20"`
// 	LastName     string    `gorm:"size:20"`
// 	BirthDate    time.Time `gorm:"type:date"`
// 	Gender       string    `gorm:"size:2"`
// 	Salary       float64   `gorm:"type:decimal(10,2)"`
// 	SuperVisorId int
// 	BranchId     int
// }

type Employee struct {
	// Value int `gorm:"column:emp_id"`
	EmpId        int
	FirstName    string
	LastName     string
	BirthDate    time.Time
	Gender       string
	Salary       float64
	SuperVisorId int
	BranchId     int
}

type Result struct {
	BranchId    int
	TotalSalary float64
}

func main() {
	// dsn := "host=192.168.2.5 user=ST795 dbname=karg password=000ST79544 sslmode=disable"
	// dsn := "deepak:root@tcp(130.344.4.6:3306)/rato?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := "ST795:000ST79544@tcp(192.168.2.5)/karg?charset=utf8mb4&parseTime=True&loc=Local"
	db, lerr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if lerr != nil {
		panic("ERROR(MAIN001): Failed to connect database")
	}

	// --- students table ---

	// var students []Student

	// sql: select * from students;
	// gorm: db.Find(&students)

	// sql: select name, major from student;
	// gorm: db.Select("name", "major").Find(&students)

	// sql: select * from students s
	// order by s.major, s.student_id desc;
	// gorm: db.Order("major, student_id desc").Find(&students)

	// sql: select * from students limit 2;
	// gorm: db.Limit(2).Find(&students)

	// sql: select * from students order by student_id desc limit 2;
	// gorm: db.Order("student_id desc").Limit(2).Find(&students)

	// sql: select * from students where major="Biology";
	// gorm: db.Where("major=?", "Biology").Find(&students)

	// sql: select name, major from students where major="chemistry" or major="Biology";
	// gorm: db.Select("name, major").Where("major=?", "chemistry").Or("major=?", "biology").Find(&students)

	// sql: select name, major from students where major<>"biology";
	// gorm: db.Select("name, major").Where("major <> ?", "biology").Find(&students)

	// sql: select name, major from students where student_id >3;
	// gorm: db.Select("name, major").Where("student_id > ?", 3).Find(&students)

	// sql: select name, major from students where student_id >= 3 and name <> "sam";
	// gorm: db.Select("name, major").Where("student_id >= ?", 3).Where("name <> ?", "sam").Find(&students)

	// sql: select * from students where name in ('cate', 'mike', 'joseph');
	// gorm: db.Where("name IN ?", []string{"cate", "mike", "joseph"}).Find(&students)

	// sql: select * from students where major in ("biology", "Chemistry") and student_id >2;
	// gorm: db.Where("major IN ?", []string{"biology", "chemistry"}).Where("student_id > ?", 2).Find(&students)

	// for _, stud := range students {
	// 	fmt.Printf("student_id: %d, Name: %s, Major:%s\n", stud.StudentId, stud.Name, stud.Major)
	// }

	// --- employees table ---

	// AutoMigrating Employee struct

	// err := db.AutoMigrate(&Employee{})
	// if err != nil {
	// 	panic("ERROR (MAIN002): Failed to AutoMigrate the Employee struct")
	// }

	// // Inserting a new record in Employees table

	// emp_1 := Employee{
	// 	FirstName: "David",
	// 	LastName:  "Wallace",
	// 	BirthDate: time.Date(1967, 11, 17, 0, 0, 0, 0, time.UTC),
	// 	Gender:    "M",
	// 	Salary:    250000,
	// 	BranchId:  1,
	// }
	// result := db.Create(&emp_1)
	// if result.Error != nil {
	// 	fmt.Println("ERROR (MAIN003): Failed to create a record in Employees table")
	// } else {
	// 	fmt.Println("New emp_id:", emp_1.EmpId)
	// }

	// Inserting more records can be done in two ways
	// 1. Use a loop over a set of objects
	// 2. Use Batch insert - Insert the set of objects at a single time

	// Inserting using a loop

	// employeesToAdd := []Employee{
	// 	{
	// 		FirstName:    "Jan",
	// 		LastName:     "Levinson",
	// 		BirthDate:    time.Date(1961, 05, 11, 0, 0, 0, 0, time.UTC),
	// 		Gender:       "F",
	// 		Salary:       110000,
	// 		SuperVisorId: 100,
	// 		BranchId:     1,
	// 	},
	// 	{
	// 		FirstName:    "Michael",
	// 		LastName:     "Scott",
	// 		BirthDate:    time.Date(1964, 03, 15, 0, 0, 0, 0, time.UTC),
	// 		Gender:       "M",
	// 		Salary:       75000,
	// 		SuperVisorId: 100,
	// 		BranchId:     2,
	// 	},
	// 	{
	// 		FirstName:    "Angela",
	// 		LastName:     "Martin",
	// 		BirthDate:    time.Date(1971, 06, 25, 0, 0, 0, 0, time.UTC),
	// 		Gender:       "F",
	// 		Salary:       63000,
	// 		SuperVisorId: 102,
	// 		BranchId:     2,
	// 	},
	// 	{
	// 		FirstName:    "Kelly",
	// 		LastName:     "Kapoor",
	// 		BirthDate:    time.Date(1980, 02, 05, 0, 0, 0, 0, time.UTC),
	// 		Gender:       "F",
	// 		Salary:       55000,
	// 		SuperVisorId: 102,
	// 		BranchId:     2,
	// 	},
	// }

	// for _, employee := range employeesToAdd {
	// 	result := db.Create(&employee)
	// 	if result.Error != nil {
	// 		fmt.Println("ERROR (MAIN004): Failed to insert record using a loop")
	// 	} else {
	// 		fmt.Println("emp_id: ", employee.EmpId)
	// 	}
	// }

	//	Use Batch insert - Insert the set of objects at a single time

	// employeesToAdd := []Employee{
	// 	{
	// 		FirstName:    "Stanley",
	// 		LastName:     "Hudson",
	// 		BirthDate:    time.Date(1958, 02, 19, 0, 0, 0, 0, time.UTC),
	// 		Gender:       "M",
	// 		Salary:       69000,
	// 		SuperVisorId: 102,
	// 		BranchId:     2,
	// 	},
	// 	{
	// 		FirstName:    "Josh",
	// 		LastName:     "Porter",
	// 		BirthDate:    time.Date(1969, 9, 05, 0, 0, 0, 0, time.UTC),
	// 		Gender:       "M",
	// 		Salary:       78000,
	// 		SuperVisorId: 100,
	// 		BranchId:     3,
	// 	},
	// 	{
	// 		FirstName:    "Andy",
	// 		LastName:     "Bernard",
	// 		BirthDate:    time.Date(1973, 07, 22, 0, 0, 0, 0, time.UTC),
	// 		Gender:       "M",
	// 		Salary:       65000,
	// 		SuperVisorId: 106,
	// 		BranchId:     3,
	// 	},
	// }

	// result := db.Create(&employeesToAdd)
	// if result.Error != nil {
	// 	fmt.Println("ERROR (MAIN005): Failed to Batch Insert")
	// } else {
	// 	for _, employee := range employeesToAdd {
	// 		fmt.Println("Inserted employee ID:", employee.EmpId)
	// 	}
	// }

	// var employees []Employee

	// sql: select * from employees
	// db.Find(&employees)

	// sql: select * from employees e order by salary desc;
	// gorm: db.Order("salary desc").Find(&employees)/

	// sql: select * from employees e order by gender, first_name ,last_name
	// gorm: db.Order("gender, first_name, last_name").Find(&employees)

	// sql: select * from employees e limit 5;
	// gorm: db.Limit(5).Find(&employees)

	// for _, employee := range employees {
	// 	// fmt.Printf("EmpId: %d, FirstName: %s, LastName: %s, Salary: %.2f \n", employee.Value, employee.FirstName, employee.LastName, employee.Salary)
	// 	fmt.Printf("EmpId: %d, FirstName: %s, LastName: %s, Salary: %.2f \n", employee.EmpId, employee.FirstName, employee.LastName, employee.Salary)
	// }

	// sql: select first_name as forename, last_name as surname from employees e
	// gorm: db.Select("first_name AS forename, last_name AS surname").Find(&employees)

	// sql: select distinct gender from employees e;
	// db.Distinct("gender").Find(&employees)

	// sql: select count(emp_id) from employees e where gender ='F' and birth_date > '1970-01-01';
	// var count int64
	// db.Model(&Employee{}).Where("gender = ? AND birth_date > ?", "F", "1970-01-01").Count(&count)
	// db.Model(&Employee{}).Count(&count)
	// fmt.Println("count:", count)

	// sql: select avg(salary) from employees e ;
	// var avgSalary float64
	// db.Model(&Employee{}).Select("Avg(salary)").Row().Scan(&avgSalary)
	// fmt.Println("Average salary: ", avgSalary)

	// sql: select avg(salary) from employees e where gender='M';
	// var avgMenSalary float64
	// db.Model(&Employee{}).Select("Avg(salary)").Where("gender = ?", "M").Row().Scan(&avgMenSalary)
	// fmt.Println("avgMenSalary: ", avgMenSalary)

	// sql: select sum(salary) from employees e;
	// var sumSalary float64
	// db.Model(&Employee{}).Select("Sum(salary)").Row().Scan(&sumSalary)
	// fmt.Println("sumSalary: ", sumSalary)

	// sql: select branch_id id, sum(salary) from employees e group by branch_id;
	var results []Result
	db.Model(&Employee{}).
		Select("branch_id, Sum(salary) AS TotalSalary").
		Group("branch_id").
		Scan(&results)

	for _, result := range results {
		fmt.Printf("BranchID: %d, Total Salary: %.2f\n", result.BranchId, result.TotalSalary)
	}

}
