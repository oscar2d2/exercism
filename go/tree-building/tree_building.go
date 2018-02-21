package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type ByID []Record

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}

	nodes := make([]*Node, len(records))
	sort.Sort(ByID(records))

	for i, rec := range records {
		nodes[i] = &Node{ID: rec.ID}

		if i == 0 && (rec.ID != 0 || rec.Parent != 0) {
			return nil, errors.New("invalid root")
		}
		if i == 0 {
			continue
		}
		if i != rec.ID || rec.ID <= rec.Parent {
			return nil, errors.New("Invalid record")
		}

		if parent := nodes[rec.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[i])
		}
	}

	return nodes[0], nil
}
