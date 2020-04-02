package main

var (
	BUFFERSIZE int64
	sourceSize int64
	sourceName string

	totalSize    int64
	totalCoppied int64
)

type file struct {
	src  string
	dst  string
	sha1 []byte
	md5  []byte
	name string
	size int64
}
