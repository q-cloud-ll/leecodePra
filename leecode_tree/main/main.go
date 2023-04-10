package main

import (
	"LeecodePra/leecode_tree"
	"fmt"
)

//func main() {
//	var t leecode_tree.TreeNode
//	//root := t.CreateTree(5)
//	arr := []int{1, 3, 5, 7, 4, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8}
//	tree := t.BuildBinaryTree(arr, 0)
//	//list := t.PreorderTraversal(tree)
//	//for _, node := range list {
//	//	fmt.Print(node, " ")
//	//}
//	//root := t.CreateTree(8)
//
//	res := t.InorderTraversal(tree)
//	fmt.Println(res)
//
//}

func main() {
	//pre := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}
	//res := leecode_tree.ConstructFromPrePost(pre, inorder)
	var t leecode_tree.TreeNode
	arr := []int{1, 3, 5, 7, 4, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8}
	tree := t.BuildBinaryTree(arr, 0)
	list := t.InorderTraversal(tree)
	fmt.Println(list)
}
