package ringbuffer

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type MyCircularQueue struct {
	Tasks []int
	Full  bool
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	m := new(MyCircularQueue)
	for i := 0; i < k; i++ {
		m.Tasks = append(m.Tasks, i)
	}
	return *m
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (t *MyCircularQueue) EnQueue(value int) bool {
	if !t.Full {
		t.Tasks = append(t.Tasks, value)
		fmt.Println(t.Tasks)
		return true
	}
	return false
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (t *MyCircularQueue) DeQueue() bool {
	size := len(t.Tasks)
	if size > 0 {
		t.Tasks = append(t.Tasks[:0], t.Tasks[1:]...)
		fmt.Println(t.Tasks)
		if len(t.Tasks) == (size - 1) {
			return true
		}
		return false
	}
	return false
}

/** Get the front item from the queue. */
func (t *MyCircularQueue) Front() int {
	if t.IsEmpty() {
		return -1
	}

	return 0
}

/** Get the last item from the queue. */
func (t *MyCircularQueue) Rear() int {
	return 0
}

/** Checks whether the circular queue is empty or not. */
func (t *MyCircularQueue) IsEmpty() bool {
	return false
}

/** Checks whether the circular queue is full or not. */
func (t *MyCircularQueue) IsFull() bool {
	return false
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
