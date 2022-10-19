package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildTreeNodeFromArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	q := &Queue{}
	root := &TreeNode{arr[0], nil, nil}
	q.enqueue(root)
	pos := 1
	for pos < len(arr) {
		cur := q.pop().(*TreeNode)
		cur.Left = &TreeNode{arr[pos], nil, nil}
		q.enqueue(cur.Left)
		pos += 1
		if pos < len(arr) {
			cur.Right = &TreeNode{arr[pos], nil, nil}
			q.enqueue(cur.Right)
			pos += 1
		}
	}
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
