// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMDocument */


/* debug [class_header]: Header for DOMDocument */
// The class instance for the [DOMDocument] class.
var (
	DOMDocumentClass     _DOMDocumentClass
	DOMDocumentClassOnce sync.Once
)

func getDOMDocumentClass() _DOMDocumentClass {
	DOMDocumentClassOnce.Do(func() {
		DOMDocumentClass = _DOMDocumentClass{objc.GetClass("DOMDocument")}
	})
	return DOMDocumentClass
}

type _DOMDocumentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMDocument */
// An interface definition for the [DOMDocument] class.
type IDOMDocument interface {
	IDOMNode
	
/* debug [class_interface_properties]: Properties for DOMDocument */
	// properties:
	ActiveElement() IDOMElement
	Anchors() IDOMHTMLCollection
	Applets() IDOMHTMLCollection
	Body() IDOMHTMLElement
	SetBody(value IDOMHTMLElement)
	CharacterSet() objc.IObject /* cross-framework: NSString */
	Charset() objc.IObject /* cross-framework: NSString */
	SetCharset(value objc.IObject /* cross-framework: NSString */)
	Cookie() objc.IObject /* cross-framework: NSString */
	SetCookie(value objc.IObject /* cross-framework: NSString */)
	DefaultCharset() objc.IObject /* cross-framework: NSString */
	DefaultView() IDOMAbstractView
	Doctype() IDOMDocumentType
	DocumentElement() IDOMElement
	DocumentURI() objc.IObject /* cross-framework: NSString */
	SetDocumentURI(value objc.IObject /* cross-framework: NSString */)
	Domain() objc.IObject /* cross-framework: NSString */
	Forms() IDOMHTMLCollection
	Images() IDOMHTMLCollection
	Implementation() IDOMImplementation
	InputEncoding() objc.IObject /* cross-framework: NSString */
	LastModified() objc.IObject /* cross-framework: NSString */
	Links() IDOMHTMLCollection
	PreferredStylesheetSet() objc.IObject /* cross-framework: NSString */
	ReadyState() objc.IObject /* cross-framework: NSString */
	Referrer() objc.IObject /* cross-framework: NSString */
	SelectedStylesheetSet() objc.IObject /* cross-framework: NSString */
	SetSelectedStylesheetSet(value objc.IObject /* cross-framework: NSString */)
	StyleSheets() IDOMStyleSheetList
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSString */
	WebFrame() IWebFrame
	XmlEncoding() objc.IObject /* cross-framework: NSString */
	XmlStandalone() bool
	SetXmlStandalone(value bool)
	XmlVersion() objc.IObject /* cross-framework: NSString */
	SetXmlVersion(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMDocument */
	// methods:
	URLWithAttributeString(string_ objc.IObject /* cross-framework: NSString */) foundation.URL
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMDocument */
// Alloc allocates a new instance without initialization.
func (dc _DOMDocumentClass) Alloc() DOMDocument {
	rv := objc.Send[DOMDocument](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMDocumentClass) New() DOMDocument {
	rv := objc.Send[DOMDocument](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMDocument) Init() DOMDocument {
	rv := objc.Send[DOMDocument](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMDocument) Autorelease() DOMDocument {
	rv := objc.Send[DOMDocument](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMDocument creates a new DOMDocument instance.
func NewDOMDocument() DOMDocument {
	return getDOMDocumentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMDocument */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument
type DOMDocument struct {
	DOMNode
}

// DOMDocumentFrom constructs a [DOMDocument] from an unsafe.Pointer.
func DOMDocumentFrom(ptr unsafe.Pointer) DOMDocument {
	return DOMDocument{
		DOMNode: DOMNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMDocument *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMDocument */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMDocument */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMDocument */

// Constructs a URL given an attribute string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/url(withAttributeString:)
func (d_ DOMDocument) URLWithAttributeString(string_ objc.IObject /* cross-framework: NSString */) foundation.URL {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("URLWithAttributeString:"), string_)
	return rv
}/* debug [instance_methods/method]: URLWithAttributeString */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMDocument */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/activeElement
func (d_ DOMDocument) ActiveElement() IDOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("activeElement"))
	return rv
}/* debug [instance_properties/getter]: activeElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/anchors
func (d_ DOMDocument) Anchors() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("anchors"))
	return rv
}/* debug [instance_properties/getter]: anchors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/applets
func (d_ DOMDocument) Applets() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("applets"))
	return rv
}/* debug [instance_properties/getter]: applets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/body
func (d_ DOMDocument) Body() IDOMHTMLElement {
	rv := objc.Send[DOMHTMLElement](d_.ID, objc.Sel("body"))
	return rv
}/* debug [instance_properties/getter]: body */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/body
func (d_ DOMDocument) SetBody(value IDOMHTMLElement) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBody:"), value)
}/* debug [instance_properties/setter]: body */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/characterSet
func (d_ DOMDocument) CharacterSet() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("characterSet"))
	return rv
}/* debug [instance_properties/getter]: characterSet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/charset
func (d_ DOMDocument) Charset() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("charset"))
	return rv
}/* debug [instance_properties/getter]: charset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/charset
func (d_ DOMDocument) SetCharset(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCharset:"), value)
}/* debug [instance_properties/setter]: charset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/cookie
func (d_ DOMDocument) Cookie() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("cookie"))
	return rv
}/* debug [instance_properties/getter]: cookie */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/cookie
func (d_ DOMDocument) SetCookie(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCookie:"), value)
}/* debug [instance_properties/setter]: cookie */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/defaultCharset
func (d_ DOMDocument) DefaultCharset() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("defaultCharset"))
	return rv
}/* debug [instance_properties/getter]: defaultCharset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/defaultView
func (d_ DOMDocument) DefaultView() IDOMAbstractView {
	rv := objc.Send[DOMAbstractView](d_.ID, objc.Sel("defaultView"))
	return rv
}/* debug [instance_properties/getter]: defaultView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/doctype
func (d_ DOMDocument) Doctype() IDOMDocumentType {
	rv := objc.Send[DOMDocumentType](d_.ID, objc.Sel("doctype"))
	return rv
}/* debug [instance_properties/getter]: doctype */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/documentElement
func (d_ DOMDocument) DocumentElement() IDOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("documentElement"))
	return rv
}/* debug [instance_properties/getter]: documentElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/documentURI
func (d_ DOMDocument) DocumentURI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("documentURI"))
	return rv
}/* debug [instance_properties/getter]: documentURI */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/documentURI
func (d_ DOMDocument) SetDocumentURI(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDocumentURI:"), value)
}/* debug [instance_properties/setter]: documentURI */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/domain
func (d_ DOMDocument) Domain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("domain"))
	return rv
}/* debug [instance_properties/getter]: domain */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/forms
func (d_ DOMDocument) Forms() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("forms"))
	return rv
}/* debug [instance_properties/getter]: forms */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/images
func (d_ DOMDocument) Images() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("images"))
	return rv
}/* debug [instance_properties/getter]: images */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/implementation
func (d_ DOMDocument) Implementation() IDOMImplementation {
	rv := objc.Send[DOMImplementation](d_.ID, objc.Sel("implementation"))
	return rv
}/* debug [instance_properties/getter]: implementation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/inputEncoding
func (d_ DOMDocument) InputEncoding() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("inputEncoding"))
	return rv
}/* debug [instance_properties/getter]: inputEncoding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/lastModified
func (d_ DOMDocument) LastModified() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("lastModified"))
	return rv
}/* debug [instance_properties/getter]: lastModified */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/links
func (d_ DOMDocument) Links() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("links"))
	return rv
}/* debug [instance_properties/getter]: links */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/preferredStylesheetSet
func (d_ DOMDocument) PreferredStylesheetSet() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("preferredStylesheetSet"))
	return rv
}/* debug [instance_properties/getter]: preferredStylesheetSet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/readyState
func (d_ DOMDocument) ReadyState() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("readyState"))
	return rv
}/* debug [instance_properties/getter]: readyState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/referrer
func (d_ DOMDocument) Referrer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("referrer"))
	return rv
}/* debug [instance_properties/getter]: referrer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/selectedStylesheetSet
func (d_ DOMDocument) SelectedStylesheetSet() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("selectedStylesheetSet"))
	return rv
}/* debug [instance_properties/getter]: selectedStylesheetSet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/selectedStylesheetSet
func (d_ DOMDocument) SetSelectedStylesheetSet(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelectedStylesheetSet:"), value)
}/* debug [instance_properties/setter]: selectedStylesheetSet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/styleSheets
func (d_ DOMDocument) StyleSheets() IDOMStyleSheetList {
	rv := objc.Send[DOMStyleSheetList](d_.ID, objc.Sel("styleSheets"))
	return rv
}/* debug [instance_properties/getter]: styleSheets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/title
func (d_ DOMDocument) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/title
func (d_ DOMDocument) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/url
func (d_ DOMDocument) URL() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// The web frame associated with the DOM document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/webFrame
func (d_ DOMDocument) WebFrame() IWebFrame {
	rv := objc.Send[WebFrame](d_.ID, objc.Sel("webFrame"))
	return rv
}/* debug [instance_properties/getter]: webFrame */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/xmlEncoding
func (d_ DOMDocument) XmlEncoding() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("xmlEncoding"))
	return rv
}/* debug [instance_properties/getter]: xmlEncoding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/xmlStandalone
func (d_ DOMDocument) XmlStandalone() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("xmlStandalone"))
	return rv
}/* debug [instance_properties/getter]: xmlStandalone */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/xmlStandalone
func (d_ DOMDocument) SetXmlStandalone(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setXmlStandalone:"), value)
}/* debug [instance_properties/setter]: xmlStandalone */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/xmlVersion
func (d_ DOMDocument) XmlVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("xmlVersion"))
	return rv
}/* debug [instance_properties/getter]: xmlVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocument/xmlVersion
func (d_ DOMDocument) SetXmlVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setXmlVersion:"), value)
}/* debug [instance_properties/setter]: xmlVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMDocument */



