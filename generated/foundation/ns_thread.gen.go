// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSThread */


/* debug [class_header]: Header for NSThread */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Thread */
// An interface definition for the [Thread] class.
type IThread interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Thread */
	// properties:
	Cancelled() bool
	Executing() bool
	Finished() bool
	IsMainThread() bool
	Name() IString
	SetName(value IString)
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	StackSize() uint
	SetStackSize(value uint)
	ThreadDictionary() IMutableDictionary
	ThreadPriority() float64
	SetThreadPriority(value float64)
	NSAssertionHandlerKey() IString
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsExecuting() bool
	SetIsExecuting(value bool)
	IsFinished() bool
	SetIsFinished(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Thread */
	// methods:
	Cancel()
	Main()
	Start()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Thread */
// Alloc allocates a new instance without initialization.
func (tc _ThreadClass) Alloc() Thread {
	rv := objc.Send[Thread](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Thread */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Thread */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/init(block:)
func NewThreadWithBlock(block unsafe.Pointer) Thread {
	instance := getThreadClass().Alloc()
	rv := objc.Send[Thread](instance.ID, objc.Sel("initWithBlock:"), block)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewThreadWithBlock */


// Returns an object initialized with the given arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/init(target:selector:object:)
func NewThreadWithTargetSelectorObject(target objc.IObject, selector objc.SEL, argument objc.IObject) Thread {
	instance := getThreadClass().Alloc()
	rv := objc.Send[Thread](instance.ID, objc.Sel("initWithTarget:selector:object:"), target, selector, argument)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewThreadWithTargetSelectorObject */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Thread */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/detachNewThread(_:)
func (tc _ThreadClass) DetachNewThreadWithBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("detachNewThreadWithBlock:"), block)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DetachNewThreadWithBlock) */


// Detaches a new thread and uses the specified selector as the thread entry point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/detachNewThreadSelector(_:toTarget:with:)
func (tc _ThreadClass) DetachNewThreadSelectorToTargetWithObject(selector objc.SEL, target objc.IObject, argument objc.IObject) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("detachNewThreadSelector:toTarget:withObject:"), selector, target, argument)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DetachNewThreadSelectorToTargetWithObject) */


// Terminates the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/exit()
func (tc _ThreadClass) Exit() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("exit"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Exit) */


// Returns whether the application is multithreaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isMultiThreaded()
func (tc _ThreadClass) IsMultiThreaded() bool {
	rv := objc.Send[bool](objc.ID(tc.class), objc.Sel("isMultiThreaded"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsMultiThreaded) */


// Sets the current thread’s priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/setThreadPriority(_:)
func (tc _ThreadClass) SetThreadPriority(p float64) bool {
	rv := objc.Send[bool](objc.ID(tc.class), objc.Sel("setThreadPriority:"), p)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetThreadPriority) */


// Sleeps the thread for a given time interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/sleep(forTimeInterval:)
func (tc _ThreadClass) SleepForTimeInterval(ti float64) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("sleepForTimeInterval:"), ti)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SleepForTimeInterval) */


// Blocks the current thread until the time specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/sleep(until:)
func (tc _ThreadClass) SleepUntilDate(date IDate) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("sleepUntilDate:"), date)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SleepUntilDate) */


// Returns the current thread’s priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/threadPriority()
func (tc _ThreadClass) ThreadPriority() float64 {
	rv := objc.Send[float64](objc.ID(tc.class), objc.Sel("threadPriority"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ThreadPriority) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Thread */

// Returns an array containing the call stack return addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/callStackReturnAddresses
func (tc _ThreadClass) CallStackReturnAddresses() []Number {
	rv := objc.Send[[]Number](objc.ID(tc.class), objc.Sel("callStackReturnAddresses"))
	return rv
}/* debug [class_properties_class/property]: callStackReturnAddresses */

// Returns an array containing the call stack symbols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/callStackSymbols
func (tc _ThreadClass) CallStackSymbols() []string {
	rv := objc.Send[[]string](objc.ID(tc.class), objc.Sel("callStackSymbols"))
	return rv
}/* debug [class_properties_class/property]: callStackSymbols */

// Returns the thread object representing the current thread of execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/current
func (tc _ThreadClass) CurrentThread() Thread {
	rv := objc.Send[Thread](objc.ID(tc.class), objc.Sel("currentThread"))
	return rv
}/* debug [class_properties_class/property]: currentThread */

// Returns the object representing the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/main
func (tc _ThreadClass) MainThread() Thread {
	rv := objc.Send[Thread](objc.ID(tc.class), objc.Sel("mainThread"))
	return rv
}/* debug [class_properties_class/property]: mainThread */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Thread */

// Changes the cancelled state of the receiver to indicate that it should exit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/cancel()
func (t_ Thread) Cancel() {
	objc.Send[objc.ID](t_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// The main entry point routine for the thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/main()
func (t_ Thread) Main() {
	objc.Send[objc.ID](t_.ID, objc.Sel("main"))
}/* debug [instance_methods/method]: Main */


// Starts the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/start()
func (t_ Thread) Start() {
	objc.Send[objc.ID](t_.ID, objc.Sel("start"))
}/* debug [instance_methods/method]: Start */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Thread */

// Returns an array containing the call stack return addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/callStackReturnAddresses
func (t_ Thread) CallStackReturnAddresses() []Number {
	rv := objc.Send[[]Number](t_.ID, objc.Sel("callStackReturnAddresses"))
	return rv
}/* debug [instance_properties/getter]: callStackReturnAddresses */


// Returns an array containing the call stack symbols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/callStackSymbols
func (t_ Thread) CallStackSymbols() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("callStackSymbols"))
	return rv
}/* debug [instance_properties/getter]: callStackSymbols */


// Returns the thread object representing the current thread of execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/current
func (t_ Thread) CurrentThread() IThread {
	rv := objc.Send[Thread](t_.ID, objc.Sel("currentThread"))
	return rv
}/* debug [instance_properties/getter]: currentThread */


// A Boolean value that indicates whether the receiver is cancelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isCancelled
func (t_ Thread) Cancelled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("cancelled"))
	return rv
}/* debug [instance_properties/getter]: cancelled */


