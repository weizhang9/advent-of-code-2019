package utils

import (
	"io/ioutil"
	"log"
	"strconv"
)

// GetFileContent takes a file path as a string
// and return the file content in format of byte slice
// and an error if any (nil if no error occurred)
func GetFileContent(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	errstr := "ERROR: can't read file content from" + filePath
	CheckError(err, errstr)
	return content
}

// CheckError streamline error checking in Go
// it takes an optional string to provide more info
// to the error to help debug
func CheckError(e error, info ...string) {
	var i string
	if e != nil {
		if len(info) > 0 {
			i = info[0]
		}
		log.Fatalln(i, e)
	}
}

// StringToIntSlice take a slice of string and convert it into
// a slice of ints. It will panic if encounter an error
func StringToIntSlice(s []string) []int {
	ints := make([]int, 0, len(s))

	for _, v := range s {
		if len(s) == 0 {
			continue
		}

		int, err := strconv.Atoi(v)
		CheckError(err, "ERROR: can't convert string slice to int slice ")

		ints = append(ints, int)
	}

	return ints
}
