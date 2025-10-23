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
	AddAttribute(attribute IXMLNode)
	AddChild(child IXMLNode)
	AddNamespace(aNamespace IXMLNode)
	AttributeForLocalNameURI(localName string, URI string) XMLNode
	AttributeForName(name string) XMLNode
	ElementsForLocalNameURI(localName string, URI string) []XMLElement
	ElementsForName(name string) []XMLElement
	InsertChildAtIndex(child IXMLNode, index uint)
	InsertChildrenAtIndex(children []XMLNode, index uint)
	NamespaceForPrefix(name string) XMLNode
	NormalizeAdjacentTextNodesPreservingCDATA(preserve bool)
	RemoveAttributeForName(name string)
	RemoveChildAtIndex(index uint)
	RemoveNamespaceForPrefix(name string)
	ReplaceChildAtIndexWithNode(index uint, node IXMLNode)
	ResolveNamespaceForName(name string) XMLNode
	ResolvePrefixForNamespaceURI(namespaceURI string) String
	SetAttributesAsDictionary(attributes objectivec.IObject)
	SetAttributesWithDictionary(attributes IDictionary)
	SetChildren(children []XMLNode)
	Attributes() []XMLNode
	SetAttributes(value []XMLNode)
	Namespaces() []XMLNode
	SetNamespaces(value []XMLNode)
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/init(kind:options:)
func NewXMLElementWithKindOptions(kind NSXMLNodeKind, options NSXMLNodeOptions) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[XMLElement](instance.ID, objc.Sel("initWithKind:options:"), kind, options)
	rv.Autorelease()
	return rv
}


// Returns an object initialized with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:)
func NewXMLElementWithName(name string) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[XMLElement](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	rv.Autorelease()
	return rv
}


// Returns an object initialized with a specified name and a single text-node child containing a specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:stringValue:)
func NewXMLElementWithNameStringValue(name string, string_ string) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[XMLElement](instance.ID, objc.Sel("initWithName:stringValue:"), objc.String(name), objc.String(string_))
	rv.Autorelease()
	return rv
}


// Returns an object initialized with the specified name and URI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:uri:)
func NewXMLElementWithNameURI(name string, URI string) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[XMLElement](instance.ID, objc.Sel("initWithName:URI:"), objc.String(name), objc.String(URI))
	rv.Autorelease()
	return rv
}


// Returns an object created from a specified string containing XML markup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/init(xmlString:)
func NewXMLElementWithXMLStringError(string_ string, error_ IError) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[XMLElement](instance.ID, objc.Sel("initWithXMLString:error:"), objc.String(string_), error_)
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


// Adds a child node at the end of the receiver’s current list of children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/addChild(_:)
func (x_ XMLElement) AddChild(child IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("addChild:"), child)
}


// Adds a namespace node to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/addNamespace(_:)
func (x_ XMLElement) AddNamespace(aNamespace IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("addNamespace:"), aNamespace)
}


// Returns the attribute node of the receiver that is identified by a local name and URI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/attribute(forLocalName:uri:)
func (x_ XMLElement) AttributeForLocalNameURI(localName string, URI string) XMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("attributeForLocalName:URI:"), objc.String(localName), objc.String(URI))
	return rv
}


// Returns the attribute node of the receiver with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/attribute(forName:)
func (x_ XMLElement) AttributeForName(name string) XMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("attributeForName:"), objc.String(name))
	return rv
}


// Returns the child element nodes (as objects) of the receiver that are matched with the specified local name and URI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/elements(forLocalName:uri:)
func (x_ XMLElement) ElementsForLocalNameURI(localName string, URI string) []XMLElement {
	rv := objc.Send[[]XMLElement](x_.ID, objc.Sel("elementsForLocalName:URI:"), objc.String(localName), objc.String(URI))
	return rv
}


// Returns the child element nodes (as objects) of the receiver that have a specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/elements(forName:)
func (x_ XMLElement) ElementsForName(name string) []XMLElement {
	rv := objc.Send[[]XMLElement](x_.ID, objc.Sel("elementsForName:"), objc.String(name))
	return rv
}


