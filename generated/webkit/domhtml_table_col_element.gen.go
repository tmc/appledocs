// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLTableColElement */

/* debug [class_header]: Header for DOMHTMLTableColElement */
// The class instance for the [DOMHTMLTableColElement] class.
var (
	DOMHTMLTableColElementClass     _DOMHTMLTableColElementClass
	DOMHTMLTableColElementClassOnce sync.Once
)

func getDOMHTMLTableColElementClass() _DOMHTMLTableColElementClass {
	DOMHTMLTableColElementClassOnce.Do(func() {
		DOMHTMLTableColElementClass = _DOMHTMLTableColElementClass{objc.GetClass("DOMHTMLTableColElement")}
	})
	return DOMHTMLTableColElementClass
}

type _DOMHTMLTableColElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLTableColElement */
// An interface definition for the [DOMHTMLTableColElement] class.
type IDOMHTMLTableColElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLTableColElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	Ch() objc.IObject /* cross-framework: NSString */
	SetCh(value objc.IObject /* cross-framework: NSString */)
	ChOff() objc.IObject /* cross-framework: NSString */
	SetChOff(value objc.IObject /* cross-framework: NSString */)
	Span() int
	SetSpan(value int)
	VAlign() objc.IObject /* cross-framework: NSString */
	SetVAlign(value objc.IObject /* cross-framework: NSString */)
	Width() objc.IObject /* cross-framework: NSString */
	SetWidth(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLTableColElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLTableColElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLTableColElementClass) Alloc() DOMHTMLTableColElement {
	rv := objc.Send[DOMHTMLTableColElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLTableColElementClass) New() DOMHTMLTableColElement {
	rv := objc.Send[DOMHTMLTableColElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLTableColElement) Init() DOMHTMLTableColElement {
	rv := objc.Send[DOMHTMLTableColElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLTableColElement) Autorelease() DOMHTMLTableColElement {
	rv := objc.Send[DOMHTMLTableColElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLTableColElement creates a new DOMHTMLTableColElement instance.
func NewDOMHTMLTableColElement() DOMHTMLTableColElement {
	return getDOMHTMLTableColElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLTableColElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement
type DOMHTMLTableColElement struct {
	DOMHTMLElement
}

// DOMHTMLTableColElementFrom constructs a [DOMHTMLTableColElement] from an unsafe.Pointer.
func DOMHTMLTableColElementFrom(ptr unsafe.Pointer) DOMHTMLTableColElement {
	return DOMHTMLTableColElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLTableColElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLTableColElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLTableColElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLTableColElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLTableColElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/align
func (d_ DOMHTMLTableColElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
} /* debug [instance_properties/getter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/align
func (d_ DOMHTMLTableColElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
} /* debug [instance_properties/setter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/ch
func (d_ DOMHTMLTableColElement) Ch() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("ch"))
	return rv
} /* debug [instance_properties/getter]: ch */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/ch
func (d_ DOMHTMLTableColElement) SetCh(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCh:"), value)
} /* debug [instance_properties/setter]: ch */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/chOff
func (d_ DOMHTMLTableColElement) ChOff() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("chOff"))
	return rv
} /* debug [instance_properties/getter]: chOff */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/chOff
func (d_ DOMHTMLTableColElement) SetChOff(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setChOff:"), value)
} /* debug [instance_properties/setter]: chOff */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/span
func (d_ DOMHTMLTableColElement) Span() int {
	rv := objc.Send[int](d_.ID, objc.Sel("span"))
	return rv
} /* debug [instance_properties/getter]: span */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/span
func (d_ DOMHTMLTableColElement) SetSpan(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSpan:"), value)
} /* debug [instance_properties/setter]: span */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/vAlign
func (d_ DOMHTMLTableColElement) VAlign() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("vAlign"))
	return rv
} /* debug [instance_properties/getter]: vAlign */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/vAlign
func (d_ DOMHTMLTableColElement) SetVAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVAlign:"), value)
} /* debug [instance_properties/setter]: vAlign */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/width
func (d_ DOMHTMLTableColElement) Width() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("width"))
	return rv
} /* debug [instance_properties/getter]: width */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableColElement/width
func (d_ DOMHTMLTableColElement) SetWidth(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
} /* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLTableColElement */
