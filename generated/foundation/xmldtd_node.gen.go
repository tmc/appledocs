// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDTDNode] class.
var xMLDTDNodeClass = _XMLDTDNodeClass{objc.GetClass("NSXMLDTDNode")}

type _XMLDTDNodeClass struct {
	class objc.Class
}

// An interface definition for the [XMLDTDNode] class.
type IXMLDTDNode interface {
	IXMLNode
}

// A representation of element, attribute-list, entity, and notation declarations in a Document Type Definition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTDNode

type XMLDTDNode struct {
	XMLNode
}

// XMLDTDNodeFrom constructs a [XMLDTDNode] from an unsafe.Pointer.
//
// A representation of element, attribute-list, entity, and notation declarations in a Document Type Definition.
func XMLDTDNodeFrom(ptr unsafe.Pointer) XMLDTDNode {
	return XMLDTDNode{
		XMLNode: XMLNodeFrom(ptr),
	}
}



