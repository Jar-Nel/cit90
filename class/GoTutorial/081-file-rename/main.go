package main

import (
	"fmt"
    "os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk("c:/temp/test", visit)
}

func visit(path string, f os.FileInfo, err error) error {
	if name := f.Name(); strings.HasSuffix(name, ".txt"){
		dir := filepath.Dir(path)

		newname := strings.Replace(name, " something.txt", ".txt", 1)
        newpath := filepath.Join(dir, newname)
        fmt.Printf("mv %q %q\n", path, newpath)
        os.Rename(path, newpath)
	}
	return nil
}