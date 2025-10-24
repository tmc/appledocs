// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMCSSPrimitiveValue */


/* debug [class_header]: Header for DOMCSSPrimitiveValue */
// The class instance for the [DOMCSSPrimitiveValue] class.
var (
	DOMCSSPrimitiveValueClass     _DOMCSSPrimitiveValueClass
	DOMCSSPrimitiveValueClassOnce sync.Once
)

func getDOMCSSPrimitiveValueClass() _DOMCSSPrimitiveValueClass {
	DOMCSSPrimitiveValueClassOnce.Do(func() {
		DOMCSSPrimitiveValueClass = _DOMCSSPrimitiveValueClass{objc.GetClass("DOMCSSPrimitiveValue")}
	})
	return DOMCSSPrimitiveValueClass
}

type _DOMCSSPrimitiveValueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMCSSPrimitiveValue */
// An interface definition for the [DOMCSSPrimitiveValue] class.
type IDOMCSSPrimitiveValue interface {
	IDOMCSSValue
	
/* debug [class_interface_properties]: Properties for DOMCSSPrimitiveValue */
	// properties:
	PrimitiveType() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMCSSPrimitiveValue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMCSSPrimitiveValue */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSPrimitiveValueClass) Alloc() DOMCSSPrimitiveValue {
	rv := objc.Send[DOMCSSPrimitiveValue](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSPrimitiveValueClass) New() DOMCSSPrimitiveValue {
	rv := objc.Send[DOMCSSPrimitiveValue](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSPrimitiveValue) Init() DOMCSSPrimitiveValue {
	rv := objc.Send[DOMCSSPrimitiveValue](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSPrimitiveValue) Autorelease() DOMCSSPrimitiveValue {
	rv := objc.Send[DOMCSSPrimitiveValue](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSPrimitiveValue creates a new DOMCSSPrimitiveValue instance.
func NewDOMCSSPrimitiveValue() DOMCSSPrimitiveValue {
	return getDOMCSSPrimitiveValueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMCSSPrimitiveValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSPrimitiveValue
type DOMCSSPrimitiveValue struct {
	DOMCSSValue
}

// DOMCSSPrimitiveValueFrom constructs a [DOMCSSPrimitiveValue] from an unsafe.Pointer.
func DOMCSSPrimitiveValueFrom(ptr unsafe.Pointer) DOMCSSPrimitiveValue {
	return DOMCSSPrimitiveValue{
		DOMCSSValue: DOMCSSValueFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMCSSPrimitiveValue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMCSSPrimitiveValue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMCSSPrimitiveValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMCSSPrimitiveValue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMCSSPrimitiveValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSPrimitiveValue/primitiveType
func (d_ DOMCSSPrimitiveValue) PrimitiveType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("primitiveType"))
	return rv
}/* debug [instance_properties/getter]: primitiveType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMCSSPrimitiveValue */



