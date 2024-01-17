package models

type Admittable interface {
	AdmitToCollege(college string) bool
}
