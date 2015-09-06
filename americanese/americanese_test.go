package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

// files we are going to read, output, or compare against; set up as constants
const (
	inFilename       = "input.txt"
	expectedFilename = "expected.txt"
	actualFilename   = "actual.txt"
)

// test americanise
func TestAmericanize(t *testing.T) {
	// set flags for output, not sure what the numbers allude to
	log.SetFlags(0)
	log.Println("TEST amerianize")

	path, _ := filepath.Split(os.Args[0])
	var inFile, outFile *os.File
	var err error

	// commented out code that was in the book, because since the file is in the same directory the Join
	// was not needed and actually caused problems of adding a _test path.  All commented out Join calls
	// reflect this comment
	//inFilename := filepath.Join(path, inFilename)
	//inFilename := inFilename
	if inFile, err = os.Open(inFilename); err != nil {
		t.Fatal(err)
	}
	defer inFile.Close()

	//outFilename := filepath.Join(path, actualFilename)
	outFilename := actualFilename
	if outFile, err = os.Create(outFilename); err != nil {
		t.Fatal(err)
	}

	defer outFile.Close()
	defer os.Remove(outFilename)

	if err := americanese(inFile, outFile); err != nil {
		t.Fatal(err)
	}

	//compare(outFilename, filepath.Join(path, expectedFilename), t)
	compare(outFilename, expectedFilename, t)

}

func compare(actual, expected string, t *testing.T) {
	if actualBytes, err := ioutil.ReadFile(actual); err != nil {
		t.Fatal(err)
	} else if expectedBytes, err := ioutil.ReadFile(expected); err != nil {
		t.Fatal(err)
	} else {
		if bytes.Compare(actualBytes, expectedBytes) != 0 {
			t.Fatal("actual != expected")
		}
	}
}
