package main

import (
	"fmt"
//	"io/ioutil"
	"github.com/gi4nks/aggo/utils"
	"github.com/gi4nks/aggo/indexes"
	"github.com/gi4nks/aggo/files"
)

func main() {
	fmt.Printf("hello, world\n")

	// Logging features
	//Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	var logger = utils.NewLogger()
	logger.InitDebug()

	logger.Debug.Println("I have something standard to say")
	logger.Info.Println("Special Information")
	logger.Warning.Println("There is something you need to know about")
	logger.Error.Println("Something has failed")




	inputFile := "c:/tmp/Activities.md"
	outputFile := "c:/tmp/output2.txt"

	/*
	content := Read(inputFile)
	Write(outputFile, content)
	*/

	info3 := indexes.NewInformation()
	fmt.Println("Information info is: ", info3)

	info1 := indexes.NewInformation()
	info1.Phrase = "prova"
	info1.Occurrencies = 1;
	fmt.Println("Information info is: ", info1)

	info2 := indexes.NewInformation()
	info2.Phrase = "test"
	info1.Occurrencies = 1;

	fmt.Println("Information info is: ", info2)

	info2.Increase()
	fmt.Println("Information info is: ", info2)

	info2.Increase()
	fmt.Println("Information info is: ", info2)

	info2.Increase()
	fmt.Println("Information info is: ", info2.Phrase, "->", info2.Occurrencies)

	info2.AddSource(inputFile)
	info2.AddSource(outputFile)

	fmt.Println("Information info is: ", info2)


	/* working with index structure */
	index1 := indexes.NewIndex(*logger)
	index1.Add("prova", inputFile)
	index1.Add("prova", outputFile)

	index1.Add("input3", inputFile)
	index1.Add("input3", inputFile)
	index1.Add("input3", inputFile)

	index1.Add("input1", inputFile)

	fmt.Println("index info is: ", index1)
	/*
	*/
	indexFile := indexes.NewIndex(*logger)

	file_1 := files.NewFile("/Users/gianluca/Projects/golang/aggo/resources/BDD.md", *logger)
	file_1.Open()
	indexFile.Scan(*file_1)
	file_1.Close()

	file_2 := files.NewFile("/Users/gianluca/Projects/golang/aggo/resources/GitHowTo.md", *logger)
	file_2.Open()
	indexFile.Scan(*file_2)
	file_2.Close()

	/*
	indexFile.Serialize("c:/tmp/test.gob")

	indexFile_1 := NewIndex(*logger)
	indexFile_1.Deserialize("c:/tmp/test.gob")


	fmt.Println("----------------------------------")
	fmt.Println("indexFile_1 info is: ", indexFile_1)
	fmt.Println("----------------------------------")
	*/

	/*
	files, _ := ioutil.ReadDir("E:/Work/Post/Activities")
	for _, f := range files {
		fmt.Println(f.Name())
	}
	*/
	var k = indexFile.Search("OUT");
	fmt.Println("keys found into the following files: ", k)
}
