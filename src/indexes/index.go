package indexes

import (
	"os"
	"strings"
	"encoding/gob"
	. "utils"
	. "files"
	. "phonetics"
)

/* Index type */
type Index struct {
	Informations map[string]Information
	logger Logger
}

func NewIndex(log Logger) *Index {
	return &Index{Informations: make(map[string]Information), logger: log}
}

func (index *Index) Add(phrase string, source string) {
	_, ok := index.Informations[phrase]

	if ok {
		current := index.Informations[phrase]

		if current.Phrase == phrase {
			current.Increase()
			current.AddSource(source)
		}

		index.Informations[phrase] = current
	} else {
		info := NewInformation()
		info.Phrase = phrase
		info.Occurrencies = 1
		info.AddSource(source)

		index.Informations[phrase] = *info
	}

	/*
	for k := range index.Informations  {
		current := index.Informations[k]

		if current.Phrase == source {
			current.Increase()
			current.AddSource(source)
		}
	}*/
}

func (index *Index) Serialize(fileName string) {

	index.logger.Debug.Println("Serializing index")
	// Create a file for IO
	encodeFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// Since this is a binary format large parts of it will be unreadable
	encoder := gob.NewEncoder(encodeFile)

	// Write to the file
	if err := encoder.Encode(index.Informations); err != nil {
		panic(err)
	}
	encodeFile.Close()
}


func (index *Index) Deserialize(fileName string) {

	index.logger.Debug.Println("Deserializing index")
	// Open a RO file
	decodeFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()

	// Create a decoder
	decoder := gob.NewDecoder(decodeFile)

	// Place to decode into
	//index.Informations :=  make(map[string]Information)

	// Decode -- We need to pass a pointer otherwise accounts2 isn't modified
	decoder.Decode(&index.Informations)
}

//= map[string]string

func (index *Index) Scan(file File) {
	index.logger.Debug.Println(">> Scan start")

	scanner := file.NewScanner()

	count := 0
	for scanner.Scan() {
		count ++;
		ucl := strings.ToUpper(scanner.Text())
		//fmt.Println(ucl)
		index.logger.Debug.Println(">> read word ", ucl)

		meta := EncodeMetaphone(ucl)
		index.logger.Debug.Println(">> encoded word ", meta)

		if meta != "" {
			index.Add(meta, file.Name)
			index.logger.Debug.Println(">> Adding word ", meta)
		} else {
			index.logger.Warning.Println(">> Word not added due to null coded result from ", ucl)
		}
	}

	index.logger.Debug.Println(">> Word Count: ", count)

	if err := scanner.Err(); err != nil {
		index.logger.Error.Println(">> Error in scamn ", err)
		os.Exit(1)
	}
	index.logger.Debug.Println(">> Scan completed")
}

func (index *Index) Search(phrase string) []string {

	meta := EncodeMetaphone(phrase)

	if meta != "" {
		index.logger.Debug.Println(">> Encoded value ", meta)
		index.logger.Debug.Println(">> Informtion content ", index.Informations)


		searched, ok := index.Informations[meta]

		if ok {
			index.logger.Info.Println(">> found in index")
			return searched.SourceAsArray()
		} else {
			index.logger.Info.Println(">> no found in index")
			return nil
		}
	} else {
		index.logger.Warning.Println(">> Word not added due to null coded result from ", phrase)
		return nil
	}


}
