package louds

import (
	"testing"
)

type OrderedTreeNode struct {
	val         string
	parent      *OrderedTreeNode
	firstChild  *OrderedTreeNode
	nextBrother *OrderedTreeNode
}

func (node *OrderedTreeNode) Val() interface{} {
	return node.val
}

func (node *OrderedTreeNode) Parent() TreeNode {
	return node.parent
}

func (node *OrderedTreeNode) FirstChild() TreeNode {
	return node.firstChild
}

func (node *OrderedTreeNode) NextBrother() TreeNode {
	return node.nextBrother
}

func genTestTree() *OrderedTreeNode {
	node0 := &OrderedTreeNode{val: "a"}
	node1 := &OrderedTreeNode{val: "b"}
	node2 := &OrderedTreeNode{val: "c"}
	node3 := &OrderedTreeNode{val: "d"}
	node4 := &OrderedTreeNode{val: "e"}
	node5 := &OrderedTreeNode{val: "f"}
	node6 := &OrderedTreeNode{val: "g"}
	node7 := &OrderedTreeNode{val: "h"}

	node0.firstChild = node1

	node1.parent = node0
	node1.firstChild = node4
	node1.nextBrother = node2

	node2.parent = node0
	node2.nextBrother = node3

	node3.parent = node0
	node3.firstChild = node5

	node4.parent = node1

	node5.parent = node3
	node5.nextBrother = node6

	node6.parent = node3
	node6.nextBrother = node7

	node7.parent = node3

	return node0
}

func TestParent(t *testing.T) {
	louds := BuildLOUDS(genTestTree())
	expected := []int{-1, 0, 0, 0, 1, 3, 3, 3}
	for i, v := range expected {
		actual := louds.Parent(i)
		if actual != v {
			t.Errorf("Parent(%v) => '%v', want '%v'", i, actual, v)
		}
	}
}

func TestFirstChild(t *testing.T) {
	louds := BuildLOUDS(genTestTree())
	expected := []int{1, 4, -1, 5, -1, -1, -1, -1}
	for i, v := range expected {
		actual := louds.FirstChild(i)
		if actual != v {
			t.Errorf("FirstChild(%v) => '%v', want '%v'", i, actual, v)
		}
	}
}

func TestNextBrother(t *testing.T) {
	louds := BuildLOUDS(genTestTree())
	expected := []int{-1, 2, 3, -1, -1, 6, 7, -1}
	for i, v := range expected {
		actual := louds.NextBrother(i)
		if actual != v {
			t.Errorf("NextBrother(%v) => '%v', want '%v'", i, actual, v)
		}
	}
}
