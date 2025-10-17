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



