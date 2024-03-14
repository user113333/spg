package main

import (
    "os"
    "fmt"
    "filepath"
)

var CurrentDir := "/dev/test/spg"

func main() {
    dirs, _ := os.ReadDir(filepath.Join(CurrentDir, "content");
    for i, j := range dirs {
        fmt.Printf("%d, %#v\n", i, j.Name());
    }
}
