package main

import (
	"testing"
)

func TestPath(t *testing.T) {
	p := Path{}
	p.Nodes = append(p.Nodes, "/")
	p.Nodes = append(p.Nodes, "a")
	p.Nodes = append(p.Nodes, "b")

	s := p.Pwd()
	t.Log(s)
}
