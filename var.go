package main

var (
	BUFFERSIZE int64
	sourceSize int64
	sourceName string
)

type files struct {
	src  string
	dst  string
	sha1 []byte
	md5  []byte
}
