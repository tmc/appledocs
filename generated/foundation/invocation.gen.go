// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Invocation] class.
var invocationClass = _InvocationClass{objc.GetClass("NSInvocation")}

type _InvocationClass struct {
	class objc.Class
}

// An interface definition for the [Invocation] class.
type IInvocation interface {
	objectivec.IObject
}

// An Objective-C message rendered as an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation

type Invocation struct {
	objectivec.Object
}

// InvocationFrom constructs a [Invocation] from an unsafe.Pointer.
//
// An Objective-C message rendered as an object.
func InvocationFrom(ptr unsafe.Pointer) Invocation {
	return Invocation{objectivec.Object{objc.ID(ptr)}}
}



