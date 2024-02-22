package shortestpath

import "fmt"

func (g *Graph[TValue]) BellmanFord(source *Node[TValue]) (Predecessors[TValue], error) {
	if source == nil {
		return nil, fmt.Errorf("BellmanFord: source node is nil")
	}

	distances := make(map[int]float64)
	predecessors := make(Predecessors[TValue])

	distances[source.Id] = 0
	relaxation := 0

	for i := 0; i < g.Size()+1; i++ {
		for _, node := range g.Nodes {
			for _, edge := range node.Edges {
				fromDist, fromDistOk := distances[edge.From.Id]
				if !fromDistOk {
					continue
				}
				toDist, toDistOk := distances[edge.To.Id]
				if !toDistOk || fromDist+edge.Distance < toDist {
					distances[edge.To.Id] = fromDist + edge.Distance
					predecessors[edge.To.Id] = edge.From
					relaxation = i
				}
			}
		}
	}

	if relaxation == g.Size() {
		return nil, fmt.Errorf("BellmanFord: negative cycle detected")
	}

	return predecessors, nil
}
