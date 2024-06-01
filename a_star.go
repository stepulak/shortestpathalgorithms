package shortestpath

import (
	"fmt"
	"math"
	"slices"
)

func infScore(m map[int]float64, idx int) float64 {
	v, ok := m[idx]
	if !ok {
		return math.Inf(0)
	}
	return v
}

func (g *Graph[TValue]) AStar(start, end *Node[TValue], h func(n *Node[TValue]) float64) (NodeList[TValue], error) {
	if start == nil || end == nil {
		return nil, fmt.Errorf("invalid start or end")
	}
	openSet := []*Node[TValue]{start}
	cameFrom := map[int]*Node[TValue]{}
	idToNode := map[int]*Node[TValue]{
		start.Id: start,
		end.Id:   end,
	}
	gScore := map[int]float64{}
	gScore[start.Id] = 0

	fScore := map[int]float64{}
	fScore[start.Id] = h(start)

	lowestFscore := func() int {
		best := fScore[openSet[0].Id]
		bestI := openSet[0].Id
		for i := range openSet {
			if infScore(fScore, openSet[i].Id) < best {
				bestI = i
				best = infScore(fScore, i)
			}
		}
		return bestI
	}
	reconstructPath := func(current int) NodeList[TValue] {
		curr, ok := idToNode[current]
		totalPath := NodeList[TValue]{}
		for ok {
			totalPath = append(totalPath, curr)
			curr, ok = cameFrom[curr.Id]
		}
		slices.Reverse(totalPath)
		return totalPath
	}

	for len(openSet) > 0 {
		current := lowestFscore()
		if current == end.Id {
			return reconstructPath(current), nil
		}
		for i, v := range openSet {
			if v.Id == current {
				openSet = append(openSet[:i], openSet[i+1:]...)
				break
			}
		}
		for _, neighbour := range idToNode[current].Edges {
			idToNode[neighbour.To.Id] = neighbour.To
			tentativeGScore := gScore[current] + neighbour.Distance
			if tentativeGScore < infScore(gScore, neighbour.To.Id) {
				cameFrom[neighbour.To.Id] = idToNode[current]
				gScore[neighbour.To.Id] = tentativeGScore
				fScore[neighbour.To.Id] = tentativeGScore + h(neighbour.To)
				inSet := false
				for _, v := range openSet {
					if neighbour.To.Id == v.Id {
						inSet = true
						break
					}
				}
				if !inSet {
					openSet = append(openSet, neighbour.To)
				}
			}
		}
	}
	return nil, fmt.Errorf("can't reach end, A* failed")
}
