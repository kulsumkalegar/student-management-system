package services

import (
	"STUDENT_MANAGEMENT/models"
	"fmt"
)

func AddStudent(students []models.Student) []models.Student {
	var name string
	var age int
	var id int
	var course string
	var email string
	fmt.Println("enter student ID: ")
	fmt.Scanln(&id)
	fmt.Println("enter your Name: ")
	fmt.Scanln(&name)
	fmt.Println("enter your age : ")
	fmt.Scanln(&age)
	fmt.Println("enter your course: ")
	fmt.Scanln(&course)
	fmt.Println("enter your Email id : ")
	fmt.Scanln(&email)

	student := models.Student{
		ID:     id,
		Name:   name,
		Age:    age,
		Course: course,
		Email:  email,
	}
	students = append(students, student)
	fmt.Println("Added Sucessfully!")
	return students
}

func ViewStudent(students []models.Student) {
	if len(students) <= 0 {
		fmt.Println("List is empty")
	} else {
		for index, student := range students {
			fmt.Println(index+1, student.ID, student.Name, student.Age, student.Course, student.Email)
		}
	}

}

func DeleteStudent(students []models.Student) []models.Student {
	var deletestudent int
	fmt.Println("enter the student Id to delete : ")
	fmt.Scanln(&deletestudent)
	if deletestudent <= 0 {
		fmt.Println("Invalid Student Id")
	} else if deletestudent > len(students) {
		fmt.Println("Invalid")
	} else {
		i := deletestudent - 1
		students = append(students[:i], students[i+1:]...)
		fmt.Println("Deleted sucessfully!")

	}
	return students

}

func SearchStudent(students []models.Student) {
	var StudentName string
	fmt.Println("enter student name : ")
	fmt.Scanln(&StudentName)
	found := false
	for _, student := range students {
		if student.Name == StudentName {
			found = true
			fmt.Println("Student Found!")
			fmt.Println(student.ID, student.Name, student.Age, student.Course, student.Email)
		}
	}
	if found == false {
		fmt.Println("Student not found ")
	}

}

func UpdateStudent(students []models.Student) []models.Student {
	var update int
	var name string
	var age int
	var course string
	var email string
	fmt.Println("enter the student ID to be updated:")
	fmt.Scanln(&update)
	fmt.Println("enter the student name to be updated:")
	fmt.Scanln(&name)
	fmt.Println("enter student age:")
	fmt.Scanln(&age)
	fmt.Println("enter your course: ")
	fmt.Scanln(&course)
	fmt.Println("enter your Email id : ")
	fmt.Scanln(&email)

	if len(students) <= 0 {
		fmt.Println("Empty list")
	} else if update > len(students) {
		fmt.Println("Student Not found")
	} else if update <= 0 {
		fmt.Println("Invalid Student Id")
	} else {
		i := update - 1
		students[i].Name = name
		students[i].Age = age
		students[i].Course = course
		students[i].Email = email
		fmt.Println("Updated Successfully")
	}
	return students
}
