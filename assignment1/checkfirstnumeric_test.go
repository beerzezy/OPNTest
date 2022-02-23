package assignment1

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCheckFirstNumeric(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	CheckFirstNumeric()
}

func TestCheckFirstNumeric2(t *testing.T) {
	content := []byte("Tom")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	if err := userInput(); err != nil {
		t.Errorf("userInput failed: %v", err)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
