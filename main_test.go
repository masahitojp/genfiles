package main

import (
	"testing"

	"github.com/spf13/afero"
)

func TestMakeEmptyFile(t *testing.T) {
	AppFs := afero.NewMemMapFs()
	fileName := "./test.py"

	MakeEmptyFile(AppFs, fileName)

	// file exists
	exists, _ := afero.Exists(AppFs, fileName)
	assert(t, exists, true)

	// file is empty
	bytes, _ := afero.ReadFile(AppFs, fileName)
	assert(t, string(bytes), "")
}

func TestMakeEmptyFileAndCreateDir(t *testing.T) {
	AppFs := afero.NewMemMapFs()
	fileName := "path/to/test.py"

	MakeEmptyFile(AppFs, fileName)

	// file exists
	exists, _ := afero.Exists(AppFs, fileName)
	assert(t, exists, true)

	// directory exists
	existsDir, _ := afero.DirExists(AppFs, "path/to")
	assert(t, existsDir, true)

	// file is empty
	bytes, _ := afero.ReadFile(AppFs, fileName)
	assert(t, string(bytes), "")
}

func TestMakeFileName(t *testing.T) {
	assert(t, MakeFileName("src", "main", "go"), "src/main.go")
}

func assert(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
