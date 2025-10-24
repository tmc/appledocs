// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMRect */

/* debug [class_header]: Header for DOMRect */
// The class instance for the [DOMRect] class.
var (
	DOMRectClass     _DOMRectClass
	DOMRectClassOnce sync.Once
)

func getDOMRectClass() _DOMRectClass {
	DOMRectClassOnce.Do(func() {
		DOMRectClass = _DOMRectClass{objc.GetClass("DOMRect")}
	})
	return DOMRectClass
}

type _DOMRectClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMRect */
// An interface definition for the [DOMRect] class.
type IDOMRect interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMRect */
	// properties:
	Bottom() IDOMCSSPrimitiveValue
	Left() IDOMCSSPrimitiveValue
	Right() IDOMCSSPrimitiveValue
	Top() IDOMCSSPrimitiveValue
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMRect */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMRect */
// Alloc allocates a new instance without initialization.
func (dc _DOMRectClass) Alloc() DOMRect {
	rv := objc.Send[DOMRect](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMRectClass) New() DOMRect {
	rv := objc.Send[DOMRect](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMRect) Init() DOMRect {
	rv := objc.Send[DOMRect](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMRect) Autorelease() DOMRect {
	rv := objc.Send[DOMRect](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMRect creates a new DOMRect instance.
func NewDOMRect() DOMRect {
	return getDOMRectClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMRect */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRect
type DOMRect struct {
	DOMObject
}

// DOMRectFrom constructs a [DOMRect] from an unsafe.Pointer.
func DOMRectFrom(ptr unsafe.Pointer) DOMRect {
	return DOMRect{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMRect */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMRect */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMRect */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMRect */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMRect */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRect/bottom
func (d_ DOMRect) Bottom() IDOMCSSPrimitiveValue {
	rv := objc.Send[DOMCSSPrimitiveValue](d_.ID, objc.Sel("bottom"))
	return rv
} /* debug [instance_properties/getter]: bottom */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRect/left
func (d_ DOMRect) Left() IDOMCSSPrimitiveValue {
	rv := objc.Send[DOMCSSPrimitiveValue](d_.ID, objc.Sel("left"))
	return rv
} /* debug [instance_properties/getter]: left */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRect/right
func (d_ DOMRect) Right() IDOMCSSPrimitiveValue {
	rv := objc.Send[DOMCSSPrimitiveValue](d_.ID, objc.Sel("right"))
	return rv
} /* debug [instance_properties/getter]: right */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMRect/top
func (d_ DOMRect) Top() IDOMCSSPrimitiveValue {
	rv := objc.Send[DOMCSSPrimitiveValue](d_.ID, objc.Sel("top"))
	return rv
} /* debug [instance_properties/getter]: top */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMRect */
