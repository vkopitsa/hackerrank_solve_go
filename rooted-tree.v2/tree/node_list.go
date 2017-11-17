package tree

// import (
// 	"fmt"
// )

type TreeNodeList struct {
	//Neighbors []*TreeNode
	Neighbors map[int64]*TreeNode
}

func NewTreeNodeList() *TreeNodeList {
	l := TreeNodeList{}
	//l.Neighbors = make([]*TreeNode, 0, 100)
	l.Neighbors = make(map[int64]*TreeNode)
	return &l
}

func (l *TreeNodeList) Insert(n *TreeNode) {
	//ok, _ := l.IsExists(n)
	_, ok := l.Neighbors[n.Number]
	if !ok {
		//l.Neighbors = append(l.Neighbors, n)
		l.Neighbors[n.Number] = n
	}
}

func (l *TreeNodeList) Remove(n *TreeNode, owner *TreeNode) {
	//ok, index := l.IsExists(n)
	_, ok := l.Neighbors[n.Number]
	if ok {
		//l.Neighbors = append(l.Neighbors[:index], l.Neighbors[index+1:]...)
		delete(l.Neighbors, n.Number)
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

func (l TreeNodeList) List(owner *TreeNode) map[int64]*TreeNode {

	newMap := make(map[int64]*TreeNode, len(l.Neighbors)-1)
	//copy(tmp, l.Neighbors)

	for k, v := range l.Neighbors {
		if k == owner.Number {
			continue
		}
		newMap[k] = v
	}

	// _, ok := tmp[owner.Number]
	// if ok {
	// 	delete(tmp, owner.Number)
	// }
	// ok, index := l.IsExists(owner)
	// if ok {
	// 	//		fmt.Println(tmp, "1111", tmp[index])

	// 	return append(tmp[:index], tmp[index+1:]...)
	// }

	return newMap
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

// func (l TreeNodeList) IsExists(n *TreeNode) (bool, int64) {
// 	isExists := false
// 	var index int64 = 0
// 	for i := range l.Neighbors {
// 		if l.Neighbors[i].Number == n.Number {
// 			isExists = true
// 			index = int64(i)
// 			break
// 		}
// 	}

// 	return isExists, index
// }

// func (l TreeNodeList) IsExistsByNumber(number int64) (bool, int64) {
// 	isExists := false
// 	var index int64 = 0
// 	for i := range l.Neighbors {
// 		if l.Neighbors[i].Number == number {
// 			isExists = true
// 			index = int64(i)
// 			break
// 		}
// 	}

// 	return isExists, index
// }

func (l *TreeNodeList) FindByNumber(number int64, owner *TreeNode) *TreeNode {
	//ok, index := l.IsExistsByNumber(number)
	//var foundNode *TreeNode
	foundNode, ok := l.Neighbors[number]
	//fmt.Println(foundNode, ok, "foundNode, ok", number, l.Neighbors)
	if ok {
		//foundNode = l.Neighbors[index]
		foundNode.SetComeNode(owner)
	} else {
		for _, node := range l.Neighbors {
			if owner.Number == node.Number || (owner.GetComeNode() != nil && owner.GetComeNode().Number == node.Number) {
				// fmt.Println(number, "number continue")
				// fmt.Println(node, "node continue")
				continue
			}

			// fmt.Println(number, "number ")
			// fmt.Println(node, "node ")

			node.SetComeNode(owner)
			//foundNode = node.FindByNumber(number)
			//fmt.Println(node, "foundNode = node")

			foundNode = node.FindByNumber(number)
			//fmt.Println(foundNode, node, "foundNode = node.FindByNumber(number)")

			if foundNode != nil {
				break
			}
		}
	}
	//fmt.Println(foundNode, "return ", number)

	return foundNode
}

func (l *TreeNodeList) ToString(owner *TreeNode) string {
	//return fmt.Sprintf("(number: %d, value: %d, distance: %d, list: %s)", node.Number, node.Value, node.Distance, node.NodeList)
	s := ""
	for _, node := range l.Neighbors {
		if owner.Number == node.Number || (owner.GetComeNode() != nil && owner.GetComeNode().Number == node.Number) {
			continue
		}
		s += node.String() + " "

	}

	return s + ""
}

// func (l *TreeNodeList) MarshalJSON() ([]byte, error) {
// 	buffer := bytes.NewBufferString("{")
// 	length := len(l.Neighbors)
// 	count := 0
// 	for key, value := range l.Neighbors {
// 		jsonValue, err := json.Marshal(value)
// 		if err != nil {
// 			return nil, err
// 		}
// 		buffer.WriteString(fmt.Sprintf("\"%d\":%s", key, string(jsonValue)))
// 		count++
// 		if count < length {
// 			buffer.WriteString(",")
// 		}
// 	}
// 	buffer.WriteString("}")
// 	return buffer.Bytes(), nil
// }

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
	//fmt.Println(toNode, "toNode")
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
