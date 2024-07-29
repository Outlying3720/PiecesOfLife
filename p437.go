package main

import (
	"fmt"
	"strconv"
	"strings"
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

func dfs(root *TreeNode, presum map[int]int, nowsum int, target int) int {
	if root == nil {
		return 0
	}
	nowsum += root.Val

	result := presum[nowsum-target]
	// if nowsum == target {
	// 	result += 1
	// }

	presum[nowsum] += 1

	fmt.Println(presum)
	fmt.Println(root.Val, nowsum, nowsum-target, presum[nowsum-target])
	result += dfs(root.Left, presum, nowsum, target) + dfs(root.Right, presum, nowsum, target)
	presum[nowsum] -= 1

	return result
}

func pathSum(root *TreeNode, targetSum int) int {
	presum := map[int]int{0: 1}
	return dfs(root, presum, 0, targetSum)
}

func main() {
	tree := generateTree("1,-2,-3")
	fmt.Println(pathSum(tree, -1))
}
