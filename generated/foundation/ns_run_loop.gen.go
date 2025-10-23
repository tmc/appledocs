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
	// properties:
	CurrentMode() unsafe.Pointer
	SetCurrentMode(value unsafe.Pointer)
	// methods:
	LimitDateForMode(mode NSRunLoopMode /* foo */) IDate
	PerformBlock(block unsafe.Pointer)
	PerformInModesBlock(modes []string /* primitive/slice/pointer */, block unsafe.Pointer)
}

// The programmatic interface to objects that manage input sources.
//
// A object processes input for sources, such as mouse and keyboard events from the window system and objects. A object also processes events. Your application neither creates nor explicitly manages objects. The system creates a object as needed for each object, including the application’s main thread. If you need to access the current thread’s run loop, use the class method . Note that from the perspective of , objects aren’t “input”—they’re a special type, and they don’t cause the run loop to return when they fire.


// The programmatic interface to objects that manage input sources.
//
// [Full Topic]
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



// Performs one pass through the run loop in the specified mode and returns the date at which the next timer is scheduled to fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/limitDate(forMode:)
func (r_ RunLoop) LimitDateForMode(mode NSRunLoopMode /* foo */) IDate {
	rv := objc.Send[Date](r_.ID, objc.Sel("limitDateForMode:"), mode)
	return rv
}


// Schedules a block that the run loop invokes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/perform(_:)
func (r_ RunLoop) PerformBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("performBlock:"), block)
}


// Schedules a block that the run loop invokes when it’s running in any of the specified modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/perform(inModes:block:)
func (r_ RunLoop) PerformInModesBlock(modes []string /* primitive/slice/pointer */, block unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("performInModes:block:"), modes, block)
}


// The receiver’s current input mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/runloop/currentmode
func (r_ RunLoop) CurrentMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("currentMode"))
	return rv
}


// The receiver’s current input mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/runloop/currentmode
func (r_ RunLoop) SetCurrentMode(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCurrentMode:"), value)
}



