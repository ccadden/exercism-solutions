package linkedlist

import "errors"

// Define the List and Element types here.
type Element struct {
	next  *Element
	value int
}

type List struct {
	head *Element
	size int
}

func NewElement(value int) *Element {
	return &Element{value: value}
}

func New(elements []int) *List {
	if len(elements) < 1 {
		return &List{head: nil, size: 0}
	}

	list := List{}

	for _, element := range elements {
		list.Push(element)
	}

	return &list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	node := NewElement(element)

	node.next = l.head
	l.head = node

	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size < 1 {
		return 0, errors.New("Cannot Pop() empty list")
	}

	result := l.head.value
	l.head = l.head.next

	l.size--

	return result, nil
}

func (l *List) Array() []int {
	if l.size < 1 {
		return []int{}
	}

	result := make([]int, l.size)

	for i, head := l.size-1, l.head; i >= 0; i, head = i-1, head.next {
		result[i] = head.value
	}

	return result
}

func (l *List) Reverse() *List {
	reversed := &List{}

	for head := l.head; head != nil; head = head.next {
		reversed.Push(head.value)

	}

	return reversed
}
