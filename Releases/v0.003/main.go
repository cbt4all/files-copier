package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	lsSrc := readTextFileLoadToMem("src.txt")
	lsSrc = clearSrc(lsSrc)
	lsDst := readTextFileLoadToMem("dst.txt")

	// Check what is SH1 of the all srouces
	// Copy files one by one
	// Check what is SH1 of the copied files
	// if SH1s dont match, log and, delete the copied file and move on

	var files []file
	totalSize = 0
	totalCoppied = 0
	for _, item := range lsSrc {

		f, err := os.Open(item)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			log.Fatal(err)
		}

		files = append(files, file{
			src:  item,
			dst:  lsDst[0],
			md5:  md5Calculator(f),
			sha1: sha1Calculator(f),
			name: fi.Name(),
			size: fi.Size(),
		})
		totalSize = totalSize + fi.Size()
	}

	for _, item := range files {
		fmt.Println(copyOneFileRWB(item, 8192))
	}

}
