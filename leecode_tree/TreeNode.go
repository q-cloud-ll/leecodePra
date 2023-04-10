package leecode_tree

type TreeMethod interface {
	CreateTree(data int) *TreeNode
	InsertNode(root *TreeNode, val int) *TreeNode
	TraverseNode(root *TreeNode)
	BuildBinaryTree(arr []int, i int) *TreeNode
	PreorderTraversal(root *TreeNode) []*TreeNode
	InorderTraversal(root *TreeNode) []int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) CreateTree(data int) *TreeNode {
	return &TreeNode{data, nil, nil}
}

func (n *TreeNode) InsertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return n.CreateTree(val)
	} else if val <= root.Val {
		root.Left = n.InsertNode(root.Left, val)
	} else {
		root.Right = n.InsertNode(root.Right, val)
	}
	return root
}

// BuildBinaryTree 1 2 3 4 5 6 7 8
func (n *TreeNode) BuildBinaryTree(arr []int, i int) *TreeNode {
	if i < len(arr) && arr[i] != 0 {
		return &TreeNode{arr[i], n.BuildBinaryTree(arr, 2*i+1), n.BuildBinaryTree(arr, 2*i+2)}
	}
	return nil
}

// TraverseNode 遍历节点（中序遍历）
//func (n *TreeNode) TraverseNode(root *TreeNode) *[]int {
//	var res *[]int
//	if root != nil {
//		n.TraverseNode(root.Left)
//		//fmt.Print(root.Val, " ")
//		*res = append(*res, root.Val)
//		n.TraverseNode(root.Right)
//	}
//	return res
//}

func (n *TreeNode) InorderTraversal(root *TreeNode) []int {
	var res []int
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

func (n *TreeNode) PreorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, root.Val)
		res = append(res, preorderTraversal(root.Left)...)
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res
}
