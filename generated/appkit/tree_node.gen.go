// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TreeNode] class.
var TreeNodeClass objc.Class

func init() {
	TreeNodeClass = objc.GetClass("NSTreeNode")
}

type TreeNode struct {
	objc.ID
}

func TreeNodeFrom(ptr unsafe.Pointer) TreeNode {
	return TreeNode{
		ID: objc.ID(ptr),
	}
}



