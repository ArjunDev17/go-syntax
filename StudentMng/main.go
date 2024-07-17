package main

import "fmt"

// Student struct to hold student data
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

// StudentService interface to define CRUD operations
type StudentService interface {
	CreateStudent(student Student) error
	GetStudent(id int) (Student, error)
	UpdateStudent(student Student) error
	DeleteStudent(id int) error
}

// StudentRepository struct to implement StudentService interface
type StudentRepository struct {
	students map[int]Student
}

// NewStudentRepository creates a new instance of StudentRepository
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		students: make(map[int]Student),
	}
}

// CreateStudent adds a new student to the repository
func (repo *StudentRepository) CreateStudent(student Student) error {
	if _, exists := repo.students[student.ID]; exists {
		return fmt.Errorf("student with ID %d already exists", student.ID)
	}
	repo.students[student.ID] = student
	return nil
}

// GetStudent retrieves a student by ID from the repository
func (repo *StudentRepository) GetStudent(id int) (Student, error) {
	student, exists := repo.students[id]
	if !exists {
		return Student{}, fmt.Errorf("student with ID %d not found", id)
	}
	return student, nil
}

// UpdateStudent updates an existing student in the repository
func (repo *StudentRepository) UpdateStudent(student Student) error {
	if _, exists := repo.students[student.ID]; !exists {
		return fmt.Errorf("student with ID %d not found", student.ID)
	}
	repo.students[student.ID] = student
	return nil
}

// DeleteStudent removes a student by ID from the repository
func (repo *StudentRepository) DeleteStudent(id int) error {
	if _, exists := repo.students[id]; !exists {
		return fmt.Errorf("student with ID %d not found", id)
	}
	delete(repo.students, id)
	return nil
}

// Function that uses the StudentService interface
func PrintStudentDetails(service StudentService, id int) {
	student, err := service.GetStudent(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Student ID: %d, Name: %s, Age: %d, Grade: %s\n", student.ID, student.Name, student.Age, student.Grade)
}

func main() {
	repo := NewStudentRepository()

	// Create a new student
	student1 := Student{ID: 1, Name: "John Doe", Age: 20, Grade: "A"}
	err := repo.CreateStudent(student1)
	if err != nil {
		fmt.Println("Error creating student:", err)
	}

	// Use the interface to print student details
	PrintStudentDetails(repo, 1)
}
