package utils

import (
	"io/ioutil"
)

func Write(outputFile string, content []byte) {
	// read whole the file
	// write whole the body
	err := ioutil.WriteFile(outputFile, content, 0644)
	if err != nil { panic(err) }
}

