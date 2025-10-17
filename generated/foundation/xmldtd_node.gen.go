// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XMLDTDNode] class.
var XMLDTDNodeClass objc.Class

func init() {
	XMLDTDNodeClass = objc.GetClass("NSXMLDTDNode")
}

type XMLDTDNode struct {
	objc.ID
}

func XMLDTDNodeFrom(ptr unsafe.Pointer) XMLDTDNode {
	return XMLDTDNode{
		ID: objc.ID(ptr),
	}
}



