// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDocument] class.
var (
	XMLDocumentClass     _XMLDocumentClass
	XMLDocumentClassOnce sync.Once
)

func getXMLDocumentClass() _XMLDocumentClass {
	XMLDocumentClassOnce.Do(func() {
		XMLDocumentClass = _XMLDocumentClass{objc.GetClass("NSXMLDocument")}
	})
	return XMLDocumentClass
}

type _XMLDocumentClass struct {
	class objc.Class
}

// An interface definition for the [XMLDocument] class.
type IXMLDocument interface {
	IXMLNode
	// properties:
	CharacterEncoding() string /* primitive/slice/pointer */
	SetCharacterEncoding(value string /* primitive/slice/pointer */)
	DocumentContentKind() XMLDocumentContentKind
	SetDocumentContentKind(value XMLDocumentContentKind)
	DTD() IXMLDTD
	SetDTD(value IXMLDTD)
	Standalone() bool /* primitive/slice/pointer */
	SetStandalone(value bool /* primitive/slice/pointer */)
	MIMEType() string /* primitive/slice/pointer */
	SetMIMEType(value string /* primitive/slice/pointer */)
	Version() string /* primitive/slice/pointer */
	SetVersion(value string /* primitive/slice/pointer */)
	XMLData() IData
	IsStandalone() bool /* primitive/slice/pointer */
	SetIsStandalone(value bool /* primitive/slice/pointer */)
	// methods:
	AddChild(child IXMLNode)
	InsertChildAtIndex(child IXMLNode, index uint /* primitive/slice/pointer */)
	InsertChildrenAtIndex(children []XMLNode /* primitive/slice/pointer */, index uint /* primitive/slice/pointer */)
	ObjectByApplyingXSLTArgumentsError(xslt IData, arguments IDictionary /* already interface */, error_ IError) objc.ID
	ObjectByApplyingXSLTStringArgumentsError(xslt string /* primitive/slice/pointer */, arguments IDictionary /* already interface */, error_ IError) objc.ID
	ObjectByApplyingXSLTAtURLArgumentsError(xsltURL IURL, argument IDictionary /* already interface */, error_ IError) objc.ID
	RemoveChildAtIndex(index uint /* primitive/slice/pointer */)
	ReplaceChildAtIndexWithNode(index uint /* primitive/slice/pointer */, node IXMLNode)
	RootElement() IXMLElement
	SetChildren(children []XMLNode /* primitive/slice/pointer */)
	SetRootElement(root IXMLElement)
	ValidateAndReturnError(error_ IError) bool /* primitive/slice/pointer */
	XMLDataWithOptions(options XMLNodeOptions) IData
}

// An XML document as internalized into a logical tree structure.
//
// An object can have multiple child nodes but only one element, the root element. Any other node must be a object representing a comment or a processing instruction. If you attempt to add any other kind of child node to an object, such as an attribute, namespace, another document object, or an element other than the root, raises an exception. If you add a valid child node and that object already has a parent, raises an exception. An object may also have document-global attributes, such as XML version, character encoding, referenced DTD, and MIME type. The initializers of the class read an external source of XML, whether it be a local file or remote website, parse it, and process it into the tree representation. You can also construct an programmatically. There are accessor methods for getting and setting document attributes, methods for transforming documents using XSLT, a method for dynamically validating a document, and methods for printing out the content of an as XML, XHTML, HTML, or plain text. The class is thread-safe as long as any given instance is used only in one thread.


// An XML document as internalized into a logical tree structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument
type XMLDocument struct {
	XMLNode
}

