package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LineType int

const (
	CommandCd LineType = iota
	CommandLs
	OutDir
	OutFile
	Unknow
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name        string
	Size        int
	Files       []File
	Directories []Directory
}

type Path struct {
	Nodes []string
}

// Print working directory
func (p *Path) Pwd() string {
	builder := strings.Builder{}

	for i, node := range p.Nodes {
		builder.WriteString(node)
		if i == 0 || i == len(p.Nodes)-1 {
			continue
		}
		builder.WriteString("/")
	}

	return builder.String()
}

// Print Filesystem
func (p *Directory) String() string {
	builder := strings.Builder{}

	printFile := func(file File, level int) {
		builder.WriteString(strings.Repeat(" ", level))
		s := fmt.Sprintf("- %s (file, size=%d)\n", file.Name, file.Size)
		builder.WriteString(s)
	}

	printDir := func(dir Directory, level int) string {
		builder.WriteString(strings.Repeat(" ", level))
		s := fmt.Sprintf("- %s (dir)\n", dir.Name)
		builder.WriteString(s)
	}

	for i, node := range p.Nodes {
		builder.WriteString(node)
		if i == 0 || i == len(p.Nodes)-1 {
			continue
		}
		builder.WriteString("/")
	}

	return builder.String()
}

func CreateFile(line string) File {
	var name string
	var size int
	fmt.Sscanf(line, "%d %s", &size, &name)
	return File{Name: name, Size: size}
}

func CreateDir(line string) Directory {
	var name string
	fmt.Sscanf(line, "dir %s", &name)
	return Directory{
		Name:        name,
		Files:       make([]File, 0),
		Directories: make([]Directory, 0),
	}
}

func (p *Path) Cd(dir string) {
	if dir == ".." {
		p.Nodes = p.Nodes[:len(p.Nodes)-1]
	}
	p.Nodes = append(p.Nodes, dir)

}

func getLinetype(line string) LineType {
	if line[0] == '$' {
		if line[2:4] == "cd" {
			return CommandCd
		} else if line[2:4] == "ls" {
			return CommandLs
		} else {
			return Unknow
		}
	} else if line[0:3] == "dir" {
		return OutDir
	} else {
		return OutFile
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//var pwd Path

	var line string
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)
		lineType := getLinetype(line)
		if lineType == CommandCd {
			//TODO:
		}
	}

}
