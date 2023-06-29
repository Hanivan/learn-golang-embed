package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed test/version.txt
var version string

//go:embed test/picture.png
var picture []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("picture_new.png", picture, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirs, _ := path.ReadDir("files")

	for _, entry := range dirs {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
