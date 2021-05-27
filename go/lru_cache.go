package main

import "fmt"

type Node struct {
	Key   string
	Value interface{}
	pre   *Node
	next  *Node
}

type LRUCache struct {
	Head     *Node
	Tail     *Node
	HashData map[string]*Node
	Cap      int
}

func (c *LRUCache) New(cap int) *LRUCache {
	c.Head, c.Tail = &Node{}, &Node{}
	c.Head.next = c.Tail
	c.Tail.pre = c.Head
	c.Cap = cap
	c.HashData = make(map[string]*Node)
	return c
}

func (c *LRUCache) Set(key string, v interface{}) {
	n, exist := c.HashData[key]
	if exist {
		n.Value = v
		c.HashData[key] = n
		n.next.pre = n.pre
		n.pre.next = n.next
		c.appendToTail(n)
		return
	}
	fmt.Println("len map:", len(c.HashData))
	if len(c.HashData) >= c.Cap {
		tmp := c.Head.next
		c.Head.next = tmp.next
		tmp.next.pre = c.Head
		delete(c.HashData, tmp.Key)
	}
	newNode := Node{Key: key, Value: v}
	c.HashData[key] = &newNode
	c.appendToTail(&newNode)
}

func (c *LRUCache) appendToTail(n *Node) {
	n.next = c.Tail
	n.pre = c.Tail.pre
	c.Tail.pre.next = n
	c.Tail.pre = n
}

func (c *LRUCache) ShowCache() {
	node := c.Tail
	for node.pre != nil {
		fmt.Println(node.Key, ":", node.Value)
		node = node.pre
	}
}

func main() {
	fmt.Println("vim-go")
}
