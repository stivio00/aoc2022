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
	Files       map[string]*File      //Key: file name
	Directories map[string]*Directory //Key: directory name
}

type FileSystem struct {
	PathNodes []string //Current working path
	Root      *Directory
}

// Print working directory
func NewFilesystem() *FileSystem {
	fs := &FileSystem{PathNodes: make([]string, 0), Root: nil}
	fs.Root = CreateDir("dir /")
	return fs
}

// Print working directory
func (fs *FileSystem) Pwd() string {
	builder := strings.Builder{}

	for i, node := range fs.PathNodes {
		builder.WriteString(node)
		if i == 0 || i == len(fs.PathNodes)-1 {
			continue
		}
		builder.WriteString("/")
	}

	return builder.String()
}

func (fs *FileSystem) AddFile(file *File) {
	fs.GetCurrentDir().Files[file.Name] = file
}

// Add a new directory in the current dir, if fs is empty then
func (fs *FileSystem) AddDirectory(dir *Directory) {
	currentDir := fs.GetCurrentDir()
	if currentDir == nil {
		fs.Root = dir
		return
	}
	currentDir.Directories[dir.Name] = dir
}

func (fs *FileSystem) IsEmpty() bool {
	return len(fs.PathNodes) == 0
}

// Print Filesystem
func (d *Directory) String() string {
	builder := strings.Builder{}

	printFile := func(file *File, level int) {
		builder.WriteString(strings.Repeat(" ", level))
		s := fmt.Sprintf("- %s (file, size=%d)\n", file.Name, file.Size)
		builder.WriteString(s)
	}

	var printDir func(*Directory, int) //fowar declaration for recursive use
	printDir = func(dir *Directory, level int) {
		builder.WriteString(strings.Repeat(" ", level))
		s := fmt.Sprintf("- %s (dir)\n", dir.Name)
		builder.WriteString(s)

		for _, dir := range dir.Directories {
			printDir(dir, level+1)
		}

		for _, file := range dir.Files {
			printFile(file, level+1)
		}
	}

	printDir(d, 0)
	return builder.String()
}

func ParseCdLine(line string) string {
	var dirName string
	fmt.Sscanf(line, "cd %s", &dirName)
	return dirName
}

func CreateFile(line string) *File {
	var name string
	var size int
	fmt.Sscanf(line, "%d %s", &size, &name)
	return &File{Name: name, Size: size}
}

func CreateDir(line string) *Directory {
	var name string
	fmt.Sscanf(line, "dir %s", &name)
	return &Directory{
		Name:        name,
		Files:       make(map[string]*File, 0),
		Directories: make(map[string]*Directory, 0),
	}
}

func (fs *FileSystem) GetCurrentDir() *Directory {
	if len(fs.PathNodes) == 0 {
		return nil
	}
	if len(fs.PathNodes) == 1 {
		return fs.Root
	}

	var dir = fs.Root
	for _, nodeName := range fs.PathNodes[1:] {
		dir = dir.Directories[nodeName]
	}
	return dir
}

func (fs *FileSystem) Cd(dir string) {
	if dir == "." {
		return
	}
	if dir == ".." {
		fs.PathNodes = fs.PathNodes[:len(fs.PathNodes)-1]
		return
	}
	fs.PathNodes = append(fs.PathNodes, dir)

}

func getLinetype(line string) LineType {
	if line == "" {
		return Unknow
	}
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
	var fs *FileSystem = NewFilesystem()

	var line string
	// Build fs
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)
		lineType := getLinetype(line)
		switch lineType {
		case CommandCd:
			dirName := ParseCdLine(line)
			fs.Cd(dirName)
		case CommandLs:
			continue
		case OutDir:
			dir := CreateDir(line)
			fs.AddDirectory(dir)
		case OutFile:
			file := CreateFile(line)
			fs.AddFile(file)
		}
	}

	//TODO: transvers tree and calculate sizes
	// transverse dir and get problem sizes

}
