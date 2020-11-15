package shortestpath

type NodeList []*Node
type EdgeList []*Edge

type Node struct {
	Id    int
	Value interface{}
	Edges EdgeList
}

type Edge struct {
	From     *Node
	To       *Node
	Distance float64
}

type Graph struct {
	NodeCounter int
	Nodes       NodeList
}

func NewGraph() *Graph {
	return &Graph{
		NodeCounter: 0,
	}
}

func (g *Graph) Size() int {
	return len(g.Nodes)
}

func (g *Graph) AddNode(value interface{}) *Node {
	node := &Node{
		Id:    g.NodeCounter,
		Value: value,
	}
	g.NodeCounter++
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *Graph) AddEdge(from, to *Node, distance float64) *Edge {
	if from == nil || to == nil {
		return nil
	}
	edge := &Edge{
		From:     from,
		To:       to,
		Distance: distance,
	}
	from.Edges = append(from.Edges, edge)
	return edge
}
