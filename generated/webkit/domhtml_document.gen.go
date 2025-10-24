// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLDocument */


/* debug [class_header]: Header for DOMHTMLDocument */
// The class instance for the [DOMHTMLDocument] class.
var (
	DOMHTMLDocumentClass     _DOMHTMLDocumentClass
	DOMHTMLDocumentClassOnce sync.Once
)

func getDOMHTMLDocumentClass() _DOMHTMLDocumentClass {
	DOMHTMLDocumentClassOnce.Do(func() {
		DOMHTMLDocumentClass = _DOMHTMLDocumentClass{objc.GetClass("DOMHTMLDocument")}
	})
	return DOMHTMLDocumentClass
}

type _DOMHTMLDocumentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLDocument */
// An interface definition for the [DOMHTMLDocument] class.
type IDOMHTMLDocument interface {
	IDOMDocument
	
/* debug [class_interface_properties]: Properties for DOMHTMLDocument */
	// properties:
	AlinkColor() objc.IObject /* cross-framework: NSString */
	SetAlinkColor(value objc.IObject /* cross-framework: NSString */)
	BgColor() objc.IObject /* cross-framework: NSString */
	SetBgColor(value objc.IObject /* cross-framework: NSString */)
	CompatMode() objc.IObject /* cross-framework: NSString */
	DesignMode() objc.IObject /* cross-framework: NSString */
	SetDesignMode(value objc.IObject /* cross-framework: NSString */)
	Dir() objc.IObject /* cross-framework: NSString */
	SetDir(value objc.IObject /* cross-framework: NSString */)
	Embeds() IDOMHTMLCollection
	FgColor() objc.IObject /* cross-framework: NSString */
	SetFgColor(value objc.IObject /* cross-framework: NSString */)
	Height() int
	LinkColor() objc.IObject /* cross-framework: NSString */
	SetLinkColor(value objc.IObject /* cross-framework: NSString */)
	Plugins() IDOMHTMLCollection
	Scripts() IDOMHTMLCollection
	VlinkColor() objc.IObject /* cross-framework: NSString */
	SetVlinkColor(value objc.IObject /* cross-framework: NSString */)
	Width() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLDocument */
	// methods:
	CreateDocumentFragmentWithMarkupStringBaseURL(markupString objc.IObject /* cross-framework: NSString */, baseURL objc.IObject /* cross-framework: NSURL */) IDOMDocumentFragment
	CreateDocumentFragmentWithText(text objc.IObject /* cross-framework: NSString */) IDOMDocumentFragment
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLDocument */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLDocumentClass) Alloc() DOMHTMLDocument {
	rv := objc.Send[DOMHTMLDocument](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLDocumentClass) New() DOMHTMLDocument {
	rv := objc.Send[DOMHTMLDocument](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLDocument) Init() DOMHTMLDocument {
	rv := objc.Send[DOMHTMLDocument](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLDocument) Autorelease() DOMHTMLDocument {
	rv := objc.Send[DOMHTMLDocument](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLDocument creates a new DOMHTMLDocument instance.
func NewDOMHTMLDocument() DOMHTMLDocument {
	return getDOMHTMLDocumentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLDocument */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument
type DOMHTMLDocument struct {
	DOMDocument
}

// DOMHTMLDocumentFrom constructs a [DOMHTMLDocument] from an unsafe.Pointer.
func DOMHTMLDocumentFrom(ptr unsafe.Pointer) DOMHTMLDocument {
	return DOMHTMLDocument{
		DOMDocument: DOMDocumentFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLDocument *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLDocument */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLDocument */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLDocument */

// Creates a document fragment containing the given HTML markup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/createDocumentFragment(withMarkupString:baseURL:)
func (d_ DOMHTMLDocument) CreateDocumentFragmentWithMarkupStringBaseURL(markupString objc.IObject /* cross-framework: NSString */, baseURL objc.IObject /* cross-framework: NSURL */) IDOMDocumentFragment {
	rv := objc.Send[DOMDocumentFragment](d_.ID, objc.Sel("createDocumentFragmentWithMarkupString:baseURL:"), markupString, baseURL)
	return rv
}/* debug [instance_methods/method]: CreateDocumentFragmentWithMarkupStringBaseURL */


// Creates a document fragment containing the given text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/createDocumentFragment(withText:)
func (d_ DOMHTMLDocument) CreateDocumentFragmentWithText(text objc.IObject /* cross-framework: NSString */) IDOMDocumentFragment {
	rv := objc.Send[DOMDocumentFragment](d_.ID, objc.Sel("createDocumentFragmentWithText:"), text)
	return rv
}/* debug [instance_methods/method]: CreateDocumentFragmentWithText */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLDocument */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/alinkColor
func (d_ DOMHTMLDocument) AlinkColor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("alinkColor"))
	return rv
}/* debug [instance_properties/getter]: alinkColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/alinkColor
func (d_ DOMHTMLDocument) SetAlinkColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlinkColor:"), value)
}/* debug [instance_properties/setter]: alinkColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/bgColor
func (d_ DOMHTMLDocument) BgColor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("bgColor"))
	return rv
}/* debug [instance_properties/getter]: bgColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/bgColor
func (d_ DOMHTMLDocument) SetBgColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBgColor:"), value)
}/* debug [instance_properties/setter]: bgColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/compatMode
func (d_ DOMHTMLDocument) CompatMode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("compatMode"))
	return rv
}/* debug [instance_properties/getter]: compatMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/designMode
func (d_ DOMHTMLDocument) DesignMode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("designMode"))
	return rv
}/* debug [instance_properties/getter]: designMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/designMode
func (d_ DOMHTMLDocument) SetDesignMode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDesignMode:"), value)
}/* debug [instance_properties/setter]: designMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/dir
func (d_ DOMHTMLDocument) Dir() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("dir"))
	return rv
}/* debug [instance_properties/getter]: dir */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/dir
func (d_ DOMHTMLDocument) SetDir(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDir:"), value)
}/* debug [instance_properties/setter]: dir */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/embeds
func (d_ DOMHTMLDocument) Embeds() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("embeds"))
	return rv
}/* debug [instance_properties/getter]: embeds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/fgColor
func (d_ DOMHTMLDocument) FgColor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("fgColor"))
	return rv
}/* debug [instance_properties/getter]: fgColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/fgColor
func (d_ DOMHTMLDocument) SetFgColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFgColor:"), value)
}/* debug [instance_properties/setter]: fgColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/height
func (d_ DOMHTMLDocument) Height() int {
	rv := objc.Send[int](d_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/linkColor
func (d_ DOMHTMLDocument) LinkColor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("linkColor"))
	return rv
}/* debug [instance_properties/getter]: linkColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/linkColor
func (d_ DOMHTMLDocument) SetLinkColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLinkColor:"), value)
}/* debug [instance_properties/setter]: linkColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/plugins
func (d_ DOMHTMLDocument) Plugins() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("plugins"))
	return rv
}/* debug [instance_properties/getter]: plugins */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/scripts
func (d_ DOMHTMLDocument) Scripts() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("scripts"))
	return rv
}/* debug [instance_properties/getter]: scripts */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/vlinkColor
func (d_ DOMHTMLDocument) VlinkColor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("vlinkColor"))
	return rv
}/* debug [instance_properties/getter]: vlinkColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/vlinkColor
func (d_ DOMHTMLDocument) SetVlinkColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVlinkColor:"), value)
}/* debug [instance_properties/setter]: vlinkColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDocument/width
func (d_ DOMHTMLDocument) Width() int {
	rv := objc.Send[int](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLDocument */



