package main

import (
	"fmt"

	"github.com/bhanusumanth/go/files/fileutils"
)

func main() {
	filePath := "data/file.txt"
	text, err := fileutils.ReadTextFile(filePath)
	if err != nil {
		fmt.Printf("Error Occured reding file: %v", err)
	} else {
		fmt.Printf("File Contents : %s", text)
	}
	const BLEACH string = "BLEACH"
	contentToWrite := fmt.Sprintf("Ichigo Kurosaki is the main character of the show called %s \n His Zanpakuto name is Zangetsu", BLEACH)
	fmt.Printf("string in bytes : %v of type:  %T", []byte(contentToWrite), []byte(contentToWrite))
	e := fileutils.WriteTextFile("data/write_output.txt", contentToWrite)
	if e != nil {
		fmt.Printf("Error Writing this file : %v", err)
	}
}
