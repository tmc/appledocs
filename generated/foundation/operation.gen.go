// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Operation] class.
var (
	operationClass     _OperationClass
	operationClassOnce sync.Once
)

func getOperationClass() _OperationClass {
	operationClassOnce.Do(func() {
		operationClass = _OperationClass{objc.GetClass("NSOperation")}
	})
	return operationClass
}

type _OperationClass struct {
	class objc.Class
}

// An interface definition for the [Operation] class.
type IOperation interface {
	objectivec.IObject
	AddDependency(op unsafe.Pointer)
	Cancel()
	Main()
	RemoveDependency(op unsafe.Pointer)
	Start()
	WaitUntilFinished()
}

// An abstract class that represents the code and data associated with a single task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation
type Operation struct {
	objectivec.Object
}

// OperationFrom constructs a [Operation] from an unsafe.Pointer.
//
// An abstract class that represents the code and data associated with a single task.
func OperationFrom(ptr unsafe.Pointer) Operation {
	return Operation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OperationClass) Alloc() Operation {
	rv := objc.Send[Operation](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OperationClass) New() Operation {
	rv := objc.Send[Operation](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Operation) Init() Operation {
	rv := objc.Send[Operation](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Operation) Autorelease() Operation {
	rv := objc.Send[Operation](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOperation creates a new Operation instance.
func NewOperation() Operation {
	return getOperationClass().New()
}


// Makes the receiver dependent on the completion of the specified operation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/addDependency(_:)
func (o_ Operation) AddDependency(op unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addDependency:"), op)
}
// Advises the operation object that it should stop executing its task. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/cancel()
func (o_ Operation) Cancel() {
	objc.Send[objc.ID](o_.ID, objc.Sel("cancel"))
}
// Performs the receiver’s non-concurrent task. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/main()
func (o_ Operation) Main() {
	objc.Send[objc.ID](o_.ID, objc.Sel("main"))
}
// Removes the receiver’s dependence on the specified operation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/removeDependency(_:)
func (o_ Operation) RemoveDependency(op unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeDependency:"), op)
}
// Begins the execution of the operation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/start()
func (o_ Operation) Start() {
	objc.Send[objc.ID](o_.ID, objc.Sel("start"))
}
// Blocks execution of the current thread until the operation object finishes its task. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/waitUntilFinished()
func (o_ Operation) WaitUntilFinished() {
	objc.Send[objc.ID](o_.ID, objc.Sel("waitUntilFinished"))
}