// XMLDocumentFrom constructs a [XMLDocument] from an unsafe.Pointer.
//
// An XML document as internalized into a logical tree structure.
func XMLDocumentFrom(ptr unsafe.Pointer) XMLDocument {
	return XMLDocument{
		XMLNode: XMLNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (xc _XMLDocumentClass) Alloc() XMLDocument {
	rv := objc.Send[XMLDocument](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XMLDocumentClass) New() XMLDocument {
	rv := objc.Send[XMLDocument](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XMLDocument) Init() XMLDocument {
	rv := objc.Send[XMLDocument](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XMLDocument) Autorelease() XMLDocument {
	rv := objc.Send[XMLDocument](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLDocument creates a new XMLDocument instance.
func NewXMLDocument() XMLDocument {
	return getXMLDocumentClass().New()
}



// Initializes and returns an NSXMLDocument object created from the XML or HTML contents of a URL-referenced source
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/init(contentsOf:options:)
func NewXMLDocumentWithContentsOfURLOptionsError(url IURL, mask XMLNodeOptions, error_ IError) XMLDocument {
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[XMLDocument](instance.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, mask, error_)
	rv.Autorelease()
	return rv
}


// Initializes and returns an object created from an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/init(data:options:)
func NewXMLDocumentWithDataOptionsError(data IData, mask XMLNodeOptions, error_ IError) XMLDocument {
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[XMLDocument](instance.ID, objc.Sel("initWithData:options:error:"), data, mask, error_)
	rv.Autorelease()
	return rv
}


// Returns an object initialized with a single child, the root element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/init(rootElement:)
func NewXMLDocumentWithRootElement(element IXMLElement) XMLDocument {
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[XMLDocument](instance.ID, objc.Sel("initWithRootElement:"), element)
	rv.Autorelease()
	return rv
}


// Initializes and returns an object created from a string containing XML markup text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/init(xmlString:options:)
func NewXMLDocumentWithXMLStringOptionsError(string_ string /* primitive/slice/pointer */, mask XMLNodeOptions, error_ IError) XMLDocument {
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[XMLDocument](instance.ID, objc.Sel("initWithXMLString:options:error:"), objc.String(string_), mask, error_)
	rv.Autorelease()
	return rv
}



// Overridden by subclasses to substitute a custom class for an NSXML class that the parser uses to create node instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/replacementClass(for:)
func (xc _XMLDocumentClass) ReplacementClassForClass(cls objc.Class) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(xc.class), objc.Sel("replacementClassForClass:"), cls)
	return rv
}


// Adds a child node after the last of the receiver’s existing children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/addChild(_:)
func (x_ XMLDocument) AddChild(child IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("addChild:"), child)
}


// Inserts a node object at specified position in the receiver’s array of children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/insertChild(_:at:)
func (x_ XMLDocument) InsertChildAtIndex(child IXMLNode, index uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("insertChild:atIndex:"), child, index)
}


// Inserts an array of children at a specified position in the receiver’s array of children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/insertChildren(_:at:)
func (x_ XMLDocument) InsertChildrenAtIndex(children []XMLNode /* primitive/slice/pointer */, index uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("insertChildren:atIndex:"), children, index)
}


// Applies the XSLT pattern rules and templates (specified as a data object) to the receiver and returns a document object containing transformed XML or HTML markup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/object(byApplyingXSLT:arguments:)
func (x_ XMLDocument) ObjectByApplyingXSLTArgumentsError(xslt IData, arguments IDictionary /* already interface */, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("objectByApplyingXSLT:arguments:error:"), xslt, arguments, error_)
	return rv
}


// Applies the XSLT pattern rules and templates (specified as a string) to the receiver and returns a document object containing transformed XML or HTML markup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/object(byApplyingXSLTString:arguments:)
func (x_ XMLDocument) ObjectByApplyingXSLTStringArgumentsError(xslt string /* primitive/slice/pointer */, arguments IDictionary /* already interface */, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("objectByApplyingXSLTString:arguments:error:"), objc.String(xslt), arguments, error_)
	return rv
}


// Applies the XSLT pattern rules and templates located at a specified URL to the receiver and returns a document object containing transformed XML markup or an object containing plain text, RTF text, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/objectByApplyingXSLT(at:arguments:)
func (x_ XMLDocument) ObjectByApplyingXSLTAtURLArgumentsError(xsltURL IURL, argument IDictionary /* already interface */, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("objectByApplyingXSLTAtURL:arguments:error:"), xsltURL, argument, error_)
	return rv
}


// Removes the child node of the receiver located at a specified position in its array of children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/removeChild(at:)
func (x_ XMLDocument) RemoveChildAtIndex(index uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("removeChildAtIndex:"), index)
}


