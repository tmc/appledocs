// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDTD] class.
var (
	xMLDTDClass     _XMLDTDClass
	xMLDTDClassOnce sync.Once
)

func getXMLDTDClass() _XMLDTDClass {
	xMLDTDClassOnce.Do(func() {
		xMLDTDClass = _XMLDTDClass{objc.GetClass("NSXMLDTD")}
	})
	return xMLDTDClass
}

type _XMLDTDClass struct {
	class objc.Class
}

// An interface definition for the [XMLDTD] class.
type IXMLDTD interface {
	IXMLNode
}

// A representation of a Document Type Definition.
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

// Alloc allocates a new instance without initialization.
func (xc _XMLDTDClass) Alloc() XMLDTD {
	rv := objc.Send[XMLDTD](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XMLDTDClass) New() XMLDTD {
	rv := objc.Send[XMLDTD](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XMLDTD) Init() XMLDTD {
	rv := objc.Send[XMLDTD](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XMLDTD) Autorelease() XMLDTD {
	rv := objc.Send[XMLDTD](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLDTD creates a new XMLDTD instance.
func NewXMLDTD() XMLDTD {
	return getXMLDTDClass().New()
}




