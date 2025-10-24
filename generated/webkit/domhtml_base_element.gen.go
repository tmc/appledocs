// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLBaseElement */

/* debug [class_header]: Header for DOMHTMLBaseElement */
// The class instance for the [DOMHTMLBaseElement] class.
var (
	DOMHTMLBaseElementClass     _DOMHTMLBaseElementClass
	DOMHTMLBaseElementClassOnce sync.Once
)

func getDOMHTMLBaseElementClass() _DOMHTMLBaseElementClass {
	DOMHTMLBaseElementClassOnce.Do(func() {
		DOMHTMLBaseElementClass = _DOMHTMLBaseElementClass{objc.GetClass("DOMHTMLBaseElement")}
	})
	return DOMHTMLBaseElementClass
}

type _DOMHTMLBaseElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLBaseElement */
// An interface definition for the [DOMHTMLBaseElement] class.
type IDOMHTMLBaseElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLBaseElement */
	// properties:
	Href() objc.IObject /* cross-framework: NSString */
	SetHref(value objc.IObject /* cross-framework: NSString */)
	Target() objc.IObject /* cross-framework: NSString */
	SetTarget(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLBaseElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLBaseElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLBaseElementClass) Alloc() DOMHTMLBaseElement {
	rv := objc.Send[DOMHTMLBaseElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLBaseElementClass) New() DOMHTMLBaseElement {
	rv := objc.Send[DOMHTMLBaseElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLBaseElement) Init() DOMHTMLBaseElement {
	rv := objc.Send[DOMHTMLBaseElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLBaseElement) Autorelease() DOMHTMLBaseElement {
	rv := objc.Send[DOMHTMLBaseElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLBaseElement creates a new DOMHTMLBaseElement instance.
func NewDOMHTMLBaseElement() DOMHTMLBaseElement {
	return getDOMHTMLBaseElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLBaseElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseElement
type DOMHTMLBaseElement struct {
	DOMHTMLElement
}

// DOMHTMLBaseElementFrom constructs a [DOMHTMLBaseElement] from an unsafe.Pointer.
func DOMHTMLBaseElementFrom(ptr unsafe.Pointer) DOMHTMLBaseElement {
	return DOMHTMLBaseElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLBaseElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLBaseElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLBaseElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLBaseElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLBaseElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseElement/href
func (d_ DOMHTMLBaseElement) Href() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("href"))
	return rv
} /* debug [instance_properties/getter]: href */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseElement/href
func (d_ DOMHTMLBaseElement) SetHref(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHref:"), value)
} /* debug [instance_properties/setter]: href */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseElement/target
func (d_ DOMHTMLBaseElement) Target() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("target"))
	return rv
} /* debug [instance_properties/getter]: target */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseElement/target
func (d_ DOMHTMLBaseElement) SetTarget(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTarget:"), value)
} /* debug [instance_properties/setter]: target */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLBaseElement */
