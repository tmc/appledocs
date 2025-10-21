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
func (c_ CaptureVideoPreviewLayer) Connection() AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](c_.ID, objc.Sel("connection"))
	return rv
}

// A capture session with visual output to preview.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/session
func (c_ CaptureVideoPreviewLayer) Session() AVCaptureSession {
	rv := objc.Send[AVCaptureSession](c_.ID, objc.Sel("session"))
	return rv
}


// SetSession sets the value of the session property.
// A capture session with visual output to preview.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/session
func (c_ CaptureVideoPreviewLayer) SetSession(value IAVCaptureSession) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSession:"), value)
}

// A
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/isdeferredstartenabled
func (c_ CaptureVideoPreviewLayer) IsDeferredStartEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDeferredStartEnabled"))
	return rv
}


// SetIsDeferredStartEnabled sets the value of the isDeferredStartEnabled property.
// A

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/isdeferredstartenabled
func (c_ CaptureVideoPreviewLayer) SetIsDeferredStartEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDeferredStartEnabled:"), value)
}

// A
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/isdeferredstartsupported
func (c_ CaptureVideoPreviewLayer) IsDeferredStartSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDeferredStartSupported"))
	return rv
}


// SetIsDeferredStartSupported sets the value of the isDeferredStartSupported property.
// A

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/isdeferredstartsupported
func (c_ CaptureVideoPreviewLayer) SetIsDeferredStartSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDeferredStartSupported:"), value)
}

// A Boolean value that indicates whether the layer is rendering video frames from its source.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/ispreviewing
func (c_ CaptureVideoPreviewLayer) IsPreviewing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPreviewing"))
	return rv
}


// SetIsPreviewing sets the value of the isPreviewing property.
// A Boolean value that indicates whether the layer is rendering video frames from its source.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/ispreviewing
func (c_ CaptureVideoPreviewLayer) SetIsPreviewing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPreviewing:"), value)
}

// A value that indicates how the layer displays video content within its bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/videogravity
func (c_ CaptureVideoPreviewLayer) VideoGravity() LayerVideoGravity {
	rv := objc.Send[LayerVideoGravity](c_.ID, objc.Sel("videoGravity"))
	return rv
}


// SetVideoGravity sets the value of the videoGravity property.
// A value that indicates how the layer displays video content within its bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/videogravity
func (c_ CaptureVideoPreviewLayer) SetVideoGravity(value ILayerVideoGravity) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoGravity:"), value)
}



