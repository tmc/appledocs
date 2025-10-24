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
	// properties:
	ChildCount() uint
	Children() []XMLNode
	Description() IString
	Index() uint
	Level() uint
	LocalName() IString
	Name() IString
	SetName(value IString)
	NextNode() IXMLNode
	NextSibling() IXMLNode
	ObjectValue() objc.ID
	SetObjectValue(value objc.ID)
	Parent() IXMLNode
	Prefix() IString
	PreviousNode() IXMLNode
	PreviousSibling() IXMLNode
	RootDocument() IXMLDocument
	StringValue() IString
	SetStringValue(value IString)
	URI() IString
	SetURI(value IString)
	XPath() IString
	XMLString() IString
	Kind() unsafe.Pointer
	SetKind(value unsafe.Pointer)
	Next() IXMLNode
	SetNext(value IXMLNode)
	Previous() IXMLNode
	SetPrevious(value IXMLNode)
	// methods:
	CanonicalXMLStringPreservingComments(comments bool) IString
	ChildAtIndex(index uint) IXMLNode
	Detach()
	NodesForXPathError(xpath IString, error_ IError) []XMLNode
	ObjectsForXQueryError(xquery IString, error_ IError) IArray
	ObjectsForXQueryConstantsError(xquery IString, constants IDictionary, error_ IError) IArray
	SetStringValueResolvingEntities(string_ IString, resolve bool)
	XMLStringWithOptions(options XMLNodeOptions) IString
}

// The nodes in the abstract, logical tree structure that represents an XML document.
//
// Node objects can be of different kinds, corresponding to the following markup constructs in an XML document: element, attribute, text, processing instruction, namespace, and comment. In addition, a document-node object (specifically, an instance of ) represents an XML document in its entirety. objects can also represent document type declarations as well as declarations in Document Type Definitions (DTDs). Class factory methods of enable you to create nodes of each kind. Only document, element, and DTD nodes may have child nodes. Among the XML family of classes (excluding ) the class is the base class. Inheriting from it are the classes , , , and . specifies the interface common to all XML node objects and defines common node behavior and attributes, for example hierarchy level, node name and value, tree traversal, and the ability to emit representative XML markup text.


// The nodes in the abstract, logical tree structure that represents an XML document.
//
// [Full Topic]
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




// Returns an object representing an attribute node with a given name and string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/attribute(withName:stringValue:)
func (xc _XMLNodeClass) AttributeWithNameStringValue(name IString, stringValue IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("attributeWithName:stringValue:"), name, stringValue)
	return rv
}


// Returns an object representing an attribute node with a given qualified name and string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/attribute(withName:uri:stringValue:)
func (xc _XMLNodeClass) AttributeWithNameURIStringValue(name IString, URI IString, stringValue IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("attributeWithName:URI:stringValue:"), name, URI, stringValue)
	return rv
}


// Returns an object representing a comment node containing given text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/comment(withStringValue:)
func (xc _XMLNodeClass) CommentWithStringValue(stringValue IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("commentWithStringValue:"), stringValue)
	return rv
}


// Returns an empty document node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/document()
func (xc _XMLNodeClass) Document() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("document"))
	return rv
}


// Returns an object initialized with a given root element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/document(withRootElement:)
func (xc _XMLNodeClass) DocumentWithRootElement(element IXMLElement) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("documentWithRootElement:"), element)
	return rv
}


// Returns a object representing the DTD declaration for an element, attribute, entity, or notation based on a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/dtdNode(withXMLString:)
func (xc _XMLNodeClass) DTDNodeWithXMLString(string_ IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("DTDNodeWithXMLString:"), string_)
	return rv
}


// Returns an object with a given tag identifier, or name
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:)
func (xc _XMLNodeClass) ElementWithName(name IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("elementWithName:"), name)
	return rv
}


// Returns an object with the given tag (name), attributes, and children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:children:attributes:)
func (xc _XMLNodeClass) ElementWithNameChildrenAttributes(name IString, children []XMLNode, attributes []XMLNode) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("elementWithName:children:attributes:"), name, children, attributes)
	return rv
}


