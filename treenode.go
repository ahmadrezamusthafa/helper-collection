package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left,omitempty"`
	Right *TreeNode `json:"right,omitempty"`
}

func getIndex(values []int, value int) int {
	for i, v := range values {
		if value == v {
			return i
		}
	}
	return -1
}

func buildTree(inOrder []int, postOrder []int) *TreeNode {
	if len(postOrder) <= 0 {
		return nil
	}
	lastValPostOrder := postOrder[len(postOrder)-1]
	inOrderValIndex := getIndex(inOrder, lastValPostOrder)
	if inOrderValIndex < 0 {
		if len(postOrder) > 0 {
			return buildTree(inOrder, postOrder[:len(postOrder)-1])
		} else {
			return nil
		}
	}
	node := &TreeNode{
		Val: lastValPostOrder,
	}

	node.Left = buildTree(inOrder[:inOrderValIndex], postOrder[:len(postOrder)-1])
	node.Right = buildTree(inOrder[inOrderValIndex:], postOrder[:len(postOrder)-1])
	return node
}

func buildTreePreOrder(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) <= 0 {
		return nil
	}
	firstValPreOrder := preOrder[0]
	inOrderValIndex := getIndex(inOrder, firstValPreOrder)
	if inOrderValIndex < 0 {
		if len(preOrder) > 0 {
			return buildTreePreOrder(preOrder[1:], inOrder)
		} else {
			return nil
		}
	}
	node := &TreeNode{
		Val: firstValPreOrder,
	}

	node.Left = buildTreePreOrder(preOrder[1:], inOrder[:inOrderValIndex])
	node.Right = buildTreePreOrder(preOrder[1:], inOrder[inOrderValIndex:])
	return node
}

func testBuildTree() {
	in := []int{9, 3, 15, 20, 7}
	post := []int{9, 15, 7, 20, 3}
	treeNode := buildTree(in, post)

	by, _ := json.Marshal(treeNode)
	fmt.Println(string(by))
}

func testBuildTreePreOrder() {
	in := []int{9, 3, 15, 20, 7}
	pre := []int{3, 9, 20, 15, 7}
	treeNode := buildTreePreOrder(pre, in)

	by, _ := json.Marshal(treeNode)
	fmt.Println(string(by))
}
