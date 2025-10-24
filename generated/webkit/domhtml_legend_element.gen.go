// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLLegendElement */

/* debug [class_header]: Header for DOMHTMLLegendElement */
// The class instance for the [DOMHTMLLegendElement] class.
var (
	DOMHTMLLegendElementClass     _DOMHTMLLegendElementClass
	DOMHTMLLegendElementClassOnce sync.Once
)

func getDOMHTMLLegendElementClass() _DOMHTMLLegendElementClass {
	DOMHTMLLegendElementClassOnce.Do(func() {
		DOMHTMLLegendElementClass = _DOMHTMLLegendElementClass{objc.GetClass("DOMHTMLLegendElement")}
	})
	return DOMHTMLLegendElementClass
}

type _DOMHTMLLegendElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLLegendElement */
// An interface definition for the [DOMHTMLLegendElement] class.
type IDOMHTMLLegendElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLLegendElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	Form() IDOMHTMLFormElement
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLLegendElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLLegendElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLLegendElementClass) Alloc() DOMHTMLLegendElement {
	rv := objc.Send[DOMHTMLLegendElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLLegendElementClass) New() DOMHTMLLegendElement {
	rv := objc.Send[DOMHTMLLegendElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLLegendElement) Init() DOMHTMLLegendElement {
	rv := objc.Send[DOMHTMLLegendElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLLegendElement) Autorelease() DOMHTMLLegendElement {
	rv := objc.Send[DOMHTMLLegendElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLLegendElement creates a new DOMHTMLLegendElement instance.
func NewDOMHTMLLegendElement() DOMHTMLLegendElement {
	return getDOMHTMLLegendElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLLegendElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLegendElement
type DOMHTMLLegendElement struct {
	DOMHTMLElement
}

// DOMHTMLLegendElementFrom constructs a [DOMHTMLLegendElement] from an unsafe.Pointer.
func DOMHTMLLegendElementFrom(ptr unsafe.Pointer) DOMHTMLLegendElement {
	return DOMHTMLLegendElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLLegendElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLLegendElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLLegendElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLLegendElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLLegendElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLegendElement/align
func (d_ DOMHTMLLegendElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
} /* debug [instance_properties/getter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLegendElement/align
func (d_ DOMHTMLLegendElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
} /* debug [instance_properties/setter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLegendElement/form
func (d_ DOMHTMLLegendElement) Form() IDOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("form"))
	return rv
} /* debug [instance_properties/getter]: form */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLLegendElement */
