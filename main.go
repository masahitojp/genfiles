package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/stoewer/go-strcase"
)

func main() {
	srcDir := flag.String("srcDir", "main", "D to store the source files")
	testDir := flag.String("testDir", "tests", "File to store the unit tests files")
	fileExtention := flag.String("fileExtention", "py", "file extention")

	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}
	fileName := flag.Args()[0]
	snake := ToSnakeCase(fileName)

	MakeEmptyFile(MakeFileName(*srcDir, snake, *fileExtention))
	MakeEmptyFile(MakeFileName(*testDir, "test_"+snake, *fileExtention))

}

// ToSnakeCase is a base function
func ToSnakeCase(s string) string {
	return strcase.SnakeCase(s)
}

// MakeEmptyFile is a ...
// https://www.golangprograms.com/create-an-empty-file.html
func MakeEmptyFile(s string) {
	emptyFile, err := os.Create(s)
	if err != nil {
		log.Fatal(err)
	}
	emptyFile.Close()
}

// MakeFileName is a ...
func MakeFileName(dir string, fileName string, fileExtention string) string {
	return fmt.Sprintf("%s/%s.%s", dir, fileName, fileExtention)
}
