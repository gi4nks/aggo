package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"github.com/gi4nks/aggo/utils"
)

const (
	DefaultBufferSize = 100
)

/* Information type*/
type File struct {
	Name string
	isOpened bool
	hasMoreLines bool
	file *os.File
	bufReader *bufio.Reader
	logger utils.Logger
}

func NewFile(n string, log utils.Logger) *File {
	return &File{Name: n, isOpened: false, hasMoreLines: false, logger: log}
}

func (file *File) Open()  {
	file.isOpened = false

	fin, err := os.Open(file.Name)
	utils.CheckAndShow(err, "The file %s does not exist!\n", file.Name)

	file.file = fin
	file.isOpened = true

	// init the buffer reader
	file.bufReader = bufio.NewReader(file.file)
	file.hasMoreLines = true

	file.logger.Debug.Println(">> Opened ", file.Name)
}

func (file *File) Close() {
	file.file.Close()
	file.logger.Debug.Println(">> Closed ", file.Name)
}

func (file *File) ReadLine() string {
	if !file.isOpened {
		fmt.Fprintf(os.Stderr, "The file %s is not opened! Open it first.\n", file.Name)
		return "io.ErrClosedPipe"
	}

	if (!file.hasMoreLines) {
		fmt.Fprintf(os.Stderr, "The file %s has no more lines.\n", file.Name)
		return "io.ErrNoProgress"
	}

	// I can read the next row
	//line, isPrefix, err := file.bufReader.ReadLine()
	fmt.Printf("olaaaaaa\n")

	line, _, err := file.bufReader.ReadLine()

	if (err != io.EOF) {
		file.hasMoreLines = true
	} else {
		file.hasMoreLines = false
	}

	fmt.Printf("Line: %s (error %v)\n", string(line), err)

	// prefix management...
	/*
	if !isPrefix {
		fmt.Printf("Line: %s (error %v)\n", string(line), err)
	}*/

	return string(line)
}

func (file *File) NewScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(file.bufReader)
	scanner.Split(bufio.ScanWords)

	return scanner;
}
