// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLElement] class.
var xMLElementClass = _XMLElementClass{objc.GetClass("NSXMLElement")}

type _XMLElementClass struct {
	class objc.Class
}

// The element nodes in an XML tree structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement

type XMLElement struct {
	XMLNode
}

// XMLElementFrom constructs a [XMLElement] from an unsafe.Pointer.
//
// The element nodes in an XML tree structure.
func XMLElementFrom(ptr unsafe.Pointer) XMLElement {
	return XMLElement{
		XMLNode: XMLNodeFrom(ptr),
	}
}



