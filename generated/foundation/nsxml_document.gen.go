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
	DocumentContentKind() unsafe.Pointer
	SetDocumentContentKind(value unsafe.Pointer)
	Dtd() XMLDTD /* not a class type */
	SetDtd(value XMLDTD /* not a class type */)
	IsStandalone() bool /* primitive/slice/pointer */
	SetIsStandalone(value bool /* primitive/slice/pointer */)
	MimeType() string /* primitive/slice/pointer */
	SetMimeType(value string /* primitive/slice/pointer */)
	Version() string /* primitive/slice/pointer */
	SetVersion(value string /* primitive/slice/pointer */)
	XmlData() IData
	SetXmlData(value IData)
	// methods:
	SetChildren(children []XMLNode /* primitive/slice/pointer */)
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



// Sets the child nodes of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/setChildren(_:)
func (x_ XMLDocument) SetChildren(children []XMLNode /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setChildren:"), children)
}


// Sets the character encoding of the receiver to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/characterencoding
func (x_ XMLDocument) CharacterEncoding() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](x_.ID, objc.Sel("characterEncoding"))
	return rv
}


// Sets the character encoding of the receiver to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/characterencoding
func (x_ XMLDocument) SetCharacterEncoding(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setCharacterEncoding:"), objc.String(value))
}


// Sets the kind of output content for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/documentcontentkind
func (x_ XMLDocument) DocumentContentKind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("documentContentKind"))
	return rv
}


// Sets the kind of output content for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/documentcontentkind
func (x_ XMLDocument) SetDocumentContentKind(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDocumentContentKind:"), value)
}


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/dtd
func (x_ XMLDocument) Dtd() XMLDTD /* not a class type */ {
	rv := objc.Send[XMLDTD](x_.ID, objc.Sel("dtd"))
	return rv
}


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/dtd
func (x_ XMLDocument) SetDtd(value XMLDTD /* not a class type */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDtd:"), value)
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


// Returns the MIME type for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/mimetype
func (x_ XMLDocument) MimeType() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](x_.ID, objc.Sel("mimeType"))
	return rv
}


// Returns the MIME type for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/mimetype
func (x_ XMLDocument) SetMimeType(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setMimeType:"), objc.String(value))
}


// Sets the version of the receiver’s XML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/version
func (x_ XMLDocument) Version() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](x_.ID, objc.Sel("version"))
	return rv
}


// Sets the version of the receiver’s XML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/version
func (x_ XMLDocument) SetVersion(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setVersion:"), objc.String(value))
}


// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/xmldata
func (x_ XMLDocument) XmlData() IData {
	rv := objc.Send[Data](x_.ID, objc.Sel("xmlData"))
	return rv
}


// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/xmldata
func (x_ XMLDocument) SetXmlData(value IData) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setXmlData:"), value)
}



