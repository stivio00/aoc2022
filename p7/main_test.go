package main

import (
	"testing"
)

func TestPath(t *testing.T) {
	fs := NewFilesystem()
	fs.Cd("/")

	fs.AddDirectory(CreateDir("dir A"))
	fs.AddDirectory(CreateDir("dir B"))
	fs.AddFile(CreateFile("1234 test.txt"))

	fs.Cd("A")

	fs.AddDirectory(CreateDir("dir C"))
	fs.AddFile(CreateFile("423424 test2.txt"))

	fs.Cd("..")
	fs.Cd("B")
	fs.AddFile(CreateFile("94234 test4.txt"))

	s := fs.Root.String()
	t.Log(s)
}
