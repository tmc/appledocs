// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VideoProcessor] class.
var (
	VideoProcessorClass     _VideoProcessorClass
	VideoProcessorClassOnce sync.Once
)

func getVideoProcessorClass() _VideoProcessorClass {
	VideoProcessorClassOnce.Do(func() {
		VideoProcessorClass = _VideoProcessorClass{objc.GetClass("VNVideoProcessor")}
	})
	return VideoProcessorClass
}

type _VideoProcessorClass struct {
	class objc.Class
}

// An interface definition for the [VideoProcessor] class.
type IVideoProcessor interface {
	objectivec.IObject
}

// An object that performs offline analysis of video content.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor
type VideoProcessor struct {
	objectivec.Object
}

// VideoProcessorFrom constructs a [VideoProcessor] from an unsafe.Pointer.
//
// An object that performs offline analysis of video content.
func VideoProcessorFrom(ptr unsafe.Pointer) VideoProcessor {
	return VideoProcessor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VideoProcessorClass) Alloc() VideoProcessor {
	rv := objc.Send[VideoProcessor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VideoProcessorClass) New() VideoProcessor {
	rv := objc.Send[VideoProcessor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoProcessor) Init() VideoProcessor {
	rv := objc.Send[VideoProcessor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoProcessor) Autorelease() VideoProcessor {
	rv := objc.Send[VideoProcessor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoProcessor creates a new VideoProcessor instance.
func NewVideoProcessor() VideoProcessor {
	return getVideoProcessorClass().New()
}




