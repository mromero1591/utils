// files_test.go
package files

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	testFile := "testfile.txt"
	testContent := "Test content"

	// Write some data to the test file
	f, err := os.Create(testFile)
	if err != nil {
		t.Fatal(err)
	}

	_, err = f.WriteString(testContent)
	if err != nil {
		t.Fatal(err)
	}

	f.Close()

	// Call the function under test
	result, err := ReadFile(fmt.Sprintf("/files/%s", testFile))

	if err != nil {
		cleanUpTestFile(t, testFile)
		t.Fatal(err)
	}

	// Test that it reads the data back correctly
	if result != testContent {
		t.Errorf("Expected %s, got %s", testContent, result)
	}

	// Afterwards, clean up and remove the test file
	cleanUpTestFile(t, testFile)
}

func cleanUpTestFile(t *testing.T, testFile string) {
	err := os.Remove(testFile)
	if err != nil {
		t.Fatal(err)
	}
}
