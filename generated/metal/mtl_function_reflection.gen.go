// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFunctionReflection */


/* debug [class_header]: Header for MTLFunctionReflection */
// The class instance for the [FunctionReflection] class.
var (
	FunctionReflectionClass     _FunctionReflectionClass
	FunctionReflectionClassOnce sync.Once
)

func getFunctionReflectionClass() _FunctionReflectionClass {
	FunctionReflectionClassOnce.Do(func() {
		FunctionReflectionClass = _FunctionReflectionClass{objc.GetClass("MTLFunctionReflection")}
	})
	return FunctionReflectionClass
}

type _FunctionReflectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FunctionReflection */
// An interface definition for the [FunctionReflection] class.
type IFunctionReflection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FunctionReflection */
	// properties:
	Bindings() []objc.ID
	UserAnnotation() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FunctionReflection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FunctionReflection */
// Alloc allocates a new instance without initialization.
func (fc _FunctionReflectionClass) Alloc() FunctionReflection {
	rv := objc.Send[FunctionReflection](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionReflectionClass) New() FunctionReflection {
	rv := objc.Send[FunctionReflection](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionReflection) Init() FunctionReflection {
	rv := objc.Send[FunctionReflection](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionReflection) Autorelease() FunctionReflection {
	rv := objc.Send[FunctionReflection](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionReflection creates a new FunctionReflection instance.
func NewFunctionReflection() FunctionReflection {
	return getFunctionReflectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FunctionReflection */
// Represents a reflection object containing information about a function in a Metal library.


// Represents a reflection object containing information about a function in a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionReflection
type FunctionReflection struct {
	objectivec.Object
}

// FunctionReflectionFrom constructs a [FunctionReflection] from an unsafe.Pointer.
//
// Represents a reflection object containing information about a function in a Metal library.
func FunctionReflectionFrom(ptr unsafe.Pointer) FunctionReflection {
	return FunctionReflection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FunctionReflection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FunctionReflection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FunctionReflection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FunctionReflection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FunctionReflection */

// Provides a list of inputs and outputs of the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionReflection/bindings
func (f_ FunctionReflection) Bindings() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("bindings"))
	return rv
}/* debug [instance_properties/getter]: bindings */


// The string passed to the user annotation attribute for this function. Null if no user annotation is present for this function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionReflection/userAnnotation
func (f_ FunctionReflection) UserAnnotation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("userAnnotation"))
	return rv
}/* debug [instance_properties/getter]: userAnnotation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFunctionReflection */



