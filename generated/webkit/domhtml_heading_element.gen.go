// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLHeadingElement */

/* debug [class_header]: Header for DOMHTMLHeadingElement */
// The class instance for the [DOMHTMLHeadingElement] class.
var (
	DOMHTMLHeadingElementClass     _DOMHTMLHeadingElementClass
	DOMHTMLHeadingElementClassOnce sync.Once
)

func getDOMHTMLHeadingElementClass() _DOMHTMLHeadingElementClass {
	DOMHTMLHeadingElementClassOnce.Do(func() {
		DOMHTMLHeadingElementClass = _DOMHTMLHeadingElementClass{objc.GetClass("DOMHTMLHeadingElement")}
	})
	return DOMHTMLHeadingElementClass
}

type _DOMHTMLHeadingElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLHeadingElement */
// An interface definition for the [DOMHTMLHeadingElement] class.
type IDOMHTMLHeadingElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLHeadingElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLHeadingElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLHeadingElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLHeadingElementClass) Alloc() DOMHTMLHeadingElement {
	rv := objc.Send[DOMHTMLHeadingElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLHeadingElementClass) New() DOMHTMLHeadingElement {
	rv := objc.Send[DOMHTMLHeadingElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLHeadingElement) Init() DOMHTMLHeadingElement {
	rv := objc.Send[DOMHTMLHeadingElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLHeadingElement) Autorelease() DOMHTMLHeadingElement {
	rv := objc.Send[DOMHTMLHeadingElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLHeadingElement creates a new DOMHTMLHeadingElement instance.
func NewDOMHTMLHeadingElement() DOMHTMLHeadingElement {
	return getDOMHTMLHeadingElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLHeadingElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHeadingElement
type DOMHTMLHeadingElement struct {
	DOMHTMLElement
}

// DOMHTMLHeadingElementFrom constructs a [DOMHTMLHeadingElement] from an unsafe.Pointer.
func DOMHTMLHeadingElementFrom(ptr unsafe.Pointer) DOMHTMLHeadingElement {
	return DOMHTMLHeadingElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLHeadingElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLHeadingElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLHeadingElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLHeadingElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLHeadingElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHeadingElement/align
func (d_ DOMHTMLHeadingElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
} /* debug [instance_properties/getter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHeadingElement/align
func (d_ DOMHTMLHeadingElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
} /* debug [instance_properties/setter]: align */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLHeadingElement */
