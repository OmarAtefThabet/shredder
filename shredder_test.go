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
    // Write some data to the temporary file
    testData := []byte("This is some test data.")
    if _, err := tempFile.Write(testData); err != nil {
        t.Fatalf("Failed to write data to temporary file: %v", err)
    }
    tempFile.Close()
    // Test the File function
    config := Config{Iterations: 3, Remove: false}
    if err := config.File(tempFile.Name()); err != nil {
        t.Fatalf("File test failed: %v", err)
    }
}

func TestShredNonExistentFile(t *testing.T) {
    err := Shred("/nonexistent/file")
    if err == nil {
        t.Error("Expected an error when shredding a non-existent file, but got none.")
    }
}

func TestShredEmptyFile(t *testing.T) {
    // Create a temporary empty file for testing
    tempFile, err := ioutil.TempFile("", "empty_file_test")
    if err != nil {
        t.Fatalf("Failed to create temporary file: %v", err)
    }
    defer tempFile.Close()
    defer os.Remove(tempFile.Name())

    // Test the Shred function on an empty file
    if err := Shred(tempFile.Name()); err != nil {
        t.Fatalf("Shred test on an empty file failed: %v", err)
    }
}

func TestFileRemoveTrue(t *testing.T) {
    // Test with Remove set to true
    config := Config{Iterations: 3, Remove: true}
    // ... create a temporary file and write data as before
    if err := config.File(tempFile.Name()); err != nil {
        t.Fatalf("File test with Remove=true failed: %v", err)
    }
}

func TestFileNonExistentFile(t *testing.T) {
    config := Config{Iterations: 3, Remove: true}
    err := config.File("/nonexistent/file")
    if err == nil {
        t.Error("Expected an error when working with a non-existent file, but got none.")
    }
}