package main

/*

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	DefaultBufferSize = 100
)

func printLines(path string) {
	fmt.Printf("[File: %s]\n", path)

	fin, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "The file %s does not exist!\n", path)
		return
	}
	defer fin.Close()

	bufReader := bufio.NewReader(fin)
	bytes := make([]byte, DefaultBufferSize)
	for line, isPrefix, err := bufReader.ReadLine(); err != io.EOF; line,
	isPrefix, err = bufReader.ReadLine() {
		bytes = append(bytes, line...)

		if !isPrefix {
			fmt.Printf("Lines: %s (error %v)\n", string(bytes), err)
			bytes = bytes[:0]
		}
	}
}

func main() {
	fmt.Println("Reading files...")

	flag.Parse()
	for _, path := range flag.Args() {
		printLines(path)
	}
}
*/
