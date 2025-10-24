// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLIFrameElement */


/* debug [class_header]: Header for DOMHTMLIFrameElement */
// The class instance for the [DOMHTMLIFrameElement] class.
var (
	DOMHTMLIFrameElementClass     _DOMHTMLIFrameElementClass
	DOMHTMLIFrameElementClassOnce sync.Once
)

func getDOMHTMLIFrameElementClass() _DOMHTMLIFrameElementClass {
	DOMHTMLIFrameElementClassOnce.Do(func() {
		DOMHTMLIFrameElementClass = _DOMHTMLIFrameElementClass{objc.GetClass("DOMHTMLIFrameElement")}
	})
	return DOMHTMLIFrameElementClass
}

type _DOMHTMLIFrameElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLIFrameElement */
// An interface definition for the [DOMHTMLIFrameElement] class.
type IDOMHTMLIFrameElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLIFrameElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	ContentDocument() IDOMDocument
	ContentFrame() IWebFrame
	ContentWindow() IDOMAbstractView
	FrameBorder() objc.IObject /* cross-framework: NSString */
	SetFrameBorder(value objc.IObject /* cross-framework: NSString */)
	Height() objc.IObject /* cross-framework: NSString */
	SetHeight(value objc.IObject /* cross-framework: NSString */)
	LongDesc() objc.IObject /* cross-framework: NSString */
	SetLongDesc(value objc.IObject /* cross-framework: NSString */)
	MarginHeight() objc.IObject /* cross-framework: NSString */
	SetMarginHeight(value objc.IObject /* cross-framework: NSString */)
	MarginWidth() objc.IObject /* cross-framework: NSString */
	SetMarginWidth(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Scrolling() objc.IObject /* cross-framework: NSString */
	SetScrolling(value objc.IObject /* cross-framework: NSString */)
	Src() objc.IObject /* cross-framework: NSString */
	SetSrc(value objc.IObject /* cross-framework: NSString */)
	Width() objc.IObject /* cross-framework: NSString */
	SetWidth(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLIFrameElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLIFrameElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLIFrameElementClass) Alloc() DOMHTMLIFrameElement {
	rv := objc.Send[DOMHTMLIFrameElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLIFrameElementClass) New() DOMHTMLIFrameElement {
	rv := objc.Send[DOMHTMLIFrameElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLIFrameElement) Init() DOMHTMLIFrameElement {
	rv := objc.Send[DOMHTMLIFrameElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLIFrameElement) Autorelease() DOMHTMLIFrameElement {
	rv := objc.Send[DOMHTMLIFrameElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLIFrameElement creates a new DOMHTMLIFrameElement instance.
func NewDOMHTMLIFrameElement() DOMHTMLIFrameElement {
	return getDOMHTMLIFrameElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLIFrameElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement
type DOMHTMLIFrameElement struct {
	DOMHTMLElement
}

// DOMHTMLIFrameElementFrom constructs a [DOMHTMLIFrameElement] from an unsafe.Pointer.
func DOMHTMLIFrameElementFrom(ptr unsafe.Pointer) DOMHTMLIFrameElement {
	return DOMHTMLIFrameElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLIFrameElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLIFrameElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLIFrameElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLIFrameElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLIFrameElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/align
func (d_ DOMHTMLIFrameElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
}/* debug [instance_properties/getter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/align
func (d_ DOMHTMLIFrameElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
}/* debug [instance_properties/setter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/contentDocument
func (d_ DOMHTMLIFrameElement) ContentDocument() IDOMDocument {
	rv := objc.Send[DOMDocument](d_.ID, objc.Sel("contentDocument"))
	return rv
}/* debug [instance_properties/getter]: contentDocument */


// The content frame of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/contentFrame
func (d_ DOMHTMLIFrameElement) ContentFrame() IWebFrame {
	rv := objc.Send[WebFrame](d_.ID, objc.Sel("contentFrame"))
	return rv
}/* debug [instance_properties/getter]: contentFrame */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/contentWindow
func (d_ DOMHTMLIFrameElement) ContentWindow() IDOMAbstractView {
	rv := objc.Send[DOMAbstractView](d_.ID, objc.Sel("contentWindow"))
	return rv
}/* debug [instance_properties/getter]: contentWindow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/frameBorder
func (d_ DOMHTMLIFrameElement) FrameBorder() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("frameBorder"))
	return rv
}/* debug [instance_properties/getter]: frameBorder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/frameBorder
func (d_ DOMHTMLIFrameElement) SetFrameBorder(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFrameBorder:"), value)
}/* debug [instance_properties/setter]: frameBorder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/height
func (d_ DOMHTMLIFrameElement) Height() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/height
func (d_ DOMHTMLIFrameElement) SetHeight(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/longDesc
func (d_ DOMHTMLIFrameElement) LongDesc() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("longDesc"))
	return rv
}/* debug [instance_properties/getter]: longDesc */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/longDesc
func (d_ DOMHTMLIFrameElement) SetLongDesc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLongDesc:"), value)
}/* debug [instance_properties/setter]: longDesc */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/marginHeight
func (d_ DOMHTMLIFrameElement) MarginHeight() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("marginHeight"))
	return rv
}/* debug [instance_properties/getter]: marginHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/marginHeight
func (d_ DOMHTMLIFrameElement) SetMarginHeight(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarginHeight:"), value)
}/* debug [instance_properties/setter]: marginHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/marginWidth
func (d_ DOMHTMLIFrameElement) MarginWidth() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("marginWidth"))
	return rv
}/* debug [instance_properties/getter]: marginWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/marginWidth
func (d_ DOMHTMLIFrameElement) SetMarginWidth(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarginWidth:"), value)
}/* debug [instance_properties/setter]: marginWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/name
func (d_ DOMHTMLIFrameElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/name
func (d_ DOMHTMLIFrameElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/scrolling
func (d_ DOMHTMLIFrameElement) Scrolling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("scrolling"))
	return rv
}/* debug [instance_properties/getter]: scrolling */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/scrolling
func (d_ DOMHTMLIFrameElement) SetScrolling(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setScrolling:"), value)
}/* debug [instance_properties/setter]: scrolling */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/src
func (d_ DOMHTMLIFrameElement) Src() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("src"))
	return rv
}/* debug [instance_properties/getter]: src */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/src
func (d_ DOMHTMLIFrameElement) SetSrc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSrc:"), value)
}/* debug [instance_properties/setter]: src */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/width
func (d_ DOMHTMLIFrameElement) Width() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLIFrameElement/width
func (d_ DOMHTMLIFrameElement) SetWidth(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLIFrameElement */



