package main

import (
    "os"
    "fmt"
    "path/filepath"
    "html/template"
)

var CurrentDir string = "/dev/test/spg"

func read_dirs() {
    dirs, _ := os.ReadDir(filepath.Join(CurrentDir, "content"))
    for i, j := range dirs {
        fmt.Printf("%d, %#v\n", i, j.Name());
    }
}

type HtmlDataIndex struct {
    Title string
    Menus []string
}

func main() {
    t, err := template.ParseFiles(filepath.Join(CurrentDir, "layout/index.html"))
    if err != nil {
        panic(err)
    }
    data := HtmlDataIndex { "title 1", []string{"Index", "Content", "Flags"} }
    err = t.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }
}
