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
	AddChild(child unsafe.Pointer)
	InsertChildrenAtIndex(children unsafe.Pointer, index uint)
	ObjectByApplyingXSLTArgumentsError(xslt unsafe.Pointer, arguments unsafe.Pointer, error unsafe.Pointer) objc.ID
	ObjectByApplyingXSLTStringArgumentsError(xslt string, arguments unsafe.Pointer, error unsafe.Pointer) objc.ID
	RemoveChildAtIndex(index uint)
	RootElement() unsafe.Pointer
	SetRootElement(root unsafe.Pointer)
}

// An XML document as internalized into a logical tree structure.
//
// An object can have multiple child nodes but only one element, the root element. Any other node must be a object representing a comment or a processing instruction. If you attempt to add any other kind of child node to an object, such as an attribute, namespace, another document object, or an element other than the root, raises an exception. If you add a valid child node and that object already has a parent, raises an exception. An object may also have document-global attributes, such as XML version, character encoding, referenced DTD, and MIME type. The initializers of the class read an external source of XML, whether it be a local file or remote website, parse it, and process it into the tree representation. You can also construct an programmatically. There are accessor methods for getting and setting document attributes, methods for transforming documents using XSLT, a method for dynamically validating a document, and methods for printing out the content of an as XML, XHTML, HTML, or plain text. The class is thread-safe as long as any given instance is used only in one thread.
//
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


// Initializes and returns an object created from an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/init(data:options:)
func NewXMLDocumentWithDataOptionsError(data unsafe.Pointer, mask unsafe.Pointer, error unsafe.Pointer) XMLDocument {
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[XMLDocument](instance.ID, objc.Sel("initWithData:options:error:"), data, mask, error)
	rv.Autorelease()
	return rv
}


// Adds a child node after the last of the receiver’s existing children.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/addChild(_:)
func (x_ XMLDocument) AddChild(child unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("addChild:"), child)
}

// Inserts an array of children at a specified position in the receiver’s array of children.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/insertChildren(_:at:)
func (x_ XMLDocument) InsertChildrenAtIndex(children unsafe.Pointer, index uint) {
	objc.Send[objc.ID](x_.ID, objc.Sel("insertChildren:atIndex:"), children, index)
}

// Applies the XSLT pattern rules and templates (specified as a data object) to the receiver and returns a document object containing transformed XML or HTML markup.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/object(byApplyingXSLT:arguments:)
func (x_ XMLDocument) ObjectByApplyingXSLTArgumentsError(xslt unsafe.Pointer, arguments unsafe.Pointer, error unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("objectByApplyingXSLT:arguments:error:"), xslt, arguments, error)
	return rv
}

// Applies the XSLT pattern rules and templates (specified as a string) to the receiver and returns a document object containing transformed XML or HTML markup.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/object(byApplyingXSLTString:arguments:)
func (x_ XMLDocument) ObjectByApplyingXSLTStringArgumentsError(xslt string, arguments unsafe.Pointer, error unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("objectByApplyingXSLTString:arguments:error:"), objc.String(xslt), arguments, error)
	return rv
}

// Removes the child node of the receiver located at a specified position in its array of children.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/removeChild(at:)
func (x_ XMLDocument) RemoveChildAtIndex(index uint) {
	objc.Send[objc.ID](x_.ID, objc.Sel("removeChildAtIndex:"), index)
}

// Returns the root element of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/rootElement()
func (x_ XMLDocument) RootElement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("rootElement"))
	return rv
}

// Set the root element of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/setRootElement(_:)
func (x_ XMLDocument) SetRootElement(root unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setRootElement:"), root)
}

// Returns the MIME type for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/mimeType
func (x_ XMLDocument) MIMEType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("MIMEType"))
	return rv
}


// SetMIMEType sets the value of the MIMEType property.
// Returns the MIME type for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/mimeType
func (x_ XMLDocument) SetMIMEType(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setMIMEType:"), value)
}
// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/XMLDocument/xmlData
func (x_ XMLDocument) XMLData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("XMLData"))
	return rv
}


