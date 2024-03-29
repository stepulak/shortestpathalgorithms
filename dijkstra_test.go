package shortestpath

import "testing"

func TestDijkstra1(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{g.AddNode(0), g.AddNode(1), g.AddNode(2), g.AddNode(3)}
	g.AddEdge(nodes[0], nodes[1], 1)
	g.AddEdge(nodes[0], nodes[2], 1)
	g.AddEdge(nodes[1], nodes[3], 1)
	g.AddEdge(nodes[2], nodes[3], 2)

	preds, err := g.Dijkstra(nodes[0])
	if err != nil || preds == nil || len(preds) != 3 {
		t.Errorf("TestDijkstra1: Dijkstra failed: %v", err)
	}
	if preds.Get(nodes[0]) != nil || preds.Get(nodes[1]) != nodes[0] ||
		preds.Get(nodes[2]) != nodes[0] || preds.Get(nodes[3]) != nodes[1] {
		t.Error("TestDijkstra1: invalid previous nodes")
	}

	if path := preds.GetShortestPath(nodes[3]); len(path) != 3 || path[0] != nodes[0] ||
		path[1] != nodes[1] || path[2] != nodes[3] {
		t.Error("TestDijkstra1: invalid path")
	}
}

func TestDijkstra2(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{g.AddNode(0), g.AddNode(1), g.AddNode(2), g.AddNode(3)}
	g.AddEdge(nodes[0], nodes[1], 1)
	g.AddEdge(nodes[1], nodes[3], 1)
	g.AddEdge(nodes[0], nodes[2], 1)
	g.AddEdge(nodes[0], nodes[3], 3)
	g.AddEdge(nodes[2], nodes[3], 2)

	preds, err := g.Dijkstra(nodes[0])
	if err != nil || preds == nil || len(preds) != 3 {
		t.Errorf("TestDijkstra2: Dijkstra failed: %v", err)
	}
	if preds.Get(nodes[0]) != nil || preds.Get(nodes[1]) != nodes[0] ||
		preds.Get(nodes[2]) != nodes[0] || preds.Get(nodes[3]) != nodes[1] {
		t.Error("TestDijkstra1: invalid previous nodes")
	}
}

func TestDijkstra3(t *testing.T) {
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

	preds, err := g.Dijkstra(nodes[0])
	if err != nil || preds == nil || len(preds) != 8 {
		t.Errorf("TestDijkstra3: Dijkstra failed: %v", err)
	}

	if path := preds.GetShortestPath(nodes[8]); len(path) != 5 || path[0] != nodes[0] ||
		path[1] != nodes[1] || path[2] != nodes[4] || path[3] != nodes[7] || path[4] != nodes[8] {
		t.Error("TestDijkstra3: invalid path")
	}
}

func TestDijkstra4(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{g.AddNode(0), g.AddNode(1), g.AddNode(2)}
	g.AddEdge(nodes[0], nodes[1], 1)
	g.AddEdge(nodes[1], nodes[0], 1)
	g.AddEdge(nodes[0], nodes[2], 3)
	g.AddEdge(nodes[2], nodes[1], 0.1)

	preds, err := g.Dijkstra(nodes[1])
	if err != nil || preds == nil || len(preds) != 2 {
		t.Errorf("TestDijkstra4: Dijkstra failed: %v", err)
	}

	if path := preds.GetShortestPath(nodes[2]); len(path) != 3 ||
		path[0] != nodes[1] || path[1] != nodes[0] || path[2] != nodes[2] {
		t.Error("TestDijkstra4: invalid path")
	}
}

func TestDijkstra5(t *testing.T) {
	g := NewGraph[int]()
	nodes := NodeList[int]{g.AddNode(0), g.AddNode(1)}

	preds, err := g.Dijkstra(nodes[0])
	if err != nil || preds == nil || len(preds) != 0 {
		t.Errorf("TestDijkstra5: Dijkstra failed: %v", err)
	}

	if path := preds.GetShortestPath(nodes[1]); len(path) != 0 {
		t.Errorf("TestDijkstra5: invalid path, %v", len(path))
	}
}

func TestDijkstra6(t *testing.T) {
	g := NewGraph[string]()

	preds, err := g.Dijkstra(nil)
	if err == nil || preds != nil {
		t.Error("TestDijkstra6: Dijkstra succeeded")
	}
}
