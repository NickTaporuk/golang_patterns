package main

import (
	"github.com/NickTaporuk/golang_patterns/structure/composite"
	"fmt"
)

func main() {

	root := composite.Tree{
		LeafValue: 0,
		Right:&composite.Tree{
			LeafValue: 5,
			Right: &composite.Tree{ 6, nil, nil },
			Left: nil,
		},
		Left:&composite.Tree{ 4, nil, nil },
	}

	fmt.Println(root.Right.Right.LeafValue)
}
