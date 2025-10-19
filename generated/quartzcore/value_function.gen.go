// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ValueFunction] class.
var valueFunctionClass = _ValueFunctionClass{objc.GetClass("CAValueFunction")}

type _ValueFunctionClass struct {
	class objc.Class
}

// An interface definition for the [ValueFunction] class.
type IValueFunction interface {
	objectivec.IObject
}

// An object that provides a flexible method of defining animated transformations. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return valueFunctionClass.New()
}




