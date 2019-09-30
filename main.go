package main

import "log"

type Node struct {
	name     string
	children []*Node
}

var root = Node{
	name: "root",
}

func addCategory(child, parent string) Node {
	category := Node{
		name: child,
	}
	if len(parent) == 0 {
		root.children = append(root.children, &category)
		return root
	}

	for i, child := range root.children {
		parentNode := findByName(child, parent)
		if parentNode != nil && len(parentNode.children) == 0 {
			parentNode.children = append(parentNode.children, &category)
			root.children[i] = parentNode
			break
		}
	}

	return root
}

func findByName(parentNode *Node, name string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, parentNode)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.name == name {
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
	root = addCategory("a", "")
	root = addCategory("b", "")
	root = addCategory("c", "b")
	log.Println(root.children[1])
}
