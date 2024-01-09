package fileutils

import "os"

func ReadTextFile(fileName string) (string, error) {
	contentBytes, err := os.ReadFile(fileName)
	if err != nil {
		return "", nil
	}
	return string(contentBytes), nil
}
