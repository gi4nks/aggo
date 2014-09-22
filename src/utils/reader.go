package utils

import (
	"io/ioutil"
)

func Read(inputFile string) []byte {
	// read whole the file
	b, err := ioutil.ReadFile(inputFile)
	if err != nil { panic(err) }

	return b
}

func OpenWrite(inputFile string, outputFile string) {
	// read whole the file
	b, err := ioutil.ReadFile(inputFile)
	if err != nil { panic(err) }

	// write whole the body
	err = ioutil.WriteFile(outputFile, b, 0644)
	if err != nil { panic(err) }
}
