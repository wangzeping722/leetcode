package _146

//
//type Node struct {
//	key, value int
//	pre, next  *Node
//}
//type LRUCache struct {
//	cap       int
//	cache     map[int]*Node
//	head, end *Node
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		cap:   capacity,
//		cache: make(map[int]*Node, capacity),
//		head:  nil,
//		end:   nil,
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	node, ok := this.cache[key]
//	if !ok {
//		return -1
//	}
//	this.remove(node)
//	this.setHeader(node)
//	return node.value
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if node, ok := this.cache[key]; ok {
//		node.value = value
//		this.remove(node)
//		this.setHeader(node)
//	} else {
//		if len(this.cache) >= this.cap {
//			delete(this.cache, this.end.key)
//			this.remove(this.end)
//		}
//		node := &Node{
//			key:   key,
//			value: value,
//		}
//		this.cache[key] = node
//		this.setHeader(node)
//	}
//}
//
//func (this *LRUCache) remove(node *Node) {
//	if node.pre == nil {
//		this.head = node.next
//	} else {
//		node.pre.next = node.next
//	}
//	if node.next == nil {
//		this.end = node.pre
//	} else {
//		node.next.pre = node.pre
//	}
//	node.next = nil
//	node.pre = nil
//}
//
//func (this *LRUCache) setHeader(node *Node) {
//	node.pre = nil
//	if this.head == nil {
//		this.head = node
//	} else {
//		this.head.pre, node.next = node, this.head
//		this.head = node
//	}
//
//	if this.end == nil {
//		this.end = node
//	}
//}

type Node struct {
	key, val int
	prev, next *Node
}

type LRUCache struct {
	head, tail *Node
	cache map[int]*Node
	cap int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		head:  nil,
		tail:  nil,
		cache: make(map[int]*Node, capacity),
		cap:   capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.setHeader(node)
	return node.val
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.cache[key]
	if !ok {
		node = &Node{key:key, val:value}
		if this.cap >= len(this.cache) {
			delete(this.cache, this.tail.key)
			this.remove(this.tail)
		}
		this.cache[key] = node
		this.setHeader(node)
	} else {
		this.remove(node)
		this.setHeader(node)
		node.val = value
	}
}

func (this *LRUCache) remove(node *Node) {
	// head
	if node.prev == nil {
		this.head = node.next
	} else {
		node.prev.next = node.next
	}

	if node.next == nil {
		this.tail = node.prev
	} else {
		node.next.prev = node.prev
	}

	node.prev = nil
	node.next = nil
}

func (this *LRUCache) setHeader(node *Node) {
	if this.head == nil {
		this.head = node
	} else {
		this.head.prev, node.next = node, this.head
		this.head = node
	}

	if this.tail == nil {
		this.tail = node
	}
}