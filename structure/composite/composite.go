/*
 * Copyright (c) 2018. Created by Nikolay Kuropatkin
 */

package composite

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}