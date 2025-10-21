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
	Cancel()
	Main()
	Start()
}

// A thread of execution.
//
// Use this class when you want to have an Objective-C method run in its own thread of execution. Threads are especially useful when you need to perform a lengthy task, but don’t want it to block the execution of the rest of the application. In particular, you can use threads to avoid blocking the main thread of the application, which handles user interface and event-related actions. Threads can also be used to divide a large job into several smaller jobs, which can lead to performance increases on multi-core computers. The class supports semantics similar to those of for monitoring the runtime condition of a thread. You can use these semantics to cancel the execution of a thread or determine if the thread is still executing or has finished its task. Canceling a thread requires support from your thread code; see the description for for more information.
//
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




// Returns an object initialized with the given arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/init(target:selector:object:)
func NewThreadWithTargetSelectorObject(target objc.ID, selector objc.SEL, argument objc.ID) Thread {
	instance := getThreadClass().Alloc()
	rv := objc.Send[Thread](instance.ID, objc.Sel("initWithTarget:selector:object:"), target, selector, argument)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/detachNewThread(_:)
func (tc _ThreadClass) DetachNewThreadWithBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("detachNewThreadWithBlock:"), block)
}

// Detaches a new thread and uses the specified selector as the thread entry point.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/detachNewThreadSelector(_:toTarget:with:)
func (tc _ThreadClass) DetachNewThreadSelectorToTargetWithObject(selector objc.SEL, target objc.ID, argument objc.ID) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("detachNewThreadSelector:toTarget:withObject:"), selector, target, argument)
}

// Terminates the current thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/exit()
func (tc _ThreadClass) Exit() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("exit"))
}

// Returns whether the application is multithreaded.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isMultiThreaded()
func (tc _ThreadClass) IsMultiThreaded() bool {
	rv := objc.Send[bool](objc.ID(tc.class), objc.Sel("isMultiThreaded"))
	return rv
}

// Sets the current thread’s priority.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/setThreadPriority(_:)
func (tc _ThreadClass) SetThreadPriority(p unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(tc.class), objc.Sel("setThreadPriority:"), p)
	return rv
}

// Sleeps the thread for a given time interval.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/sleep(forTimeInterval:)
func (tc _ThreadClass) SleepForTimeInterval(ti TimeInterval) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("sleepForTimeInterval:"), ti)
}

// Blocks the current thread until the time specified.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/sleep(until:)
func (tc _ThreadClass) SleepUntilDate(date unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("sleepUntilDate:"), date)
}

// Returns the current thread’s priority.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/threadPriority()
func (tc _ThreadClass) ThreadPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("threadPriority"))
	return rv
}

// Returns an array containing the call stack return addresses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/callStackReturnAddresses
func (tc _ThreadClass) CallStackReturnAddresses() []Number {
	rv := objc.Send[[]Number](objc.ID(tc.class), objc.Sel("callStackReturnAddresses"))
	return rv
}
// Returns an array containing the call stack symbols.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/callStackSymbols
func (tc _ThreadClass) CallStackSymbols() []string {
	rv := objc.Send[[]string](objc.ID(tc.class), objc.Sel("callStackSymbols"))
	return rv
}
// Returns the thread object representing the current thread of execution.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/current
func (tc _ThreadClass) CurrentThread() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("currentThread"))
	return rv
}
// Returns the object representing the main thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/main
func (tc _ThreadClass) MainThread() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("mainThread"))
	return rv
}
// Changes the cancelled state of the receiver to indicate that it should exit.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/cancel()
func (t_ Thread) Cancel() {
	objc.Send[objc.ID](t_.ID, objc.Sel("cancel"))
}

// The main entry point routine for the thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/main()
func (t_ Thread) Main() {
	objc.Send[objc.ID](t_.ID, objc.Sel("main"))
}

// Starts the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/start()
func (t_ Thread) Start() {
	objc.Send[objc.ID](t_.ID, objc.Sel("start"))
}

