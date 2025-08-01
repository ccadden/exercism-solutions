package linkedlist

import "errors"

// Define the List and Element types here.
type Element struct {
	next  *Element
	value int
}

type List struct {
	head *Element
	tail *Element
	size int
}

func NewElement(value int) *Element {
	return &Element{value: value}
}

func New(elements []int) *List {
	if len(elements) < 1 {
		return &List{head: nil, tail: nil, size: 0}
	}

	head := NewElement(elements[0])
	list := List{head: head, size: len(elements)}
	prevNode := head

	for _, element := range elements[1:] {
		e := NewElement(element)
		prevNode.next = e

		prevNode = prevNode.next
	}

	list.tail = prevNode

	return &list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	e := NewElement(element)

	if l.Size() < 1 {
		l.head = e
		l.tail = e
	} else {
		l.tail.next = e
		l.tail = l.tail.next
	}

	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size < 1 {
		return 0, errors.New("Cannot Pop() empty list")
	}

	var result = l.tail.value
	node := l.head

	if l.size == 1 {
		l.head = nil
	} else {
		for i := 0; i < l.size-2; i++ {
			node = node.next
		}

		node.next = nil
		l.tail = node
	}

	l.size--

	return result, nil
}

func (l *List) Array() []int {
	if l.size < 1 {
		return []int{}
	}

	result := []int{}
	node := l.head

	for node != nil {
		result = append(result, node.value)
		node = node.next
	}

	return result
}

func (l *List) Reverse() *List {
	arr := l.Array()
	for i := 0; i < len(arr)/2; i++ {
		var j = len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}

	return New(arr)
}
