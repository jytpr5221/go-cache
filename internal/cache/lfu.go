package cache

type LFUNode struct {
    key  string
    prev *LFUNode
    next *LFUNode
    freq int
}

type LFUEviction struct {
    head *LFUNode
    tail *LFUNode
    size int
}

func NewLFUEviction() *LFUEviction {
    return &LFUEviction{}
}

func (lfu *LFUEviction) OnSet(key string, meta *Metadata) (string, bool) {
    if meta != nil {
        if node, ok := meta.EvictionNode.(*LFUNode); ok && node != nil {
            node.key = key
            node.freq++
            lfu.reposition(node)
            return "", false
        }
    }

    node := &LFUNode{
        key:  key,
        freq: 1,
    }

    lfu.insertNode(node)

    if meta != nil {
        meta.EvictionNode = node
    }

    if lfu.size > MAX_CACHE_SIZE {
        evicted := lfu.removeHead()
        if evicted != nil {
            return evicted.key, true
        }
    }

    return "", false
}

func (lfu *LFUEviction) OnGet(key string, meta *Metadata) {
    if meta == nil {
        return
    }

    node, ok := meta.EvictionNode.(*LFUNode)
    if !ok || node == nil {
        return
    }

    node.key = key
    node.freq++
    lfu.reposition(node)
}

func (lfu *LFUEviction) OnDel(key string, meta *Metadata) {
    if meta == nil {
        return
    }

    node, ok := meta.EvictionNode.(*LFUNode)
    if !ok || node == nil {
        return
    }

    lfu.removeNode(node)
    meta.EvictionNode = nil
}

func (lfu *LFUEviction) insertNode(node *LFUNode) {
    if lfu.head == nil {
        lfu.head = node
        lfu.tail = node
        lfu.size = 1
        return
    }

    cur := lfu.head
    for cur != nil && cur.freq <= node.freq {
        cur = cur.next
    }

    if cur == nil {
        node.prev = lfu.tail
        lfu.tail.next = node
        lfu.tail = node
        lfu.size++
        return
    }

    if cur.prev == nil {
        node.next = cur
        cur.prev = node
        lfu.head = node
        lfu.size++
        return
    }

    prev := cur.prev
    prev.next = node
    node.prev = prev
    node.next = cur
    cur.prev = node
    lfu.size++
}

func (lfu *LFUEviction) reposition(node *LFUNode) {
    lfu.detach(node)
    lfu.insertNode(node)
}

func (lfu *LFUEviction) detach(node *LFUNode) {
    if node.prev != nil {
        node.prev.next = node.next
    } else {
        lfu.head = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    } else {
        lfu.tail = node.prev
    }

    node.prev = nil
    node.next = nil
}

func (lfu *LFUEviction) removeNode(node *LFUNode) {
    lfu.detach(node)
    lfu.size--
}

func (lfu *LFUEviction) removeHead() *LFUNode {
    if lfu.head == nil {
        return nil
    }

    node := lfu.head
    lfu.detach(node)
    lfu.size--
    return node
}