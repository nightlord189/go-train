package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) ListToString() string {
	var result string
	next := n
	for next != nil {
		result += strconv.Itoa(next.Value)
		next = next.Next
	}
	return result
}

func (n *Node) ReverseList() *Node {
	var prev *Node
	current := n
	next := n.Next
	for next != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func CreateList(values []int) *Node {
	nodes := make([]Node, len(values))
	for index, val := range values {
		nodes[index] = Node{
			Value: val,
		}
	}
	for i := 0; i+1 < len(values); i++ {
		nodes[i].Next = &nodes[i+1]
	}
	return &nodes[0]
}

func main() {
	fmt.Println("start")
	list := CreateList([]int{1, 2, 3, 4})
	fmt.Println(list.ListToString())
	fmt.Println(list.ReverseList().ListToString())
}
