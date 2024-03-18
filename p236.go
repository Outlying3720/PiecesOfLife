/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import (
	"strings"
	"strconv"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTree(str string) (root *TreeNode) {
	s := strings.TrimLeft(str, "[")
	s = strings.TrimRight(s, "]")
	arr := strings.Split(s, ",")
	if len(arr) == 0 || arr[0] == "null" {
		return
	}
	root = new(TreeNode)
	root.Val, _ = strconv.Atoi(arr[0])
	arr = arr[1:]
	queue := []*TreeNode{root}
	for len(queue) > 0 && len(arr) > 0 {
		node := queue[0]
		queue = queue[1:]

		if arr[0] != "null" {
			node.Left = new(TreeNode)
			node.Left.Val, _ = strconv.Atoi(arr[0])
			queue = append(queue, node.Left)
		}
		arr = arr[1:]
		if len(arr) > 0 {
			if arr[0] != "null" {
				node.Right = new(TreeNode)
				node.Right.Val, _ = strconv.Atoi(arr[0])
				queue = append(queue, node.Right)
			}
			arr = arr[1:]
		}
	}
	return
}

var result *TreeNode

func dfs(root *TreeNode, p int, q int) bool {
	r, leftresult, rightresult := false, false, false
	if root.Val == p || root.Val == q {
		r = true
	}
	if root.Left != nil {
		leftresult = dfs(root.Left, p, q)
	}
	if root.Right != nil {
		rightresult = dfs(root.Right, p, q)
	}
	if (leftresult && rightresult) || (r && leftresult) || (r && rightresult) {
		if result == nil {
			result = root
		}
	}

	r = r || leftresult || rightresult
	return r
}
 
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
   dfs(root, p.Val, q.Val)
   return result
}

func main() {
	s := "[3,5,1,6,2,0,8,null,null,7,4]"

	tree := generateTree(s)

	dfs(tree, 5, 4)

	fmt.Println(result.Val)
}