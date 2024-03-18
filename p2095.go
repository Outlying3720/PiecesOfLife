package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
    var left *ListNode
    var right *ListNode
	var pre *ListNode
    left = head
    right = head

    for right != nil && right.Next != nil {
		right = right.Next.Next
		pre = left
		left = left.Next
    }
	if pre != nil {
		pre.Next = pre.Next.Next
	} else {
		head = head.Next
	}
    return head
}

func main() {
	nums := []int{1}
	last := &ListNode{Val: nums[len(nums)-1], Next: nil}
	for i:=len(nums)-2; i>=0; i-- {
		v := nums[i]
		// fmt.Print(v, " ")
		tmp := ListNode{Val: v, Next: last}
		last = &tmp
	}
	head := last
	deleteMiddle(head)
	for head != nil {
        fmt.Print(head.Val, " ")
		head = head.Next
    }
}