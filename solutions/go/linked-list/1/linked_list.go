package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

type Node struct {
	Value any
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
}

var EmptyListError = errors.New("Empty List")

func NewList(elements ...any) *List {
	list := &List{
		head: nil,
		tail: nil,
	}

	if len(elements) < 1 {
		return list
	}

	list.head = NewNode(elements[0])

	currentNode := list.head

	for _, val := range elements[1:] {
		currentNode.next = NewNode(val)
		currentNode.next.prev = currentNode

		currentNode = currentNode.next
	}

	list.tail = currentNode

	return list
}

func NewNode(value any) *Node {
	return &Node{
		Value: value,
		prev:  nil,
		next:  nil,
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	node := NewNode(v)

	if l.head == nil || l.tail == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = l.head.prev

	return
}

func (l *List) Push(v any) {
	node := NewNode(v)

	if l.head == nil || l.tail == nil {
		l.head = node
		l.tail = node
		return
	}

	l.tail.next = node
	node.prev = l.tail
	l.tail = l.tail.next

	return
}

func (l *List) Shift() (any, error) {
	if l.head == nil && l.tail == nil {
		return nil, EmptyListError
	}

	res := l.head.Value

	if l.head.next != nil {
		l.head = l.head.next
		l.head.prev = nil
	} else {
		l.head = nil
		l.tail = nil
	}

	return res, nil
}

func (l *List) Pop() (any, error) {
	if l.head == nil && l.tail == nil {
		return nil, EmptyListError
	}

	res := l.tail.Value

	if l.tail.prev != nil {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		l.head = nil
		l.tail = nil
	}

	return res, nil
}

func (l *List) Reverse() {
	if l.head == nil || l.tail == nil {
		return
	}

	currentNode := l.tail

	for currentNode != nil {
		currentNode.prev, currentNode.next = currentNode.next, currentNode.prev
		currentNode = currentNode.next
	}

	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
