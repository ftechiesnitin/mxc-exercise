package main

import (
	"log"

	"mxc-exercise/ntree"
)

func main() {
	err := ntree.AddCategory("a", "")
	if err != nil {
		log.Println(err)
	}
	err = ntree.AddCategory("b", "")
	if err != nil {
		log.Println(err)
	}
	err = ntree.AddCategory("c", "b")
	if err != nil {
		log.Println(err)
	}
	err = ntree.AddCategory("d", "a")
	if err != nil {
		log.Println(err)
	}
	err = ntree.AddCategory("e", "a")
	if err != nil {
		log.Println(err)
	}
	err = ntree.AddCategory("f", "a")
	if err != nil {
		log.Println(err)
	}
	err = ntree.AddCategory("g", "")
	if err != nil {
		log.Println(err)
	}
	err = ntree.GetAllChildren("a")
	if err != nil {
		log.Println(err)
	}
	err = ntree.GetAllChildren("root")
	if err != nil {
		log.Println(err)
	}
}
