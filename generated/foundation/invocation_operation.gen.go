// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InvocationOperation] class.
var invocationOperationClass = _InvocationOperationClass{objc.GetClass("NSInvocationOperation")}

type _InvocationOperationClass struct {
	class objc.Class
}

// An operation that manages the execution of a single encapsulated task specified as an invocation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation

type InvocationOperation struct {
	Operation
}

// InvocationOperationFrom constructs a [InvocationOperation] from an unsafe.Pointer.
//
// An operation that manages the execution of a single encapsulated task specified as an invocation.
func InvocationOperationFrom(ptr unsafe.Pointer) InvocationOperation {
	return InvocationOperation{
		Operation: OperationFrom(ptr),
	}
}



