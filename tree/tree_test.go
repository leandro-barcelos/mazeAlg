package tree

import (
	"testing"
)

func TestTree(t *testing.T) {
	t1 := new(Tree)
	t2 := new(Tree)
	t3 := new(Tree)

	t2.Connect(t1)
	t3.Connect(t2)

	if !t3.Is_connected(t1) {
		t.Error("Tree struct FAILED")
	} else {
		t.Log("Tree struct PASSED")
	}
}
