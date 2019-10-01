package main

import (
	"errors"
	"log"
)

type Node struct {
	name     string
	children []*Node
}

var root = Node{
	name: "root",
}

func addCategory(child, parent string) error {
	category := Node{
		name: child,
	}
	if len(parent) == 0 {
		root.children = append(root.children, &category)
		return nil
	}

	childNode := findByName(&root, child)
	if childNode != nil {
		return errors.New("Category already exists")
	}

	parentNode := findByName(&root, parent)
	if parentNode != nil {
		parentNode.children = append(parentNode.children, &category)
	}

	return nil
}

func getAllChildren(name string) error {
	node := findByName(&root, name)
	if node == nil {
		return errors.New("Category does not exists")
	}

	if len(node.children) > 0 {
		log.Println("Children of Category:", name)
		for _, child := range node.children {
			log.Println(child.name)
		}
	}

	return nil
}

func findByName(parentNode *Node, name string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, parentNode)

	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
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
			n--
		}
	}
	return nil
}

func main() {
	err := addCategory("a", "")
	if err != nil {
		log.Println(err)
	}
	err = addCategory("b", "")
	if err != nil {
		log.Println(err)
	}
	err = addCategory("c", "b")
	if err != nil {
		log.Println(err)
	}
	err = addCategory("d", "a")
	if err != nil {
		log.Println(err)
	}
	err = addCategory("e", "a")
	if err != nil {
		log.Println(err)
	}
	err = addCategory("f", "a")
	if err != nil {
		log.Println(err)
	}
	err = addCategory("g", "")
	if err != nil {
		log.Println(err)
	}
	err = getAllChildren("a")
	if err != nil {
		log.Println(err)
	}
	err = getAllChildren("root")
	if err != nil {
		log.Println(err)
	}
}
