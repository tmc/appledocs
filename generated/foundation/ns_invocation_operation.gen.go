// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InvocationOperation] class.
var (
	InvocationOperationClass     _InvocationOperationClass
	InvocationOperationClassOnce sync.Once
)

func getInvocationOperationClass() _InvocationOperationClass {
	InvocationOperationClassOnce.Do(func() {
		InvocationOperationClass = _InvocationOperationClass{objc.GetClass("NSInvocationOperation")}
	})
	return InvocationOperationClass
}

type _InvocationOperationClass struct {
	class objc.Class
}

// An interface definition for the [InvocationOperation] class.
type IInvocationOperation interface {
	IOperation
}

// An operation that manages the execution of a single encapsulated task specified as an invocation.
//
// The class is a concrete subclass of that you use to initiate an operation that consists of invoking a selector on a specified object. This class implements a non-concurrent operation. For more information on concurrent versus non-concurrent operations, see .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getInvocationOperationClass().New()
}


// Returns an object initialized with the specified target and selector.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/initWithTarget:selector:object:
func NewInvocationOperationWithTargetSelectorObject(target objc.ID, sel objc.SEL, arg objc.ID) InvocationOperation {
	instance := getInvocationOperationClass().Alloc()
	rv := objc.Send[InvocationOperation](instance.ID, objc.Sel("initWithTarget:selector:object:"), target, sel, arg)
	rv.Autorelease()
	return rv
}

// Returns an object initialized with the specified invocation object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/initWithInvocation:
func NewInvocationOperationWithInvocation(inv unsafe.Pointer) InvocationOperation {
	instance := getInvocationOperationClass().Alloc()
	rv := objc.Send[InvocationOperation](instance.ID, objc.Sel("initWithInvocation:"), inv)
	rv.Autorelease()
	return rv
}


// The receiver’s invocation object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/invocation
func (i_ InvocationOperation) Invocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("invocation"))
	return rv
}

// The result of the invocation or method.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/result
func (i_ InvocationOperation) Result() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("result"))
	return rv
}


