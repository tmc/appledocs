// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class function */


/* debug [class_header]: Header for function */
// The class instance for the [function] class.
var (
	FunctionClass     _functionClass
	FunctionClassOnce sync.Once
)

func getfunctionClass() _functionClass {
	FunctionClassOnce.Do(func() {
		FunctionClass = _functionClass{objc.GetClass("function")}
	})
	return FunctionClass
}

type _functionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for function */
// An interface definition for the [function] class.
type Ifunction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for function */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for function */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for function */
// Alloc allocates a new instance without initialization.
func (fc _functionClass) Alloc() function {
	rv := objc.Send[function](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _functionClass) New() function {
	rv := objc.Send[function](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ function) Init() function {
	rv := objc.Send[function](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ function) Autorelease() function {
	rv := objc.Send[function](f_.ID, objc.Sel("autorelease"))
	return rv
}

// Newfunction creates a new function instance.
func Newfunction() function {
	return getfunctionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for function */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/function-c.ivar
type function struct {
	objectivec.Object
}

// functionFrom constructs a [function] from an unsafe.Pointer.
func functionFrom(ptr unsafe.Pointer) function {
	return function{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for function *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for function */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for function */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for function */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for function */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class function */



