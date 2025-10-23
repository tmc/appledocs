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
	OperationClass     _OperationClass
	OperationClassOnce sync.Once
)

func getOperationClass() _OperationClass {
	OperationClassOnce.Do(func() {
		OperationClass = _OperationClass{objc.GetClass("NSOperation")}
	})
	return OperationClass
}

type _OperationClass struct {
	class objc.Class
}

// An interface definition for the [Operation] class.
type IOperation interface {
	objectivec.IObject
	// properties:
	CompletionBlock() unsafe.Pointer
	SetCompletionBlock(value unsafe.Pointer)
	Dependencies() IOperation
	SetDependencies(value IOperation)
	IsAsynchronous() bool /* primitive/slice/pointer */
	SetIsAsynchronous(value bool /* primitive/slice/pointer */)
	IsCancelled() bool /* primitive/slice/pointer */
	SetIsCancelled(value bool /* primitive/slice/pointer */)
	IsConcurrent() bool /* primitive/slice/pointer */
	SetIsConcurrent(value bool /* primitive/slice/pointer */)
	IsExecuting() bool /* primitive/slice/pointer */
	SetIsExecuting(value bool /* primitive/slice/pointer */)
	IsFinished() bool /* primitive/slice/pointer */
	SetIsFinished(value bool /* primitive/slice/pointer */)
	IsReady() bool /* primitive/slice/pointer */
	SetIsReady(value bool /* primitive/slice/pointer */)
	Name() string /* primitive/slice/pointer */
	SetName(value string /* primitive/slice/pointer */)
	QualityOfService() unsafe.Pointer
	SetQualityOfService(value unsafe.Pointer)
	QueuePriority() unsafe.Pointer
	SetQueuePriority(value unsafe.Pointer)
	ThreadPriority() float64 /* primitive/slice/pointer */
	SetThreadPriority(value float64 /* primitive/slice/pointer */)
	// methods:
}

// An abstract class that represents the code and data associated with a single task.
//
// Because the class is an abstract class, you do not use it directly but instead subclass or use one of the system-defined subclasses ( or ) to perform the actual task. Despite being abstract, the base implementation of does include significant logic to coordinate the safe execution of your task. The presence of this built-in logic allows you to focus on the actual implementation of your task, rather than on the glue code needed to ensure it works correctly with other system objects. An operation object is a single-shot object—that is, it executes its task once and cannot be used to execute it again. You typically execute operations by adding them to an operation queue (an instance of the class). An operation queue executes its operations either directly, by running them on secondary threads, or indirectly using the library (also known as Grand Central Dispatch). For more information about how queues execute operations, see . If you do not want to use an operation queue, you can execute an operation yourself by calling its method directly from your code. Executing operations manually does put more of a burden on your code, because starting an operation that is not in the ready state triggers an exception. The property reports on the operation’s readiness.


// An abstract class that represents the code and data associated with a single task.
//
// [Full Topic]
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



// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/completionblock
func (o_ Operation) CompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("completionBlock"))
	return rv
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/completionblock
func (o_ Operation) SetCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCompletionBlock:"), value)
}


// An array of the operation objects that must finish executing before the current object can begin executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/dependencies
func (o_ Operation) Dependencies() IOperation {
	rv := objc.Send[Operation](o_.ID, objc.Sel("dependencies"))
	return rv
}


// An array of the operation objects that must finish executing before the current object can begin executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/dependencies
func (o_ Operation) SetDependencies(value IOperation) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDependencies:"), value)
}


// A Boolean value indicating whether the operation executes its task asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isasynchronous
func (o_ Operation) IsAsynchronous() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAsynchronous"))
	return rv
}


// A Boolean value indicating whether the operation executes its task asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isasynchronous
func (o_ Operation) SetIsAsynchronous(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsAsynchronous:"), value)
}


// A Boolean value indicating whether the operation has been cancelled
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/iscancelled
func (o_ Operation) IsCancelled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isCancelled"))
	return rv
}


// A Boolean value indicating whether the operation has been cancelled
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/iscancelled
func (o_ Operation) SetIsCancelled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsCancelled:"), value)
}


// A Boolean value indicating whether the operation executes its task asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isconcurrent
func (o_ Operation) IsConcurrent() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isConcurrent"))
	return rv
}


// A Boolean value indicating whether the operation executes its task asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isconcurrent
func (o_ Operation) SetIsConcurrent(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsConcurrent:"), value)
}


// A Boolean value indicating whether the operation is currently executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isexecuting
func (o_ Operation) IsExecuting() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isExecuting"))
	return rv
}


// A Boolean value indicating whether the operation is currently executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isexecuting
func (o_ Operation) SetIsExecuting(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsExecuting:"), value)
}


// A Boolean value indicating whether the operation has finished executing its task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isfinished
func (o_ Operation) IsFinished() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isFinished"))
	return rv
}


// A Boolean value indicating whether the operation has finished executing its task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isfinished
func (o_ Operation) SetIsFinished(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsFinished:"), value)
}


// A Boolean value indicating whether the operation can be performed now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isready
func (o_ Operation) IsReady() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isReady"))
	return rv
}


// A Boolean value indicating whether the operation can be performed now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isready
func (o_ Operation) SetIsReady(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsReady:"), value)
}


// The name of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/name
func (o_ Operation) Name() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](o_.ID, objc.Sel("name"))
	return rv
}


// The name of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/name
func (o_ Operation) SetName(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setName:"), objc.String(value))
}


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/qualityofservice
func (o_ Operation) QualityOfService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("qualityOfService"))
	return rv
}


// The relative amount of importance for granting system resources to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/qualityofservice
func (o_ Operation) SetQualityOfService(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQualityOfService:"), value)
}


// The execution priority of the operation in an operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/queuepriority-swift.property
func (o_ Operation) QueuePriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("queuePriority"))
	return rv
}


// The execution priority of the operation in an operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/queuepriority-swift.property
func (o_ Operation) SetQueuePriority(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQueuePriority:"), value)
}


// The thread priority to use when executing the operation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/threadpriority
func (o_ Operation) ThreadPriority() float64 /* primitive/slice/pointer */ {
	rv := objc.Send[float64](o_.ID, objc.Sel("threadPriority"))
	return rv
}


// The thread priority to use when executing the operation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/threadpriority
func (o_ Operation) SetThreadPriority(value float64 /* primitive/slice/pointer */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setThreadPriority:"), value)
}