// Replaces the child node of the receiver located at a specified position in its array of children with another node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/replaceChild(at:with:)
func (x_ XMLDocument) ReplaceChildAtIndexWithNode(index uint /* primitive/slice/pointer */, node IXMLNode) {
	objc.Send[objc.ID](x_.ID, objc.Sel("replaceChildAtIndex:withNode:"), index, node)
}


// Returns the root element of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/rootElement()
func (x_ XMLDocument) RootElement() IXMLElement {
	rv := objc.Send[XMLElement](x_.ID, objc.Sel("rootElement"))
	return rv
}


// Sets the child nodes of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/setChildren(_:)
func (x_ XMLDocument) SetChildren(children []XMLNode /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setChildren:"), children)
}


// Set the root element of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/setRootElement(_:)
func (x_ XMLDocument) SetRootElement(root IXMLElement) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setRootElement:"), root)
}


// Validates the document against the governing schema and returns whether the document conforms to the schema.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/validate()
func (x_ XMLDocument) ValidateAndReturnError(error_ IError) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](x_.ID, objc.Sel("validateAndReturnError:"), error_)
	return rv
}


// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/xmlData(options:)
func (x_ XMLDocument) XMLDataWithOptions(options XMLNodeOptions) IData {
	rv := objc.Send[Data](x_.ID, objc.Sel("XMLDataWithOptions:"), options)
	return rv
}


// Sets the character encoding of the receiver to ,
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/characterEncoding
func (x_ XMLDocument) CharacterEncoding() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](x_.ID, objc.Sel("characterEncoding"))
	return rv
}


// Sets the character encoding of the receiver to ,
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/characterEncoding
func (x_ XMLDocument) SetCharacterEncoding(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setCharacterEncoding:"), objc.String(value))
}


// Sets the kind of output content for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/documentContentKind
func (x_ XMLDocument) DocumentContentKind() XMLDocumentContentKind {
	rv := objc.Send[XMLDocumentContentKind](x_.ID, objc.Sel("documentContentKind"))
	return rv
}


// Sets the kind of output content for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/documentContentKind
func (x_ XMLDocument) SetDocumentContentKind(value XMLDocumentContentKind) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDocumentContentKind:"), value)
}


// Returns an object representing the internal DTD associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/dtd
func (x_ XMLDocument) DTD() IXMLDTD {
	rv := objc.Send[XMLDTD](x_.ID, objc.Sel("DTD"))
	return rv
}


// Returns an object representing the internal DTD associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/dtd
func (x_ XMLDocument) SetDTD(value IXMLDTD) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDTD:"), value)
}


// Sets a Boolean value that specifies whether the receiver represents a standalone XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/isStandalone
func (x_ XMLDocument) Standalone() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](x_.ID, objc.Sel("standalone"))
	return rv
}


// Sets a Boolean value that specifies whether the receiver represents a standalone XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/isStandalone
func (x_ XMLDocument) SetStandalone(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setStandalone:"), value)
}


// Returns the MIME type for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/mimeType
func (x_ XMLDocument) MIMEType() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](x_.ID, objc.Sel("MIMEType"))
	return rv
}


// Returns the MIME type for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/mimeType
func (x_ XMLDocument) SetMIMEType(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setMIMEType:"), objc.String(value))
}


// Sets the version of the receiver’s XML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/version
func (x_ XMLDocument) Version() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](x_.ID, objc.Sel("version"))
	return rv
}


// Sets the version of the receiver’s XML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/version
func (x_ XMLDocument) SetVersion(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setVersion:"), objc.String(value))
}


// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/xmlData
func (x_ XMLDocument) XMLData() IData {
	rv := objc.Send[Data](x_.ID, objc.Sel("XMLData"))
	return rv
}


// Sets a Boolean value that specifies whether the receiver represents a standalone XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/isstandalone
func (x_ XMLDocument) IsStandalone() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](x_.ID, objc.Sel("isStandalone"))
	return rv
}


// Sets a Boolean value that specifies whether the receiver represents a standalone XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/isstandalone
func (x_ XMLDocument) SetIsStandalone(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setIsStandalone:"), value)
}


