package main

import (
	"fmt"

	"github.com/bhanusumanth/go/types/models"
)

func main() {
	var ichigo models.Student
	ichigo.Name = "Ichigo Kurosaki"
	ichigo.PrintStudent()
	rukia := models.Student{}
	rukia.PrintStudent()
	renji := models.NewStudent("Renji", 25)
	renji.PrintStudent()
	soulReaper := models.Course{Name: "Shinigami", Dur: 99999.99999}
	fmt.Print(soulReaper)

	byakuya := models.NewTransferStudent("Byakuya", "Soul Society")
	fmt.Println(byakuya)

	var students [4]models.Admittable
	students[0] = ichigo
	students[1] = rukia
	students[2] = renji
	students[3] = byakuya

	for _, std := range students {
		fmt.Println(std)
	}
}
