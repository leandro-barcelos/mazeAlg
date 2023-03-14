package t2

type Tree struct {
	parent *Tree
}

func (tree Tree) Root() (root Tree) {
	if tree.parent == nil {
		return tree
	}

	return tree.parent.Root()
}

func (t1 Tree) Is_connected(t2 Tree) (o bool) {
	return t1.Root() == t2.Root()
}

func (t1 *Tree) Connect(t2 Tree) {
	t2Root := t2.Root()
	t1Root := t1.Root()
	t2Root.parent = &t1Root
}
