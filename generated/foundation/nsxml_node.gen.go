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
	ChildCount() int /* primitive/slice/pointer. */
	SetChildCount(value int /* primitive/slice/pointer. */)
	Children() IXMLNode
	SetChildren(value IXMLNode)
	Description() IString
	SetDescription(value IString)
	Index() int /* primitive/slice/pointer. */
	SetIndex(value int /* primitive/slice/pointer. */)
	Kind() unsafe.Pointer
	SetKind(value unsafe.Pointer)
	Level() int /* primitive/slice/pointer. */
	SetLevel(value int /* primitive/slice/pointer. */)
	LocalName() IString
	SetLocalName(value IString)
	Name() IString
	SetName(value IString)
	Next() IXMLNode
	SetNext(value IXMLNode)
	NextSibling() IXMLNode
	SetNextSibling(value IXMLNode)
	ObjectValue() unsafe.Pointer
	SetObjectValue(value unsafe.Pointer)
	Parent() IXMLNode
	SetParent(value IXMLNode)
	Prefix() IString
	SetPrefix(value IString)
	Previous() IXMLNode
	SetPrevious(value IXMLNode)
	PreviousSibling() IXMLNode
	SetPreviousSibling(value IXMLNode)
	RootDocument() IXMLDocument
	SetRootDocument(value IXMLDocument)
	StringValue() IString
	SetStringValue(value IString)
	Uri() IString
	SetUri(value IString)
	XPath() IString
	SetXPath(value IString)
	XmlString() IString
	SetXmlString(value IString)
	// methods:
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



// Returns the number of child nodes the receiver has.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/childcount
func (x_ XMLNode) ChildCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](x_.ID, objc.Sel("childCount"))
	return rv
}


// Returns the number of child nodes the receiver has.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/childcount
func (x_ XMLNode) SetChildCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setChildCount:"), value)
}


// Returns an immutable array containing the child nodes of the receiver (as
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/children
func (x_ XMLNode) Children() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("children"))
	return rv
}


// Returns an immutable array containing the child nodes of the receiver (as
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/children
func (x_ XMLNode) SetChildren(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setChildren:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/description
func (x_ XMLNode) Description() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("description"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/description
func (x_ XMLNode) SetDescription(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDescription:"), value)
}


// Returns the index of the receiver identifying its position relative to its sibling nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/index
func (x_ XMLNode) Index() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](x_.ID, objc.Sel("index"))
	return rv
}


// Returns the index of the receiver identifying its position relative to its sibling nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/index
func (x_ XMLNode) SetIndex(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setIndex:"), value)
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


// Returns the nesting level of the receiver within the tree hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/level
func (x_ XMLNode) Level() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](x_.ID, objc.Sel("level"))
	return rv
}


// Returns the nesting level of the receiver within the tree hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/level
func (x_ XMLNode) SetLevel(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setLevel:"), value)
}


// Returns the local name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/localname
func (x_ XMLNode) LocalName() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("localName"))
	return rv
}


// Returns the local name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/localname
func (x_ XMLNode) SetLocalName(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setLocalName:"), value)
}


// Returns the name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/name
func (x_ XMLNode) Name() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("name"))
	return rv
}


// Returns the name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/name
func (x_ XMLNode) SetName(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setName:"), value)
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


// Returns the next
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/nextsibling
func (x_ XMLNode) NextSibling() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("nextSibling"))
	return rv
}


// Returns the next
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/nextsibling
func (x_ XMLNode) SetNextSibling(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNextSibling:"), value)
}


// Returns the object value of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/objectvalue
func (x_ XMLNode) ObjectValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("objectValue"))
	return rv
}


// Returns the object value of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/objectvalue
func (x_ XMLNode) SetObjectValue(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setObjectValue:"), value)
}


// Returns the parent node of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/parent
func (x_ XMLNode) Parent() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("parent"))
	return rv
}


// Returns the parent node of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/parent
func (x_ XMLNode) SetParent(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setParent:"), value)
}


// Returns the prefix of the receiver’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/prefix
func (x_ XMLNode) Prefix() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("prefix"))
	return rv
}


// Returns the prefix of the receiver’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/prefix
func (x_ XMLNode) SetPrefix(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setPrefix:"), value)
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


// Returns the previous
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/previoussibling
func (x_ XMLNode) PreviousSibling() IXMLNode {
	rv := objc.Send[XMLNode](x_.ID, objc.Sel("previousSibling"))
	return rv
}


// Returns the previous
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/previoussibling
func (x_ XMLNode) SetPreviousSibling(value IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setPreviousSibling:"), value)
}


// Returns the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/rootdocument
func (x_ XMLNode) RootDocument() IXMLDocument {
	rv := objc.Send[XMLDocument](x_.ID, objc.Sel("rootDocument"))
	return rv
}


// Returns the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/rootdocument
func (x_ XMLNode) SetRootDocument(value IXMLDocument) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setRootDocument:"), value)
}


// Returns the content of the receiver as a string value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/stringvalue
func (x_ XMLNode) StringValue() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("stringValue"))
	return rv
}


// Returns the content of the receiver as a string value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/stringvalue
func (x_ XMLNode) SetStringValue(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setStringValue:"), value)
}


// Returns the URI associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/uri
func (x_ XMLNode) Uri() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("uri"))
	return rv
}


// Returns the URI associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/uri
func (x_ XMLNode) SetUri(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setUri:"), value)
}


// Returns the XPath expression identifying the receiver’s location in the document tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/xpath
func (x_ XMLNode) XPath() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("xPath"))
	return rv
}


// Returns the XPath expression identifying the receiver’s location in the document tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/xpath
func (x_ XMLNode) SetXPath(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setXPath:"), value)
}


// Returns the string representation of the receiver as it would appear in an XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/xmlstring
func (x_ XMLNode) XmlString() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("xmlString"))
	return rv
}


// Returns the string representation of the receiver as it would appear in an XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmlnode/xmlstring
func (x_ XMLNode) SetXmlString(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setXmlString:"), value)
}



