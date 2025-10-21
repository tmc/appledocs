// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [CaptureVideoPreviewLayer] class.
var (
	CaptureVideoPreviewLayerClass     _CaptureVideoPreviewLayerClass
	CaptureVideoPreviewLayerClassOnce sync.Once
)

func getCaptureVideoPreviewLayerClass() _CaptureVideoPreviewLayerClass {
	CaptureVideoPreviewLayerClassOnce.Do(func() {
		CaptureVideoPreviewLayerClass = _CaptureVideoPreviewLayerClass{objc.GetClass("AVCaptureVideoPreviewLayer")}
	})
	return CaptureVideoPreviewLayerClass
}

type _CaptureVideoPreviewLayerClass struct {
	class objc.Class
}

// An interface definition for the [CaptureVideoPreviewLayer] class.
type ICaptureVideoPreviewLayer interface {
	quartzcore.ILayer
	RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates coregraphics.CGRect) coregraphics.CGRect
}

// A Core Animation layer that displays video from a camera device.
//
// Use this layer to provide a preview of the content the camera captures. A convenient way to use this class in iOS is to set it as the backing layer for a view as shown below.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer
type CaptureVideoPreviewLayer struct {
	quartzcore.Layer
}

// CaptureVideoPreviewLayerFrom constructs a [CaptureVideoPreviewLayer] from an unsafe.Pointer.
//
// A Core Animation layer that displays video from a camera device.
func CaptureVideoPreviewLayerFrom(ptr unsafe.Pointer) CaptureVideoPreviewLayer {
	return CaptureVideoPreviewLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureVideoPreviewLayerClass) Alloc() CaptureVideoPreviewLayer {
	rv := objc.Send[CaptureVideoPreviewLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureVideoPreviewLayerClass) New() CaptureVideoPreviewLayer {
	rv := objc.Send[CaptureVideoPreviewLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureVideoPreviewLayer) Init() CaptureVideoPreviewLayer {
	rv := objc.Send[CaptureVideoPreviewLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureVideoPreviewLayer) Autorelease() CaptureVideoPreviewLayer {
	rv := objc.Send[CaptureVideoPreviewLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureVideoPreviewLayer creates a new CaptureVideoPreviewLayer instance.
func NewCaptureVideoPreviewLayer() CaptureVideoPreviewLayer {
	return getCaptureVideoPreviewLayerClass().New()
}


// Converts a rectangle from metadata output coordinates to the coordinate space of the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/layerRectConverted(fromMetadataOutputRect:)
func (c_ CaptureVideoPreviewLayer) RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("rectForMetadataOutputRectOfInterest:"), rectInMetadataOutputCoordinates)
	return rv
}

// An object that describes the connection from the layer to a particular input port.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/connection
func (c_ CaptureVideoPreviewLayer) Connection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("connection"))
	return rv
}

// A capture session with visual output to preview.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/session
func (c_ CaptureVideoPreviewLayer) Session() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("session"))
	return rv
}


// SetSession sets the value of the session property.
// A capture session with visual output to preview.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/session
func (c_ CaptureVideoPreviewLayer) SetSession(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSession:"), value)
}



