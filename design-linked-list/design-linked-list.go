type MyLinkedList struct {
    head *Node
    length int
}

type Node struct {
    val int
    next *Node
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
    if !(0 <= index && index < this.length) { // index out of range check
        return -1
    }
    node := this.head
    for i := 0; i < index; i++ {
        node = node.next
    }
    return node.val
}

func (this *MyLinkedList) AddAtHead(val int)  {
    if this.head == nil { // if head not present, create head w/ val
        this.head = &Node{val, nil}
    } else { // if head is present, head is new node with old list as this.next
        this.head = &Node{val, this.head}
    }
    this.length++
}

func (this *MyLinkedList) AddAtTail(val int)  {
    if this.head == nil { // if head not present, create head (same as tail) w/ val
        this.head = &Node{val, this.head}
    } else {
        node := this.head
        for node.next != nil { // iterate until node.next is nil. since head was non-nil, node is too.
            node = node.next
        }
        node.next = &Node{val, nil}
    }
    this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if !(0 <= index && index <= this.length) { // index out of range check
        return
    }
    // delegate to simpler functions for head and tail cases
    if index == 0 {
        this.AddAtHead(val)
        return
    } else if index == this.length {
        this.AddAtTail(val)
        return
    }
    // handle not-head-nor-tail case
    node := this.head
    for i := 0; i < index; i++ {
        node = node.next
    }
    if node != nil {
        tail := *node
        *node = Node{val, &tail}
    } else {
        panic("node should not be nil here if bounds checks are working!")
    }
    this.length++
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if !(0 <= index && index < this.length) { // index out of range check
        return
    }
    if index == 0 {
        this.DeleteAtHead()
        return
    }
    if index == this.length - 1 {
        this.DeleteAtTail()
        return
    }
    node := this.head
    for i := 0; i < index; i++ {
        node = node.next
    }
    if node == nil {
        panic("this shouldn't be nil because bounds check")
    }
    if node.next != nil { // node is not tail
        *node = *node.next
    } else { // node.next == nil so node is tail
        panic("should be handled by DeleteAtTail")
    }
    this.length--
}

func (this *MyLinkedList) DeleteAtHead() {
    if this.head == nil {
        panic("should have been caught by bounds-checking?")
    } else { // if head is present, head is new node with old tail.
        if this.length == 1 {
            this.head = nil
        } else {
            *this.head = *this.head.next
        }
    }
    this.length--
}

func (this *MyLinkedList) DeleteAtTail() {
    if this.head == nil { // if head not present, do nothing
        return
    } else {
        node := this.head
        for node.next.next != nil { // iterate until node.next.next is nil.
            node = node.next
        } // node has the pointer to kill to deref node.next (the tail)
        node.next = nil
    }
    this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */