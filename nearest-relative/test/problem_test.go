package types_test

import (
	"testing"

	"github.com/SeedJobs/devops-go-problems/nearest-relative/problem"
	"github.com/SeedJobs/devops-go-problems/nearest-relative/test"
)

func TestTreeOfOne(t *testing.T) {
	n := &types.SimpleNode{}
	n.SimpleNode = n
	if n != problem.FindNearestRelative(n, n) {
		t.Fatal("Does not pass simple test case")
	}
}

func TestAbstractTreeOfFive(t *testing.T) {
	tree := []*types.SimpleNode{}
	for i := 0; i < 5; i++ {
		tree = append(tree, &types.SimpleNode{})
	}
	tree[0].SimpleNode = tree[0]
	tree[1].SimpleNode, tree[2].SimpleNode = tree[0], tree[0]
	tree[3].SimpleNode, tree[4].SimpleNode = tree[2], tree[2]
	if tree[0] != problem.FindNearestRelative(tree[1], tree[4]) {
		t.Fatal("Does not return the root node")
	}
	if tree[2] != problem.FindNearestRelative(tree[3], tree[4]) {
		t.Fatal("Does not return correct value")
	}
	if tree[3] != problem.FindNearestRelative(tree[3], tree[3]) {
		t.Fatal("Does not return correct value")
	}
	tree[4].SimpleNode = tree[3]
	if tree[3] != problem.FindNearestRelative(tree[3], tree[4]) {
		t.Fatal("Does not return correct value")
	}
}
