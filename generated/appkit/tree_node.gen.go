
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TreeNode] class.
var TreeNodeClass _TreeNodeClass

func init() {
	TreeNodeClass = _TreeNodeClass{objc.GetClass("NSTreeNode")}
}

type _TreeNodeClass struct {
	objc.Class
}

// An interface definition for the [TreeNode] class.
type ITreeNode interface {
	ID() objc.ID
}

type TreeNode struct {
	id objc.ID
}

func TreeNodeFrom(ptr unsafe.Pointer) TreeNode {
	return TreeNode{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TreeNode) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TreeNodeClass) Alloc() TreeNode {
	rv := objc.Send[TreeNode](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TreeNodeClass) New() TreeNode {
	rv := objc.Send[TreeNode](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTreeNode creates and returns a new initialized instance.
func NewTreeNode() TreeNode {
	return TreeNodeClass.New()
}

// Init initializes the instance.
func (t_ TreeNode) Init() TreeNode {
	rv := objc.Send[TreeNode](t_.ID(), selInit)
	return rv
}
