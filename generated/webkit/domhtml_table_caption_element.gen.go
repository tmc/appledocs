// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLTableCaptionElement */

/* debug [class_header]: Header for DOMHTMLTableCaptionElement */
// The class instance for the [DOMHTMLTableCaptionElement] class.
var (
	DOMHTMLTableCaptionElementClass     _DOMHTMLTableCaptionElementClass
	DOMHTMLTableCaptionElementClassOnce sync.Once
)

func getDOMHTMLTableCaptionElementClass() _DOMHTMLTableCaptionElementClass {
	DOMHTMLTableCaptionElementClassOnce.Do(func() {
		DOMHTMLTableCaptionElementClass = _DOMHTMLTableCaptionElementClass{objc.GetClass("DOMHTMLTableCaptionElement")}
	})
	return DOMHTMLTableCaptionElementClass
}

type _DOMHTMLTableCaptionElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLTableCaptionElement */
// An interface definition for the [DOMHTMLTableCaptionElement] class.
type IDOMHTMLTableCaptionElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLTableCaptionElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLTableCaptionElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLTableCaptionElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLTableCaptionElementClass) Alloc() DOMHTMLTableCaptionElement {
	rv := objc.Send[DOMHTMLTableCaptionElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLTableCaptionElementClass) New() DOMHTMLTableCaptionElement {
	rv := objc.Send[DOMHTMLTableCaptionElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLTableCaptionElement) Init() DOMHTMLTableCaptionElement {
	rv := objc.Send[DOMHTMLTableCaptionElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLTableCaptionElement) Autorelease() DOMHTMLTableCaptionElement {
	rv := objc.Send[DOMHTMLTableCaptionElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLTableCaptionElement creates a new DOMHTMLTableCaptionElement instance.
func NewDOMHTMLTableCaptionElement() DOMHTMLTableCaptionElement {
	return getDOMHTMLTableCaptionElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLTableCaptionElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCaptionElement
type DOMHTMLTableCaptionElement struct {
	DOMHTMLElement
}

// DOMHTMLTableCaptionElementFrom constructs a [DOMHTMLTableCaptionElement] from an unsafe.Pointer.
func DOMHTMLTableCaptionElementFrom(ptr unsafe.Pointer) DOMHTMLTableCaptionElement {
	return DOMHTMLTableCaptionElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLTableCaptionElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLTableCaptionElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLTableCaptionElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLTableCaptionElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLTableCaptionElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCaptionElement/align
func (d_ DOMHTMLTableCaptionElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
} /* debug [instance_properties/getter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCaptionElement/align
func (d_ DOMHTMLTableCaptionElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
} /* debug [instance_properties/setter]: align */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLTableCaptionElement */
