// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSXMLDocument */


/* debug [class_header]: Header for NSXMLDocument */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for XMLDocument */
// An interface definition for the [XMLDocument] class.
type IXMLDocument interface {
	IXMLNode
	
/* debug [class_interface_properties]: Properties for XMLDocument */
	// properties:
	DocumentContentKind() XMLDocumentContentKind /* not a class type */
	SetDocumentContentKind(value XMLDocumentContentKind /* not a class type */)
	CharacterEncoding() IString
	SetCharacterEncoding(value IString)
	Dtd() IXMLDTD
	SetDtd(value IXMLDTD)
	IsStandalone() bool
	SetIsStandalone(value bool)
	MimeType() IString
	SetMimeType(value IString)
	Version() IString
	SetVersion(value IString)
	XmlData() IData
	SetXmlData(value IData)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for XMLDocument */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for XMLDocument */
// Alloc allocates a new instance without initialization.
func (xc _XMLDocumentClass) Alloc() XMLDocument {
	rv := objc.Send[XMLDocument](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for XMLDocument */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for XMLDocument */

// Initializes and returns an object created from an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/init(data:options:)
func NewXMLDocumentWithDataOptionsError(data IData, mask XMLNodeOptions, error_ IError) XMLDocument {
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[XMLDocument](instance.ID, objc.Sel("initWithData:options:error:"), data, mask, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewXMLDocumentWithDataOptionsError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for XMLDocument */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for XMLDocument */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for XMLDocument */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for XMLDocument */

// Sets the kind of output content for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/documentContentKind
func (x_ XMLDocument) DocumentContentKind() XMLDocumentContentKind /* not a class type */ {
	rv := objc.Send[XMLDocumentContentKind](x_.ID, objc.Sel("documentContentKind"))
	return rv
}/* debug [instance_properties/getter]: documentContentKind */


// Sets the kind of output content for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/documentContentKind
func (x_ XMLDocument) SetDocumentContentKind(value XMLDocumentContentKind /* not a class type */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDocumentContentKind:"), value)
}/* debug [instance_properties/setter]: documentContentKind */


// Sets the character encoding of the receiver to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/characterencoding
func (x_ XMLDocument) CharacterEncoding() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("characterEncoding"))
	return rv
}/* debug [instance_properties/getter]: characterEncoding */


// Sets the character encoding of the receiver to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/characterencoding
func (x_ XMLDocument) SetCharacterEncoding(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setCharacterEncoding:"), value)
}/* debug [instance_properties/setter]: characterEncoding */


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/dtd
func (x_ XMLDocument) Dtd() IXMLDTD {
	rv := objc.Send[XMLDTD](x_.ID, objc.Sel("dtd"))
	return rv
}/* debug [instance_properties/getter]: dtd */


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/dtd
func (x_ XMLDocument) SetDtd(value IXMLDTD) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDtd:"), value)
}/* debug [instance_properties/setter]: dtd */


// Sets a Boolean value that specifies whether the receiver represents a standalone XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/isstandalone
func (x_ XMLDocument) IsStandalone() bool {
	rv := objc.Send[bool](x_.ID, objc.Sel("isStandalone"))
	return rv
}/* debug [instance_properties/getter]: isStandalone */


// Sets a Boolean value that specifies whether the receiver represents a standalone XML document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/isstandalone
func (x_ XMLDocument) SetIsStandalone(value bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setIsStandalone:"), value)
}/* debug [instance_properties/setter]: isStandalone */


// Returns the MIME type for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/mimetype
func (x_ XMLDocument) MimeType() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("mimeType"))
	return rv
}/* debug [instance_properties/getter]: mimeType */


// Returns the MIME type for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/mimetype
func (x_ XMLDocument) SetMimeType(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setMimeType:"), value)
}/* debug [instance_properties/setter]: mimeType */


// Sets the version of the receiver’s XML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/version
func (x_ XMLDocument) Version() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("version"))
	return rv
}/* debug [instance_properties/getter]: version */


// Sets the version of the receiver’s XML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/version
func (x_ XMLDocument) SetVersion(value IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setVersion:"), value)
}/* debug [instance_properties/setter]: version */


// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/xmldata
func (x_ XMLDocument) XmlData() IData {
	rv := objc.Send[Data](x_.ID, objc.Sel("xmlData"))
	return rv
}/* debug [instance_properties/getter]: xmlData */


// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/xmldocument/xmldata
func (x_ XMLDocument) SetXmlData(value IData) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setXmlData:"), value)
}/* debug [instance_properties/setter]: xmlData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSXMLDocument */


