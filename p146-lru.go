package main

import "fmt"

type LinkList struct {
	Key  int
	Val  int
	Next *LinkList
	Prev *LinkList
}

type LRUCache struct {
	hash     map[int]*LinkList
	headprev *LinkList
	endnext  *LinkList
	Capacity int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.hash = map[int]*LinkList{}
	lru.headprev = &LinkList{}
	lru.endnext = &LinkList{Prev: lru.headprev}
	lru.headprev.Next = lru.endnext
	lru.Capacity = capacity

	return lru
}

func (this *LRUCache) get(key int) *LinkList {
	if node, ok := this.hash[key]; ok {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		node.Prev = this.headprev
		this.headprev.Next.Prev = node
		this.headprev.Next, node.Next = node, this.headprev.Next

		// fmt.Print("after get: ")
		// for c := this.headprev.Next; c != this.endnext; c = c.Next {
		// 	fmt.Print(c.Key, " ")
		// }
		// fmt.Println("")
		// fmt.Print("rev   get: ")
		// for c := this.endnext.Prev; c != this.headprev; c = c.Prev {
		// 	fmt.Print(c.Key, " ")
		// }
		// fmt.Println("")

		return node
	}
	return nil
}

func (this *LRUCache) Get(key int) int {
	if node := this.get(key); node != nil {
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node := this.get(key); node != nil {
		node.Val = value
		return
	}
	newnode := &LinkList{Key: key, Val: value}
	this.hash[key] = newnode

	this.headprev.Next.Prev = newnode
	this.headprev.Next, newnode.Next = newnode, this.headprev.Next
	newnode.Prev = this.headprev

	if len(this.hash) > this.Capacity {
		e := this.endnext.Prev
		delete(this.hash, e.Key)
		this.endnext.Prev.Prev.Next = this.endnext
		this.endnext.Prev = this.endnext.Prev.Prev
		e.Next = nil
		e.Prev = nil
	}

	// fmt.Print("after put: ")
	// for c := this.headprev.Next; c != this.endnext; c = c.Next {
	// 	fmt.Print(c.Key, " ")
	// }
	// fmt.Println("")

	// fmt.Print("rev   get: ")
	// for c := this.endnext.Prev; c != this.headprev; c = c.Prev {
	// 	fmt.Print(c.Key, " ")
	// }
	// fmt.Println("")
}

func main() {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
}
