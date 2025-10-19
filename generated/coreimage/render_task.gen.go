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

// Waits until the finishes and returns. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderTask/waitUntilCompleted()
func (r_ RenderTask) WaitUntilCompletedAndReturnError(error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("waitUntilCompletedAndReturnError:"), error)
	return rv
}


