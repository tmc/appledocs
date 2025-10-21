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
	TreeNodeClass     _TreeNodeClass
	TreeNodeClassOnce sync.Once
)

func getTreeNodeClass() _TreeNodeClass {
	TreeNodeClassOnce.Do(func() {
		TreeNodeClass = _TreeNodeClass{objc.GetClass("NSTreeNode")}
	})
	return TreeNodeClass
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
// simplifies the creation and management of trees of objects. Each tree node represents a model object. A tree node with as its parent node is considered the root of the tree.
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


// Creates and returns a tree node that represents the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/treeNodeWithRepresentedObject:
func (tc _TreeNodeClass) TreeNodeWithRepresentedObject(modelObject objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("treeNodeWithRepresentedObject:"), modelObject)
	return rv
}

// An array containing receiver’s child nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/children
func (t_ TreeNode) ChildNodes() []TreeNode {
	rv := objc.Send[[]TreeNode](t_.ID, objc.Sel("childNodes"))
	return rv
}



