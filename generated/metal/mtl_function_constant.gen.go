// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFunctionConstant */


/* debug [class_header]: Header for MTLFunctionConstant */
// The class instance for the [FunctionConstant] class.
var (
	FunctionConstantClass     _FunctionConstantClass
	FunctionConstantClassOnce sync.Once
)

func getFunctionConstantClass() _FunctionConstantClass {
	FunctionConstantClassOnce.Do(func() {
		FunctionConstantClass = _FunctionConstantClass{objc.GetClass("MTLFunctionConstant")}
	})
	return FunctionConstantClass
}

type _FunctionConstantClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FunctionConstant */
// An interface definition for the [FunctionConstant] class.
type IFunctionConstant interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FunctionConstant */
	// properties:
	Index() uint
	Name() objc.IObject /* cross-framework: NSString */
	Required() bool
	Type() DataType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FunctionConstant */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FunctionConstant */
// Alloc allocates a new instance without initialization.
func (fc _FunctionConstantClass) Alloc() FunctionConstant {
	rv := objc.Send[FunctionConstant](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionConstantClass) New() FunctionConstant {
	rv := objc.Send[FunctionConstant](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionConstant) Init() FunctionConstant {
	rv := objc.Send[FunctionConstant](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionConstant) Autorelease() FunctionConstant {
	rv := objc.Send[FunctionConstant](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionConstant creates a new FunctionConstant instance.
func NewFunctionConstant() FunctionConstant {
	return getFunctionConstantClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FunctionConstant */
// A constant that specializes the behavior of a shader.
//
// Don’t create an instance directly. Instead, the list of function constants for a function by querying the property of an instance. An instance should only be obtained from a nonspecialized function created with the method. You only need an instance if you don’t have sufficient information to create an instance used to create a specialized function with the or method.


// A constant that specializes the behavior of a shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant
type FunctionConstant struct {
	objectivec.Object
}

// FunctionConstantFrom constructs a [FunctionConstant] from an unsafe.Pointer.
//
// A constant that specializes the behavior of a shader.
func FunctionConstantFrom(ptr unsafe.Pointer) FunctionConstant {
	return FunctionConstant{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FunctionConstant *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FunctionConstant */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FunctionConstant */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FunctionConstant */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FunctionConstant */

// The index of the function constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/index
func (f_ FunctionConstant) Index() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("index"))
	return rv
}/* debug [instance_properties/getter]: index */


// The name of the function constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/name
func (f_ FunctionConstant) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// A Boolean value indicating whether the function constant must be provided to specialize the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/required
func (f_ FunctionConstant) Required() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("required"))
	return rv
}/* debug [instance_properties/getter]: required */


// The data type of the function constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/type
func (f_ FunctionConstant) Type() DataType {
	rv := objc.Send[DataType](f_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFunctionConstant */



