package shortestpath

import "testing"

func TestBellmanFord1(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{}
	for i := 0; i < 6; i++ {
		nodes = append(nodes, g.AddNode(i))
	}
	g.AddEdge(nodes[0], nodes[1], 10)
	g.AddEdge(nodes[0], nodes[2], 20)
	g.AddEdge(nodes[1], nodes[3], 50)
	g.AddEdge(nodes[2], nodes[3], 20)
	g.AddEdge(nodes[2], nodes[4], 33)
	g.AddEdge(nodes[1], nodes[4], 10)
	g.AddEdge(nodes[3], nodes[4], -20)
	g.AddEdge(nodes[4], nodes[5], 1)
	g.AddEdge(nodes[3], nodes[5], -2)

	preds, err := g.BellmanFord(nodes[0])
	if err != nil || preds == nil || len(preds) != 5 {
		t.Error("TestBellmanFord1: BellmanFord failed")
	}

	if path := preds.GetShortestPath(nodes[5]); len(path) != 4 || path[0] != nodes[0] ||
		path[1] != nodes[1] || path[2] != nodes[4] || path[3] != nodes[5] {
		t.Error("TestBellmanFord1: invalid path")
	}
}

func TestBellmanFord2(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{}
	for i := 0; i < 6; i++ {
		nodes = append(nodes, g.AddNode(i))
	}
	g.AddEdge(nodes[0], nodes[1], 10)
	g.AddEdge(nodes[1], nodes[2], 10)
	g.AddEdge(nodes[2], nodes[4], 7)
	g.AddEdge(nodes[4], nodes[3], 5)
	g.AddEdge(nodes[3], nodes[1], 5)
	g.AddEdge(nodes[4], nodes[5], 15)

	preds, err := g.BellmanFord(nodes[0])
	if err != nil || preds == nil || len(preds) != 5 {
		t.Error("TestBellmanFord2: BellmanFord failed")
	}

	if path := preds.GetShortestPath(nodes[5]); len(path) != 5 || path[0] != nodes[0] ||
		path[1] != nodes[1] || path[2] != nodes[2] || path[3] != nodes[4] || path[4] != nodes[5] {
		t.Error("TestBellmanFord2: invalid path")
	}
}

func TestBellmanFord3(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{}
	for i := 0; i < 6; i++ {
		nodes = append(nodes, g.AddNode(i))
	}
	g.AddEdge(nodes[0], nodes[1], 10)
	g.AddEdge(nodes[1], nodes[2], 1)
	g.AddEdge(nodes[2], nodes[4], 3)
	g.AddEdge(nodes[4], nodes[3], -10)
	g.AddEdge(nodes[3], nodes[1], 4)
	g.AddEdge(nodes[4], nodes[5], 22)

	if preds, err := g.BellmanFord(nodes[0]); err == nil || preds != nil {
		t.Error("TestBellmanFord3: BellmanFord succeeded")
	}
}

func TestBellmanFord4(t *testing.T) {
	g := NewGraph[string]()
	if preds, err := g.BellmanFord(nil); preds != nil || err == nil {
		t.Error("TestBellmanFord4: BellmanFord succeeded")
	}
}

func TestBellmanFord5(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{g.AddNode(0), g.AddNode(1)}

	preds, err := g.BellmanFord(nodes[0])
	if err != nil || preds == nil || len(preds) != 0 {
		t.Error("TestBellmanFord5: BellmanFord failed")
	}

	if path := preds.GetShortestPath(nodes[1]); path != nil {
		t.Error("TestBellmanFord5: invalid path")
	}
}
