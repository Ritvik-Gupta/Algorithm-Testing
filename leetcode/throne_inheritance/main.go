package throne_inheritance

type HeirarchyNode struct {
	name     string
	isAlive  bool
	children []*HeirarchyNode
}

func NewHeirarchyNode(name string) *HeirarchyNode {
	return &HeirarchyNode{name: name, isAlive: true, children: []*HeirarchyNode{}}
}

type ThroneInheritance struct {
	root             *HeirarchyNode
	heirarchy_record map[string]*HeirarchyNode
}

func Constructor(kingName string) ThroneInheritance {
	root, heirarchy_record := NewHeirarchyNode(kingName), make(map[string]*HeirarchyNode)
	heirarchy_record[kingName] = root
	return ThroneInheritance{root, heirarchy_record}
}

func (throne *ThroneInheritance) Birth(parentName string, childName string) {
	child := NewHeirarchyNode(childName)
	throne.heirarchy_record[childName] = child
	throne.heirarchy_record[parentName].children = append(throne.heirarchy_record[parentName].children, child)
}

func (throne *ThroneInheritance) Death(name string) {
	throne.heirarchy_record[name].isAlive = false
}

func (throne *ThroneInheritance) GetInheritanceOrder() (result []string) {
	node_stack := []*HeirarchyNode{throne.root}
	for len(node_stack) > 0 {
		node := node_stack[len(node_stack)-1]
		node_stack = node_stack[:len(node_stack)-1]

		if node.isAlive {
			result = append(result, node.name)
		}
		for i := len(node.children) - 1; i >= 0; i-- {
			node_stack = append(node_stack, node.children[i])
		}
	}

	return
}
