// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RenderTask] class.
var (
	RenderTaskClass     _RenderTaskClass
	RenderTaskClassOnce sync.Once
)

func getRenderTaskClass() _RenderTaskClass {
	RenderTaskClassOnce.Do(func() {
		RenderTaskClass = _RenderTaskClass{objc.GetClass("CIRenderTask")}
	})
	return RenderTaskClass
}

type _RenderTaskClass struct {
	class objc.Class
}

// An interface definition for the [RenderTask] class.
type IRenderTask interface {
	objectivec.IObject
	// properties:
	// methods:
	WaitUntilCompletedAndReturnError(error_ unsafe.Pointer) IRenderInfo
}

// A single render task.
//
// A single render task issued in conjunction with . A object appears in Xcode Quick Look as a graph.


// A single render task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderTask
type RenderTask struct {
	objectivec.Object
}

// RenderTaskFrom constructs a [RenderTask] from an unsafe.Pointer.
//
// A single render task.
func RenderTaskFrom(ptr unsafe.Pointer) RenderTask {
	return RenderTask{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RenderTaskClass) Alloc() RenderTask {
	rv := objc.Send[RenderTask](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RenderTaskClass) New() RenderTask {
	rv := objc.Send[RenderTask](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderTask) Init() RenderTask {
	rv := objc.Send[RenderTask](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderTask) Autorelease() RenderTask {
	rv := objc.Send[RenderTask](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderTask creates a new RenderTask instance.
func NewRenderTask() RenderTask {
	return getRenderTaskClass().New()
}



// Waits until the finishes and returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderTask/waitUntilCompleted()
func (r_ RenderTask) WaitUntilCompletedAndReturnError(error_ unsafe.Pointer) IRenderInfo {
	rv := objc.Send[RenderInfo](r_.ID, objc.Sel("waitUntilCompletedAndReturnError:"), error_)
	return rv
}



