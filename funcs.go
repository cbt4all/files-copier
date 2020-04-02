package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func copyOneFileRWB(src, dst string, BUFFERSIZE int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Size of the Source in Bytes
	sourceSize := sourceFileStat.Size()
	sourceName := sourceFileStat.Name()

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	fmt.Printf("source MD5 %x\n", md5Calculator(source))
	fmt.Printf("source SHA1 %x\n", sha1Calculator(source))

	dstf := dst + sourceName
	_, err = os.Stat(dstf)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dstf)
	}

	destination, err := os.Create(dstf)
	if err != nil {
		//return err
		log.Fatal(err)
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	var tw int64

	for {
		n1, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n1 == 0 {
			break
		}

		n2, err := destination.Write(buf[:n1])
		if err != nil {
			return err
		}

		if n1 != n2 {
			log.Fatal("Number of bytes read don't match what is copied. Some IO problems!...")
		}

		tw = tw + int64(n2)
		fmt.Print("\r", sourceName, " is being copied...", (100*int64(tw))/sourceSize, "%")
	}

	fmt.Printf("\ndestination MD5 %x\n", md5Calculator(destination))
	fmt.Printf("destination SHA1 %x\n\n", sha1Calculator(destination))

	return err
}

func sha1Calculator(f *os.File) []byte {
	sha := sha1.New()
	io.TeeReader(f, sha)
	return sha.Sum(nil)
}

func md5Calculator(f *os.File) []byte {
	md := md5.New()
	io.TeeReader(f, md)
	return md.Sum(nil)
}

//https://github.com/cbt4all/networking/tree/master/golang-copy-paste/read-text-file-load-to-mem/_v0.001
func readTextFileLoadToMem(filename string) []string {

	var sliceStr []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		sliceStr = append(sliceStr, scanner.Text())
	}
	return sliceStr
}
