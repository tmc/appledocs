// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDTD] class.
var xMLDTDClass = _XMLDTDClass{objc.GetClass("NSXMLDTD")}

type _XMLDTDClass struct {
	class objc.Class
}

// An interface definition for the [XMLDTD] class.
type IXMLDTD interface {
	IXMLNode
}

// A representation of a Document Type Definition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD

type XMLDTD struct {
	XMLNode
}

// XMLDTDFrom constructs a [XMLDTD] from an unsafe.Pointer.
//
// A representation of a Document Type Definition.
func XMLDTDFrom(ptr unsafe.Pointer) XMLDTD {
	return XMLDTD{
		XMLNode: XMLNodeFrom(ptr),
	}
}



