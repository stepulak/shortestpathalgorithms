package shortestpath

// Look up table to find a node's predecessor
// Key: Node.Id of node A
// Value: *Node of node B
// Result: Edge from node A to node B
type Predecessors[TValue any] map[int]*Node[TValue]

func (p Predecessors[TValue]) Get(node *Node[TValue]) *Node[TValue] {
	if node, ok := p[node.Id]; ok {
		return node
	}
	return nil
}

func (p Predecessors[TValue]) Set(from, to *Node[TValue]) {
	p[to.Id] = from
}

// Get the longest shortest path from starting node to nodeEnd
func (p Predecessors[TValue]) GetShortestPath(nodeEnd *Node[TValue]) NodeList[TValue] {
	if nodeEnd == nil || p[nodeEnd.Id] == nil {
		return nil
	}
	path := NodeList[TValue]{}
	for node := nodeEnd; node != nil; node = p[node.Id] {
		path = append(NodeList[TValue]{node}, path...)
	}
	return path
}
