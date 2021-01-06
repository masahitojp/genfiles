package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/stoewer/go-strcase"
)

func main() {
	srcDir := flag.String("srcDir", "main", "File to store the source files")
	testDir := flag.String("testDir", "tests", "File to store the unit tests files")
	fileExtension := flag.String("fileExtention", "py", "file extension")

	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}
	fileName := flag.Args()[0]
	snake := ToSnakeCase(fileName)

	fileNames := []string{
		MakeFileName(*srcDir, snake, *fileExtension),
		MakeFileName(*testDir, "test_"+snake, *fileExtension),
	}
	AppFs := afero.NewOsFs()
	for _, fileName := range fileNames {
		MakeEmptyFile(AppFs, fileName)
		fmt.Println("created: " + fileName)
	}
}

// ToSnakeCase is a base function
func ToSnakeCase(s string) string {
	return strcase.SnakeCase(s)
}

// MakeEmptyFile is a ...
// https://www.golangprograms.com/create-an-empty-file.html
func MakeEmptyFile(fs afero.Fs, s string) {
	dirName := filepath.Dir(s)
	exists, err := afero.DirExists(fs, dirName)
	if !exists && err != nil {
		mkdirErr := fs.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			log.Fatal(mkdirErr)
		}
	}
	emptyFile, err := fs.Create(s)
	if err != nil {
		log.Fatal(err)
	}
	emptyFile.Close()
}

// MakeFileName is a ...
func MakeFileName(dir, fileName, fileExtention string) string {
	return fmt.Sprintf("%s/%s.%s", dir, fileName, fileExtention)
}
