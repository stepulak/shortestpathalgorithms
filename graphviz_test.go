package shortestpath

import (
	"fmt"
	"testing"
)

func TestDotViz(t *testing.T) {
	type node struct {
		v int
	}
	g := NewGraph[node]()
	nodes := []*Node[node]{}
	for i := range 10 {
		nodes = append(nodes, g.AddNode(node{i}))
	}
	out := g.AddNode(node{11})
	for _, n := range nodes {
		g.AddEdge(n, out, float64(n.Value.v))
	}
	str, err := g.ToDot(func(n node) string {
		return fmt.Sprintf("%v", n.v)
	})
	expected := `digraph{
	"0" -> "11" [label="0.00"]
	"1" -> "11" [label="1.00"]
	"2" -> "11" [label="2.00"]
	"3" -> "11" [label="3.00"]
	"4" -> "11" [label="4.00"]
	"5" -> "11" [label="5.00"]
	"6" -> "11" [label="6.00"]
	"7" -> "11" [label="7.00"]
	"8" -> "11" [label="8.00"]
	"9" -> "11" [label="9.00"]
}`
	if err != nil {
		t.Error(err)
	}
	if str != expected {
		t.Error("Wrong digraph", str, "expected", expected)
	}
}
