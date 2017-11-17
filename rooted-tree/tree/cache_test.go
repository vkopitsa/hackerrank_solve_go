package tree

import (
	//"fmt"
	//"math/rand"
	"testing"
)

func TestCreate(t *testing.T) {
	cache := NewTreeNodeCache()
	if cache == nil {
		t.Errorf("Expected %d got %d", cache, cache)
	}
}

func TestAddGetRemoveTo(t *testing.T) {
	cache := NewTreeNodeCache()

	cache.Add(1, 1, 1)
	path := cache.GetPath(1, 1)

	if path != 1 {
		t.Errorf("Expected %d got %d", 1, path)
	}

	path2 := cache.GetPath(2, 2)
	if path2 != 0 {
		t.Errorf("Expected %d got %d", 0, path2)
	}

	cache.Remove(1)
	path3 := cache.GetPath(1, 1)
	if path3 != 0 {
		t.Errorf("Expected %d got %d", 0, path3)
	}
}

func TestBuild(t *testing.T) {
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

	cache := NewTreeNodeCache()
	cache.Build(secondNode)

	path5 := cache.GetPath(2, 5)
	if path5 != 5 {
		t.Errorf("Expected %d got %d", 5, path5)
	}

	path7 := cache.GetPath(2, 7)
	if path7 != 5 {
		t.Errorf("Expected %d got %d", 5, path7)
	}
}
