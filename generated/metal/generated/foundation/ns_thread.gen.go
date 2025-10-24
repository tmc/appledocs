// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Thread] class.
var (
	ThreadClass     _ThreadClass
	ThreadClassOnce sync.Once
)

func getThreadClass() _ThreadClass {
	ThreadClassOnce.Do(func() {
		ThreadClass = _ThreadClass{objc.GetClass("NSThread")}
	})
	return ThreadClass
}

type _ThreadClass struct {
	class objc.Class
}

// An interface definition for the [Thread] class.
type IThread interface {
	objectivec.IObject
	// properties:
	NSAssertionHandlerKey() IString
	IsCancelled() bool /* primitive/slice/pointer. */
	SetIsCancelled(value bool /* primitive/slice/pointer. */)
	IsExecuting() bool /* primitive/slice/pointer. */
	SetIsExecuting(value bool /* primitive/slice/pointer. */)
	IsFinished() bool /* primitive/slice/pointer. */
	SetIsFinished(value bool /* primitive/slice/pointer. */)
	IsMainThread() bool /* primitive/slice/pointer. */
	SetIsMainThread(value bool /* primitive/slice/pointer. */)
	Name() IString
	SetName(value IString)
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	StackSize() int /* primitive/slice/pointer. */
	SetStackSize(value int /* primitive/slice/pointer. */)
	ThreadDictionary() IMutableDictionary
	SetThreadDictionary(value IMutableDictionary)
	ThreadPriority() float64 /* primitive/slice/pointer. */
	SetThreadPriority(value float64 /* primitive/slice/pointer. */)
	// methods:
}

// A thread of execution.
//
// Use this class when you want to have an Objective-C method run in its own thread of execution. Threads are especially useful when you need to perform a lengthy task, but don’t want it to block the execution of the rest of the application. In particular, you can use threads to avoid blocking the main thread of the application, which handles user interface and event-related actions. Threads can also be used to divide a large job into several smaller jobs, which can lead to performance increases on multi-core computers. The class supports semantics similar to those of for monitoring the runtime condition of a thread. You can use these semantics to cancel the execution of a thread or determine if the thread is still executing or has finished its task. Canceling a thread requires support from your thread code; see the description for for more information.


// A thread of execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread
type Thread struct {
	objectivec.Object
}

// ThreadFrom constructs a [Thread] from an unsafe.Pointer.
//
// A thread of execution.
func ThreadFrom(ptr unsafe.Pointer) Thread {
	return Thread{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ThreadClass) Alloc() Thread {
	rv := objc.Send[Thread](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ThreadClass) New() Thread {
	rv := objc.Send[Thread](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Thread) Init() Thread {
	rv := objc.Send[Thread](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Thread) Autorelease() Thread {
	rv := objc.Send[Thread](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewThread creates a new Thread instance.
func NewThread() Thread {
	return getThreadClass().New()
}



// A key with a corresponding value in the thread dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsassertionhandlerkey
func (t_ Thread) NSAssertionHandlerKey() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("NSAssertionHandlerKey"))
	return rv
}


// A Boolean value that indicates whether the receiver is cancelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/iscancelled
func (t_ Thread) IsCancelled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isCancelled"))
	return rv
}


// A Boolean value that indicates whether the receiver is cancelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/iscancelled
func (t_ Thread) SetIsCancelled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsCancelled:"), value)
}


// A Boolean value that indicates whether the receiver is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isexecuting
func (t_ Thread) IsExecuting() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isExecuting"))
	return rv
}


// A Boolean value that indicates whether the receiver is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isexecuting
func (t_ Thread) SetIsExecuting(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsExecuting:"), value)
}


// A Boolean value that indicates whether the receiver has finished execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isfinished
func (t_ Thread) IsFinished() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFinished"))
	return rv
}


// A Boolean value that indicates whether the receiver has finished execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isfinished
func (t_ Thread) SetIsFinished(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFinished:"), value)
}


// A Boolean value that indicates whether the receiver is the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/ismainthread-swift.property
func (t_ Thread) IsMainThread() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isMainThread"))
	return rv
}


// A Boolean value that indicates whether the receiver is the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/ismainthread-swift.property
func (t_ Thread) SetIsMainThread(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsMainThread:"), value)
}


// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/name
func (t_ Thread) Name() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("name"))
	return rv
}


// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/name
func (t_ Thread) SetName(value IString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/qualityofservice
func (t_ Thread) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](t_.ID, objc.Sel("qualityOfService"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/qualityofservice
func (t_ Thread) SetQualityOfService(value QualityOfService) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setQualityOfService:"), value)
}


// The stack size of the receiver, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/stacksize
func (t_ Thread) StackSize() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](t_.ID, objc.Sel("stackSize"))
	return rv
}


// The stack size of the receiver, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/stacksize
func (t_ Thread) SetStackSize(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStackSize:"), value)
}


// The thread object’s dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/threaddictionary
func (t_ Thread) ThreadDictionary() IMutableDictionary {
	rv := objc.Send[MutableDictionary](t_.ID, objc.Sel("threadDictionary"))
	return rv
}


// The thread object’s dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/threaddictionary
func (t_ Thread) SetThreadDictionary(value IMutableDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setThreadDictionary:"), value)
}


// The receiver’s priority
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/threadpriority
func (t_ Thread) ThreadPriority() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](t_.ID, objc.Sel("threadPriority"))
	return rv
}


// The receiver’s priority
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/threadpriority
func (t_ Thread) SetThreadPriority(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setThreadPriority:"), value)
}



