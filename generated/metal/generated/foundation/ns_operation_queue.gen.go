// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OperationQueue] class.
var (
	OperationQueueClass     _OperationQueueClass
	OperationQueueClassOnce sync.Once
)

func getOperationQueueClass() _OperationQueueClass {
	OperationQueueClassOnce.Do(func() {
		OperationQueueClass = _OperationQueueClass{objc.GetClass("NSOperationQueue")}
	})
	return OperationQueueClass
}

type _OperationQueueClass struct {
	class objc.Class
}

// An interface definition for the [OperationQueue] class.
type IOperationQueue interface {
	objectivec.IObject
	// properties:
	Suspended() bool /* primitive/slice/pointer. */
	SetSuspended(value bool /* primitive/slice/pointer. */)
	MaxConcurrentOperationCount() int /* primitive/slice/pointer. */
	SetMaxConcurrentOperationCount(value int /* primitive/slice/pointer. */)
	Name() IString
	SetName(value IString)
	OperationCount() uint /* primitive/slice/pointer. */
	Operations() []Operation /* primitive/slice/pointer. */
	Progress() IProgress
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	UnderlyingQueue() unsafe.Pointer
	SetUnderlyingQueue(value unsafe.Pointer)
	IsReady() bool /* primitive/slice/pointer. */
	SetIsReady(value bool /* primitive/slice/pointer. */)
	QueuePriority() unsafe.Pointer
	SetQueuePriority(value unsafe.Pointer)
	IsSuspended() bool /* primitive/slice/pointer. */
	SetIsSuspended(value bool /* primitive/slice/pointer. */)
	// methods:
	AddBarrierBlock(barrier unsafe.Pointer)
	AddOperationWithBlock(block unsafe.Pointer)
	AddOperation(op IOperation)
	AddOperationsWaitUntilFinished(ops []Operation /* primitive/slice/pointer. */, wait bool /* primitive/slice/pointer. */)
	CancelAllOperations()
	WaitUntilAllOperationsAreFinished()
}

// A queue that regulates the execution of operations.
//
// An operation queue invokes its queued objects based on their priority and readiness. After you add an operation to a queue, it remains in the queue until the operation finishes its task. You can’t directly remove an operation from a queue after you add it. For more information about using operation queues, see the .


// A queue that regulates the execution of operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue
type OperationQueue struct {
	objectivec.Object
}

// OperationQueueFrom constructs a [OperationQueue] from an unsafe.Pointer.
//
// A queue that regulates the execution of operations.
func OperationQueueFrom(ptr unsafe.Pointer) OperationQueue {
	return OperationQueue{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OperationQueueClass) Alloc() OperationQueue {
	rv := objc.Send[OperationQueue](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OperationQueueClass) New() OperationQueue {
	rv := objc.Send[OperationQueue](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OperationQueue) Init() OperationQueue {
	rv := objc.Send[OperationQueue](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OperationQueue) Autorelease() OperationQueue {
	rv := objc.Send[OperationQueue](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOperationQueue creates a new OperationQueue instance.
func NewOperationQueue() OperationQueue {
	return getOperationQueueClass().New()
}



// Returns the operation queue that launched the current operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/current
func (oc _OperationQueueClass) CurrentQueue() OperationQueue {
	rv := objc.Send[OperationQueue](objc.ID(oc.class), objc.Sel("currentQueue"))
	return rv
}

// Returns the operation queue associated with the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/main
func (oc _OperationQueueClass) MainQueue() OperationQueue {
	rv := objc.Send[OperationQueue](objc.ID(oc.class), objc.Sel("mainQueue"))
	return rv
}

// Invokes a block when the queue finishes all enqueued operations, and prevents subsequent operations from starting until the block has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addBarrierBlock(_:)
func (o_ OperationQueue) AddBarrierBlock(barrier unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addBarrierBlock:"), barrier)
}


// Wraps the specified block in an operation and adds it to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperation(_:)-5s294
func (o_ OperationQueue) AddOperationWithBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addOperationWithBlock:"), block)
}


// Adds the specified operation to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperation(_:)-64o8a
func (o_ OperationQueue) AddOperation(op IOperation) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addOperation:"), op)
}


// Adds the specified operations to the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperations(_:waitUntilFinished:)
func (o_ OperationQueue) AddOperationsWaitUntilFinished(ops []Operation /* primitive/slice/pointer. */, wait bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addOperations:waitUntilFinished:"), ops, wait)
}


// Cancels all queued and executing operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/cancelAllOperations()
func (o_ OperationQueue) CancelAllOperations() {
	objc.Send[objc.ID](o_.ID, objc.Sel("cancelAllOperations"))
}


// Blocks the current thread until all the receiver’s queued and executing operations finish executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/waitUntilAllOperationsAreFinished()
func (o_ OperationQueue) WaitUntilAllOperationsAreFinished() {
	objc.Send[objc.ID](o_.ID, objc.Sel("waitUntilAllOperationsAreFinished"))
}


// Returns the operation queue that launched the current operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/current
func (o_ OperationQueue) CurrentQueue() IOperationQueue {
	rv := objc.Send[OperationQueue](o_.ID, objc.Sel("currentQueue"))
	return rv
}


// A Boolean value indicating whether the queue is actively scheduling operations for execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/isSuspended
func (o_ OperationQueue) Suspended() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("suspended"))
	return rv
}


// A Boolean value indicating whether the queue is actively scheduling operations for execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/isSuspended
func (o_ OperationQueue) SetSuspended(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSuspended:"), value)
}


// Returns the operation queue associated with the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/main
func (o_ OperationQueue) MainQueue() IOperationQueue {
	rv := objc.Send[OperationQueue](o_.ID, objc.Sel("mainQueue"))
	return rv
}


// The maximum number of queued operations that can run at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/maxConcurrentOperationCount
func (o_ OperationQueue) MaxConcurrentOperationCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](o_.ID, objc.Sel("maxConcurrentOperationCount"))
	return rv
}