// Returns an object with a single text-node child containing the specified text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:stringValue:)
func (xc _XMLNodeClass) ElementWithNameStringValue(name IString, string_ IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("elementWithName:stringValue:"), name, string_)
	return rv
}


// Returns an element whose fully qualified name is specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:uri:)
func (xc _XMLNodeClass) ElementWithNameURI(name IString, URI IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("elementWithName:URI:"), name, URI)
	return rv
}


// Returns the local name from the specified qualified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/localName(forName:)
func (xc _XMLNodeClass) LocalNameForName(name IString) IString {
	rv := objc.Send[String](objc.ID(xc.class), objc.Sel("localNameForName:"), name)
	return rv
}


// Returns an object representing a namespace with a specified name and URI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/namespace(withName:stringValue:)
func (xc _XMLNodeClass) NamespaceWithNameStringValue(name IString, stringValue IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("namespaceWithName:stringValue:"), name, stringValue)
	return rv
}


// Returns an object representing one of the predefined namespaces with the specified prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/predefinedNamespace(forPrefix:)
func (xc _XMLNodeClass) PredefinedNamespaceForPrefix(name IString) IXMLNode {
	rv := objc.Send[XMLNode](objc.ID(xc.class), objc.Sel("predefinedNamespaceForPrefix:"), name)
	return rv
}


// Returns the prefix from the specified qualified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/prefix(forName:)
func (xc _XMLNodeClass) PrefixForName(name IString) IString {
	rv := objc.Send[String](objc.ID(xc.class), objc.Sel("prefixForName:"), name)
	return rv
}


// Returns an object representing a processing instruction with a specified name and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/processingInstruction(withName:stringValue:)
func (xc _XMLNodeClass) ProcessingInstructionWithNameStringValue(name IString, stringValue IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("processingInstructionWithName:stringValue:"), name, stringValue)
	return rv
}


// Returns an object representing a text node with specified content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/text(withStringValue:)
func (xc _XMLNodeClass) TextWithStringValue(stringValue IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(xc.class), objc.Sel("textWithStringValue:"), stringValue)
	return rv
}


// Returns a string object encapsulating the receiver’s XML in canonical form.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/canonicalXMLStringPreservingComments(_:)
func (x_ XMLNode) CanonicalXMLStringPreservingComments(comments bool) IString {
	rv := objc.Send[String](x_.ID, objc.Sel("canonicalXMLStringPreservingComments:"), comments)
	return rv
}


// Returns the child node of the receiver at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/child(at:)
func (x_ XMLNode) ChildAtIndex(index uint) IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("childAtIndex:"), index)
	return rv
}


// Detaches the receiver from its parent node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/detach()
func (x_ XMLNode) Detach() {
	objc.Send[objc.ID](x_.ID, objc.Sel("detach"))
}


// Returns the nodes resulting from executing an XPath query upon the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/nodes(forXPath:)
func (x_ XMLNode) NodesForXPathError(xpath IString, error_ IError) []XMLNode {
	rv := objc.Send[[]XMLNode](x_.ID, objc.Sel("nodesForXPath:error:"), xpath, error_)
	return rv
}


// Returns the objects resulting from executing an XQuery query upon the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/objects(forXQuery:)
func (x_ XMLNode) ObjectsForXQueryError(xquery IString, error_ IError) IArray {
	rv := objc.Send[Array](x_.ID, objc.Sel("objectsForXQuery:error:"), xquery, error_)
	return rv
}


// Returns the objects resulting from executing an XQuery query upon the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/objects(forXQuery:constants:)
func (x_ XMLNode) ObjectsForXQueryConstantsError(xquery IString, constants IDictionary, error_ IError) IArray {
	rv := objc.Send[Array](x_.ID, objc.Sel("objectsForXQuery:constants:error:"), xquery, constants, error_)
	return rv
}


// Sets the content of the receiver as a string value and, optionally, resolves character references, predefined entities, and user-defined entities as declared in the associated DTD.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/setStringValue(_:resolvingEntities:)
func (x_ XMLNode) SetStringValueResolvingEntities(string_ IString, resolve bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setStringValue:resolvingEntities:"), string_, resolve)
}


