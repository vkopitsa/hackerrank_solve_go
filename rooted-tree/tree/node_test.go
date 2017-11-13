package tree

import (
	//"fmt"
	"math/rand"
	"testing"
)

func TestNewNode(t *testing.T) {
	var solve int64 = 1
	n := NewTreeNode(solve)
	if n.Number != solve {
		t.Errorf("Expected %d got %d", n.Number, solve)
	}
}

func TestAddNodeTo(t *testing.T) {
	rootNode := NewTreeNode(1)

	var solve int64 = 7

	var newNode *TreeNode
	for v := range rand.Perm(7)[1:] {
		newNode = NewTreeNode(int64(v + 2))
		rootNode.Insert(newNode)
	}

	if rootNode.Len() != solve {
		t.Errorf("Expected %d got %d", rootNode.Len(), solve)
	}
}

func TestRemoveNodeFrom(t *testing.T) {
	rootNode := NewTreeNode(1)

	secondNode := NewTreeNode(2)
	rootNode.Insert(secondNode)

	thirtNode := NewTreeNode(3)
	fourthNode := NewTreeNode(4)
	fifthNode := NewTreeNode(5)
	secondNode.Insert(thirtNode)
	secondNode.Insert(fourthNode)
	secondNode.Insert(fifthNode)

	sixthNode := NewTreeNode(6)
	fifthNode.Insert(sixthNode)

	seventhNode := NewTreeNode(7)
	sixthNode.Insert(seventhNode)

	if rootNode.Len() != 7 {
		t.Errorf("Expected %d got %d", rootNode.Len(), 7)
	}

	sixthNode.Remove(seventhNode)
	if rootNode.Len() != 6 {
		t.Errorf("Expected %d got %d", rootNode.Len(), 6)
	}

	rootNode.Remove(sixthNode)
	if rootNode.Len() != 5 {
		t.Errorf("Expected %d got %d", rootNode.Len(), 5)
	}

	rootNode.Remove(secondNode)
	if rootNode.Len() != 1 {
		t.Errorf("Expected %d got %d", rootNode.Len(), 1)
	}
}

func TestList(t *testing.T) {
	rootNode := NewTreeNode(1)

	secondNode := NewTreeNode(2)
	rootNode.Insert(secondNode)

	thirtNode := NewTreeNode(3)
	fourthNode := NewTreeNode(4)
	fifthNode := NewTreeNode(5)
	secondNode.Insert(thirtNode)
	secondNode.Insert(fourthNode)
	secondNode.Insert(fifthNode)

	sixthNode := NewTreeNode(6)
	fifthNode.Insert(sixthNode)

	seventhNode := NewTreeNode(7)
	sixthNode.Insert(seventhNode)

	if len(rootNode.List()) != 1 {
		t.Errorf("Expected %d got %d", len(rootNode.List()), 1)
	}

	if len(secondNode.List()) != 4 {
		t.Errorf("Expected %d got %d", len(secondNode.List()), 4)
	}

	rootNode.Remove(fifthNode)
	if len(secondNode.List()) != 3 {
		t.Errorf("Expected %d got %d", len(secondNode.List()), 3)
	}
}

func TestTreeNodes(t *testing.T) {
	rootNode := NewTreeNode(1)

	secondNode := NewTreeNode(2)
	rootNode.Insert(secondNode)

	thirtNode := NewTreeNode(3)
	fourthNode := NewTreeNode(4)
	fifthNode := NewTreeNode(5)
	secondNode.Insert(thirtNode)
	secondNode.Insert(fourthNode)
	secondNode.Insert(fifthNode)

	sixthNode := NewTreeNode(6)
	fifthNode.Insert(sixthNode)

	seventhNode := NewTreeNode(7)
	sixthNode.Insert(seventhNode)

	if rootNode.Len() != 7 {
		t.Errorf("Expected %d got %d", rootNode.Len(), 7)
	}
}