// The thread object’s dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/threaddictionary
func (t_ Thread) ThreadDictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("threadDictionary"))
	return rv
}


// SetThreadDictionary sets the value of the threadDictionary property.
// The thread object’s dictionary.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/threaddictionary
func (t_ Thread) SetThreadDictionary(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setThreadDictionary:"), value)
}

// A key with a corresponding value in the thread dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsassertionhandlerkey
func (t_ Thread) NSAssertionHandlerKey() string {
	rv := objc.Send[string](t_.ID, objc.Sel("NSAssertionHandlerKey"))
	return rv
}

// A Boolean value that indicates whether the receiver is executing.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isexecuting
func (t_ Thread) IsExecuting() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isExecuting"))
	return rv
}


// SetIsExecuting sets the value of the isExecuting property.
// A Boolean value that indicates whether the receiver is executing.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isexecuting
func (t_ Thread) SetIsExecuting(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsExecuting:"), value)
}

// A Boolean value that indicates whether the receiver is cancelled.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/iscancelled
func (t_ Thread) IsCancelled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isCancelled"))
	return rv
}


// SetIsCancelled sets the value of the isCancelled property.
// A Boolean value that indicates whether the receiver is cancelled.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/iscancelled
func (t_ Thread) SetIsCancelled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsCancelled:"), value)
}

// A Boolean value that indicates whether the receiver has finished execution.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isfinished
func (t_ Thread) IsFinished() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFinished"))
	return rv
}


// SetIsFinished sets the value of the isFinished property.
// A Boolean value that indicates whether the receiver has finished execution.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isfinished
func (t_ Thread) SetIsFinished(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFinished:"), value)
}

// Returns an array containing the call stack return addresses.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/callStackReturnAddresses
func (t_ Thread) CallStackReturnAddresses() []Number {
	rv := objc.Send[[]Number](t_.ID, objc.Sel("callStackReturnAddresses"))
	return rv
}

// Returns an array containing the call stack symbols.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/callStackSymbols
func (t_ Thread) CallStackSymbols() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("callStackSymbols"))
	return rv
}

// Returns the thread object representing the current thread of execution.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/current
func (t_ Thread) CurrentThread() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("currentThread"))
	return rv
}

// A Boolean value that indicates whether the receiver is cancelled.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isCancelled
func (t_ Thread) Cancelled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("cancelled"))
	return rv
}

// A Boolean value that indicates whether the receiver is executing.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isExecuting
func (t_ Thread) Executing() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("executing"))
	return rv
}

// A Boolean value that indicates whether the receiver has finished execution.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isFinished
func (t_ Thread) Finished() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("finished"))
	return rv
}

// A Boolean value that indicates whether the receiver is the main thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isMainThread-swift.property
func (t_ Thread) IsMainThread() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isMainThread"))
	return rv
}

// Returns the object representing the main thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/main
func (t_ Thread) MainThread() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("mainThread"))
	return rv
}

// The name of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/name
func (t_ Thread) Name() string {
	rv := objc.Send[string](t_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/name
func (t_ Thread) SetName(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/qualityOfService
func (t_ Thread) QualityOfService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("qualityOfService"))
	return rv
}


// SetQualityOfService sets the value of the qualityOfService property.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/qualityOfService
func (t_ Thread) SetQualityOfService(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setQualityOfService:"), value)
}

// The stack size of the receiver, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/stackSize
func (t_ Thread) StackSize() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("stackSize"))
	return rv
}


// SetStackSize sets the value of the stackSize property.
// The stack size of the receiver, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/stackSize
func (t_ Thread) SetStackSize(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStackSize:"), value)
}

// The receiver’s priority
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/threadPriority
func (t_ Thread) ThreadPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("threadPriority"))
	return rv
}


// SetThreadPriority sets the value of the threadPriority property.
// The receiver’s priority

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/threadPriority
func (t_ Thread) SetThreadPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setThreadPriority:"), value)
}


