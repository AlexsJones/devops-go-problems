package types_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/AlexsJones/devops-go-problems/nearest-relative/problem"
	"github.com/AlexsJones/devops-go-problems/nearest-relative/test"
)

func BenchmarkGetNearestRelative(b *testing.B) {
	tree := []*types.SimpleNode{}
	// Creating the entire tree
	for i := 0; i < b.N; i++ {
		tree = append(tree, &types.SimpleNode{})
		tree[i].SimpleNode = tree[0]
	}
	for i, node := range tree {
		tree[(i<<2)%len(tree)].SimpleNode = node
	}
	tree[0].SimpleNode = tree[0]
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b.ResetTimer()
	problem.FindNearestRelative(tree[r.Int()%len(tree)], tree[r.Int()%len(tree)])
}
