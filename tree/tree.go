package tree

type Tree struct {
	parent *Tree
}

func (tree *Tree) Root() *Tree {
	if tree.parent == nil {
		return tree
	}

	return tree.parent.Root()
}

func (t1 *Tree) Is_connected(t2 *Tree) bool {
	return t1.Root() == t2.Root()
}

func (t1 *Tree) Connect(t2 *Tree) {
	t2.Root().parent = t1
}
