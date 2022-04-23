package main

import "fmt"

// 获取二叉树节点最大距离
// node 为根节点，若最大距离和node有关 那么就等于node.left.maxDistance +node.rightMaxDistance+1
//若无关久等于 max(left.maxDistance,right.maxDistance)
type Info struct {
	MaxDistance int
	height      int
}

func New(maxDistance, height int) *Info {
	return &Info{maxDistance, height}

}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func process(node *Node) *Info {
	if node == nil {
		return New(0, 0)
	}
	leftInfo := process(node.Left)
	rightInfo := process(node.Right)
	height := Max(leftInfo.height, rightInfo.height) + 1
	maxDistance := Max(Max(leftInfo.MaxDistance, rightInfo.MaxDistance), leftInfo.height+rightInfo.height+1)
	return &Info{height: height, MaxDistance: maxDistance}
}
func main() {
	node := GetBinaryTree()
	info := process(node)
	fmt.Println(info.MaxDistance, info.height)
}
