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

// An interface definition for the [InvocationOperation] class.
type IInvocationOperation interface {
	IOperation
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
// Alloc allocates a new instance without initialization.
func (ic _InvocationOperationClass) Alloc() InvocationOperation {
	rv := objc.Send[InvocationOperation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ic _InvocationOperationClass) New() InvocationOperation {
	rv := objc.Send[InvocationOperation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InvocationOperation) Init() InvocationOperation {
	rv := objc.Send[InvocationOperation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InvocationOperation) Autorelease() InvocationOperation {
	rv := objc.Send[InvocationOperation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInvocationOperation creates a new InvocationOperation instance.
func NewInvocationOperation() InvocationOperation {
	return invocationOperationClass.New()
}




