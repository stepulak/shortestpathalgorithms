package shortestpath

import (
	"fmt"
	"math"
	"slices"
)

func (g *Graph[TValue]) AStar(start, end *Node[TValue], h func(n *Node[TValue]) float64) (NodeList[TValue], error) {
	if start == nil || end == nil {
		return nil, fmt.Errorf("invalid start or end")
	}
	type nodeScore map[*Node[TValue]]float64

	gScore := nodeScore{start: 0}
	fScore := nodeScore{start: h(start)}
	cameFrom := map[*Node[TValue]]*Node[TValue]{}

	scoreOrInf := func(scoreMap nodeScore, n *Node[TValue]) float64 {
		val, ok := scoreMap[n]
		if !ok {
			return math.Inf(0)
		}
		return val
	}

	reconstructPath := func(current *Node[TValue]) NodeList[TValue] {
		totalPath := NodeList[TValue]{}
		for {
			totalPath = append(totalPath, current)
			next, ok := cameFrom[current]
			if !ok {
				break
			}
			current = next
		}
		slices.Reverse(totalPath)
		return totalPath
	}

	openSet := NewBinaryHeap[*Node[TValue]](func(a, b *Node[TValue]) bool {
		return scoreOrInf(fScore, a) < scoreOrInf(fScore, b)
	})
	openSet.Push(start)

	for openSet.Size() > 0 {
		current, _ := openSet.Pop()
		if current == end {
			return reconstructPath(current), nil
		}
		for _, neighbour := range current.Edges {
			tentativeGScore := gScore[current] + neighbour.Distance
			if tentativeGScore >= scoreOrInf(gScore, neighbour.To) {
				continue
			}
			cameFrom[neighbour.To] = current
			gScore[neighbour.To] = tentativeGScore
			fScore[neighbour.To] = tentativeGScore + h(neighbour.To)
			if !openSet.Contains(neighbour.To, func(a, b *Node[TValue]) bool { return a == b }) {
				openSet.Push(neighbour.To)
			}
		}
	}
	return nil, fmt.Errorf("can't reach end, A* failed")
}