// Returns the string representation of the receiver as it would appear in an XML document, with one or more output options specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/xmlString(options:)
func (x_ XMLNode) XMLStringWithOptions(options XMLNodeOptions) IString {
	rv := objc.Send[String](x_.ID, objc.Sel("XMLStringWithOptions:"), options)
	return rv
}


// Returns the number of child nodes the receiver has.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/childCount
func (x_ XMLNode) ChildCount() uint {
	rv := objc.Send[uint](x_.ID, objc.Sel("childCount"))
	return rv
}


// Returns an immutable array containing the child nodes of the receiver (as objects).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/children
func (x_ XMLNode) Children() []XMLNode {
	rv := objc.Send[[]XMLNode](x_.ID, objc.Sel("children"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/description
func (x_ XMLNode) Description() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("description"))
	return rv
}


// Returns the index of the receiver identifying its position relative to its sibling nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/index
func (x_ XMLNode) Index() uint {
	rv := objc.Send[uint](x_.ID, objc.Sel("index"))
	return rv
}


// Returns the nesting level of the receiver within the tree hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/level
func (x_ XMLNode) Level() uint {
	rv := objc.Send[uint](x_.ID, objc.Sel("level"))
	return rv
}


// Returns the local name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/localName
func (x_ XMLNode) LocalName() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("localName"))
	return rv
}


// Returns the name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/name
func (x_ XMLNode) Name() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("name"))
	return rv
}


// Returns the name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/name
func (x_ XMLNode) SetName(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setName:"), value)
}


// Returns the next object in document order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/next
func (x_ XMLNode) NextNode() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("nextNode"))
	return rv
}


// Returns the next object that is a sibling node to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/nextSibling
func (x_ XMLNode) NextSibling() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("nextSibling"))
	return rv
}


// Returns the object value of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/objectValue
func (x_ XMLNode) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("objectValue"))
	return rv
}


// Returns the object value of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/objectValue
func (x_ XMLNode) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setObjectValue:"), value)
}


// Returns the parent node of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/parent
func (x_ XMLNode) Parent() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("parent"))
	return rv
}


// Returns the prefix of the receiver’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/prefix
func (x_ XMLNode) Prefix() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("prefix"))
	return rv
}


// Returns the previous object in document order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/previous
func (x_ XMLNode) PreviousNode() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("previousNode"))
	return rv
}


// Returns the previous object that is a sibling node to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/previousSibling
func (x_ XMLNode) PreviousSibling() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("previousSibling"))
	return rv
}


// Returns the object containing the root element and representing the XML document as a whole.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/rootDocument
func (x_ XMLNode) RootDocument() IXMLDocument {
	rv := objc.Send[XMLDocument](x_.ID, objc.Sel("rootDocument"))
	return rv
}


// Returns the content of the receiver as a string value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/stringValue
func (x_ XMLNode) StringValue() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("stringValue"))
	return rv
}


// Returns the content of the receiver as a string value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/stringValue
func (x_ XMLNode) SetStringValue(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setStringValue:"), value)
}


// Returns the URI associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/uri
func (x_ XMLNode) URI() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("URI"))
	return rv
}


// Returns the URI associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/uri
func (x_ XMLNode) SetURI(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setURI:"), value)
}


// Returns the XPath expression identifying the receiver’s location in the document tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/xPath
func (x_ XMLNode) XPath() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("XPath"))
	return rv
}


// Returns the string representation of the receiver as it would appear in an XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLNode/xmlString
func (x_ XMLNode) XMLString() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("XMLString"))
	return rv
}


// Returns the kind of node the receiver is as a constant of type
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/kind-swift.property
func (x_ XMLNode) Kind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("kind"))
	return rv
}


// Returns the kind of node the receiver is as a constant of type
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/kind-swift.property
func (x_ XMLNode) SetKind(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setKind:"), value)
}


// Returns the next
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/next
func (x_ XMLNode) Next() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("next"))
	return rv
}


// Returns the next
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/next
func (x_ XMLNode) SetNext(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNext:"), value)
}


// Returns the previous
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/previous
func (x_ XMLNode) Previous() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("previous"))
	return rv
}


// Returns the previous
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/previous
func (x_ XMLNode) SetPrevious(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setPrevious:"), value)
}


