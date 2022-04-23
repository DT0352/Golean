package main

import (
	"fmt"
)

// 查找二叉树的后续节点
type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

//				 1
//  		2			  3
//		4	  5	   		6	  7
//	  8	  9	10	11	12	 13	14	15
// 8,4,9,2,10,5,11,1,12,6,13,3,14,7,15
//如果一个节点有右node，那么他的后续node就是他右节点的左节点的左节点。。。
//如果一个节点没有右节点，那么他的后续节点就是他的上级节点，且上级节点的左节点为他的上级节点
//有右节点 则是
func getBinaryTree() *Node {
	tree1 := Node{1, nil, nil, nil}
	tree2 := Node{2, nil, nil, nil}
	tree3 := Node{3, nil, nil, nil}
	tree4 := Node{4, nil, nil, nil}
	tree5 := Node{5, nil, nil, nil}
	tree6 := Node{6, nil, nil, nil}
	tree7 := Node{7, nil, nil, nil}
	tree8 := Node{8, nil, nil, nil}
	tree9 := Node{9, nil, nil, nil}
	tree10 := Node{10, nil, nil, nil}
	tree11 := Node{11, nil, nil, nil}
	tree12 := Node{12, nil, nil, nil}
	tree13 := Node{13, nil, nil, nil}
	tree14 := Node{14, nil, nil, nil}
	tree15 := Node{15, nil, nil, nil}

	tree1.Left = &tree2
	tree1.Right = &tree3
	tree1.Parent = nil
	tree2.Left = &tree4
	tree2.Right = &tree5
	tree2.Parent = &tree1
	tree3.Left = &tree6
	tree3.Right = &tree7
	tree3.Parent = &tree1
	tree4.Left = &tree8
	tree4.Right = &tree9
	tree4.Parent = &tree2
	tree5.Left = &tree10
	tree5.Right = &tree11
	tree5.Parent = &tree2
	tree6.Left = &tree12
	tree6.Right = &tree13
	tree6.Parent = &tree3
	tree7.Left = &tree14
	tree7.Right = &tree15
	tree7.Parent = &tree3
	tree8.Parent = &tree4
	tree9.Parent = &tree4
	tree10.Parent = &tree5
	tree11.Parent = &tree5
	tree12.Parent = &tree6
	tree13.Parent = &tree6
	tree14.Parent = &tree7
	tree15.Parent = &tree7
	return &tree8
}

// 获得后续节点 中序遍历的下一个节点
func getSuccessorNode(node *Node) *Node {
	if node == nil {
		return node
	}
	if node.Right != nil {
		return getMostLeft(node.Right)
	} else {
		parent := node.Parent
		for parent != nil && parent.Right == node {
			node = parent
			parent = node.Parent
		}
		if parent == nil {
			return &Node{-1, nil, nil, nil}

		}
		return parent
	}
}
func getMostLeft(node *Node) *Node {
	if node == nil {
		return node
	}
	for node.Left != nil {
		node = node.Left
	}
	return node

}

//获得前置节点，中序遍历的上一个节点
//节点没有节点，节点的上级节点的右节点是其本身否则就往上寻找
//节点有子节点且
func getPerNode(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.Right == nil {
		parent := node.Parent
		for node.Parent != nil && parent.Right != node {
			node = parent
			parent = node.Parent
		}
		return parent
	} else {
		left := node.Left
		for node.Right != nil && left.Right != nil {
			node = left
			left = node.Right
		}
		if left == nil {
			return &Node{-1, nil, nil, nil}

		}
		return left
	}
}

func main() {
	node := getBinaryTree()
	fmt.Printf("node[%v]的后续节点：%v\n", node.Value, getSuccessorNode(node).Value)
	fmt.Printf("node[%v]的前序节点：%v\n", node.Value, getPerNode(node).Value)
}