func TestTreeFoundNodeByNumber(t *testing.T) {
	rootNode := NewTreeNode(1)

	secondNode := NewTreeNode(2)
	rootNode.Insert(secondNode)

	thirtNode := NewTreeNode(3)
	fourthNode := NewTreeNode(4)
	fifthNode := NewTreeNode(5)
	secondNode.Insert(thirtNode)
	secondNode.Insert(fourthNode)
	secondNode.Insert(fifthNode)

	sixthNode := NewTreeNode(6)
	fifthNode.Insert(sixthNode)

	seventhNode := NewTreeNode(7)
	sixthNode.Insert(seventhNode)

	foundNode6 := rootNode.FindByNumber(6)

	if foundNode6.Number != 6 {
		t.Errorf("Expected %d got %d", foundNode6.Number, 6)
	}

	foundNode4 := rootNode.FindByNumber(4)

	if foundNode4.Number != 4 {
		t.Errorf("Expected %d got %d", foundNode4.Number, 4)
	}

	foundNode7 := fifthNode.FindByNumber(7)

	if foundNode7.Number != 7 {
		t.Errorf("Expected %d got %d", foundNode7.Number, 7)
	}
}

func TestSolve(t *testing.T) {
	rootNode := NewTreeNode(1)

	secondNode := NewTreeNode(2)
	rootNode.Insert(secondNode)

	thirtNode := NewTreeNode(3)
	fourthNode := NewTreeNode(4)
	fifthNode := NewTreeNode(5)
	secondNode.Insert(thirtNode)
	secondNode.Insert(fourthNode)
	secondNode.Insert(fifthNode)

	sixthNode := NewTreeNode(6)
	fifthNode.Insert(sixthNode)

	seventhNode := NewTreeNode(7)
	sixthNode.Insert(seventhNode)

	//Values of Nodes after U 5 10 2: [0 0 0 0 10 12 14].
	foundNode5 := fifthNode.FindByNumber(5)
	foundNode5.Reset().Update(10, 2)

	if foundNode5.Value != 10 {
		t.Errorf("Expected %d got %d", foundNode5.Value, 10)
	}

	foundNode6 := fifthNode.FindByNumber(6)
	if foundNode6.Value != 12 {
		t.Errorf("Expected %d got %d", foundNode6.Value, 12)
	}

	foundNode7 := fifthNode.FindByNumber(7)
	if foundNode7.Value != 14 {
		t.Errorf("Expected %d got %d", foundNode7.Value, 14)
	}

	// Values of Nodes after U 4 5 3: [0 0 0 5 10 12 14].
	foundNode4 := fifthNode.FindByNumber(4)
	foundNode4.Reset().Update(5, 3)

	if foundNode4.Value != 5 {
		t.Errorf("Expected %d got %d", foundNode4.Value, 5)
	}

	if foundNode5.Value != 10 {
		t.Errorf("Expected %d got %d", foundNode5.Value, 10)
	}

	if foundNode6.Value != 12 {
		t.Errorf("Expected %d got %d", foundNode6.Value, 12)
	}

	if foundNode7.Value != 14 {
		t.Errorf("Expected %d got %d", foundNode7.Value, 14)
	}

	// Q 1 7
	// Sum of the Nodes from 1 to 7: 0 + 0 + 10 + 12 + 14 = 36.
	sum1_7 := rootNode.FindByNumber(1).Sum(7)
	if sum1_7 != 36 {
		t.Errorf("Expected %d got %d", sum1_7, 36)
	}

	// Values of Nodes after U 6 7 4: [0 0 0 5 10 19 25].
	foundNode6.Reset().Update(7, 4)

	if foundNode4.Value != 5 {
		t.Errorf("Expected %d got %d", foundNode4.Value, 5)
	}

	if foundNode5.Value != 10 {
		t.Errorf("Expected %d got %d", foundNode5.Value, 10)
	}

	if foundNode6.Value != 19 {
		t.Errorf("Expected %d got %d", foundNode6.Value, 12)
	}

	if foundNode7.Value != 25 {
		t.Errorf("Expected %d got %d", foundNode7.Value, 14)
	}

	// TODO:
	// Q 2 7
	// Sum of the Nodes from 2 to 7: 0 + 10 + 19 + 25 = 54.
	sum2_7 := rootNode.FindByNumber(2).Sum(7)
	if sum2_7 != 54 {
		t.Errorf("Expected %d got %d", sum2_7, 54)
	}

	// Q 1 4
	// Sum of the Nodes from 1 to 4: 0 + 0 + 5 = 5.
	sum1_4 := rootNode.FindByNumber(1).Sum(4)
	if sum1_4 != 5 {
		t.Errorf("Expected %d got %d", sum1_4, 5)
	}

	// Q 2 4
	// Sum of the Nodes from 2 to 4: 0 + 5 = 5.
	sum2_4 := rootNode.FindByNumber(2).Sum(4)
	if sum2_4 != 5 {
		t.Errorf("Expected %d got %d", sum2_4, 5)
	}
}
