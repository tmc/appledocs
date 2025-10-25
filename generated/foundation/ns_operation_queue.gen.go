// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSOperationQueue */


/* debug [class_header]: Header for NSOperationQueue */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OperationQueue */
// An interface definition for the [OperationQueue] class.
type IOperationQueue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OperationQueue */
	// properties:
	Suspended() bool
	SetSuspended(value bool)
	MaxConcurrentOperationCount() int
	SetMaxConcurrentOperationCount(value int)
	Name() IString
	SetName(value IString)
	OperationCount() uint
	Operations() []objc.IObject /* cross-framework: Operation */
	Progress() IProgress
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	UnderlyingQueue() objectivec.IObject
	SetUnderlyingQueue(value objectivec.IObject)
	IsReady() bool
	SetIsReady(value bool)
	QueuePriority() objectivec.IObject
	SetQueuePriority(value objectivec.IObject)
	IsSuspended() bool
	SetIsSuspended(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OperationQueue */
	// methods:
	AddBarrierBlock(barrier unsafe.Pointer)
	AddOperationWithBlock(block unsafe.Pointer)
	AddOperation(op objc.IObject /* cross-framework: Operation */)
	AddOperationsWaitUntilFinished(ops []objc.IObject /* cross-framework: Operation */, wait bool)
	CancelAllOperations()
	WaitUntilAllOperationsAreFinished()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OperationQueue */
// Alloc allocates a new instance without initialization.
func (oc _OperationQueueClass) Alloc() OperationQueue {
	rv := objc.Send[OperationQueue](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OperationQueue */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OperationQueue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OperationQueue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OperationQueue */

// Returns the operation queue that launched the current operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/current
func (oc _OperationQueueClass) CurrentQueue() OperationQueue {
	rv := objc.Send[OperationQueue](objc.ID(oc.class), objc.Sel("currentQueue"))
	return rv
}/* debug [class_properties_class/property]: currentQueue */

// Returns the operation queue associated with the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/main
func (oc _OperationQueueClass) MainQueue() OperationQueue {
	rv := objc.Send[OperationQueue](objc.ID(oc.class), objc.Sel("mainQueue"))
	return rv
}/* debug [class_properties_class/property]: mainQueue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OperationQueue */

// Invokes a block when the queue finishes all enqueued operations, and prevents subsequent operations from starting until the block has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addBarrierBlock(_:)
func (o_ OperationQueue) AddBarrierBlock(barrier unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addBarrierBlock:"), barrier)
}/* debug [instance_methods/method]: AddBarrierBlock */


// Wraps the specified block in an operation and adds it to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperation(_:)-5s294
func (o_ OperationQueue) AddOperationWithBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addOperationWithBlock:"), block)
}/* debug [instance_methods/method]: AddOperationWithBlock */


// Adds the specified operation to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperation(_:)-64o8a
func (o_ OperationQueue) AddOperation(op objc.IObject /* cross-framework: Operation */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addOperation:"), op)
}/* debug [instance_methods/method]: AddOperation */


// Adds the specified operations to the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperations(_:waitUntilFinished:)
func (o_ OperationQueue) AddOperationsWaitUntilFinished(ops []objc.IObject /* cross-framework: Operation */, wait bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addOperations:waitUntilFinished:"), ops, wait)
}/* debug [instance_methods/method]: AddOperationsWaitUntilFinished */


// Cancels all queued and executing operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/cancelAllOperations()
func (o_ OperationQueue) CancelAllOperations() {
	objc.Send[objc.ID](o_.ID, objc.Sel("cancelAllOperations"))
}/* debug [instance_methods/method]: CancelAllOperations */


// Blocks the current thread until all the receiver’s queued and executing operations finish executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/waitUntilAllOperationsAreFinished()
func (o_ OperationQueue) WaitUntilAllOperationsAreFinished() {
	objc.Send[objc.ID](o_.ID, objc.Sel("waitUntilAllOperationsAreFinished"))
}/* debug [instance_methods/method]: WaitUntilAllOperationsAreFinished */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OperationQueue */

// Returns the operation queue that launched the current operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/current
func (o_ OperationQueue) CurrentQueue() IOperationQueue {
	rv := objc.Send[OperationQueue](o_.ID, objc.Sel("currentQueue"))
	return rv
}/* debug [instance_properties/getter]: currentQueue */


// A Boolean value indicating whether the queue is actively scheduling operations for execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/isSuspended
func (o_ OperationQueue) Suspended() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("suspended"))
	return rv
}/* debug [instance_properties/getter]: suspended */


