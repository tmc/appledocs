// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	AddNamespace(aNamespace unsafe.Pointer)
	AttributeForLocalNameURI(localName string, URI string) unsafe.Pointer
	AttributeForName(name string) unsafe.Pointer
	ElementsForName(name string) []XMLElement
	InsertChildAtIndex(child unsafe.Pointer, index uint)
	NamespaceForPrefix(name string) unsafe.Pointer
	RemoveChildAtIndex(index uint)
	RemoveNamespaceForPrefix(name string)
	ReplaceChildAtIndexWithNode(index uint, node unsafe.Pointer)
}

// The element nodes in an XML tree structure.
//
// An object may have child nodes, specifically comment nodes, processing-instruction nodes, text nodes, and other nodes. It may also have attribute nodes and namespace nodes associated with it (however, namespace and attribute nodes are not considered children). Any attempt to add a node, node, namespace node, or attribute node as a child raises an exception. If you add a child node to an object and that child already has a parent, raises an exception; the child must be detached or copied first.
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

// Alloc allocates a new instance without initialization.
func (xc _XMLElementClass) Alloc() XMLElement {
	rv := objc.Send[XMLElement](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Returns an object initialized with a specified name and a single text-node child containing a specified value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:stringValue:)
func NewXMLElementWithNameStringValue(name string, string_ string) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[XMLElement](instance.ID, objc.Sel("initWithName:stringValue:"), objc.String(name), objc.String(string_))
	rv.Autorelease()
	return rv
}



// Returns an object initialized with the specified name and URI.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:uri:)
func NewXMLElementWithNameURI(name string, URI string) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[XMLElement](instance.ID, objc.Sel("initWithName:URI:"), objc.String(name), objc.String(URI))
	rv.Autorelease()
	return rv
}


// Adds a namespace node to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/addNamespace(_:)
func (x_ XMLElement) AddNamespace(aNamespace unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("addNamespace:"), aNamespace)
}

// Returns the attribute node of the receiver that is identified by a local name and URI.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/attribute(forLocalName:uri:)
func (x_ XMLElement) AttributeForLocalNameURI(localName string, URI string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("attributeForLocalName:URI:"), objc.String(localName), objc.String(URI))
	return rv
}

// Returns the attribute node of the receiver with the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/attribute(forName:)
func (x_ XMLElement) AttributeForName(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("attributeForName:"), objc.String(name))
	return rv
}

// Returns the child element nodes (as objects) of the receiver that have a specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/elements(forName:)
func (x_ XMLElement) ElementsForName(name string) []XMLElement {
	rv := objc.Send[[]XMLElement](x_.ID, objc.Sel("elementsForName:"), objc.String(name))
	return rv
}

// Inserts a new child node at a specified location in the receiver’s list of child nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/insertChild(_:at:)
func (x_ XMLElement) InsertChildAtIndex(child unsafe.Pointer, index uint) {
	objc.Send[objc.ID](x_.ID, objc.Sel("insertChild:atIndex:"), child, index)
}

// Returns the namespace node with a specified prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/namespace(forPrefix:)
func (x_ XMLElement) NamespaceForPrefix(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("namespaceForPrefix:"), objc.String(name))
	return rv
}

// Removes the child node of the receiver identified by a given index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/removeChild(at:)
func (x_ XMLElement) RemoveChildAtIndex(index uint) {
	objc.Send[objc.ID](x_.ID, objc.Sel("removeChildAtIndex:"), index)
}

// Removes a namespace node that is identified by a given prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/removeNamespace(forPrefix:)
func (x_ XMLElement) RemoveNamespaceForPrefix(name string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("removeNamespaceForPrefix:"), objc.String(name))
}

// Replaces a child node at a specified location with another child node.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/replaceChild(at:with:)
func (x_ XMLElement) ReplaceChildAtIndexWithNode(index uint, node unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("replaceChildAtIndex:withNode:"), index, node)
}

// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/namespaces
func (x_ XMLElement) Namespaces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("namespaces"))
	return rv
}


// SetNamespaces sets the value of the namespaces property.
// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/namespaces
func (x_ XMLElement) SetNamespaces(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNamespaces:"), value)
}

// Sets all attributes of the receiver at once, replacing any existing attribute nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/attributes
func (x_ XMLElement) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("attributes"))
	return rv
}


// SetAttributes sets the value of the attributes property.
// Sets all attributes of the receiver at once, replacing any existing attribute nodes.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/attributes
func (x_ XMLElement) SetAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setAttributes:"), value)
}


