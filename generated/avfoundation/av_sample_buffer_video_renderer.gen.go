// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVSampleBufferVideoRenderer] class.
var aVSampleBufferVideoRendererClass = _AVSampleBufferVideoRendererClass{objc.GetClass("AVSampleBufferVideoRenderer")}

type _AVSampleBufferVideoRendererClass struct {
	class objc.Class
}

// An interface definition for the [AVSampleBufferVideoRenderer] class.
type IAVSampleBufferVideoRenderer interface {
	objectivec.IObject
}

// An object that enqueues video sample buffers for rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer

type AVSampleBufferVideoRenderer struct {
	objectivec.Object
}

// AVSampleBufferVideoRendererFrom constructs a [AVSampleBufferVideoRenderer] from an unsafe.Pointer.
//
// An object that enqueues video sample buffers for rendering.
func AVSampleBufferVideoRendererFrom(ptr unsafe.Pointer) AVSampleBufferVideoRenderer {
	return AVSampleBufferVideoRenderer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVSampleBufferVideoRendererClass) Alloc() AVSampleBufferVideoRenderer {
	rv := objc.Send[AVSampleBufferVideoRenderer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVSampleBufferVideoRendererClass) New() AVSampleBufferVideoRenderer {
	rv := objc.Send[AVSampleBufferVideoRenderer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVSampleBufferVideoRenderer) Init() AVSampleBufferVideoRenderer {
	rv := objc.Send[AVSampleBufferVideoRenderer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVSampleBufferVideoRenderer) Autorelease() AVSampleBufferVideoRenderer {
	rv := objc.Send[AVSampleBufferVideoRenderer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferVideoRenderer creates a new AVSampleBufferVideoRenderer instance.
func NewAVSampleBufferVideoRenderer() AVSampleBufferVideoRenderer {
	return aVSampleBufferVideoRendererClass.New()
}




