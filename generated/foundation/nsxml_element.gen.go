// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSXMLElement */


/* debug [class_header]: Header for NSXMLElement */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for XMLElement */
// An interface definition for the [XMLElement] class.
type IXMLElement interface {
	IXMLNode
	
/* debug [class_interface_properties]: Properties for XMLElement */
	// properties:
	Attributes() IXMLNode
	SetAttributes(value IXMLNode)
	Namespaces() IXMLNode
	SetNamespaces(value IXMLNode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for XMLElement */
	// methods:
	AddAttribute(attribute IXMLNode)
	RemoveNamespaceForPrefix(name IString)
	ResolvePrefixForNamespaceURI(namespaceURI IString) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for XMLElement */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for XMLElement */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for XMLElement */

// Returns an object initialized with the specified name and URI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:uri:)
func NewXMLElementWithNameURI(name IString, URI IString) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[XMLElement](instance.ID, objc.Sel("initWithName:URI:"), name, URI)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewXMLElementWithNameURI */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for XMLElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for XMLElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for XMLElement */

// Adds an attribute node to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/addAttribute(_:)
func (x_ XMLElement) AddAttribute(attribute IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("addAttribute:"), attribute)
}/* debug [instance_methods/method]: AddAttribute */


// Removes a namespace node that is identified by a given prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/removeNamespace(forPrefix:)
func (x_ XMLElement) RemoveNamespaceForPrefix(name IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("removeNamespaceForPrefix:"), name)
}/* debug [instance_methods/method]: RemoveNamespaceForPrefix */


// Returns the prefix associated with the specified URI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/resolvePrefix(forNamespaceURI:)
func (x_ XMLElement) ResolvePrefixForNamespaceURI(namespaceURI IString) IString {
	rv := objc.Send[String](x_.ID, objc.Sel("resolvePrefixForNamespaceURI:"), namespaceURI)
	return rv
}/* debug [instance_methods/method]: ResolvePrefixForNamespaceURI */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for XMLElement */

// Sets all attributes of the receiver at once, replacing any existing attribute nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/attributes
func (x_ XMLElement) Attributes() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("attributes"))
	return rv
}/* debug [instance_properties/getter]: attributes */


// Sets all attributes of the receiver at once, replacing any existing attribute nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/attributes
func (x_ XMLElement) SetAttributes(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setAttributes:"), value)
}/* debug [instance_properties/setter]: attributes */


// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/namespaces
func (x_ XMLElement) Namespaces() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("namespaces"))
	return rv
}/* debug [instance_properties/getter]: namespaces */


// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlelement/namespaces
func (x_ XMLElement) SetNamespaces(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNamespaces:"), value)
}/* debug [instance_properties/setter]: namespaces */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSXMLElement */


