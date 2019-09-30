package main

import (
	"fmt"
)

type Node struct {
	name     string
	children []*Node
}

// func addCategory(child, parent string) {
// 	if len(parent) == 0 {

// 	}
// }

func findByName(root *Node, id string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.name == id {
			return nextUp
		}
		if len(nextUp.children) > 0 {
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	return nil
}

func main() {
	b := Node{
		name: "b",
	}
	a := Node{
		name:     "a",
		children: []*Node{&b},
	}
	c := Node{
		name: "c",
	}
	d := Node{
		name: "d",
	}
	e := Node{
		name: "e",
	}
	f := Node{
		name: "f",
	}
	a.children = append(a.children, &c)
	c.children = append(c.children, &d)
	c.children = append(c.children, &e)
	c.children = append(c.children, &f)
	data := findByName(&a, "c")
	fmt.Println(data)
	fmt.Println(a, c)
}
