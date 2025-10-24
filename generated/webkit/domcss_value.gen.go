// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCSSValue */

/* debug [class_header]: Header for DOMCSSValue */
// The class instance for the [DOMCSSValue] class.
var (
	DOMCSSValueClass     _DOMCSSValueClass
	DOMCSSValueClassOnce sync.Once
)

func getDOMCSSValueClass() _DOMCSSValueClass {
	DOMCSSValueClassOnce.Do(func() {
		DOMCSSValueClass = _DOMCSSValueClass{objc.GetClass("DOMCSSValue")}
	})
	return DOMCSSValueClass
}

type _DOMCSSValueClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCSSValue */
// An interface definition for the [DOMCSSValue] class.
type IDOMCSSValue interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMCSSValue */
	// properties:
	CssText() objc.IObject /* cross-framework: NSString */
	SetCssText(value objc.IObject /* cross-framework: NSString */)
	CssValueType() unsafe.Pointer
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCSSValue */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCSSValue */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSValueClass) Alloc() DOMCSSValue {
	rv := objc.Send[DOMCSSValue](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSValueClass) New() DOMCSSValue {
	rv := objc.Send[DOMCSSValue](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSValue) Init() DOMCSSValue {
	rv := objc.Send[DOMCSSValue](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSValue) Autorelease() DOMCSSValue {
	rv := objc.Send[DOMCSSValue](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSValue creates a new DOMCSSValue instance.
func NewDOMCSSValue() DOMCSSValue {
	return getDOMCSSValueClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCSSValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSValue
type DOMCSSValue struct {
	DOMObject
}

// DOMCSSValueFrom constructs a [DOMCSSValue] from an unsafe.Pointer.
func DOMCSSValueFrom(ptr unsafe.Pointer) DOMCSSValue {
	return DOMCSSValue{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCSSValue */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCSSValue */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCSSValue */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCSSValue */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCSSValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSValue/cssText
func (d_ DOMCSSValue) CssText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("cssText"))
	return rv
} /* debug [instance_properties/getter]: cssText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSValue/cssText
func (d_ DOMCSSValue) SetCssText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCssText:"), value)
} /* debug [instance_properties/setter]: cssText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSValue/cssValueType
func (d_ DOMCSSValue) CssValueType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("cssValueType"))
	return rv
} /* debug [instance_properties/getter]: cssValueType */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCSSValue */
