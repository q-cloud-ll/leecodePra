package leecode_tree

// RightSideView 二叉树的右侧视图
func RightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(ans) == depth {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return ans
}

func preorderTraversal(root *TreeNode) []int {
	var list []int
	if root != nil {
		list = append(list, root.Val)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorder1(root, &res)
	return res
}

func inorder1(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder1(root.Left, res)
	*res = append(*res, root.Val)
	inorder1(root.Right, res)
}

func ConstructFromPrePost(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}

	root.Left = ConstructFromPrePost(preorder[1:1+rootIndex], inorder[:rootIndex])
	root.Right = ConstructFromPrePost(preorder[1+rootIndex:], inorder[rootIndex+1:])

	return root
}