// The maximum number of queued operations that can run at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/maxConcurrentOperationCount
func (o_ OperationQueue) SetMaxConcurrentOperationCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setMaxConcurrentOperationCount:"), value)
}


// The name of the operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/name
func (o_ OperationQueue) Name() IString {
	rv := objc.Send[String](o_.ID, objc.Sel("name"))
	return rv
}


// The name of the operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/name
func (o_ OperationQueue) SetName(value IString) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setName:"), value)
}


// The number of operations currently in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/operationCount
func (o_ OperationQueue) OperationCount() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](o_.ID, objc.Sel("operationCount"))
	return rv
}


// The operations currently in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/operations
func (o_ OperationQueue) Operations() []Operation /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Operation](o_.ID, objc.Sel("operations"))
	return rv
}


// An object that represents the total progress of the operations executing in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/progress
func (o_ OperationQueue) Progress() IProgress {
	rv := objc.Send[Progress](o_.ID, objc.Sel("progress"))
	return rv
}


// The default service level to apply to operations that the queue invokes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/qualityOfService
func (o_ OperationQueue) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](o_.ID, objc.Sel("qualityOfService"))
	return rv
}


// The default service level to apply to operations that the queue invokes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/qualityOfService
func (o_ OperationQueue) SetQualityOfService(value QualityOfService) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQualityOfService:"), value)
}


// The dispatch queue that the operation queue uses to invoke operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/underlyingQueue
func (o_ OperationQueue) UnderlyingQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("underlyingQueue"))
	return rv
}


// The dispatch queue that the operation queue uses to invoke operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/underlyingQueue
func (o_ OperationQueue) SetUnderlyingQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUnderlyingQueue:"), value)
}


// A Boolean value indicating whether the operation can be performed now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isready
func (o_ OperationQueue) IsReady() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isReady"))
	return rv
}


// A Boolean value indicating whether the operation can be performed now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isready
func (o_ OperationQueue) SetIsReady(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsReady:"), value)
}


// The execution priority of the operation in an operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/queuepriority-swift.property
func (o_ OperationQueue) QueuePriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("queuePriority"))
	return rv
}


// The execution priority of the operation in an operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/queuepriority-swift.property
func (o_ OperationQueue) SetQueuePriority(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQueuePriority:"), value)
}


// A Boolean value indicating whether the queue is actively scheduling operations for execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/issuspended
func (o_ OperationQueue) IsSuspended() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isSuspended"))
	return rv
}


// A Boolean value indicating whether the queue is actively scheduling operations for execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/issuspended
func (o_ OperationQueue) SetIsSuspended(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsSuspended:"), value)
}



