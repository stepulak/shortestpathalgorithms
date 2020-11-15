package shortestpath

import (
	"fmt"
)

func (g *Graph) Dijkstra(source *Node) (Predecessors, error) {
	if source == nil {
		return nil, fmt.Errorf("Dijkstra: source node is nil")
	}

	distances := make(map[int]float64)
	predecessors := make(Predecessors)

	bh := NewBinaryHeap(func(a, b interface{}) bool {
		distA, okA := distances[a.(*Node).Id]
		distB, okB := distances[b.(*Node).Id]
		if okA && okB {
			return distA <= distB
		} else if okA {
			return true
		} else {
			return false
		}
	})

	distances[source.Id] = 0
	bh.Push(source)

	for bh.Size() > 0 {
		val, ok := bh.Pop()
		if !ok {
			return nil, fmt.Errorf("Dijkstra: unable to pop")
		}
		// For each neighbour check if you can improve the distance
		for _, edge := range val.(*Node).Edges {
			newD := distances[edge.From.Id] + edge.Distance
			if oldD, ok := distances[edge.To.Id]; !ok || newD < oldD {
				distances[edge.To.Id] = newD
				predecessors.Set(edge.From, edge.To)
				bh.Push(edge.To)
			}
		}
	}

	return predecessors, nil
}
