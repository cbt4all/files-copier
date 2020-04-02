package main

import "fmt"

func main() {

	lsSrc := readTextFileLoadToMem("src.txt")
	dst := readTextFileLoadToMem("dst.txt")

	// Check what is SH1 of the all srouces
	// Copy files one by one
	// Check what is SH1 of the copied files
	// if SH1s dont match, log and, delete the copied file and move on
	for _, item := range lsSrc {
		fmt.Println(copyOneFileRWB(item, dst[0], 8192))
	}
}
