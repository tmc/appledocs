// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class customQueryFunction */


/* debug [class_header]: Header for customQueryFunction */
// The class instance for the [customQueryFunction] class.
var (
	CustomQueryFunctionClass     _customQueryFunctionClass
	CustomQueryFunctionClassOnce sync.Once
)

func getcustomQueryFunctionClass() _customQueryFunctionClass {
	CustomQueryFunctionClassOnce.Do(func() {
		CustomQueryFunctionClass = _customQueryFunctionClass{objc.GetClass("customQueryFunction")}
	})
	return CustomQueryFunctionClass
}

type _customQueryFunctionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for customQueryFunction */
// An interface definition for the [customQueryFunction] class.
type IcustomQueryFunction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for customQueryFunction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for customQueryFunction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for customQueryFunction */
// Alloc allocates a new instance without initialization.
func (cc _customQueryFunctionClass) Alloc() customQueryFunction {
	rv := objc.Send[customQueryFunction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _customQueryFunctionClass) New() customQueryFunction {
	rv := objc.Send[customQueryFunction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ customQueryFunction) Init() customQueryFunction {
	rv := objc.Send[customQueryFunction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ customQueryFunction) Autorelease() customQueryFunction {
	rv := objc.Send[customQueryFunction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcustomQueryFunction creates a new customQueryFunction instance.
func NewcustomQueryFunction() customQueryFunction {
	return getcustomQueryFunctionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for customQueryFunction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customQueryFunction-c.ivar
type customQueryFunction struct {
	objectivec.Object
}

// customQueryFunctionFrom constructs a [customQueryFunction] from an unsafe.Pointer.
func customQueryFunctionFrom(ptr unsafe.Pointer) customQueryFunction {
	return customQueryFunction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for customQueryFunction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for customQueryFunction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for customQueryFunction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for customQueryFunction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for customQueryFunction */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class customQueryFunction */



