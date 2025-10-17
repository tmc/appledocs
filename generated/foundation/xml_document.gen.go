// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDocument] class.
var xMLDocumentClass = _XMLDocumentClass{objc.GetClass("NSXMLDocument")}

type _XMLDocumentClass struct {
	class objc.Class
}

// An interface definition for the [XMLDocument] class.
type IXMLDocument interface {
	IXMLNode
}

// An XML document as internalized into a logical tree structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument

type XMLDocument struct {
	XMLNode
}

// XMLDocumentFrom constructs a [XMLDocument] from an unsafe.Pointer.
//
// An XML document as internalized into a logical tree structure.
func XMLDocumentFrom(ptr unsafe.Pointer) XMLDocument {
	return XMLDocument{
		XMLNode: XMLNodeFrom(ptr),
	}
}



