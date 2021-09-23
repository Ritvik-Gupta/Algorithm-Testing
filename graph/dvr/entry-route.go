package dvr

type EntryRoute interface {
	Distance() uint
	Copy() EntryRoute
}

type ViaEntry struct {
	distance  uint
	ViaNodeId string
}

func CreateViaEntry(viaNodeId string, distance uint) *ViaEntry {
	return &ViaEntry{ViaNodeId: viaNodeId, distance: distance}
}

func (via *ViaEntry) Distance() uint {
	return via.distance
}

func (via *ViaEntry) Copy() EntryRoute {
	return &ViaEntry{ViaNodeId: via.ViaNodeId, distance: via.distance}
}

func (via *ViaEntry) CloneWith(weight uint, viaNodeId string) *ViaEntry {
	return &ViaEntry{ViaNodeId: viaNodeId, distance: via.distance + weight}
}

type _UnreachableEntry struct{}

func (*_UnreachableEntry) Distance() uint {
	return 0
}

func (unreachable *_UnreachableEntry) Copy() EntryRoute {
	return unreachable
}

var UnreachableEntry = &_UnreachableEntry{}
