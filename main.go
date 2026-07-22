package main

import (
	"STUDENT_MANAGEMENT/services"
	"STUDENT_MANAGEMENT/storage"
	"fmt"
)

func menu() {
	fmt.Println("=====Students Management System=====")

	fmt.Println("1.Add Student")
	fmt.Println("2.View Students")
	fmt.Println("3. Delete Student")
	fmt.Println("4. Search Student")
	fmt.Println("5. Update Student")
	fmt.Println("6. Exit")
}

func main() {
	students, err := storage.LoadStudents()
	if err != nil {
		fmt.Println("No existing books found. Starting with empty library.")
	}
	// id := uuid.New()
	for {
		menu()
		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			students = services.AddStudent(students)
			err = storage.SaveStudents(students)
			if err != nil {
				fmt.Println(err)
			}

		case 2:
			services.ViewStudent(students)
		case 3:
			students = services.DeleteStudent(students)
			err = storage.SaveStudents(students)
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			services.SearchStudent(students)
		case 5:
			students = services.UpdateStudent(students)
			err = storage.SaveStudents(students)
			if err != nil {
				fmt.Println(err)
			}

		case 6:
			fmt.Println("GoodBye!")
			return
		default:
			fmt.Println("Invalid")
			return
		}

	}

}