// A Boolean value that indicates whether the receiver is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isExecuting
func (t_ Thread) Executing() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("executing"))
	return rv
}/* debug [instance_properties/getter]: executing */


// A Boolean value that indicates whether the receiver has finished execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isFinished
func (t_ Thread) Finished() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("finished"))
	return rv
}/* debug [instance_properties/getter]: finished */


// A Boolean value that indicates whether the receiver is the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/isMainThread-swift.property
func (t_ Thread) IsMainThread() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isMainThread"))
	return rv
}/* debug [instance_properties/getter]: isMainThread */


// Returns the object representing the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/main
func (t_ Thread) MainThread() IThread {
	rv := objc.Send[Thread](t_.ID, objc.Sel("mainThread"))
	return rv
}/* debug [instance_properties/getter]: mainThread */


// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/name
func (t_ Thread) Name() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/name
func (t_ Thread) SetName(value IString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/qualityOfService
func (t_ Thread) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](t_.ID, objc.Sel("qualityOfService"))
	return rv
}/* debug [instance_properties/getter]: qualityOfService */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/qualityOfService
func (t_ Thread) SetQualityOfService(value QualityOfService) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setQualityOfService:"), value)
}/* debug [instance_properties/setter]: qualityOfService */


// The stack size of the receiver, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/stackSize
func (t_ Thread) StackSize() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("stackSize"))
	return rv
}/* debug [instance_properties/getter]: stackSize */


// The stack size of the receiver, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/stackSize
func (t_ Thread) SetStackSize(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStackSize:"), value)
}/* debug [instance_properties/setter]: stackSize */


// The thread object’s dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/threadDictionary
func (t_ Thread) ThreadDictionary() IMutableDictionary {
	rv := objc.Send[MutableDictionary](t_.ID, objc.Sel("threadDictionary"))
	return rv
}/* debug [instance_properties/getter]: threadDictionary */


// The receiver’s priority
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/threadPriority
func (t_ Thread) ThreadPriority() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("threadPriority"))
	return rv
}/* debug [instance_properties/getter]: threadPriority */


// The receiver’s priority
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Thread/threadPriority
func (t_ Thread) SetThreadPriority(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setThreadPriority:"), value)
}/* debug [instance_properties/setter]: threadPriority */


// A key with a corresponding value in the thread dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsassertionhandlerkey
func (t_ Thread) NSAssertionHandlerKey() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("NSAssertionHandlerKey"))
	return rv
}/* debug [instance_properties/getter]: NSAssertionHandlerKey */


// A Boolean value that indicates whether the receiver is cancelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/iscancelled
func (t_ Thread) IsCancelled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isCancelled"))
	return rv
}/* debug [instance_properties/getter]: isCancelled */


// A Boolean value that indicates whether the receiver is cancelled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/iscancelled
func (t_ Thread) SetIsCancelled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsCancelled:"), value)
}/* debug [instance_properties/setter]: isCancelled */


// A Boolean value that indicates whether the receiver is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isexecuting
func (t_ Thread) IsExecuting() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isExecuting"))
	return rv
}/* debug [instance_properties/getter]: isExecuting */


// A Boolean value that indicates whether the receiver is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isexecuting
func (t_ Thread) SetIsExecuting(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsExecuting:"), value)
}/* debug [instance_properties/setter]: isExecuting */


// A Boolean value that indicates whether the receiver has finished execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isfinished
func (t_ Thread) IsFinished() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFinished"))
	return rv
}/* debug [instance_properties/getter]: isFinished */


// A Boolean value that indicates whether the receiver has finished execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/thread/isfinished
func (t_ Thread) SetIsFinished(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFinished:"), value)
}/* debug [instance_properties/setter]: isFinished */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSThread */


