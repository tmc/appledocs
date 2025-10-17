// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RunLoop] class.
var runLoopClass = _RunLoopClass{objc.GetClass("NSRunLoop")}

type _RunLoopClass struct {
	class objc.Class
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


