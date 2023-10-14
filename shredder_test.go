package shredder

import (
    "testing"
    "os"
    "io/ioutil"
)

func TestShred(t *testing.T) {
    // Create a temporary file for testing
    tempFile, err := ioutil.TempFile("", "shred_test")
    if err != nil {
        t.Fatalf("Failed to create temporary file: %v", err)
    }
    defer tempFile.Close()
    defer os.Remove(tempFile.Name())

    // Write some data to the temporary file
    testData := []byte("This is some test data.")
    if _, err := tempFile.Write(testData); err != nil {
        t.Fatalf("Failed to write data to temporary file: %v", err)
    }

    // Test the Shred function
    if err := Shred(tempFile.Name()); err != nil {
        t.Fatalf("Shred test failed: %v", err)
    }
    // You can add assertions to verify the file was properly shredded here.
}

func TestFile(t *testing.T) {
    // Create a temporary file for testing
    tempFile, err := ioutil.TempFile("", "file_test")
    if err != nil {
        t.Fatalf("Failed to create temporary file: %v", err)
    }
    defer tempFile.Close()
    defer os.Remove(tempFile.Name())

    // Write some data to the temporary file
    testData := []byte("This is some test data.")
    if _, err := tempFile.Write(testData); err != nil {
        t.Fatalf("Failed to write data to temporary file: %v", err)
    }

    // Test the File function
    config := Config{Iterations: 3, Remove: true}
    if err := config.File(tempFile.Name()); err != nil {
        t.Fatalf("File test failed: %v", err)
    }
    // You can add assertions to verify the file was properly shredded and removed here.
}