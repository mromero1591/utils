package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// ReadFile reads a file from the provided path and returns its content.
// The path is relative to the root directory of the project.
// It assumes that the path is valid. The function returns the file content
// as a string if successful, otherwise it returns the encountered error.
// Note that the error includes the filepath if reading the file fails.
func ReadFile(file string) (string, error) {
	filepath := CreateFilepath(file)
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to read file from: %s", filepath)
	}

	return string(bytes), nil
}

func WriteToFile(fileName string, input []string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range input {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Failed writing to file: %s", err)
		}
	}
	writer.Flush()
}

func CreateFilepath(filename string) string {
	return filepath.Join(GetRootDirectory(), filename)
}

func GetRootDirectory() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
