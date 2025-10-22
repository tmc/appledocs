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
	AttributeDeclarationForNameElementName(name string, elementName string) XMLDTDNode
	ElementDeclarationForName(name string) XMLDTDNode
	EntityDeclarationForName(name string) XMLDTDNode
	InsertChildAtIndex(child IXMLNode, index uint)
	InsertChildrenAtIndex(children []XMLNode, index uint)
	NotationDeclarationForName(name string) XMLDTDNode
	ReplaceChildAtIndexWithNode(index uint, node IXMLNode)
	PublicID() string
	SetPublicID(value string)
	SystemID() string
	SetSystemID(value string)
	Dtd() NSXMLDTD
	SetDtd(value IXMLDTD)
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



// Initializes and returns an object created from the DTD declarations encapsulated in an object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/init(data:options:)
func NewXMLDTDWithDataOptionsError(data NSData, mask NSXMLNodeOptions, error_ unsafe.Pointer) XMLDTD {
	instance := getXMLDTDClass().Alloc()
	rv := objc.Send[XMLDTD](instance.ID, objc.Sel("initWithData:options:error:"), data, mask, error_)
	rv.Autorelease()
	return rv
}



// Returns a DTD node representing the predefined entity declaration with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/predefinedEntityDeclaration(forName:)
func (xc _XMLDTDClass) PredefinedEntityDeclarationForName(name string) XMLDTDNode {
	rv := objc.Send[XMLDTDNode](objc.ID(xc.class), objc.Sel("predefinedEntityDeclarationForName:"), objc.String(name))
	return rv
}


// Returns the DTD node representing an attribute-list declaration for a given attribute and its element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/attributeDeclaration(forName:elementName:)
func (x_ XMLDTD) AttributeDeclarationForNameElementName(name string, elementName string) XMLDTDNode {
	rv := objc.Send[XMLDTDNode](x_.ID, objc.Sel("attributeDeclarationForName:elementName:"), objc.String(name), objc.String(elementName))
	return rv
}


// Returns the DTD node representing an element declaration for a specified element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/elementDeclaration(forName:)
func (x_ XMLDTD) ElementDeclarationForName(name string) XMLDTDNode {
	rv := objc.Send[XMLDTDNode](x_.ID, objc.Sel("elementDeclarationForName:"), objc.String(name))
	return rv
}


// Returns the DTD node representing the entity declaration for a specified entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/entityDeclaration(forName:)
func (x_ XMLDTD) EntityDeclarationForName(name string) XMLDTDNode {
	rv := objc.Send[XMLDTDNode](x_.ID, objc.Sel("entityDeclarationForName:"), objc.String(name))
	return rv
}


// Inserts a child node in the receiver’s list of children at a specific location in the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/insertChild(_:at:)
func (x_ XMLDTD) InsertChildAtIndex(child IXMLNode, index uint) {
	objc.Send[objc.ID](x_.ID, objc.Sel("insertChild:atIndex:"), child, index)
}


// Inserts an array of child nodes at a specified location in the receiver’s list of children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/insertChildren(_:at:)
func (x_ XMLDTD) InsertChildrenAtIndex(children []XMLNode, index uint) {
	objc.Send[objc.ID](x_.ID, objc.Sel("insertChildren:atIndex:"), children, index)
}


// Returns the DTD node representing the notation declaration identified by the specified notation name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/notationDeclaration(forName:)
func (x_ XMLDTD) NotationDeclarationForName(name string) XMLDTDNode {
	rv := objc.Send[XMLDTDNode](x_.ID, objc.Sel("notationDeclarationForName:"), objc.String(name))
	return rv
}


// Replaces a child at a particular index with another child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/replaceChild(at:with:)
func (x_ XMLDTD) ReplaceChildAtIndexWithNode(index uint, node IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("replaceChildAtIndex:withNode:"), index, node)
}


// Returns the receiver’s public identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/publicID
func (x_ XMLDTD) PublicID() string {
	rv := objc.Send[string](x_.ID, objc.Sel("publicID"))
	return rv
}


// Returns the receiver’s public identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/publicID
func (x_ XMLDTD) SetPublicID(value string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setPublicID:"), objc.String(value))
}


// Returns the receiver’s system identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/systemID
func (x_ XMLDTD) SystemID() string {
	rv := objc.Send[string](x_.ID, objc.Sel("systemID"))
	return rv
}


// Returns the receiver’s system identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDTD/systemID
func (x_ XMLDTD) SetSystemID(value string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setSystemID:"), objc.String(value))
}


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/dtd
func (x_ XMLDTD) Dtd() NSXMLDTD {
	rv := objc.Send[NSXMLDTD](x_.ID, objc.Sel("dtd"))
	return rv
}


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/dtd
func (x_ XMLDTD) SetDtd(value IXMLDTD) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDtd:"), value)
}


