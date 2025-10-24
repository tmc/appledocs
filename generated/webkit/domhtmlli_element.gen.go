// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLLIElement */

/* debug [class_header]: Header for DOMHTMLLIElement */
// The class instance for the [DOMHTMLLIElement] class.
var (
	DOMHTMLLIElementClass     _DOMHTMLLIElementClass
	DOMHTMLLIElementClassOnce sync.Once
)

func getDOMHTMLLIElementClass() _DOMHTMLLIElementClass {
	DOMHTMLLIElementClassOnce.Do(func() {
		DOMHTMLLIElementClass = _DOMHTMLLIElementClass{objc.GetClass("DOMHTMLLIElement")}
	})
	return DOMHTMLLIElementClass
}

type _DOMHTMLLIElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLLIElement */
// An interface definition for the [DOMHTMLLIElement] class.
type IDOMHTMLLIElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLLIElement */
	// properties:
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	Value() int
	SetValue(value int)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLLIElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLLIElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLLIElementClass) Alloc() DOMHTMLLIElement {
	rv := objc.Send[DOMHTMLLIElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLLIElementClass) New() DOMHTMLLIElement {
	rv := objc.Send[DOMHTMLLIElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLLIElement) Init() DOMHTMLLIElement {
	rv := objc.Send[DOMHTMLLIElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLLIElement) Autorelease() DOMHTMLLIElement {
	rv := objc.Send[DOMHTMLLIElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLLIElement creates a new DOMHTMLLIElement instance.
func NewDOMHTMLLIElement() DOMHTMLLIElement {
	return getDOMHTMLLIElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLLIElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLIElement
type DOMHTMLLIElement struct {
	DOMHTMLElement
}

// DOMHTMLLIElementFrom constructs a [DOMHTMLLIElement] from an unsafe.Pointer.
func DOMHTMLLIElementFrom(ptr unsafe.Pointer) DOMHTMLLIElement {
	return DOMHTMLLIElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLLIElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLLIElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLLIElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLLIElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLLIElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLIElement/type
func (d_ DOMHTMLLIElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
} /* debug [instance_properties/getter]: type */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLIElement/type
func (d_ DOMHTMLLIElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
} /* debug [instance_properties/setter]: type */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLIElement/value
func (d_ DOMHTMLLIElement) Value() int {
	rv := objc.Send[int](d_.ID, objc.Sel("value"))
	return rv
} /* debug [instance_properties/getter]: value */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLIElement/value
func (d_ DOMHTMLLIElement) SetValue(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
} /* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLLIElement */
