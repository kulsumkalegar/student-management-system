package storage

import (
	"STUDENT_MANAGEMENT/models"
	"encoding/json"
	"fmt"
	"os"
)

func SaveStudents(students []models.Student) error {
	data, err := json.Marshal(students)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	file, err := os.Create("students.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	_, err = file.WriteString(string(data))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	file.Close()
	return nil
}

func LoadStudents() ([]models.Student, error) {
	var books []models.Student
	data, err := os.ReadFile("books.json")
	if err != nil {
		return books, nil
	}
	json.Unmarshal(data, &books)
	return books, nil
}
