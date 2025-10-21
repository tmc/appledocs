// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ValueFunction] class.
type IValueFunction interface {
	objectivec.IObject
}

// An object that provides a flexible method of defining animated transformations.
//
// You can use a value function to specify the individual components of an animated transform. For example, to create a basic animation that rotates a layer from 0° to 180° around its z-axis, you would create a object with a of , a of , and a of a with a function name of . The following code shows how you would create such a rotation and apply it to a named . The value functions and require 3 values, for the individual , and components. When working with these value functions, you specify the animation’s and as arrays. The following code shows how you could animate a layer’s scale from to using a value function.
//
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

// Alloc allocates a new instance without initialization.
func (vc _ValueFunctionClass) Alloc() ValueFunction {
	rv := objc.Send[ValueFunction](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the value function object identified by the name.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAValueFunction/init(name:)
func NewValueFunctionWithName(name unsafe.Pointer) ValueFunction {
	rv := objc.Send[ValueFunction](objc.ID(getValueFunctionClass().class), objc.Sel("functionWithName:"), name)
	return rv
}


// Returns the value function object identified by the name.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAValueFunction/init(name:)
func (vc _ValueFunctionClass) FunctionWithName(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("functionWithName:"), name)
	return rv
}

// Returns the name of the value function.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAValueFunction/name
func (v_ ValueFunction) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("name"))
	return rv
}


