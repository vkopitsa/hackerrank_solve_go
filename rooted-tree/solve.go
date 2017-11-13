//
// https://www.hackerrank.com/challenges/rooted-tree/problem
//
package main

import "fmt"

// X and Y
type row_node struct {
	X int64
	Y int64
}

// U T V K - 85
// or
// Q A B - 81
type row_update_query struct {
	Type       rune
	TorA       int64
	VorB       int64
	KorNothing int64
}

func main() {

	var nodeNumber, queryNumber, rootNumber int64
	fmt.Scanf("%v %v %v", &nodeNumber, &queryNumber, &rootNumber)

	//fmt.Println(nodeNumber, queryNumber, rootNumber, "init")

	nodes := make([]row_node, nodeNumber-1)
	for i := range nodes {
		fmt.Scanf("%v %v", &nodes[i].X, &nodes[i].Y)
		//fmt.Println(nodes[i], "nodes[i]")
	}

	queries := make([]row_update_query, queryNumber)
	for i := range queries {
		fmt.Scanf("%c %v %v %v", &queries[i].Type, &queries[i].TorA, &queries[i].VorB, &queries[i].KorNothing)
		//fmt.Println(queries[i], "queries[i]")
	}

	solve := solve(rootNumber, nodes, queries)

	fmt.Printf("%s", solve)
}

func solve(rootNumber int64, nodes []row_node, queries []row_update_query) string {
	r := ""
	var tree *TreeNode
	for _, n := range nodes {
		var x, y *TreeNode
		// X
		if tree != nil {
			x = tree.FindByNumber(n.X)
		}

		if x == nil {
			x = NewTreeNode(n.X)
		}

		// Y
		if tree != nil {
			y = tree.FindByNumber(n.Y)
		}

		if y == nil {
			y = NewTreeNode(n.Y)
		}

		if tree == nil {
			x.Insert(y)
			tree = x
		} else {
			x.Insert(y)
		}
	}

	rootNode := tree.FindByNumber(rootNumber)

	for _, q := range queries {
		// update
		if q.Type == 85 {
			foundNode := rootNode.FindByNumber(q.TorA)
			foundNode.Reset().Update(q.VorB, q.KorNothing)
		} else {
			// query
			sum := rootNode.FindByNumber(q.TorA).Sum(q.VorB)
			r += fmt.Sprintf("%d\n", sum)
		}
	}

	return r
}

//// all

type TreeNode struct {
	NodeList *TreeNodeList
	ComeNode *TreeNode
	Distance int64
	Value    int64
	Number   int64
}

func NewTreeNode(number int64) *TreeNode {
	n := TreeNode{}
	n.NodeList = NewTreeNodeList()
	n.NodeList.Insert(&n)
	n.Number = number
	return &n
}

func (node *TreeNode) Insert(n *TreeNode) {
	node.NodeList.Insert(n)
	n.NodeList.Insert(node)
}

func (node *TreeNode) Remove(n *TreeNode) {
	node.NodeList.Remove(n, node)
}

func (node *TreeNode) List() []*TreeNode {
	return node.NodeList.List(node)
}

func (node *TreeNode) Len() int64 {
	//fmt.Println(node, node.GetComeNode(), "1111")
	return node.NodeList.Len(node) + 1
}

func (node *TreeNode) SetComeNode(n *TreeNode) {
	node.ComeNode = n
}

func (node *TreeNode) GetComeNode() *TreeNode {
	return node.ComeNode
}

func (node *TreeNode) SetDistance(distance int64) {
	node.Distance = distance + 1
}

func (node *TreeNode) GetDistance() int64 {
	return node.Distance
}

func (node *TreeNode) FindByNumber(number int64) *TreeNode {
	return node.NodeList.FindByNumber(number, node)
}

func (node *TreeNode) Update(valueV int64, valueK int64) *TreeNode {
	node.Calculate(valueV, valueK)
	node.UpdateQuery(valueV, valueK)
	return node
}

func (node *TreeNode) UpdateQuery(valueV int64, valueK int64) *TreeNode {
	node.NodeList.UpdateQuery(node, valueV, valueK)
	return node
}

func (node *TreeNode) Sum(toNumber int64) int64 {
	toNode := node.FindByNumber(toNumber)

	return node.NodeList.Sum(node, toNode)
}

