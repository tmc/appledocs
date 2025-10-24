// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLFrameElement */

/* debug [class_header]: Header for DOMHTMLFrameElement */
// The class instance for the [DOMHTMLFrameElement] class.
var (
	DOMHTMLFrameElementClass     _DOMHTMLFrameElementClass
	DOMHTMLFrameElementClassOnce sync.Once
)

func getDOMHTMLFrameElementClass() _DOMHTMLFrameElementClass {
	DOMHTMLFrameElementClassOnce.Do(func() {
		DOMHTMLFrameElementClass = _DOMHTMLFrameElementClass{objc.GetClass("DOMHTMLFrameElement")}
	})
	return DOMHTMLFrameElementClass
}

type _DOMHTMLFrameElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLFrameElement */
// An interface definition for the [DOMHTMLFrameElement] class.
type IDOMHTMLFrameElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLFrameElement */
	// properties:
	ContentDocument() IDOMDocument
	ContentFrame() IWebFrame
	ContentWindow() IDOMAbstractView
	FrameBorder() objc.IObject /* cross-framework: NSString */
	SetFrameBorder(value objc.IObject /* cross-framework: NSString */)
	Height() int
	Location() objc.IObject /* cross-framework: NSString */
	SetLocation(value objc.IObject /* cross-framework: NSString */)
	LongDesc() objc.IObject /* cross-framework: NSString */
	SetLongDesc(value objc.IObject /* cross-framework: NSString */)
	MarginHeight() objc.IObject /* cross-framework: NSString */
	SetMarginHeight(value objc.IObject /* cross-framework: NSString */)
	MarginWidth() objc.IObject /* cross-framework: NSString */
	SetMarginWidth(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	NoResize() bool
	SetNoResize(value bool)
	Scrolling() objc.IObject /* cross-framework: NSString */
	SetScrolling(value objc.IObject /* cross-framework: NSString */)
	Src() objc.IObject /* cross-framework: NSString */
	SetSrc(value objc.IObject /* cross-framework: NSString */)
	Width() int
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLFrameElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLFrameElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLFrameElementClass) Alloc() DOMHTMLFrameElement {
	rv := objc.Send[DOMHTMLFrameElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLFrameElementClass) New() DOMHTMLFrameElement {
	rv := objc.Send[DOMHTMLFrameElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLFrameElement) Init() DOMHTMLFrameElement {
	rv := objc.Send[DOMHTMLFrameElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLFrameElement) Autorelease() DOMHTMLFrameElement {
	rv := objc.Send[DOMHTMLFrameElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLFrameElement creates a new DOMHTMLFrameElement instance.
func NewDOMHTMLFrameElement() DOMHTMLFrameElement {
	return getDOMHTMLFrameElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLFrameElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement
type DOMHTMLFrameElement struct {
	DOMHTMLElement
}

// DOMHTMLFrameElementFrom constructs a [DOMHTMLFrameElement] from an unsafe.Pointer.
func DOMHTMLFrameElementFrom(ptr unsafe.Pointer) DOMHTMLFrameElement {
	return DOMHTMLFrameElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLFrameElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLFrameElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLFrameElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLFrameElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLFrameElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/contentDocument
func (d_ DOMHTMLFrameElement) ContentDocument() IDOMDocument {
	rv := objc.Send[DOMDocument](d_.ID, objc.Sel("contentDocument"))
	return rv
} /* debug [instance_properties/getter]: contentDocument */

// The content frame of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/contentFrame
func (d_ DOMHTMLFrameElement) ContentFrame() IWebFrame {
	rv := objc.Send[WebFrame](d_.ID, objc.Sel("contentFrame"))
	return rv
} /* debug [instance_properties/getter]: contentFrame */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/contentWindow
func (d_ DOMHTMLFrameElement) ContentWindow() IDOMAbstractView {
	rv := objc.Send[DOMAbstractView](d_.ID, objc.Sel("contentWindow"))
	return rv
} /* debug [instance_properties/getter]: contentWindow */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/frameBorder
func (d_ DOMHTMLFrameElement) FrameBorder() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("frameBorder"))
	return rv
} /* debug [instance_properties/getter]: frameBorder */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/frameBorder
func (d_ DOMHTMLFrameElement) SetFrameBorder(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFrameBorder:"), value)
} /* debug [instance_properties/setter]: frameBorder */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/height
func (d_ DOMHTMLFrameElement) Height() int {
	rv := objc.Send[int](d_.ID, objc.Sel("height"))
	return rv
} /* debug [instance_properties/getter]: height */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/location
func (d_ DOMHTMLFrameElement) Location() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("location"))
	return rv
} /* debug [instance_properties/getter]: location */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/location
func (d_ DOMHTMLFrameElement) SetLocation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocation:"), value)
} /* debug [instance_properties/setter]: location */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/longDesc
func (d_ DOMHTMLFrameElement) LongDesc() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("longDesc"))
	return rv
} /* debug [instance_properties/getter]: longDesc */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/longDesc
func (d_ DOMHTMLFrameElement) SetLongDesc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLongDesc:"), value)
} /* debug [instance_properties/setter]: longDesc */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/marginHeight
func (d_ DOMHTMLFrameElement) MarginHeight() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("marginHeight"))
	return rv
} /* debug [instance_properties/getter]: marginHeight */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/marginHeight
func (d_ DOMHTMLFrameElement) SetMarginHeight(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarginHeight:"), value)
} /* debug [instance_properties/setter]: marginHeight */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/marginWidth
func (d_ DOMHTMLFrameElement) MarginWidth() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("marginWidth"))
	return rv
} /* debug [instance_properties/getter]: marginWidth */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/marginWidth
func (d_ DOMHTMLFrameElement) SetMarginWidth(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMarginWidth:"), value)
} /* debug [instance_properties/setter]: marginWidth */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/name
func (d_ DOMHTMLFrameElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
} /* debug [instance_properties/getter]: name */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/name
func (d_ DOMHTMLFrameElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
} /* debug [instance_properties/setter]: name */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/noResize
func (d_ DOMHTMLFrameElement) NoResize() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("noResize"))
	return rv
} /* debug [instance_properties/getter]: noResize */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/noResize
func (d_ DOMHTMLFrameElement) SetNoResize(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNoResize:"), value)
} /* debug [instance_properties/setter]: noResize */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/scrolling
func (d_ DOMHTMLFrameElement) Scrolling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("scrolling"))
	return rv
} /* debug [instance_properties/getter]: scrolling */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/scrolling
func (d_ DOMHTMLFrameElement) SetScrolling(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setScrolling:"), value)
} /* debug [instance_properties/setter]: scrolling */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/src
func (d_ DOMHTMLFrameElement) Src() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("src"))
	return rv
} /* debug [instance_properties/getter]: src */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/src
func (d_ DOMHTMLFrameElement) SetSrc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSrc:"), value)
} /* debug [instance_properties/setter]: src */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameElement/width
func (d_ DOMHTMLFrameElement) Width() int {
	rv := objc.Send[int](d_.ID, objc.Sel("width"))
	return rv
} /* debug [instance_properties/getter]: width */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLFrameElement */
