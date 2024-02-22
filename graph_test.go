package shortestpath

import (
	"math"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph[byte]()
	if g == nil || g.Size() != 0 || len(g.Nodes) != 0 {
		t.Error("TestGraph: construction failed")
	}

	nodeA := g.AddNode('A')
	if nodeA == nil || g.Size() != 1 || g.Nodes[0] != nodeA {
		t.Error("TestGraph: adding nodeA failed")
	}
	nodeB := g.AddNode('B')
	if nodeB == nil || g.Size() != 2 || g.Nodes[1] != nodeB {
		t.Error("TestGraph: adding nodeB failed")
	}

	edgeAB := g.AddEdge(nodeA, nodeB, 1)
	if edgeAB == nil || len(nodeA.Edges) != 1 || nodeA.Edges[0] != edgeAB {
		t.Error("TestGraph: adding edgeAB failed")
	}
	if edgeAB.From != nodeA || edgeAB.To != nodeB || math.Abs(edgeAB.Distance-1.0) >= 0.0001 {
		t.Error("TestGraph: edgeAB is invalid")
	}
}
