package tree

import (
	//"encoding/json"
	"fmt"
)

type TreeNode struct {
	NodeList *TreeNodeList `json:"list"`
	ComeNode *TreeNode     `json:"-"`
	Distance int64
	Value    int64
	Number   int64
}

func NewTreeNode(number int64) *TreeNode {
	n := TreeNode{}
	n.NodeList = NewTreeNodeList()
	n.Number = number
	n.NodeList.Insert(&n)

	return &n
}

func (node *TreeNode) Insert(n *TreeNode) {
	node.NodeList.Insert(n)
	n.NodeList.Insert(node)
}

func (node *TreeNode) Remove(n *TreeNode) {
	node.NodeList.Remove(n, node)
}

func (node *TreeNode) List() map[int64]*TreeNode {
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
	//fmt.Println(node, "node")
	toNode := node.FindByNumber(toNumber)
	//fmt.Println(toNode, "toNode")

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
	// b, err := json.Marshal(node)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return ""
	// }
	// return string(b)
	return fmt.Sprintf("{\"number\": %d, \"value\": %d, \"distance\": %d}", node.Number, node.Value, node.Distance)
}
