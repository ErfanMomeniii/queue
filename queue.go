package queue

// Queue is an instantiation of the queue
type Queue struct {
	front *Node
	back  *Node
}

// Node is an instantiation that contains value and pointer to the previous and next node of the queue
type Node struct {
	value    any
	previous *Node
	next     *Node
}

// PushFront adds a new element at the beginning of the queue.
func (q *Queue) PushFront(value any) {
	n := Node{value: value, next: q.front}
	if q.front != nil {
		q.front.previous = &n
		q.front = &n
	} else {
		q.front = &n
		q.back = &n
	}
}

// PushBack adds a new element at the end of the queue.
func (q *Queue) PushBack(value any) {
	n := Node{Value: value, previous: q.back}
	if q.back != nil {
		q.back.next = &n
		q.back = &n
	} else {
		q.back = &n
		q.front = &n
	}
}

// PopFront retrieves and removes the value of the first element in the queue.
func (q *Queue) PopFront() (value any) {
	if q.front == q.back {
		if q.front == nil {
			value = nil
		} else {
			value = q.front.value
			q.front = nil
			q.back = nil
		}
	} else {
		value = q.front.value
		q.front = q.front.next
		q.front.previous = nil
	}
	return
}

// PopBack retrieves and removes the value of the last element in the queue.
func (q *Queue) PopBack() (value any) {
	if q.back == q.front {
		if q.back == nil {
			value = nil
		} else {
			value = q.back.value
			q.back = nil
			q.front = nil
		}
	} else {
		value = q.back.value
		q.back = q.back.previous
		q.back.next = nil
	}
	return
}

// GetFront returns the value of front elements of the queue.
func (q *Queue) GetFront() (value any) {
	if q.front != nil {
		value = q.front.value
	}

	return
}

// GetBack returns the value of last elements of the queue
func (q *Queue) GetBack() (value any) {
	if q.back != nil {
		value = q.back.value
	}

	return
}

// New creates a new instance of a queue.
func New() *Queue {
	return &Queue{front: nil, back: nil}
}