// A Boolean value indicating whether the queue is actively scheduling operations for execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/isSuspended
func (o_ OperationQueue) SetSuspended(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSuspended:"), value)
}/* debug [instance_properties/setter]: suspended */


// Returns the operation queue associated with the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/main
func (o_ OperationQueue) MainQueue() IOperationQueue {
	rv := objc.Send[OperationQueue](o_.ID, objc.Sel("mainQueue"))
	return rv
}/* debug [instance_properties/getter]: mainQueue */


// The maximum number of queued operations that can run at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/maxConcurrentOperationCount
func (o_ OperationQueue) MaxConcurrentOperationCount() int {
	rv := objc.Send[int](o_.ID, objc.Sel("maxConcurrentOperationCount"))
	return rv
}/* debug [instance_properties/getter]: maxConcurrentOperationCount */


// The maximum number of queued operations that can run at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/maxConcurrentOperationCount
func (o_ OperationQueue) SetMaxConcurrentOperationCount(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setMaxConcurrentOperationCount:"), value)
}/* debug [instance_properties/setter]: maxConcurrentOperationCount */


// The name of the operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/name
func (o_ OperationQueue) Name() IString {
	rv := objc.Send[String](o_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/name
func (o_ OperationQueue) SetName(value IString) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The number of operations currently in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/operationCount
func (o_ OperationQueue) OperationCount() uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("operationCount"))
	return rv
}/* debug [instance_properties/getter]: operationCount */


// The operations currently in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/operations
func (o_ OperationQueue) Operations() []objc.IObject /* cross-framework: Operation */ {
	rv := objc.Send[[]Operation](o_.ID, objc.Sel("operations"))
	// Slice of concrete type to slice of interface - needs conversion
	result := make([]objc.IObject /* cross-framework: Operation */, len(rv))
	for i, v := range rv {
		result[i] = v
	}
	return result
}/* debug [instance_properties/getter]: operations */


// An object that represents the total progress of the operations executing in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/progress
func (o_ OperationQueue) Progress() IProgress {
	rv := objc.Send[Progress](o_.ID, objc.Sel("progress"))
	return rv
}/* debug [instance_properties/getter]: progress */


// The default service level to apply to operations that the queue invokes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/qualityOfService
func (o_ OperationQueue) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](o_.ID, objc.Sel("qualityOfService"))
	return rv
}/* debug [instance_properties/getter]: qualityOfService */


// The default service level to apply to operations that the queue invokes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/qualityOfService
func (o_ OperationQueue) SetQualityOfService(value QualityOfService) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQualityOfService:"), value)
}/* debug [instance_properties/setter]: qualityOfService */


// The dispatch queue that the operation queue uses to invoke operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/underlyingQueue
func (o_ OperationQueue) UnderlyingQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("underlyingQueue"))
	return rv
}/* debug [instance_properties/getter]: underlyingQueue */


// The dispatch queue that the operation queue uses to invoke operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/underlyingQueue
func (o_ OperationQueue) SetUnderlyingQueue(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUnderlyingQueue:"), value)
}/* debug [instance_properties/setter]: underlyingQueue */


// A Boolean value indicating whether the operation can be performed now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isready
func (o_ OperationQueue) IsReady() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isReady"))
	return rv
}/* debug [instance_properties/getter]: isReady */


// A Boolean value indicating whether the operation can be performed now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/isready
func (o_ OperationQueue) SetIsReady(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsReady:"), value)
}/* debug [instance_properties/setter]: isReady */


// The execution priority of the operation in an operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/queuepriority-swift.property
func (o_ OperationQueue) QueuePriority() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("queuePriority"))
	return rv
}/* debug [instance_properties/getter]: queuePriority */


// The execution priority of the operation in an operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operation/queuepriority-swift.property
func (o_ OperationQueue) SetQueuePriority(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQueuePriority:"), value)
}/* debug [instance_properties/setter]: queuePriority */


// A Boolean value indicating whether the queue is actively scheduling operations for execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/issuspended
func (o_ OperationQueue) IsSuspended() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isSuspended"))
	return rv
}/* debug [instance_properties/getter]: isSuspended */


// A Boolean value indicating whether the queue is actively scheduling operations for execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/operationqueue/issuspended
func (o_ OperationQueue) SetIsSuspended(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsSuspended:"), value)
}/* debug [instance_properties/setter]: isSuspended */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSOperationQueue */



