package shortestpath

import (
	"testing"
)

func TestAStar1(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{g.AddNode(0), g.AddNode(1), g.AddNode(2), g.AddNode(3)}
	g.AddEdge(nodes[0], nodes[1], 1)
	g.AddEdge(nodes[0], nodes[2], 1)
	g.AddEdge(nodes[1], nodes[3], 1)
	g.AddEdge(nodes[2], nodes[3], 2)

	path, err := g.AStar(nodes[0], nodes[3], func(n *Node[int]) float64 { return float64(0) })
	if err != nil || path == nil || len(path) != 3 {
		t.Errorf("TestAStar1: AStar failed: %v", err)
	}
	if path == nil || path[0] != nodes[0] ||
		path[1] != nodes[1] {
		t.Error("TestAStar1: invalid previous nodes")
	}

	if len(path) != 3 || path[0] != nodes[0] ||
		path[1] != nodes[1] || path[2] != nodes[3] {
		t.Error("TestAStar1: invalid path")
	}
}

func TestAStar2(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{g.AddNode(0), g.AddNode(1), g.AddNode(2), g.AddNode(3)}
	g.AddEdge(nodes[0], nodes[1], 1)
	g.AddEdge(nodes[1], nodes[3], 1)
	g.AddEdge(nodes[0], nodes[2], 1)
	g.AddEdge(nodes[0], nodes[3], 3)
	g.AddEdge(nodes[2], nodes[3], 2)

	path, err := g.AStar(nodes[0], nodes[3], func(n *Node[int]) float64 { return float64(n.Value) })
	if err != nil || path == nil || len(path) != 3 {
		t.Errorf("TestAStar2: AStar failed: %v", err)
	}
	if path[0] == nil || path[0] != nodes[0] ||
		path[1] != nodes[1] || path[len(path)-2] != nodes[1] {
		t.Error("TestAStar1: invalid previous nodes")
	}
}

func TestAStar3(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{}
	for i := 0; i <= 8; i++ {
		nodes = append(nodes, g.AddNode(i))
	}
	g.AddEdge(nodes[0], nodes[1], 0.5)
	g.AddEdge(nodes[0], nodes[2], 1)
	g.AddEdge(nodes[1], nodes[4], 0.1)
	g.AddEdge(nodes[1], nodes[5], 0.8)
	g.AddEdge(nodes[2], nodes[3], 3)
	g.AddEdge(nodes[2], nodes[4], 2)
	g.AddEdge(nodes[5], nodes[6], 3)
	g.AddEdge(nodes[4], nodes[6], 1)
	g.AddEdge(nodes[4], nodes[7], 1.1)
	g.AddEdge(nodes[3], nodes[7], 1)
	g.AddEdge(nodes[6], nodes[8], 5)
	g.AddEdge(nodes[7], nodes[8], 0.4)

	path, err := g.AStar(nodes[0], nodes[len(nodes)-1], func(n *Node[int]) float64 { return float64(n.Value) })
	if err != nil || path == nil || len(path) != 5 {
		t.Errorf("TestAStar3: AStar failed: %v", err)
	}
	expectedPath := []int{0, 1, 4, 7, 8}
	// fmt.Errorf("3: %v", path)
	for i, v := range path {
		if expectedPath[i] != v.Value {
			t.Errorf("invalid path")
		}
	}

}

func TestAStar4(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{g.AddNode(0), g.AddNode(1), g.AddNode(2)}
	g.AddEdge(nodes[0], nodes[1], 1)
	g.AddEdge(nodes[1], nodes[0], 1)
	g.AddEdge(nodes[0], nodes[2], 3)
	g.AddEdge(nodes[2], nodes[1], 0.1)

	path, err := g.AStar(nodes[1], nodes[len(nodes)-1], func(n *Node[int]) float64 { return float64(n.Value) })
	if err != nil || path == nil || len(path) != 3 {
		t.Errorf("TestAStar4: AStar failed: %v", err)
	}
	expectedPath := []int{1, 0, 2}
	for i, v := range path {
		if expectedPath[i] != v.Value {
			t.Errorf("invalid path")
		}
	}
}

func TestAStar5(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{g.AddNode(0), g.AddNode(1)}

	path, err := g.AStar(nodes[0], nodes[1], func(n *Node[int]) float64 { return float64(n.Value) })
	if err == nil || path != nil || len(path) != 0 {
		t.Errorf("TestAStar5: AStar failed: %v", err)
	}
}

func TestAStar6(t *testing.T) {
	g := NewGraph[string]()

	path, err := g.AStar(nil, nil, func(n *Node[string]) float64 { return float64(n.Id) })
	if err == nil || path != nil {
		t.Error("TestAStar6: AStar succeeded")
	}
}
