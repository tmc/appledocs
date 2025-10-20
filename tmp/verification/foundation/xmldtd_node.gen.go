// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var xMLDTDNodeClass _XMLDTDNodeClass

func init() {
	xMLDTDNodeClass = _XMLDTDNodeClass{objc.GetClass("NSXMLDTDNode")}
}

type _XMLDTDNodeClass struct {
	class objc.Class
}

type XMLDTDNode struct {
	objc.ID
}

func XMLDTDNodeFrom(ptr unsafe.Pointer) XMLDTDNode {
	return XMLDTDNode{
		ID: objc.ID(ptr),
	}
}




