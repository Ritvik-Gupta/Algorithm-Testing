package graph

import (
	"fmt"
	"sync"

	"github.com/algorithm-testing/graph/dvr"
)

type Graph struct {
	nodes map[string]*GraphNode
}

func CreateGraph() *Graph {
	return &Graph{nodes: make(map[string]*GraphNode)}
}

func (graph *Graph) AddNode(nodeId string) error {
	if _, containsNode := graph.nodes[nodeId]; containsNode {
		return fmt.Errorf("node with Id already created")
	}
	graph.nodes[nodeId] = createGraphNode(nodeId)
	return nil
}

func (graph *Graph) Connect(nodeAId, nodeBId string, weight uint) error {
	if nodeA, ok := graph.nodes[nodeAId]; !ok {
		return fmt.Errorf("first node cannot be found in the Graph")
	} else if nodeB, ok := graph.nodes[nodeBId]; !ok {
		return fmt.Errorf("second node cannot be found in the Graph")
	} else {
		nodeA.addNeighbour(nodeB, weight)
		nodeB.addNeighbour(nodeA, weight)
		return nil
	}
}

func (graph *Graph) Print() {
	for _, node := range graph.nodes {
		fmt.Print(node.id, "\t->\t")
		for _, neighbour := range node.neighbours {
			fmt.Print(neighbour.node.id, "[", neighbour.weight, "]\t")
		}
		fmt.Println()
	}
}

func (graph *Graph) InitiateDVR() {
	totalNodes := len(graph.nodes)

	updatingChannel := make(chan string, totalNodes)
	cachedUpdates := make(map[string][]*dvr.UpdateVector)

	for _, node := range graph.nodes {
		node.routingTable.Entries[node.id] = dvr.CreateViaEntry(node.id, 0)

		for _, otherNode := range graph.nodes {
			if neighbour, ok := node.neighbours[otherNode.id]; ok {
				node.routingTable.Entries[otherNode.id] = dvr.CreateViaEntry(otherNode.id, neighbour.weight)
			} else {
				node.routingTable.Entries[otherNode.id] = dvr.UnreachableEntry
			}
		}

		cachedUpdates[node.id] = make([]*dvr.UpdateVector, len(node.neighbours))
		for _, neighbour := range node.neighbours {
			cachedUpdates[node.id] = append(
				cachedUpdates[node.id],
				dvr.CreateUpdateVector(neighbour.node.routingTable.Copy(), neighbour.weight),
			)
		}
	}

	var waitForUpdates sync.WaitGroup
	waitForUpdates.Add(totalNodes)

	for hasCache := true; hasCache; {
		hasCache = false
		for nodeId, updates := range cachedUpdates {
			node := graph.nodes[nodeId]

			go func(receivedUpdates []*dvr.UpdateVector) {
				if hasUpdated := node.routingTable.IssueUpdate(receivedUpdates); hasUpdated {
					updatingChannel <- nodeId
				}
				waitForUpdates.Done()
			}(updates)

			cachedUpdates[nodeId] = make([]*dvr.UpdateVector, len(node.neighbours))
			waitForUpdates.Wait()
		}
	}

	for {
		select {
		case nodeId := <-updatingChannel:
			{
				node := graph.nodes[nodeId]
				for _, neighbour := range node.neighbours {
					cachedUpdates[neighbour.node.id] = append(
						cachedUpdates[neighbour.node.id],
						dvr.CreateUpdateVector(node.routingTable.Copy(), neighbour.weight),
					)
				}
			}
		default:
			break // TODO :
		}
	}
}
