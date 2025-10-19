// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [AVSampleBufferDisplayLayer] class.
var (
	aVSampleBufferDisplayLayerClass     _AVSampleBufferDisplayLayerClass
	aVSampleBufferDisplayLayerClassOnce sync.Once
)

func getAVSampleBufferDisplayLayerClass() _AVSampleBufferDisplayLayerClass {
	aVSampleBufferDisplayLayerClassOnce.Do(func() {
		aVSampleBufferDisplayLayerClass = _AVSampleBufferDisplayLayerClass{objc.GetClass("AVSampleBufferDisplayLayer")}
	})
	return aVSampleBufferDisplayLayerClass
}

type _AVSampleBufferDisplayLayerClass struct {
	class objc.Class
}

// An interface definition for the [AVSampleBufferDisplayLayer] class.
type IAVSampleBufferDisplayLayer interface {
	quartzcore.ILayer
}

// An object that displays compressed or uncompressed video frames.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer
type AVSampleBufferDisplayLayer struct {
	quartzcore.Layer
}

// AVSampleBufferDisplayLayerFrom constructs a [AVSampleBufferDisplayLayer] from an unsafe.Pointer.
//
// An object that displays compressed or uncompressed video frames.
func AVSampleBufferDisplayLayerFrom(ptr unsafe.Pointer) AVSampleBufferDisplayLayer {
	return AVSampleBufferDisplayLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVSampleBufferDisplayLayerClass) Alloc() AVSampleBufferDisplayLayer {
	rv := objc.Send[AVSampleBufferDisplayLayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVSampleBufferDisplayLayerClass) New() AVSampleBufferDisplayLayer {
	rv := objc.Send[AVSampleBufferDisplayLayer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVSampleBufferDisplayLayer) Init() AVSampleBufferDisplayLayer {
	rv := objc.Send[AVSampleBufferDisplayLayer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVSampleBufferDisplayLayer) Autorelease() AVSampleBufferDisplayLayer {
	rv := objc.Send[AVSampleBufferDisplayLayer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferDisplayLayer creates a new AVSampleBufferDisplayLayer instance.
func NewAVSampleBufferDisplayLayer() AVSampleBufferDisplayLayer {
	return getAVSampleBufferDisplayLayerClass().New()
}