func (node *TreeNode) Reset() *TreeNode {
	node.Distance = 0
	node.ComeNode = nil

	return node
}

func (node *TreeNode) Calculate(valueV int64, valueK int64) *TreeNode {
	node.Value += valueV + node.GetDistance()*valueK

	return node
}

func (node *TreeNode) String() string {
	return fmt.Sprintf("(number: %d, value: %d, distance: %d)", node.Number, node.Value, node.Distance)
}

type TreeNodeList struct {
	Neighbors []*TreeNode
}

func NewTreeNodeList() *TreeNodeList {
	l := TreeNodeList{}
	l.Neighbors = make([]*TreeNode, 0, 100)
	return &l
}

func (l *TreeNodeList) Insert(n *TreeNode) {
	ok, _ := l.IsExists(n)
	if !ok {
		l.Neighbors = append(l.Neighbors, n)
	}
}

func (l *TreeNodeList) Remove(n *TreeNode, owner *TreeNode) {
	ok, index := l.IsExists(n)
	if ok {
		l.Neighbors = append(l.Neighbors[:index], l.Neighbors[index+1:]...)
	} else {
		for _, node := range l.Neighbors {
			if owner.Number == node.Number || (owner.GetComeNode() != nil && owner.GetComeNode().Number == node.Number) {
				continue
			}

			node.SetComeNode(owner)
			node.Remove(n)
		}
	}
}

func (l TreeNodeList) List(owner *TreeNode) []*TreeNode {

	tmp := make([]*TreeNode, len(l.Neighbors))
	copy(tmp, l.Neighbors)

	ok, index := l.IsExists(owner)
	if ok {
		//		fmt.Println(tmp, "1111", tmp[index])

		return append(tmp[:index], tmp[index+1:]...)
	}

	return tmp
}

func (l *TreeNodeList) Len(owner *TreeNode) int64 {
	var number int64 = 0
	for _, node := range l.Neighbors {
		if owner.Number == node.Number || (owner.GetComeNode() != nil && owner.GetComeNode().Number == node.Number) {
			continue
		}

		node.SetComeNode(owner)
		number += node.Len()
	}

	return number
}

func (l TreeNodeList) IsExists(n *TreeNode) (bool, int64) {
	isExists := false
	var index int64 = 0
	for i := range l.Neighbors {
		if l.Neighbors[i].Number == n.Number {
			isExists = true
			index = int64(i)
			break
		}
	}

	return isExists, index
}

func (l TreeNodeList) IsExistsByNumber(number int64) (bool, int64) {
	isExists := false
	var index int64 = 0
	for i := range l.Neighbors {
		if l.Neighbors[i].Number == number {
			isExists = true
			index = int64(i)
			break
		}
	}

	return isExists, index
}

func (l *TreeNodeList) FindByNumber(number int64, owner *TreeNode) *TreeNode {
	ok, index := l.IsExistsByNumber(number)
	var foundNode *TreeNode
	if ok {
		foundNode = l.Neighbors[index]
		foundNode.SetComeNode(owner)
	} else {
		for _, node := range l.Neighbors {
			if owner.Number == node.Number || (owner.GetComeNode() != nil && owner.GetComeNode().Number == node.Number) {
				continue
			}

			node.SetComeNode(owner)
			foundNode = node.FindByNumber(number)
			if foundNode != nil {
				break
			}
		}
	}

	return foundNode
}

// Solution

func (l *TreeNodeList) UpdateQuery(owner *TreeNode, valueV int64, valueK int64) {

	for _, node := range l.Neighbors {
		if node.Number > owner.Number {
			node.SetComeNode(owner)
			node.SetDistance(owner.GetDistance())
			node.Calculate(valueV, valueK)

			node.UpdateQuery(valueV, valueK)
		}
	}
}

func (l *TreeNodeList) Sum(owner *TreeNode, toNode *TreeNode) int64 {
	sum := toNode.Value

	var tmpNode *TreeNode = toNode.ComeNode
	for {
		if tmpNode == nil || owner.Number == tmpNode.Number {
			break
		}

		sum += tmpNode.Value
		tmpNode = tmpNode.ComeNode
	}

	return sum
}
