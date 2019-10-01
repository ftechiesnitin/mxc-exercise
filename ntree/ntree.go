package ntree

import (
	"errors"
	"log"
)

// Node : Defining a N-ary based struct
type Node struct {
	name     string
	children []*Node
}

// root declaring a predefined root Node
var root = Node{
	name: "root",
}

// AddCategory : Add a category node to the tree
func AddCategory(child, parent string) error {
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

// RemoveCategory : Add a category node to the tree
// func RemoveCategory(child, parent string) error {
// 	category := Node{
// 		name: child,
// 	}
// 	if len(parent) == 0 {
// 		root.children = append(root.children, &category)
// 		return nil
// 	}

// 	childNode := findByName(&root, child)
// 	if childNode != nil {
// 		return errors.New("Category already exists")
// 	}

// 	parentNode := findByName(&root, parent)
// 	if parentNode != nil {
// 		parentNode.children = append(parentNode.children, &category)
// 	}

// 	return nil
// }

// GetAllChildren : Get All Children of a specific category node
func GetAllChildren(name string) error {
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

// findByName : Find a particular node in tree
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
