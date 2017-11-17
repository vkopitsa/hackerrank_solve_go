package tree

import (
	"fmt"
)

type TreeNodeCache struct {
	Items    map[int64]map[int64]int64
	Count    int64
	allItems map[int64]bool
}

func NewTreeNodeCache() *TreeNodeCache {
	c := TreeNodeCache{}
	c.Items = make(map[int64]map[int64]int64, 10000)
	c.allItems = make(map[int64]bool, 10000)
	return &c
}

func (cache *TreeNodeCache) Add(owner int64, target int64, path int64) bool {

	//cache.allItems = append(cache.allItems, owner, target, path)
	cache.allItems[owner] = true
	cache.allItems[target] = true
	cache.allItems[path] = true

	_, ok := cache.Items[owner]
	if !ok {
		cache.Items[owner] = make(map[int64]int64, 10000)
	}

	_, ok2 := cache.Items[owner][target]
	if !ok2 {
		cache.Items[owner][target] = path
		cache.Count++
		return true
	}

	return false
}

func (cache *TreeNodeCache) GetPath(owner int64, target int64) int64 {
	return cache.Items[owner][target]
}

func (cache *TreeNodeCache) IsPath(root int64, target int64) (ok bool) {
	_, ok = cache.Items[root][target]
	return
}

func (cache *TreeNodeCache) IsRoot(root int64) (ok bool) {
	_, ok = cache.Items[root]
	return
}

func (cache *TreeNodeCache) GetList(owner int64) map[int64]int64 {
	return cache.Items[owner]
}

func (cache *TreeNodeCache) Remove(owner int64) {
	delete(cache.Items, owner)
}

func (cache *TreeNodeCache) Build(owner *TreeNode) {
	for number, _ := range owner.List() {
		cache.Add(owner.Number, number, number)
		// 		if !ok {
		// 			break
		// 		}
		// 		cache.buildDeep(node, owner)
		//cache.buildDeep2(node, owner)
		//		cache.buildDeep2(owner, node)
	}

	for _, node := range owner.List() {
		//fmt.Println(cache.GetList(owner.Number), owner.Number, "Items1")

		cache.buildDeep(node, owner)

		// 		i := 0
		// 		for number, _ := range cache.allItems {
		// 			// 			if cache.IsPath(owner.Number, number) {
		// 			// 				continue
		// 			// 			}
		// 			if i >= len(cache.allItems) {
		// 				//break
		// 			}
		// 			cache.Add(owner.Number, number, node.Number)
		// 			i++
		// 		}

	}

	// 	for number := range cache.GetList(owner.Number) {
	// 		cache.Add(node.Number, number, owner.Number)
	// 	}

	fmt.Println(cache.Items, "Items")
	fmt.Println(cache.Count, "Count")
	fmt.Println(cache.allItems, "cache.allItems")

}

// private method

func (cache *TreeNodeCache) buildDeep(node *TreeNode, owner *TreeNode) {
	if cache.IsRoot(node.Number) {
		return
	}
	// 	fmt.Println(node, "node")
	//fmt.Println(cache.Items, "Items")
	// 	fmt.Println(cache.Count, "Count")
	for number, _ := range node.List() {
		cache.Add(node.Number, number, number)

		// cache.buildDeep2(owner, node)

		// cache.buildDeep2(node, owner)

		// 		if !ok {
		// 			break
		// 		}
		//cache.buildDeep(n, node)
		// 		cache.buildDeep2(n, node)
		// 		cache.buildDeep2(node, n)
	}

	//fmt.Println(node.List(), "node")

	for _, n := range node.List() {
		fmt.Println(cache.GetList(owner.Number), owner.Number, "Items")
		cache.buildDeep(n, node)

		// 		i := 0
		// 		for number, _ := range cache.allItems {
		// 			// 			if cache.IsPath(node.Number, number) {
		// 			// 				continue
		// 			// 			}

		// 			if i >= len(cache.allItems) {
		// 				//break
		// 			}
		// 			cache.Add(node.Number, number, owner.Number)
		// 			i++
		// 		}
	}

	// 	for _, number := range cache.allItems {
	// 		if cache.IsPath(node.Number, number) {
	// 			continue
	// 		}
	// 		cache.Add(node.Number, number, owner.Number)
	// 	}

	//cache.buildDeep2(node, owner)
	//cache.buildDeep2(owner, node)
	//cache.buildDeep2(node, owner)
}

func (cache *TreeNodeCache) buildDeep2(node *TreeNode, owner *TreeNode) {
	for number, _ := range cache.GetList(node.Number) {
		ok := cache.Add(owner.Number, number, node.Number)
		if !ok {
			break
		}
		//cache.buildDeep3(n, node)
		//cache.buildDeep2(n, node)
		// 		cache.buildDeep2(node, n)
	}
}
