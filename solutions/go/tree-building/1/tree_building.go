package tree

import (
	"cmp"
	"errors"
	"slices"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func newNode(id int) *Node {
	return &Node{ID: id, Children: []*Node{}}
}

func Build(records []Record) (*Node, error) {
	seenNodes := map[int]bool{}
	expectedId := 0

	if len(records) < 1 {
		return nil, nil
	}

	var nodes = make(map[int]*Node, len(records))

	slices.SortFunc(records, func(a, b Record) int {
		return cmp.Compare(a.ID, b.ID)
	})

	for i, record := range records {
		if record.ID < record.Parent {
			return nil, errors.New("Parent Node greater than Child Node")
		}

		if _, ok := seenNodes[record.ID]; ok {
			return nil, errors.New("Duplicate node")
		}

		if record.ID != expectedId {
			return nil, errors.New("Non-contiguous")
		}

		expectedId++

		nodes[i] = newNode(record.ID)
		seenNodes[record.ID] = true

		if i == 0 {
			if record.Parent != record.ID {
				return nil, errors.New("Invalid root node")
			}

			continue
		}

		if record.ID == record.Parent {
			return nil, errors.New("Cyclic definition")
		}

		parentNode := nodes[record.Parent]
		parentNode.Children = append(parentNode.Children, nodes[i])
	}

	return nodes[0], nil
}
