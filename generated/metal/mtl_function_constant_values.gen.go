// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFunctionConstantValues */


/* debug [class_header]: Header for MTLFunctionConstantValues */
// The class instance for the [FunctionConstantValues] class.
var (
	FunctionConstantValuesClass     _FunctionConstantValuesClass
	FunctionConstantValuesClassOnce sync.Once
)

func getFunctionConstantValuesClass() _FunctionConstantValuesClass {
	FunctionConstantValuesClassOnce.Do(func() {
		FunctionConstantValuesClass = _FunctionConstantValuesClass{objc.GetClass("MTLFunctionConstantValues")}
	})
	return FunctionConstantValuesClass
}

type _FunctionConstantValuesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FunctionConstantValues */
// An interface definition for the [FunctionConstantValues] class.
type IFunctionConstantValues interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FunctionConstantValues */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FunctionConstantValues */
	// methods:
	Reset()
	SetConstantValueTypeAtIndex(value objectivec.IObject, type_ DataType, index uint)
	SetConstantValueTypeWithName(value objectivec.IObject, type_ DataType, name objc.IObject /* cross-framework: NSString */)
	SetConstantValuesTypeWithRange(values objectivec.IObject, type_ DataType, range_ corefoundation.Range)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FunctionConstantValues */
// Alloc allocates a new instance without initialization.
func (fc _FunctionConstantValuesClass) Alloc() FunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionConstantValuesClass) New() FunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionConstantValues) Init() FunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionConstantValues) Autorelease() FunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionConstantValues creates a new FunctionConstantValues instance.
func NewFunctionConstantValues() FunctionConstantValues {
	return getFunctionConstantValuesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FunctionConstantValues */
// A set of constant values that specialize a graphics or compute GPU function.
//
// An instance sets constant values for function constants. You declare function constants with the attribute in MSL (Metal Shading Language) source code. See the for more information. With an instance, you can set each constant value individually with an index or a name, or set multiple constant values with an index range. You can apply a single instance to multiple instances of any kind, such as a vertex function and a fragment function. When you create a specialized function, subsequent changes to its constant values have no effect. However, you can reset, add, or modify a constant value in your instance and reuse it to create another instance.


// A set of constant values that specialize a graphics or compute GPU function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues
type FunctionConstantValues struct {
	objectivec.Object
}

// FunctionConstantValuesFrom constructs a [FunctionConstantValues] from an unsafe.Pointer.
//
// A set of constant values that specialize a graphics or compute GPU function.
func FunctionConstantValuesFrom(ptr unsafe.Pointer) FunctionConstantValues {
	return FunctionConstantValues{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FunctionConstantValues *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FunctionConstantValues */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FunctionConstantValues */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FunctionConstantValues */

// Deletes all previously set constant values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/reset()
func (f_ FunctionConstantValues) Reset() {
	objc.Send[objc.ID](f_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */


// Sets a value for a function constant at a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/setConstantValue(_:type:index:)
func (f_ FunctionConstantValues) SetConstantValueTypeAtIndex(value objectivec.IObject, type_ DataType, index uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConstantValue:type:atIndex:"), value, type_, index)
}/* debug [instance_methods/method]: SetConstantValueTypeAtIndex */


// Sets a value for a function constant with a specific name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/setConstantValue(_:type:withName:)
func (f_ FunctionConstantValues) SetConstantValueTypeWithName(value objectivec.IObject, type_ DataType, name objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConstantValue:type:withName:"), value, type_, name)
}/* debug [instance_methods/method]: SetConstantValueTypeWithName */


// Sets values for a group of function constants within a specific index range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/setConstantValues:type:withRange:
func (f_ FunctionConstantValues) SetConstantValuesTypeWithRange(values objectivec.IObject, type_ DataType, range_ corefoundation.Range) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConstantValues:type:withRange:"), values, type_, range_)
}/* debug [instance_methods/method]: SetConstantValuesTypeWithRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FunctionConstantValues */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFunctionConstantValues */



