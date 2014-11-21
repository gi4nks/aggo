package renamer

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var flagPath = flag.String("path", "", "path to traverse in search of png files.")

func visit(path string, f os.FileInfo, err error) (e error) {
	if filepath.Ext(path) != ".png" {
		return
	}
	log.Println(path)
	return
}

func init() {
	flag.Parse()
}

func visit(path string, f os.FileInfo, err error) (e error) {
	if strings.HasPrefix(f.Name(), "name_") {
		dir := filepath.Dir(path)
		base := filepath.Base(path)
		newname := filepath.Join(dir, strings.Replace(base, "name_", "name1_", 1))
		log.Printf("mv \"%s\" \"%s\"\n", path, newname)
		os.Rename(path, newname)
	}
	return
}

func main() {
	if *flagPath == "" {
		flag.Usage()
		return
	}
	filepath.Walk(*flagPath, visit)
}



