// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLModElement */

/* debug [class_header]: Header for DOMHTMLModElement */
// The class instance for the [DOMHTMLModElement] class.
var (
	DOMHTMLModElementClass     _DOMHTMLModElementClass
	DOMHTMLModElementClassOnce sync.Once
)

func getDOMHTMLModElementClass() _DOMHTMLModElementClass {
	DOMHTMLModElementClassOnce.Do(func() {
		DOMHTMLModElementClass = _DOMHTMLModElementClass{objc.GetClass("DOMHTMLModElement")}
	})
	return DOMHTMLModElementClass
}

type _DOMHTMLModElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLModElement */
// An interface definition for the [DOMHTMLModElement] class.
type IDOMHTMLModElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLModElement */
	// properties:
	Cite() objc.IObject /* cross-framework: NSString */
	SetCite(value objc.IObject /* cross-framework: NSString */)
	DateTime() objc.IObject /* cross-framework: NSString */
	SetDateTime(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLModElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLModElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLModElementClass) Alloc() DOMHTMLModElement {
	rv := objc.Send[DOMHTMLModElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLModElementClass) New() DOMHTMLModElement {
	rv := objc.Send[DOMHTMLModElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLModElement) Init() DOMHTMLModElement {
	rv := objc.Send[DOMHTMLModElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLModElement) Autorelease() DOMHTMLModElement {
	rv := objc.Send[DOMHTMLModElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLModElement creates a new DOMHTMLModElement instance.
func NewDOMHTMLModElement() DOMHTMLModElement {
	return getDOMHTMLModElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLModElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLModElement
type DOMHTMLModElement struct {
	DOMHTMLElement
}

// DOMHTMLModElementFrom constructs a [DOMHTMLModElement] from an unsafe.Pointer.
func DOMHTMLModElementFrom(ptr unsafe.Pointer) DOMHTMLModElement {
	return DOMHTMLModElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLModElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLModElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLModElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLModElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLModElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLModElement/cite
func (d_ DOMHTMLModElement) Cite() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("cite"))
	return rv
} /* debug [instance_properties/getter]: cite */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLModElement/cite
func (d_ DOMHTMLModElement) SetCite(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCite:"), value)
} /* debug [instance_properties/setter]: cite */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLModElement/dateTime
func (d_ DOMHTMLModElement) DateTime() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("dateTime"))
	return rv
} /* debug [instance_properties/getter]: dateTime */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLModElement/dateTime
func (d_ DOMHTMLModElement) SetDateTime(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateTime:"), value)
} /* debug [instance_properties/setter]: dateTime */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLModElement */
