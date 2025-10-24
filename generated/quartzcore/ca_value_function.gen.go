// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAValueFunction */


/* debug [class_header]: Header for CAValueFunction */
// The class instance for the [ValueFunction] class.
var (
	ValueFunctionClass     _ValueFunctionClass
	ValueFunctionClassOnce sync.Once
)

func getValueFunctionClass() _ValueFunctionClass {
	ValueFunctionClassOnce.Do(func() {
		ValueFunctionClass = _ValueFunctionClass{objc.GetClass("CAValueFunction")}
	})
	return ValueFunctionClass
}

type _ValueFunctionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ValueFunction */
// An interface definition for the [ValueFunction] class.
type IValueFunction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ValueFunction */
	// properties:
	Name() ValueFunctionName /* typedef */
	FromValue() objectivec.IObject
	SetFromValue(value objectivec.IObject)
	ToValue() objectivec.IObject
	SetToValue(value objectivec.IObject)
	ValueFunction() IValueFunction
	SetValueFunction(value IValueFunction)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ValueFunction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ValueFunction */
// Alloc allocates a new instance without initialization.
func (vc _ValueFunctionClass) Alloc() ValueFunction {
	rv := objc.Send[ValueFunction](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _ValueFunctionClass) New() ValueFunction {
	rv := objc.Send[ValueFunction](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ ValueFunction) Init() ValueFunction {
	rv := objc.Send[ValueFunction](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ ValueFunction) Autorelease() ValueFunction {
	rv := objc.Send[ValueFunction](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewValueFunction creates a new ValueFunction instance.
func NewValueFunction() ValueFunction {
	return getValueFunctionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ValueFunction */
// An object that provides a flexible method of defining animated transformations.
//
// You can use a value function to specify the individual components of an animated transform. For example, to create a basic animation that rotates a layer from 0° to 180° around its z-axis, you would create a object with a of , a of , and a of a with a function name of . The following code shows how you would create such a rotation and apply it to a named . The value functions and require 3 values, for the individual , and components. When working with these value functions, you specify the animation’s and as arrays. The following code shows how you could animate a layer’s scale from to using a value function.


// An object that provides a flexible method of defining animated transformations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAValueFunction
type ValueFunction struct {
	objectivec.Object
}

// ValueFunctionFrom constructs a [ValueFunction] from an unsafe.Pointer.
//
// An object that provides a flexible method of defining animated transformations.
func ValueFunctionFrom(ptr unsafe.Pointer) ValueFunction {
	return ValueFunction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ValueFunction */

// Returns the value function object identified by the name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAValueFunction/init(name:)
func NewValueFunctionWithName(name ValueFunctionName /* typedef */) ValueFunction {
	rv := objc.Send[ValueFunction](objc.ID(getValueFunctionClass().class), objc.Sel("functionWithName:"), name)
	return rv
}/* debug [class_init_methods/constructor]: NewValueFunctionWithName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ValueFunction */

// Returns the value function object identified by the name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAValueFunction/init(name:)
func (vc _ValueFunctionClass) FunctionWithName(name ValueFunctionName /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("functionWithName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FunctionWithName) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ValueFunction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ValueFunction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ValueFunction */

// Returns the name of the value function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAValueFunction/name
func (v_ ValueFunction) Name() ValueFunctionName /* typedef */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// Defines the value the receiver uses to start interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/fromvalue
func (v_ ValueFunction) FromValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("fromValue"))
	return rv
}/* debug [instance_properties/getter]: fromValue */


// Defines the value the receiver uses to start interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/fromvalue
func (v_ ValueFunction) SetFromValue(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFromValue:"), value)
}/* debug [instance_properties/setter]: fromValue */


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/tovalue
func (v_ ValueFunction) ToValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("toValue"))
	return rv
}/* debug [instance_properties/getter]: toValue */


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/tovalue
func (v_ ValueFunction) SetToValue(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setToValue:"), value)
}/* debug [instance_properties/setter]: toValue */


// An optional value function that is applied to interpolated values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/valuefunction
func (v_ ValueFunction) ValueFunction() IValueFunction {
	rv := objc.Send[ValueFunction](v_.ID, objc.Sel("valueFunction"))
	return rv
}/* debug [instance_properties/getter]: valueFunction */


// An optional value function that is applied to interpolated values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/valuefunction
func (v_ ValueFunction) SetValueFunction(value IValueFunction) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setValueFunction:"), value)
}/* debug [instance_properties/setter]: valueFunction */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAValueFunction */


