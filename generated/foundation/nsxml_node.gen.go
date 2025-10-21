// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XMLNode] class.
var (
	XMLNodeClass     _XMLNodeClass
	XMLNodeClassOnce sync.Once
)

func getXMLNodeClass() _XMLNodeClass {
	XMLNodeClassOnce.Do(func() {
		XMLNodeClass = _XMLNodeClass{objc.GetClass("NSXMLNode")}
	})
	return XMLNodeClass
}

type _XMLNodeClass struct {
	class objc.Class
}

// An interface definition for the [XMLNode] class.
type IXMLNode interface {
	objectivec.IObject
	ChildAtIndex(index uint) unsafe.Pointer
	SetStringValueResolvingEntities(string_ string, resolve bool)
	XMLStringWithOptions(options unsafe.Pointer) string
}

// The nodes in the abstract, logical tree structure that represents an XML document.
//
// Node objects can be of different kinds, corresponding to the following markup constructs in an XML document: element, attribute, text, processing instruction, namespace, and comment. In addition, a document-node object (specifically, an instance of ) represents an XML document in its entirety. objects can also represent document type declarations as well as declarations in Document Type Definitions (DTDs). Class factory methods of enable you to create nodes of each kind. Only document, element, and DTD nodes may have child nodes. Among the XML family of classes (excluding ) the class is the base class. Inheriting from it are the classes , , , and . specifies the interface common to all XML node objects and defines common node behavior and attributes, for example hierarchy level, node name and value, tree traversal, and the ability to emit representative XML markup text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode
type XMLNode struct {
	objectivec.Object
}

// XMLNodeFrom constructs a [XMLNode] from an unsafe.Pointer.
//
// The nodes in the abstract, logical tree structure that represents an XML document.
func XMLNodeFrom(ptr unsafe.Pointer) XMLNode {
	return XMLNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (xc _XMLNodeClass) Alloc() XMLNode {
	rv := objc.Send[XMLNode](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XMLNodeClass) New() XMLNode {
	rv := objc.Send[XMLNode](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XMLNode) Init() XMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XMLNode) Autorelease() XMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLNode creates a new XMLNode instance.
func NewXMLNode() XMLNode {
	return getXMLNodeClass().New()
}




// Returns an instance initialized with the constant indicating node kind.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/init(kind:)
func NewXMLNodeWithKind(kind unsafe.Pointer) XMLNode {
	instance := getXMLNodeClass().Alloc()
	rv := objc.Send[XMLNode](instance.ID, objc.Sel("initWithKind:"), kind)
	rv.Autorelease()
	return rv
}



// Returns an instance initialized with the constant indicating node kind and one or more initialization options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/init(kind:options:)
func NewXMLNodeWithKindOptions(kind unsafe.Pointer, options unsafe.Pointer) XMLNode {
	instance := getXMLNodeClass().Alloc()
	rv := objc.Send[XMLNode](instance.ID, objc.Sel("initWithKind:options:"), kind, options)
	rv.Autorelease()
	return rv
}


// Returns an object representing an attribute node with a given name and string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/attribute(withName:stringValue:)
func (xc _XMLNodeClass) AttributeWithNameStringValue(name string, stringValue string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("attributeWithName:stringValue:"), objc.String(name), objc.String(stringValue))
	return rv
}

// Returns an object representing an attribute node with a given qualified name and string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/attribute(withName:uri:stringValue:)
func (xc _XMLNodeClass) AttributeWithNameURIStringValue(name string, URI string, stringValue string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("attributeWithName:URI:stringValue:"), objc.String(name), objc.String(URI), objc.String(stringValue))
	return rv
}

// Returns an object representing a comment node containing given text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/comment(withStringValue:)
func (xc _XMLNodeClass) CommentWithStringValue(stringValue string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("commentWithStringValue:"), objc.String(stringValue))
	return rv
}

// Returns an empty document node.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/document()
func (xc _XMLNodeClass) Document() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("document"))
	return rv
}

// Returns an object initialized with a given root element.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/document(withRootElement:)
func (xc _XMLNodeClass) DocumentWithRootElement(element unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("documentWithRootElement:"), element)
	return rv
}

// Returns a object representing the DTD declaration for an element, attribute, entity, or notation based on a given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/dtdNode(withXMLString:)
func (xc _XMLNodeClass) DTDNodeWithXMLString(string_ string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("DTDNodeWithXMLString:"), objc.String(string_))
	return rv
}

// Returns an object with a given tag identifier, or name
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:)
func (xc _XMLNodeClass) ElementWithName(name string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("elementWithName:"), objc.String(name))
	return rv
}

