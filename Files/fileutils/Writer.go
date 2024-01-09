package fileutils

import (
	"log"
	"os"
)

func WriteTextFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0666)
	if err != nil {
		log.Fatal("Cant Write to file")
		return err
	}
	return nil
}
