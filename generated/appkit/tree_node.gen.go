// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TreeNode] class.
var treeNodeClass = _TreeNodeClass{objc.GetClass("NSTreeNode")}

type _TreeNodeClass struct {
	class objc.Class
}

// An interface definition for the [TreeNode] class.
type ITreeNode interface {
	objectivec.IObject
}

// A node in a tree of nodes. [Full Topic]
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



