package shortestpath

type NodeList[TValue any] []*Node[TValue]
type EdgeList[TValue any] []*Edge[TValue]

type Node[TValue any] struct {
	Id    int
	Value TValue
	Edges EdgeList[TValue]
}

type Edge[TValue any] struct {
	From     *Node[TValue]
	To       *Node[TValue]
	Distance float64
}

type Graph[TValue any] struct {
	NodeCounter int
	Nodes       NodeList[TValue]
}

func NewGraph[TValue any]() *Graph[TValue] {
	return &Graph[TValue]{
		NodeCounter: 0,
	}
}

func (g *Graph[TValue]) Size() int {
	return len(g.Nodes)
}

func (g *Graph[TValue]) AddNode(value TValue) *Node[TValue] {
	node := &Node[TValue]{
		Id:    g.NodeCounter,
		Value: value,
	}
	g.NodeCounter++
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *Graph[TValue]) AddEdge(from, to *Node[TValue], distance float64) *Edge[TValue] {
	if from == nil || to == nil {
		return nil
	}
	edge := &Edge[TValue]{
		From:     from,
		To:       to,
		Distance: distance,
	}
	from.Edges = append(from.Edges, edge)
	return edge
}
