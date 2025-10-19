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
	runLoopClass     _RunLoopClass
	runLoopClassOnce sync.Once
)

func getRunLoopClass() _RunLoopClass {
	runLoopClassOnce.Do(func() {
		runLoopClass = _RunLoopClass{objc.GetClass("NSRunLoop")}
	})
	return runLoopClass
}

type _RunLoopClass struct {
	class objc.Class
}

// An interface definition for the [RunLoop] class.
type IRunLoop interface {
	objectivec.IObject
	AddTimerForMode(timer unsafe.Pointer, mode unsafe.Pointer)
	GetCFRunLoop() unsafe.Pointer
}

// The programmatic interface to objects that manage input sources. [Full Topic]
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


// Registers a given timer with a given input mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/add(_:forMode:)-392ag
func (r_ RunLoop) AddTimerForMode(timer unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addTimer:forMode:"), timer, mode)
}
// Returns the receiver’s underlying run loop object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/getCFRunLoop()
func (r_ RunLoop) GetCFRunLoop() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("getCFRunLoop"))
	return rv
}


