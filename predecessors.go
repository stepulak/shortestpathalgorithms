package shortestpath

// Look up table to find a node's predecessor
// Key: Node.Id of node A
// Value: *Node of node B
// Result: Edge from node A to node B
type Predecessors map[int]*Node

func (p Predecessors) Get(node *Node) *Node {
	if node, ok := p[node.Id]; ok {
		return node
	}
	return nil
}

func (p Predecessors) Set(from, to *Node) {
	p[to.Id] = from
}

// Get the longest shortest path from starting node to nodeEnd
func (p Predecessors) GetShortestPath(nodeEnd *Node) NodeList {
	if nodeEnd == nil || p[nodeEnd.Id] == nil {
		return nil
	}
	path := NodeList{}
	for node := nodeEnd; node != nil; node = p[node.Id] {
		path = append(NodeList{node}, path...)
	}
	return path
}
