package graph

import (
	"github.com/algorithm-testing/graph/dvr"
)

type GraphNode struct {
	id           string
	neighbours   map[string]*Neighbour
	routingTable *dvr.RoutingTable
}

func createGraphNode(id string) *GraphNode {
	return &GraphNode{
		id:         id,
		neighbours: make(map[string]*Neighbour),
		routingTable: &dvr.RoutingTable{
			ForNodeId: id,
			Entries:   make(map[string]dvr.EntryRoute),
		},
	}
}

type Neighbour struct {
	node   *GraphNode
	weight uint
}

func (graphNode *GraphNode) addNeighbour(node *GraphNode, weight uint) {
	graphNode.neighbours[node.id] = &Neighbour{node: node, weight: weight}
}
