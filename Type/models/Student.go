package models

import "fmt"

type Student struct {
	Id      string
	Name    string
	age     int32
	College string
	active  bool
}

func (s Student) AdmitToCollege(college string) bool {
	s.College = college
	return true
}

// factory pattern for the type construction
func NewStudent(name string, a int32) Student {
	return Student{Name: name, age: a}
}

func (s Student) PrintStudent() {
	fmt.Printf("\n------Student Details-----------------\n")
	fmt.Printf("Name : %s Age : %d College : %s", s.Name, s.age, s.College)
}