// Returns an object with the given tag (name), attributes, and children.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:children:attributes:)
func (xc _XMLNodeClass) ElementWithNameChildrenAttributes(name string, children unsafe.Pointer, attributes unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("elementWithName:children:attributes:"), objc.String(name), children, attributes)
	return rv
}

// Returns an object with a single text-node child containing the specified text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:stringValue:)
func (xc _XMLNodeClass) ElementWithNameStringValue(name string, string_ string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("elementWithName:stringValue:"), objc.String(name), objc.String(string_))
	return rv
}

// Returns an element whose fully qualified name is specified.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:uri:)
func (xc _XMLNodeClass) ElementWithNameURI(name string, URI string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("elementWithName:URI:"), objc.String(name), objc.String(URI))
	return rv
}

// Returns an object representing a namespace with a specified name and URI.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/namespace(withName:stringValue:)
func (xc _XMLNodeClass) NamespaceWithNameStringValue(name string, stringValue string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("namespaceWithName:stringValue:"), objc.String(name), objc.String(stringValue))
	return rv
}

// Returns an object representing one of the predefined namespaces with the specified prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/predefinedNamespace(forPrefix:)
func (xc _XMLNodeClass) PredefinedNamespaceForPrefix(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(xc.class), objc.Sel("predefinedNamespaceForPrefix:"), objc.String(name))
	return rv
}

// Returns an object representing a processing instruction with a specified name and value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/processingInstruction(withName:stringValue:)
func (xc _XMLNodeClass) ProcessingInstructionWithNameStringValue(name string, stringValue string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("processingInstructionWithName:stringValue:"), objc.String(name), objc.String(stringValue))
	return rv
}

// Returns an object representing a text node with specified content.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/text(withStringValue:)
func (xc _XMLNodeClass) TextWithStringValue(stringValue string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("textWithStringValue:"), objc.String(stringValue))
	return rv
}

// Returns the child node of the receiver at the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/child(at:)
func (x_ XMLNode) ChildAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("childAtIndex:"), index)
	return rv
}

// Sets the content of the receiver as a string value and, optionally, resolves character references, predefined entities, and user-defined entities as declared in the associated DTD.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/setStringValue(_:resolvingEntities:)
func (x_ XMLNode) SetStringValueResolvingEntities(string_ string, resolve bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setStringValue:resolvingEntities:"), objc.String(string_), resolve)
}

// Returns the string representation of the receiver as it would appear in an XML document, with one or more output options specified.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/xmlString(options:)
func (x_ XMLNode) XMLStringWithOptions(options unsafe.Pointer) string {
	rv := objc.Send[string](x_.ID, objc.Sel("XMLStringWithOptions:"), options)
	return rv
}

// Returns the index of the receiver identifying its position relative to its sibling nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/index
func (x_ XMLNode) Index() uint {
	rv := objc.Send[uint](x_.ID, objc.Sel("index"))
	return rv
}

// Returns the kind of node the receiver is as a constant of type .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/kind-swift.property
func (x_ XMLNode) Kind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("kind"))
	return rv
}

// Returns the nesting level of the receiver within the tree hierarchy.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/level
func (x_ XMLNode) Level() uint {
	rv := objc.Send[uint](x_.ID, objc.Sel("level"))
	return rv
}

// Returns the name of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/name
func (x_ XMLNode) Name() string {
	rv := objc.Send[string](x_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// Returns the name of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/name
func (x_ XMLNode) SetName(value string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setName:"), objc.String(value))
}

// Returns the object value of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/objectValue
func (x_ XMLNode) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("objectValue"))
	return rv
}


// SetObjectValue sets the value of the objectValue property.
// Returns the object value of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/objectValue
func (x_ XMLNode) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setObjectValue:"), value)
}

// Returns the parent node of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/parent
func (x_ XMLNode) Parent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("parent"))
	return rv
}

// Returns the content of the receiver as a string value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/stringValue
func (x_ XMLNode) StringValue() string {
	rv := objc.Send[string](x_.ID, objc.Sel("stringValue"))
	return rv
}


// SetStringValue sets the value of the stringValue property.
// Returns the content of the receiver as a string value.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/stringValue
func (x_ XMLNode) SetStringValue(value string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setStringValue:"), objc.String(value))
}

// Returns the URI associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/uri
func (x_ XMLNode) URI() string {
	rv := objc.Send[string](x_.ID, objc.Sel("URI"))
	return rv
}


// SetURI sets the value of the URI property.
// Returns the URI associated with the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/uri
func (x_ XMLNode) SetURI(value string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setURI:"), objc.String(value))
}


