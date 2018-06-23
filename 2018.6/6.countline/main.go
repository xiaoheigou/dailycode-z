package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	// inspect the input args
	dirCount := len(os.Args)
	if dirCount < 2 {
		log.Fatal(errors.New("please input at least one dirPath"))
	}
	//walk through the dir ,count the line number
	totalCount := 0
	for i := 1; i < dirCount; i++ {
		dirPath := os.Args[i]
		sumCount := walkDir(dirPath)
		totalCount += sumCount
	}
	fmt.Println(totalCount)
}

//walk the dir
func walkDir(dirName string) int {
	var sumCount int
	filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if getExt(path) == ".go" {
			sumCount += countFileLine(path)
			fmt.Println(path)
		}
		return nil
	})
	return sumCount
}

// get go file code line
func countFileLine(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lineCount int
	reader := bufio.NewReader(file)
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {
			lineCount++
		}
	}
	return lineCount
}

//get file ext
func getExt(fileName string) string {
	return path.Ext(fileName)
}
