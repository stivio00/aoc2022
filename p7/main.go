package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fs *FileSystem = NewFilesystem()

	var line string
	// Build fs
	for scanner.Scan() {
		line = scanner.Text()
		fs.ProccesLine(line)
	}

	FillSum(fs.Root)

}
