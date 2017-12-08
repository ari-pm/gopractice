package main

import (
	"fmt"
)

// Since this BTree is a good candidate that can be added to the other packages, the fields need to be capitalized so theyca n be exported.
type BTree struct {
	Val   int
	Left  *BTree // does the left and right here needs to be upper or lower case?
	Right *BTree
}

func (bt *BTree) init() {
	bt.Val = 0
	bt.Left = nil
	bt.Right = nil
}

func main() {
	btree := new(BTree)
	btree.init()
	leftbt := new(BTree)
	leftbt.init()
	leftbt.Val = 1
	rightbt := new(BTree)
	rightbt.init()
	rightbt.Val = 3

	btree.Left = leftbt
	btree.Right = rightbt
	// Now traverse the tree in different ways:
	fmt.Println("inOrderTrav")
	inOrderTrav(btree)
	fmt.Println("preOrderTrav")
	preOrderTrav(btree)
	fmt.Println("postOrderTrav")
	postOrderTrav(btree)
}

func inOrderTrav(tree *BTree) {
	if tree == nil {
		return
	}
	inOrderTrav(tree.Left)
	fmt.Println(tree.Val)
	inOrderTrav(tree.Right)
}

func preOrderTrav(tree *BTree) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	inOrderTrav(tree.Left)
	inOrderTrav(tree.Right)
}

func postOrderTrav(tree *BTree) {
	if tree == nil {
		return
	}
	inOrderTrav(tree.Left)
	inOrderTrav(tree.Right)
	fmt.Println(tree.Val)
}