// Inserts a new child node at a specified location in the receiver’s list of child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/insertChild(_:at:)
func (x_ XMLElement) InsertChildAtIndex(child IXMLNode, index uint) {
	objc.Send[objc.ID](x_.ID, objc.Sel("insertChild:atIndex:"), child, index)
}


// Inserts an array of child nodes at a specified location in the receiver’s list of children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/insertChildren(_:at:)
func (x_ XMLElement) InsertChildrenAtIndex(children []XMLNode, index uint) {
	objc.Send[objc.ID](x_.ID, objc.Sel("insertChildren:atIndex:"), children, index)
}


// Returns the namespace node with a specified prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/namespace(forPrefix:)
func (x_ XMLElement) NamespaceForPrefix(name string) XMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("namespaceForPrefix:"), objc.String(name))
	return rv
}


// Coalesces adjacent text nodes of the receiver that you have explicitly added, optionally including CDATA sections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/normalizeAdjacentTextNodesPreservingCDATA(_:)
func (x_ XMLElement) NormalizeAdjacentTextNodesPreservingCDATA(preserve bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("normalizeAdjacentTextNodesPreservingCDATA:"), preserve)
}


// Removes an attribute node identified by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/removeAttribute(forName:)
func (x_ XMLElement) RemoveAttributeForName(name string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("removeAttributeForName:"), objc.String(name))
}


// Removes the child node of the receiver identified by a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/removeChild(at:)
func (x_ XMLElement) RemoveChildAtIndex(index uint) {
	objc.Send[objc.ID](x_.ID, objc.Sel("removeChildAtIndex:"), index)
}


// Removes a namespace node that is identified by a given prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/removeNamespace(forPrefix:)
func (x_ XMLElement) RemoveNamespaceForPrefix(name string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("removeNamespaceForPrefix:"), objc.String(name))
}


// Replaces a child node at a specified location with another child node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/replaceChild(at:with:)
func (x_ XMLElement) ReplaceChildAtIndexWithNode(index uint, node IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("replaceChildAtIndex:withNode:"), index, node)
}


// Returns the namespace node with the prefix matching the given qualified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/resolveNamespace(forName:)
func (x_ XMLElement) ResolveNamespaceForName(name string) XMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("resolveNamespaceForName:"), objc.String(name))
	return rv
}


// Returns the prefix associated with the specified URI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/resolvePrefix(forNamespaceURI:)
func (x_ XMLElement) ResolvePrefixForNamespaceURI(namespaceURI string) String {
	rv := objc.Send[String](x_.ID, objc.Sel("resolvePrefixForNamespaceURI:"), objc.String(namespaceURI))
	return rv
}


// Sets the attributes of the receiver based on the key-value pairs specified in the passed-in dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/setAttributesAs(_:)
func (x_ XMLElement) SetAttributesAsDictionary(attributes objectivec.IObject) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setAttributesAsDictionary:"), attributes)
}


// Sets the attributes of the receiver based on the key-value pairs specified in the passed dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/setAttributesWith(_:)
func (x_ XMLElement) SetAttributesWithDictionary(attributes IDictionary) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setAttributesWithDictionary:"), attributes)
}


// Sets all child nodes of the receiver at once, replacing any existing children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/setChildren(_:)
func (x_ XMLElement) SetChildren(children []XMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setChildren:"), children)
}


// Sets all attributes of the receiver at once, replacing any existing attribute nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/attributes
func (x_ XMLElement) Attributes() []XMLNode {
	rv := objc.Send[[]XMLNode](x_.ID, objc.Sel("attributes"))
	return rv
}


// Sets all attributes of the receiver at once, replacing any existing attribute nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/attributes
func (x_ XMLElement) SetAttributes(value []XMLNode) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](x_.ID, objc.Sel("setAttributes:"), nsArray)
}


// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/namespaces
func (x_ XMLElement) Namespaces() []XMLNode {
	rv := objc.Send[[]XMLNode](x_.ID, objc.Sel("namespaces"))
	return rv
}


// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLElement/namespaces
func (x_ XMLElement) SetNamespaces(value []XMLNode) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](x_.ID, objc.Sel("setNamespaces:"), nsArray)
}


