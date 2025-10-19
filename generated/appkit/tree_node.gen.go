// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TreeNode] class.
var (
	treeNodeClass     _TreeNodeClass
	treeNodeClassOnce sync.Once
)

func getTreeNodeClass() _TreeNodeClass {
	treeNodeClassOnce.Do(func() {
		treeNodeClass = _TreeNodeClass{objc.GetClass("NSTreeNode")}
	})
	return treeNodeClass
}

type _TreeNodeClass struct {
	class objc.Class
}

// An interface definition for the [TreeNode] class.
type ITreeNode interface {
	objectivec.IObject
}

// A node in a tree of nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode
type TreeNode struct {
	objectivec.Object
}

// TreeNodeFrom constructs a [TreeNode] from an unsafe.Pointer.
//
// A node in a tree of nodes.
func TreeNodeFrom(ptr unsafe.Pointer) TreeNode {
	return TreeNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TreeNodeClass) Alloc() TreeNode {
	rv := objc.Send[TreeNode](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TreeNodeClass) New() TreeNode {
	rv := objc.Send[TreeNode](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TreeNode) Init() TreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TreeNode) Autorelease() TreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTreeNode creates a new TreeNode instance.
func NewTreeNode() TreeNode {
	return getTreeNodeClass().New()
}




