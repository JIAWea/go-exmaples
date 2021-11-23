package main

import "fmt"

// Node 节点
type Node struct {
	Val   int
	Left  int
	Right int
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorderTraversalRecusive 前序遍历递归法
func preorderTraversalRecusive(root *TreeNode) {
	if root == nil || root.Val == 0 {
		// fmt.Println("该二叉树没有节点")
		return
	}
	fmt.Printf("%v ", root.Val)
	preorderTraversalRecusive(root.Left)
	preorderTraversalRecusive(root.Right)
}

// preorderTraversal 前序遍历非递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 先保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]

		// pop
		stack = stack[:len(stack)-1]
		root = node.Right
	}

	return result
}

// inorderTraversal 中序遍历非递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			// 一直向左添加
			stack = append(stack, root)
			root = root.Left
		}

		val := stack[len(stack)-1]
		result = append(result, val.Val)
		// pop
		stack = stack[:len(stack)-1]
		root = val.Right
	}

	return result
}

// postorderTraversal 非递归后续遍历
func postorderTraversal(root *TreeNode) []int {
	// 通过lastVisit标识右子节点是否已经弹出
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 先看看，不弹出
		node := stack[len(stack)-1]

		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// 标记当前节点已经弹出过了
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

// DFSPreorderTraversal 深度优先 从上到下
func DFSPreorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// BFSPreorderTraversal 广度优先
func BFSPreorderTraversal(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		valList := make([]int, 0)

		// 记录添加子节点之前的长度，先遍历当前层的所有元素，再添加下一层
		lenQueue := len(queue)
		for i := 0; i < lenQueue; i++ {
			level := queue[0]
			// 出队
			queue = queue[1:]
			valList = append(valList, level.Val)

			if level.Left != nil {
				queue = append(queue, level.Left)
			}

			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, valList)
	}

	return result
}

func main() {
	// 递归前序遍历
	t1 := TreeNode{10, &TreeNode{5, &TreeNode{2, nil, nil}, nil}, &TreeNode{20, nil, nil}}
	fmt.Printf("递归前序遍历：")

	preorderTraversalRecusive(&t1)

	fmt.Printf("\n*************************\n")

	// 非递归前序遍历
	t2 := TreeNode{10, &TreeNode{5, &TreeNode{2, nil, nil}, nil}, &TreeNode{20, nil, nil}}
	fmt.Println("非递归前序遍历：", preorderTraversal(&t2))

	// 非递归中序遍历
	fmt.Println("非递归中序遍历：", inorderTraversal(&t2))

	// 非递归后序遍历
	fmt.Println("非递归后序遍历：", postorderTraversal(&t2))

	// 深度优先遍历
	fmt.Println("深度优先遍历：", DFSPreorderTraversal(&t2))

	// 广度优先遍历
	fmt.Println("广度优先遍历：", BFSPreorderTraversal(&t2))
}
