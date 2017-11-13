package tree

// import (
// 	"fmt"
// )

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
