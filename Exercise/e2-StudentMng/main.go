package main

type Student struct {
	Name       string
	USN        string
	Sunbject   string
	FatherName string
}
type StudenteService interface {
	CreateStudent(student Student) error
	GetStudent(usn string) (Student, error)
}
func StudentRepo(s StudenteService)  {
	
}
func (stu *Student) insert() {

}
