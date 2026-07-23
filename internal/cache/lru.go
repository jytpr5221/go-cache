package cache

type LRUNode struct {
	key  string
	prev *LRUNode
	next *LRUNode
}

type LRUEviction struct {
	head *LRUNode
	tail *LRUNode
	size int
}

func NewLRUEviction() *LRUEviction {
	return &LRUEviction{}
}

func (lru *LRUEviction) OnSet(key string, meta *Metadata) (string, bool) {
	if meta != nil {
		if node, ok := meta.EvictionNode.(*LRUNode); ok && node != nil {
			node.key = key
			lru.moveToFront(node)
			return "", false
		}
	}  

	node := &LRUNode{key: key}
	lru.insertFront(node)
	lru.size++

	if meta != nil {
		meta.EvictionNode = node
	}

	if lru.size > MAX_CACHE_SIZE {
		evicted := lru.removeTail()
		if evicted != nil {
			return evicted.key, true
		}
	}

	return "", false
}

func (lru *LRUEviction) OnGet(key string, meta *Metadata) {
	if meta == nil {
		return
	}

	node, ok := meta.EvictionNode.(*LRUNode)
	if !ok || node == nil {
		return
	}

	node.key = key
	lru.moveToFront(node)
}

func (lru *LRUEviction) OnDel(key string, meta *Metadata) {
	if meta == nil {
		return
	}

	node, ok := meta.EvictionNode.(*LRUNode)
	if !ok || node == nil {
		return
	}

	lru.removeNode(node)
	meta.EvictionNode = nil
}

func (lru *LRUEviction) insertFront(node *LRUNode) {
	node.prev = nil
	node.next = lru.head

	if lru.head != nil {
		lru.head.prev = node
	} else {
		lru.tail = node
	}

	lru.head = node
}

func (lru *LRUEviction) moveToFront(node *LRUNode) {
	if node == nil || node == lru.head {
		return
	}

	lru.detach(node)
	lru.insertFront(node)
}

func (lru *LRUEviction) detach(node *LRUNode) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lru.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lru.tail = node.prev
	}

	node.prev = nil
	node.next = nil
}

func (lru *LRUEviction) removeNode(node *LRUNode) {
	lru.detach(node)
	lru.size--
}

func (lru *LRUEviction) removeTail() *LRUNode {
	if lru.tail == nil {
		return nil
	}

	node := lru.tail
	lru.detach(node)
	lru.size--
	return node
}
