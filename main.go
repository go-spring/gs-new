package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed tmpl.zip
var b []byte

func main() {
	fmt.Println("new")
	size := int64(len(b))
	r, err := zip.NewReader(bytes.NewReader(b), size)
	if err != nil {
		panic(err)
	}
	for _, f := range r.File {
		fmt.Println(f.Name)
	}
}
