package dvr

type RoutingTable struct {
	ForNodeId string
	Entries   map[string]EntryRoute
}

func (routingTable *RoutingTable) Copy() *RoutingTable {
	copyRoutingTable := &RoutingTable{
		ForNodeId: routingTable.ForNodeId,
		Entries:   make(map[string]EntryRoute),
	}
	for key, entry := range routingTable.Entries {
		copyRoutingTable.Entries[key] = entry.Copy()
	}
	return copyRoutingTable
}

func (routingTable *RoutingTable) IssueUpdate(receivedUpdates []*UpdateVector) bool {
	hasUpdated := false

	for nodeId, bestEntry := range routingTable.Entries {
		for _, updateVector := range receivedUpdates {
			entry := updateVector.routingTable.Entries[nodeId]

			if viaEntry, ok := entry.(*ViaEntry); ok {
				viaEntryClone := viaEntry.CloneWith(updateVector.weight, nodeId)

				if viaEntryClone.Distance() < bestEntry.Distance() {
					bestEntry = viaEntryClone
					hasUpdated = true
				}
			} else if entry == UnreachableEntry {
				continue
			}
		}
	}

	return hasUpdated
}

type UpdateVector struct {
	routingTable *RoutingTable
	weight       uint
}

func CreateUpdateVector(routingTable *RoutingTable, weight uint) *UpdateVector {
	return &UpdateVector{routingTable: routingTable, weight: weight}
}
