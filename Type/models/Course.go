package models

import "fmt"

type Duration float32

type Course struct {
	Id   string
	Name string
	Dur  Duration
	Std  Student
}

func (c Course) String() string {
	return fmt.Sprintf("\n ID : %s \t Name : %s \t %f Duration \t %v", c.Id, c.Name, c.Dur, c.Std)
}
