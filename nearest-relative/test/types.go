package types

import "github.com/SeedJobs/devops-go-tests/nearest-relative/problem"

type SimpleNode struct {
	*SimpleNode
}

func (s *SimpleNode) GetParent() problem.Node {
	return s.SimpleNode
}
