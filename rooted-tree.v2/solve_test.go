package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func _TestSolve(t *testing.T) {
	nodes := []row_node{
		row_node{X: 1, Y: 2},
		row_node{X: 2, Y: 3},
		row_node{X: 2, Y: 4},
		row_node{X: 2, Y: 5},
		row_node{X: 5, Y: 6},
		row_node{X: 6, Y: 7},
	}

	queries := []row_update_query{
		row_update_query{Type: rune('U'), TorA: 5, VorB: 10, KorNothing: 2},
		row_update_query{Type: rune('U'), TorA: 4, VorB: 5, KorNothing: 3},
		row_update_query{Type: rune('Q'), TorA: 1, VorB: 7},
		row_update_query{Type: rune('U'), TorA: 6, VorB: 7, KorNothing: 4},
		row_update_query{Type: rune('Q'), TorA: 2, VorB: 7},
		row_update_query{Type: rune('Q'), TorA: 1, VorB: 4},
		row_update_query{Type: rune('Q'), TorA: 2, VorB: 4},
	}

	solve := solve(1, nodes, queries)
	result := `36
54
5
5
`
	if solve != result {
		t.Errorf("Expected %s got %s", result, solve)
	}
}

func TestSolve2(t *testing.T) {
	fmt.Println(time.Now(), "Start")
	file, err := os.Open("input011.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var nodeNumber, queryNumber, rootNumber int64
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%v %v %v", &nodeNumber, &queryNumber, &rootNumber)

	//fmt.Println(nodeNumber, queryNumber, rootNumber, "init")

	nodes := make([]row_node, nodeNumber-1)
	for i := range nodes {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%v %v", &nodes[i].X, &nodes[i].Y)
		//fmt.Println(nodes[i], "nodes[i]")
	}

	queries := make([]row_update_query, queryNumber)
	for i := range queries {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%c %v %v %v", &queries[i].Type, &queries[i].TorA, &queries[i].VorB, &queries[i].KorNothing)
		//fmt.Println(queries[i], "queries[i]")
	}

	// fmt.Println(nodes[len(nodes)-1], len(nodes), "len(nodes)")
	// fmt.Println(nodes[0], len(nodes), "nodes[0]")
	// fmt.Println(queries[len(queries)-1], len(queries), "len(queries)")
	// fmt.Println(queries[0], len(queries), "queries[0]")

	solve := solve(rootNumber, nodes, queries)
	fmt.Printf("%s", solve)

	// var result int64 = 2510535321
	// if solve != result {
	// 	t.Errorf("Expected %d got %d", result, solve)
	// }
}
