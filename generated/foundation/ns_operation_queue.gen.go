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
	AddBarrierBlock(barrier unsafe.Pointer)
	AddOperation(op IOperation)
	CancelAllOperations()
	OperationCount() uint
	Operations() []Operation
	IsReady() bool
	SetIsReady(value bool)
	QueuePriority() unsafe.Pointer
	SetQueuePriority(value unsafe.Pointer)
	IsSuspended() bool
	SetIsSuspended(value bool)
	MaxConcurrentOperationCount() int
	SetMaxConcurrentOperationCount(value int)
	Name() string
	SetName(value string)
	Progress() NSProgress
	SetProgress(value IProgress)
	QualityOfService() QualityOfService
	SetQualityOfService(value IQualityOfService)
	UnderlyingQueue() unsafe.Pointer
	SetUnderlyingQueue(value unsafe.Pointer)
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



// Invokes a block when the queue finishes all enqueued operations, and prevents subsequent operations from starting until the block has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addBarrierBlock(_:)

func (o_ OperationQueue) AddBarrierBlock(barrier unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addBarrierBlock:"), barrier)
}


// Adds the specified operation to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperation(_:)-64o8a

func (o_ OperationQueue) AddOperation(op IOperation) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addOperation:"), op)
}


// Cancels all queued and executing operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/cancelAllOperations()

func (o_ OperationQueue) CancelAllOperations() {
	objc.Send[objc.ID](o_.ID, objc.Sel("cancelAllOperations"))
}


// The number of operations currently in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/operationCount

func (o_ OperationQueue) OperationCount() uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("operationCount"))
	return rv
}


// The operations currently in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/operations

func (o_ OperationQueue) Operations() []Operation {
	rv := objc.Send[[]Operation](o_.ID, objc.Sel("operations"))
	return rv
}


// A Boolean value indicating whether the operation can be performed now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isready

func (o_ OperationQueue) IsReady() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isReady"))
	return rv
}


// A Boolean value indicating whether the operation can be performed now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isready

func (o_ OperationQueue) SetIsReady(value bool) {
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

func (o_ OperationQueue) IsSuspended() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isSuspended"))
	return rv
}


// A Boolean value indicating whether the queue is actively scheduling operations for execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/issuspended

func (o_ OperationQueue) SetIsSuspended(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsSuspended:"), value)
}


// The maximum number of queued operations that can run at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/maxconcurrentoperationcount

func (o_ OperationQueue) MaxConcurrentOperationCount() int {
	rv := objc.Send[int](o_.ID, objc.Sel("maxConcurrentOperationCount"))
	return rv
}


// The maximum number of queued operations that can run at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/maxconcurrentoperationcount

func (o_ OperationQueue) SetMaxConcurrentOperationCount(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setMaxConcurrentOperationCount:"), value)
}


// The name of the operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/name

func (o_ OperationQueue) Name() string {
	rv := objc.Send[string](o_.ID, objc.Sel("name"))
	return rv
}


// The name of the operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/name

func (o_ OperationQueue) SetName(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setName:"), objc.String(value))
}


// An object that represents the total progress of the operations executing in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/progress

func (o_ OperationQueue) Progress() NSProgress {
	rv := objc.Send[NSProgress](o_.ID, objc.Sel("progress"))
	return rv
}


// An object that represents the total progress of the operations executing in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/progress

func (o_ OperationQueue) SetProgress(value IProgress) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setProgress:"), value)
}


// The default service level to apply to operations that the queue invokes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/qualityofservice

func (o_ OperationQueue) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](o_.ID, objc.Sel("qualityOfService"))
	return rv
}


// The default service level to apply to operations that the queue invokes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/qualityofservice

func (o_ OperationQueue) SetQualityOfService(value IQualityOfService) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQualityOfService:"), value)
}


// The dispatch queue that the operation queue uses to invoke operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/underlyingqueue

func (o_ OperationQueue) UnderlyingQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("underlyingQueue"))
	return rv
}


// The dispatch queue that the operation queue uses to invoke operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/underlyingqueue

func (o_ OperationQueue) SetUnderlyingQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUnderlyingQueue:"), value)
}



