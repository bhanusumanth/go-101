package models

type TransferStudent struct {
	Student
	College string
}

func (ts TransferStudent) AdmitToCollege(college string) bool {
	ts.College = college
	return true
}

func NewTransferStudent(name string, college string) TransferStudent {
	ts := TransferStudent{}
	ts.Name = name
	ts.College = college
	return ts

}
