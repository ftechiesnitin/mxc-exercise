# N-ARY Tree Representation in Golang
N-ary tree representation in giolang, with following functions for example. It is just a basic example, open to PR's.

```
                   1
                 / | \
                /  |   \
              2    3     4
             / \       / | \
            5    6    7  8  9
           /   / | \ 
          10  11 12 13
```

### Methods
```go
package main

import (
	"log"

	"mxc-exercise/ntree"
)

func main() {
    // AddCategory: to Add a Node to the Root Node with Parent Node as empty String
	err := ntree.AddCategory("a", "")
	if err != nil {
		log.Println(err)
    }
	err = ntree.AddCategory("b", "")
	if err != nil {
		log.Println(err)
    }
    // AddCategory: to Add a Node to the Parent Node as a second Argument
	err = ntree.AddCategory("c", "b")
	if err != nil {
		log.Println(err)
	}
	err = ntree.AddCategory("d", "a")
	if err != nil {
		log.Println(err)
    }
    
    // GetAllChildren: Get All Children of a Node
	err = ntree.GetAllChildren("a")
	if err != nil {
		log.Println(err)
	}
}
```