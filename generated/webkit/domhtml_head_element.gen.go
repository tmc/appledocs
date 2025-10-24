// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLHeadElement */

/* debug [class_header]: Header for DOMHTMLHeadElement */
// The class instance for the [DOMHTMLHeadElement] class.
var (
	DOMHTMLHeadElementClass     _DOMHTMLHeadElementClass
	DOMHTMLHeadElementClassOnce sync.Once
)

func getDOMHTMLHeadElementClass() _DOMHTMLHeadElementClass {
	DOMHTMLHeadElementClassOnce.Do(func() {
		DOMHTMLHeadElementClass = _DOMHTMLHeadElementClass{objc.GetClass("DOMHTMLHeadElement")}
	})
	return DOMHTMLHeadElementClass
}

type _DOMHTMLHeadElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLHeadElement */
// An interface definition for the [DOMHTMLHeadElement] class.
type IDOMHTMLHeadElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLHeadElement */
	// properties:
	Profile() objc.IObject /* cross-framework: NSString */
	SetProfile(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLHeadElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLHeadElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLHeadElementClass) Alloc() DOMHTMLHeadElement {
	rv := objc.Send[DOMHTMLHeadElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLHeadElementClass) New() DOMHTMLHeadElement {
	rv := objc.Send[DOMHTMLHeadElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLHeadElement) Init() DOMHTMLHeadElement {
	rv := objc.Send[DOMHTMLHeadElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLHeadElement) Autorelease() DOMHTMLHeadElement {
	rv := objc.Send[DOMHTMLHeadElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLHeadElement creates a new DOMHTMLHeadElement instance.
func NewDOMHTMLHeadElement() DOMHTMLHeadElement {
	return getDOMHTMLHeadElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLHeadElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHeadElement
type DOMHTMLHeadElement struct {
	DOMHTMLElement
}

// DOMHTMLHeadElementFrom constructs a [DOMHTMLHeadElement] from an unsafe.Pointer.
func DOMHTMLHeadElementFrom(ptr unsafe.Pointer) DOMHTMLHeadElement {
	return DOMHTMLHeadElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLHeadElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLHeadElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLHeadElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLHeadElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLHeadElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHeadElement/profile
func (d_ DOMHTMLHeadElement) Profile() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("profile"))
	return rv
} /* debug [instance_properties/getter]: profile */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHeadElement/profile
func (d_ DOMHTMLHeadElement) SetProfile(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setProfile:"), value)
} /* debug [instance_properties/setter]: profile */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLHeadElement */
