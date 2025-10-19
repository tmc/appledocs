// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RenderTask] class.
var renderTaskClass = _RenderTaskClass{objc.GetClass("CIRenderTask")}

type _RenderTaskClass struct {
	class objc.Class
}

// An interface definition for the [RenderTask] class.
type IRenderTask interface {
	objectivec.IObject
	WaitUntilCompletedAndReturnError(error unsafe.Pointer) unsafe.Pointer
}

// A single render task. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return renderTaskClass.New()
}


// Waits until the finishes and returns. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderTask/waitUntilCompleted()
func (r_ RenderTask) WaitUntilCompletedAndReturnError(error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("waitUntilCompletedAndReturnError:"), error)
	return rv
}


