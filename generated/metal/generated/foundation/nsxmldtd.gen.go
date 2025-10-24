// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDTD] class.
var (
	XMLDTDClass     _XMLDTDClass
	XMLDTDClassOnce sync.Once
)

func getXMLDTDClass() _XMLDTDClass {
	XMLDTDClassOnce.Do(func() {
		XMLDTDClass = _XMLDTDClass{objc.GetClass("NSXMLDTD")}
	})
	return XMLDTDClass
}

type _XMLDTDClass struct {
	class objc.Class
}

// An interface definition for the [XMLDTD] class.
type IXMLDTD interface {
	IXMLNode
	// properties:
	PublicID() IString
	SetPublicID(value IString)
	SystemID() IString
	SetSystemID(value IString)
	Dtd() IXMLDTD
	SetDtd(value IXMLDTD)
	// methods:
}

// A representation of a Document Type Definition.
//
// An instance of the class is held as a property of an instance, accessed through the property . In the data model, an object is conceptually similar to namespace and attribute nodes: it is not considered to be a child of the object although it is closely associated with it. It is at the “root” of a shallow tree consisting primarily of nodes representing DTD declarations. Acceptable child nodes are instances of the class as well as objects representing comment nodes and processing-instruction nodes. You create an object in one of three ways: By processing an XML document with its own internal (in-line) DTD By process a standalone (external) DTD Programmatically Once an instance is in place, you can add, remove, and change the objects representing various DTD declarations. When you write the document out as XML, the new or modified internal DTD is included (assuming you set the DTD in the instance). You may also programmatically create an external DTD and write that out to its own file.


// A representation of a Document Type Definition.
//
// [Full Topic]
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



// Returns the receiver’s public identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtd/publicid
func (x_ XMLDTD) PublicID() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("publicID"))
	return rv
}


// Returns the receiver’s public identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtd/publicid
func (x_ XMLDTD) SetPublicID(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setPublicID:"), value)
}


// Returns the receiver’s system identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtd/systemid
func (x_ XMLDTD) SystemID() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("systemID"))
	return rv
}


// Returns the receiver’s system identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldtd/systemid
func (x_ XMLDTD) SetSystemID(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setSystemID:"), value)
}


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/dtd
func (x_ XMLDTD) Dtd() IXMLDTD {
	rv := objc.Send[XMLDTD](x_.ID, objc.Sel("dtd"))
	return rv
}


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/dtd
func (x_ XMLDTD) SetDtd(value IXMLDTD) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDtd:"), value)
}



