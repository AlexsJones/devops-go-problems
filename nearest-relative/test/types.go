package types

import "github.com/AlexsJones/devops-go-problems/nearest-relative/problem"

type SimpleNode struct {
	*SimpleNode
}

func (s *SimpleNode) GetParent() problem.Node {
	return s.SimpleNode
}
