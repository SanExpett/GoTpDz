package stack

import (
	"errors"
)

var errStack = errors.New("coudnt top or pop form empty stack")

type Stack struct {
	Nodes []string
}

func Create() *Stack {
	return &Stack{make([]string, 0)}
}

func (s *Stack) Len() int {
	return len(s.Nodes)
}

func (s *Stack) Top() (string, error) {
	if len(s.Nodes) == 0 {
		return "", errStack
	}

	return s.Nodes[len(s.Nodes)-1], nil
}

func (s *Stack) Push(node string) {
	s.Nodes = append(s.Nodes, node)
}

func (s *Stack) Pop() (string, error) {
	top, err := s.Top()
	if err != nil {
		return "", errStack
	}

	s.Nodes = s.Nodes[:s.Len()-1]

	return top, nil
}
