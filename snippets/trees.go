package snippets

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) Insert(val int) {
	if root == nil {
		return
	}

	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {

			root.Left.Insert(val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			root.Right.Insert(val)
		}
	}

}
func (root *TreeNode) Inorder() {
	if root.Left != nil {
		root.Left.Inorder()
	}
	fmt.Println(root.Val)
	if root.Right != nil {
		root.Right.Inorder()
	}
}

func (root *TreeNode) IterativeInorder() {
	stack := []*TreeNode{}
	node := root

	for node != nil || len(stack) > 0 {
		// process all left nodes first
		for node != nil {
			stack = append(stack, node)
			node = node.Left

		}
		//pop the topmost which is your leftmost leaf node
		node = stack[len(stack)-1]
		//resize the stack
		stack = stack[:len(stack)-1]

		fmt.Println(node.Val)
		//process right node
		node = node.Right
	}
}

func (root *TreeNode) IterativePreorder() {
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		fmt.Println(node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
}

func (root *TreeNode) Search(val int) bool {
	if root == nil {
		return false
	}
	if val == root.Val {
		return true
	} else if val < root.Val {
		return root.Left.Search(val)
	} else {
		return root.Right.Search(val)
	}
}

func BST() {
	elements := []int{4, 6, 1, 0, 7, 9}
	root := &TreeNode{Val: elements[0], Left: nil, Right: nil}

	for i := 1; i < len(elements); i++ {
		root.Insert(elements[i])
	}
	//fmt.Println(root.Search(9))
	root.IterativeInorder()

}
