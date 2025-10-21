// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RunLoop] class.
var (
	RunLoopClass     _RunLoopClass
	RunLoopClassOnce sync.Once
)

func getRunLoopClass() _RunLoopClass {
	RunLoopClassOnce.Do(func() {
		RunLoopClass = _RunLoopClass{objc.GetClass("NSRunLoop")}
	})
	return RunLoopClass
}

type _RunLoopClass struct {
	class objc.Class
}

// An interface definition for the [RunLoop] class.
type IRunLoop interface {
	objectivec.IObject
	AcceptInputForModeBeforeDate(mode unsafe.Pointer, limitDate unsafe.Pointer)
	AddTimerForMode(timer unsafe.Pointer, mode unsafe.Pointer)
	GetCFRunLoop() unsafe.Pointer
	LimitDateForMode(mode unsafe.Pointer) unsafe.Pointer
	Run()
	RunModeBeforeDate(mode unsafe.Pointer, limitDate unsafe.Pointer) bool
	RunUntilDate(limitDate unsafe.Pointer)
}

// The programmatic interface to objects that manage input sources.
//
// A object processes input for sources, such as mouse and keyboard events from the window system and objects. A object also processes events. Your application neither creates nor explicitly manages objects. The system creates a object as needed for each object, including the application’s main thread. If you need to access the current thread’s run loop, use the class method . Note that from the perspective of , objects aren’t “input”—they’re a special type, and they don’t cause the run loop to return when they fire.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop
type RunLoop struct {
	objectivec.Object
}

// RunLoopFrom constructs a [RunLoop] from an unsafe.Pointer.
//
// The programmatic interface to objects that manage input sources.
func RunLoopFrom(ptr unsafe.Pointer) RunLoop {
	return RunLoop{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RunLoopClass) Alloc() RunLoop {
	rv := objc.Send[RunLoop](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RunLoopClass) New() RunLoop {
	rv := objc.Send[RunLoop](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RunLoop) Init() RunLoop {
	rv := objc.Send[RunLoop](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RunLoop) Autorelease() RunLoop {
	rv := objc.Send[RunLoop](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRunLoop creates a new RunLoop instance.
func NewRunLoop() RunLoop {
	return getRunLoopClass().New()
}

// Returns the run loop for the current thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/current
func (rc _RunLoopClass) CurrentRunLoop() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("currentRunLoop"))
	return rv
}

// Returns the run loop of the main thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/main
func (rc _RunLoopClass) MainRunLoop() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("mainRunLoop"))
	return rv
}

// Runs the loop once or until the specified date, accepting input only for the specified mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/acceptInput(forMode:before:)
func (r_ RunLoop) AcceptInputForModeBeforeDate(mode unsafe.Pointer, limitDate unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("acceptInputForMode:beforeDate:"), mode, limitDate)
}

// Registers a given timer with a given input mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/add(_:forMode:)-392ag
func (r_ RunLoop) AddTimerForMode(timer unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addTimer:forMode:"), timer, mode)
}

// Returns the receiver’s underlying run loop object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/getCFRunLoop()
func (r_ RunLoop) GetCFRunLoop() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("getCFRunLoop"))
	return rv
}

// Performs one pass through the run loop in the specified mode and returns the date at which the next timer is scheduled to fire.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/limitDate(forMode:)
func (r_ RunLoop) LimitDateForMode(mode unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("limitDateForMode:"), mode)
	return rv
}

// Puts the receiver into a permanent loop, during which time it processes data from all attached input sources.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/run()
func (r_ RunLoop) Run() {
	objc.Send[objc.ID](r_.ID, objc.Sel("run"))
}

// Runs the loop once, blocking for input in the specified mode until a given date.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/run(mode:before:)
func (r_ RunLoop) RunModeBeforeDate(mode unsafe.Pointer, limitDate unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("runMode:beforeDate:"), mode, limitDate)
	return rv
}

// Runs the loop until the specified date, during which time it processes data from all attached input sources.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/run(until:)
func (r_ RunLoop) RunUntilDate(limitDate unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("runUntilDate:"), limitDate)
}

// Returns the run loop for the current thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/current
func (r_ RunLoop) CurrentRunLoop() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("currentRunLoop"))
	return rv
}

// The receiver’s current input mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/currentMode
func (r_ RunLoop) CurrentMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("currentMode"))
	return rv
}

// Returns the run loop of the main thread.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/main
func (r_ RunLoop) MainRunLoop() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("mainRunLoop"))
	return rv
}
