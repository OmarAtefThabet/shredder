package shredder

import (
	"fmt"
    "os"
    "testing"
)

func TestShred(t *testing.T) {
	
	if len(os.Args) != 2 {
        os.Exit(1)
    }
    Input := Config{iterations: 3, remove: false}
    Path := os.Args[1]
    err := Input.File(Path)
    if err != nil {
        t.Fatalf("Error: %v", err)
    } else {
        fmt.Println("File shredded")
    }
}