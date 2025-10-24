// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLOptGroupElement */

/* debug [class_header]: Header for DOMHTMLOptGroupElement */
// The class instance for the [DOMHTMLOptGroupElement] class.
var (
	DOMHTMLOptGroupElementClass     _DOMHTMLOptGroupElementClass
	DOMHTMLOptGroupElementClassOnce sync.Once
)

func getDOMHTMLOptGroupElementClass() _DOMHTMLOptGroupElementClass {
	DOMHTMLOptGroupElementClassOnce.Do(func() {
		DOMHTMLOptGroupElementClass = _DOMHTMLOptGroupElementClass{objc.GetClass("DOMHTMLOptGroupElement")}
	})
	return DOMHTMLOptGroupElementClass
}

type _DOMHTMLOptGroupElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLOptGroupElement */
// An interface definition for the [DOMHTMLOptGroupElement] class.
type IDOMHTMLOptGroupElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLOptGroupElement */
	// properties:
	Disabled() bool
	SetDisabled(value bool)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLOptGroupElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLOptGroupElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLOptGroupElementClass) Alloc() DOMHTMLOptGroupElement {
	rv := objc.Send[DOMHTMLOptGroupElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLOptGroupElementClass) New() DOMHTMLOptGroupElement {
	rv := objc.Send[DOMHTMLOptGroupElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLOptGroupElement) Init() DOMHTMLOptGroupElement {
	rv := objc.Send[DOMHTMLOptGroupElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLOptGroupElement) Autorelease() DOMHTMLOptGroupElement {
	rv := objc.Send[DOMHTMLOptGroupElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLOptGroupElement creates a new DOMHTMLOptGroupElement instance.
func NewDOMHTMLOptGroupElement() DOMHTMLOptGroupElement {
	return getDOMHTMLOptGroupElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLOptGroupElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptGroupElement
type DOMHTMLOptGroupElement struct {
	DOMHTMLElement
}

// DOMHTMLOptGroupElementFrom constructs a [DOMHTMLOptGroupElement] from an unsafe.Pointer.
func DOMHTMLOptGroupElementFrom(ptr unsafe.Pointer) DOMHTMLOptGroupElement {
	return DOMHTMLOptGroupElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLOptGroupElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLOptGroupElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLOptGroupElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLOptGroupElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLOptGroupElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptGroupElement/disabled
func (d_ DOMHTMLOptGroupElement) Disabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("disabled"))
	return rv
} /* debug [instance_properties/getter]: disabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptGroupElement/disabled
func (d_ DOMHTMLOptGroupElement) SetDisabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisabled:"), value)
} /* debug [instance_properties/setter]: disabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptGroupElement/label
func (d_ DOMHTMLOptGroupElement) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("label"))
	return rv
} /* debug [instance_properties/getter]: label */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptGroupElement/label
func (d_ DOMHTMLOptGroupElement) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLabel:"), value)
} /* debug [instance_properties/setter]: label */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLOptGroupElement */
