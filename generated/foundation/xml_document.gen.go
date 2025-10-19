// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDocument] class.
var (
	xMLDocumentClass     _XMLDocumentClass
	xMLDocumentClassOnce sync.Once
)

func getXMLDocumentClass() _XMLDocumentClass {
	xMLDocumentClassOnce.Do(func() {
		xMLDocumentClass = _XMLDocumentClass{objc.GetClass("NSXMLDocument")}
	})
	return xMLDocumentClass
}

type _XMLDocumentClass struct {
	class objc.Class
}

// An interface definition for the [XMLDocument] class.
type IXMLDocument interface {
	IXMLNode
}

// An XML document as internalized into a logical tree structure.
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

// Alloc allocates a new instance without initialization.
func (xc _XMLDocumentClass) Alloc() XMLDocument {
	rv := objc.Send[XMLDocument](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XMLDocumentClass) New() XMLDocument {
	rv := objc.Send[XMLDocument](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XMLDocument) Init() XMLDocument {
	rv := objc.Send[XMLDocument](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XMLDocument) Autorelease() XMLDocument {
	rv := objc.Send[XMLDocument](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLDocument creates a new XMLDocument instance.
func NewXMLDocument() XMLDocument {
	return getXMLDocumentClass().New()
}




