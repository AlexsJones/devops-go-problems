package problem

// Node a simple interface that is only required
// to return its parent's node.
type Node interface {
	// GetParent returns the parent to the node
	// If the Node doesn't have a parent, it returns itself.
	GetParent() Node
}
