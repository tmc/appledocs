// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	ChildNodes() []TreeNode
	Children() NSTreeNode
	SetChildren(value ITreeNode)
	IndexPath() foundation.IndexPath
	SetIndexPath(value foundation.IIndexPath)
	IsLeaf() bool
	SetIsLeaf(value bool)
	MutableChildren() foundation.MutableArray
	SetMutableChildren(value foundation.IMutableArray)
	Parent() NSTreeNode
	SetParent(value ITreeNode)
	RepresentedObject() unsafe.Pointer
	SetRepresentedObject(value unsafe.Pointer)
}

// A node in a tree of nodes.
//
// simplifies the creation and management of trees of objects. Each tree node represents a model object. A tree node with as its parent node is considered the root of the tree.


// A node in a tree of nodes.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/treeNodeWithRepresentedObject:
func (tc _TreeNodeClass) TreeNodeWithRepresentedObject(modelObject objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("treeNodeWithRepresentedObject:"), modelObject)
	return rv
}


// An array containing receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/children
func (t_ TreeNode) ChildNodes() []TreeNode {
	rv := objc.Send[[]TreeNode](t_.ID, objc.Sel("childNodes"))
	return rv
}


// An array containing receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/children
func (t_ TreeNode) Children() NSTreeNode {
	rv := objc.Send[NSTreeNode](t_.ID, objc.Sel("children"))
	return rv
}


// An array containing receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/children
func (t_ TreeNode) SetChildren(value ITreeNode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChildren:"), value)
}


// The position of the receiver relative to its root parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/indexpath
func (t_ TreeNode) IndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](t_.ID, objc.Sel("indexPath"))
	return rv
}


// The position of the receiver relative to its root parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/indexpath
func (t_ TreeNode) SetIndexPath(value foundation.IIndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIndexPath:"), value)
}


// A Boolean that indicates whether the receiver is a leaf node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/isleaf
func (t_ TreeNode) IsLeaf() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isLeaf"))
	return rv
}


// A Boolean that indicates whether the receiver is a leaf node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/isleaf
func (t_ TreeNode) SetIsLeaf(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsLeaf:"), value)
}


// A mutable array that provides read-write access to the receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/mutablechildren
func (t_ TreeNode) MutableChildren() foundation.MutableArray {
	rv := objc.Send[foundation.MutableArray](t_.ID, objc.Sel("mutableChildren"))
	return rv
}


// A mutable array that provides read-write access to the receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/mutablechildren
func (t_ TreeNode) SetMutableChildren(value foundation.IMutableArray) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMutableChildren:"), value)
}


// The receiver’s parent node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/parent
func (t_ TreeNode) Parent() NSTreeNode {
	rv := objc.Send[NSTreeNode](t_.ID, objc.Sel("parent"))
	return rv
}


// The receiver’s parent node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/parent
func (t_ TreeNode) SetParent(value ITreeNode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParent:"), value)
}


// The object the tree node represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/representedobject
func (t_ TreeNode) RepresentedObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("representedObject"))
	return rv
}


// The object the tree node represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/representedobject
func (t_ TreeNode) SetRepresentedObject(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRepresentedObject:"), value)
}



