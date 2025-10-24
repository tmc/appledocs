// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [XMLElement] class.
var (
	XMLElementClass     _XMLElementClass
	XMLElementClassOnce sync.Once
)

func getXMLElementClass() _XMLElementClass {
	XMLElementClassOnce.Do(func() {
		XMLElementClass = _XMLElementClass{objc.GetClass("NSXMLElement")}
	})
	return XMLElementClass
}

type _XMLElementClass struct {
	class objc.Class
}





// An interface definition for the [XMLElement] class.
type IXMLElement interface {
	IXMLNode
	

	// properties:
	Attributes() IXMLNode
	SetAttributes(value IXMLNode)
	Namespaces() IXMLNode
	SetNamespaces(value IXMLNode)


	

	// methods:
	AddAttribute(attribute IXMLNode)
	RemoveNamespaceForPrefix(name IString)
	ResolvePrefixForNamespaceURI(namespaceURI IString) IString


}





// Alloc allocates a new instance without initialization.
func (xc _XMLElementClass) Alloc() XMLElement {
	rv := objc.Send[XMLElement](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (xc _XMLElementClass) New() XMLElement {
	rv := objc.Send[XMLElement](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XMLElement) Init() XMLElement {
	rv := objc.Send[XMLElement](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XMLElement) Autorelease() XMLElement {
	rv := objc.Send[XMLElement](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLElement creates a new XMLElement instance.
func NewXMLElement() XMLElement {
	return getXMLElementClass().New()
}





// The element nodes in an XML tree structure.
//
// An object may have child nodes, specifically comment nodes, processing-instruction nodes, text nodes, and other nodes. It may also have attribute nodes and namespace nodes associated with it (however, namespace and attribute nodes are not considered children). Any attempt to add a node, node, namespace node, or attribute node as a child raises an exception. If you add a child node to an object and that child already has a parent, raises an exception; the child must be detached or copied first.


// The element nodes in an XML tree structure.
//
// [Full Topic]
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






// Returns an object initialized with the specified name and URI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:uri:)
func NewXMLElementWithNameURI(name IString, URI IString) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[XMLElement](instance.ID, objc.Sel("initWithName:URI:"), name, URI)
	rv.Autorelease()
	return rv
}

















// Adds an attribute node to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/addAttribute(_:)
func (x_ XMLElement) AddAttribute(attribute IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("addAttribute:"), attribute)
}


// Removes a namespace node that is identified by a given prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/removeNamespace(forPrefix:)
func (x_ XMLElement) RemoveNamespaceForPrefix(name IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("removeNamespaceForPrefix:"), name)
}


// Returns the prefix associated with the specified URI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/resolvePrefix(forNamespaceURI:)
func (x_ XMLElement) ResolvePrefixForNamespaceURI(namespaceURI IString) IString {
	rv := objc.Send[String](x_.ID, objc.Sel("resolvePrefixForNamespaceURI:"), namespaceURI)
	return rv
}







// Sets all attributes of the receiver at once, replacing any existing attribute nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/attributes
func (x_ XMLElement) Attributes() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("attributes"))
	return rv
}


// Sets all attributes of the receiver at once, replacing any existing attribute nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/attributes
func (x_ XMLElement) SetAttributes(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setAttributes:"), value)
}


// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/namespaces
func (x_ XMLElement) Namespaces() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("namespaces"))
	return rv
}


// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/namespaces
func (x_ XMLElement) SetNamespaces(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNamespaces:"), value)
}







