// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [SampleBufferDisplayLayer] class.
var (
	SampleBufferDisplayLayerClass     _SampleBufferDisplayLayerClass
	SampleBufferDisplayLayerClassOnce sync.Once
)

func getSampleBufferDisplayLayerClass() _SampleBufferDisplayLayerClass {
	SampleBufferDisplayLayerClassOnce.Do(func() {
		SampleBufferDisplayLayerClass = _SampleBufferDisplayLayerClass{objc.GetClass("AVSampleBufferDisplayLayer")}
	})
	return SampleBufferDisplayLayerClass
}

type _SampleBufferDisplayLayerClass struct {
	class objc.Class
}

// An interface definition for the [SampleBufferDisplayLayer] class.
type ISampleBufferDisplayLayer interface {
	quartzcore.ILayer
}

// An object that displays compressed or uncompressed video frames.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer
type SampleBufferDisplayLayer struct {
	quartzcore.Layer
}

// SampleBufferDisplayLayerFrom constructs a [SampleBufferDisplayLayer] from an unsafe.Pointer.
//
// An object that displays compressed or uncompressed video frames.
func SampleBufferDisplayLayerFrom(ptr unsafe.Pointer) SampleBufferDisplayLayer {
	return SampleBufferDisplayLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SampleBufferDisplayLayerClass) Alloc() SampleBufferDisplayLayer {
	rv := objc.Send[SampleBufferDisplayLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SampleBufferDisplayLayerClass) New() SampleBufferDisplayLayer {
	rv := objc.Send[SampleBufferDisplayLayer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferDisplayLayer) Init() SampleBufferDisplayLayer {
	rv := objc.Send[SampleBufferDisplayLayer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferDisplayLayer) Autorelease() SampleBufferDisplayLayer {
	rv := objc.Send[SampleBufferDisplayLayer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferDisplayLayer creates a new SampleBufferDisplayLayer instance.
func NewSampleBufferDisplayLayer() SampleBufferDisplayLayer {
	return getSampleBufferDisplayLayerClass().New()
}




