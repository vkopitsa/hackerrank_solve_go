//
// https://www.hackerrank.com/challenges/rooted-tree/problem
//
package main

import (
	. "./tree"
	"fmt"
	"time"
)

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
	fmt.Println(time.Now(), "Start")

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
	//var tree *TreeNode

	m := make(map[int64]*TreeNode, len(nodes))

	for _, n := range nodes {
		//var x, y *TreeNode
		// X
		x, ok := m[n.X]
		if !ok {
			x = NewTreeNode(n.X)
			m[n.X] = x
		}
		// if tree != nil {
		// 	x = tree.FindByNumber(n.X)
		// }

		// if x == nil {
		// 	x = NewTreeNode(n.X)
		// }

		// Y
		y, ok := m[n.Y]
		if !ok {
			y = NewTreeNode(n.Y)
			m[n.Y] = y
		}
		// if tree != nil {
		// 	y = tree.FindByNumber(n.Y)
		// }

		// if y == nil {
		// 	y = NewTreeNode(n.Y)
		// }

		// if tree == nil {
		// 	x.Insert(y)
		// 	tree = x
		// } else {
		// 	x.Insert(y)
		// }
		x.Insert(y)

		// fmt.Println(x, x.List(), "-------x")
		// fmt.Println(y, y.List(), "--------y")
	}

	cache := NewTreeNodeCache()
	cache.Build(m[rootNumber])

	// rootNode := tree.FindByNumber(rootNumber)
	//rootNode := m[rootNumber]
	fmt.Println(time.Now(), "Build end")
	for i, q := range queries {

		if i%100 == 0 {
			fmt.Println(time.Now(), "hundred", i)
		}

		// update
		if q.Type == 85 {
			// fmt.Println(tree, "tree")
			// fmt.Println(rootNumber, "rootNumber")
			// fmt.Println(rootNode, "rootNode")
			//foundNode := rootNode.FindByNumber(q.TorA)
			foundNode := m[q.TorA]
			//fmt.Println(q, "Update Start")
			foundNode.Reset().Update(q.VorB, q.KorNothing)
			//fmt.Println("Update End")
		} else {
			// query
			//sum := rootNode.FindByNumber(q.TorA).Sum(q.VorB)
			foundNode := m[q.TorA]
			//fmt.Println(q, "Sum Start")
			sum := foundNode.FindByNumber(q.TorA).Sum(q.VorB)
			//fmt.Println(sum, "Sum End")
			r += fmt.Sprintf("%d\n", sum)
		}
	}

	return r
}
