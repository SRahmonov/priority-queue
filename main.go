package main

type queue struct {
	firstElement *queueNode
	lastElement  *queueNode
	size         int
}
type queueNode struct {
	next  *queueNode
	prev  *queueNode
	value interface{}
}
func (receiver *queue) len() int {
	return receiver.size
}
func (receiver *queue) first() interface{} {
	return receiver.firstElement
}
func (receiver *queue) last() interface{} {
	return receiver.lastElement
}
func (receiver *queue) equeue(value interface{}) {
	if receiver.len() == 0 {
		receiver.firstElement = &queueNode{
			next:  nil,
			prev:  nil,
			value: value,
		}
		receiver.lastElement = receiver.firstElement
		receiver.size++
		return
	}
	receiver.size++
	current := receiver.firstElement
	if current != nil {
		for {
			if current.next == nil {
				current.next = &queueNode{
					next:  nil,
					prev:  current,
					value: value,
				}
			}
			receiver.lastElement = current.next
			break
		}
		current = current.next
	}
}
func (receiver *queue) dequeue() interface{}  {
	if receiver.len() == 0 {
		return queue{}
	}
	returnFirst := queue{
		firstElement: receiver.firstElement,
		lastElement:  nil,
		size:         0,
	}
	receiver.firstElement = receiver.firstElement.next
	receiver.size--
	if receiver.size == 0 {
		receiver.lastElement = receiver.firstElement
	}
	return returnFirst.firstElement.value
}

func main() {}