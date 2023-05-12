type MyCircularQueue struct {
    arr []int
    left int
    right int
    length int
}

func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{make([]int, k), 0, 0, 0}
}

func (t *MyCircularQueue) EnQueue(value int) bool {
    if t.length == len(t.arr) {
        return false
    } else if t.IsEmpty() {
        t.arr[t.left] = value
        t.length++
    } else {
        t.right = (t.right + 1) % len(t.arr)
        t.arr[t.right] = value
        t.length++
    }
    return true
}

func (t *MyCircularQueue) DeQueue() bool {
    if t.length == 0 {
        return false
    }
    t.arr[t.left] = -1
    if t.left != t.right {
        t.left = (t.left + 1) % len(t.arr)
    }
    t.length--
    return true
}

func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.arr[this.left]
}

func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.arr[this.right]
}

func (this *MyCircularQueue) IsEmpty() bool {
    return this.length == 0
}

func (this *MyCircularQueue) IsFull() bool {
    return this.length == len(this.arr)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */