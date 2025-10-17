// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RunLoop] class.
var RunLoopClass objc.Class

func init() {
	RunLoopClass = objc.GetClass("NSRunLoop")
}

type RunLoop struct {
	objc.ID
}

func RunLoopFrom(ptr unsafe.Pointer) RunLoop {
	return RunLoop{
		ID: objc.ID(ptr),
	}
}


// Registers a given timer with a given input mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/RunLoop/add(_:forMode:)-392ag
func (r_ RunLoop) AddTimerForMode(timer unsafe.Pointer, mode unsafe.Pointer) {
	sel := objc.RegisterName("addTimer:forMode:")
	r_.ID.Send(sel, timer, mode)
}
// Returns the receiver’s underlying run loop object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/RunLoop/getCFRunLoop()
func (r_ RunLoop) GetCFRunLoop() unsafe.Pointer {
	sel := objc.RegisterName("getCFRunLoop")
	ret := r_.ID.Send(sel)
	return unsafe.Pointer(ret)
}


