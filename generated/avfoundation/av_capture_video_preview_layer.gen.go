// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [AVCaptureVideoPreviewLayer] class.
var (
	aVCaptureVideoPreviewLayerClass     _AVCaptureVideoPreviewLayerClass
	aVCaptureVideoPreviewLayerClassOnce sync.Once
)

func getAVCaptureVideoPreviewLayerClass() _AVCaptureVideoPreviewLayerClass {
	aVCaptureVideoPreviewLayerClassOnce.Do(func() {
		aVCaptureVideoPreviewLayerClass = _AVCaptureVideoPreviewLayerClass{objc.GetClass("AVCaptureVideoPreviewLayer")}
	})
	return aVCaptureVideoPreviewLayerClass
}

type _AVCaptureVideoPreviewLayerClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureVideoPreviewLayer] class.
type IAVCaptureVideoPreviewLayer interface {
	quartzcore.ILayer
}

// A Core Animation layer that displays video from a camera device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer
type AVCaptureVideoPreviewLayer struct {
	quartzcore.Layer
}

// AVCaptureVideoPreviewLayerFrom constructs a [AVCaptureVideoPreviewLayer] from an unsafe.Pointer.
//
// A Core Animation layer that displays video from a camera device.
func AVCaptureVideoPreviewLayerFrom(ptr unsafe.Pointer) AVCaptureVideoPreviewLayer {
	return AVCaptureVideoPreviewLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureVideoPreviewLayerClass) Alloc() AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureVideoPreviewLayerClass) New() AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureVideoPreviewLayer) Init() AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureVideoPreviewLayer) Autorelease() AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureVideoPreviewLayer creates a new AVCaptureVideoPreviewLayer instance.
func NewAVCaptureVideoPreviewLayer() AVCaptureVideoPreviewLayer {
	return getAVCaptureVideoPreviewLayerClass().New()
}




