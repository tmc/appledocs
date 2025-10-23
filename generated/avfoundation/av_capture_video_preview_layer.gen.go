// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Connection() IAVCaptureConnection
	SetConnection(value IAVCaptureConnection)
	IsDeferredStartEnabled() bool
	SetIsDeferredStartEnabled(value bool)
	IsDeferredStartSupported() bool
	SetIsDeferredStartSupported(value bool)
	IsPreviewing() bool
	SetIsPreviewing(value bool)
	Session() IAVCaptureSession
	SetSession(value IAVCaptureSession)
	VideoGravity() unsafe.Pointer
	SetVideoGravity(value unsafe.Pointer)
}

// A Core Animation layer that displays video from a camera device.
//
// Use this layer to provide a preview of the content the camera captures. A convenient way to use this class in iOS is to set it as the backing layer for a view as shown below.


// A Core Animation layer that displays video from a camera device.
//
// [Full Topic]
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



// An object that describes the connection from the layer to a particular input port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/connection
func (c_ CaptureVideoPreviewLayer) Connection() IAVCaptureConnection {
	rv := objc.Send[CaptureConnection](c_.ID, objc.Sel("connection"))
	return rv
}


// An object that describes the connection from the layer to a particular input port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/connection
func (c_ CaptureVideoPreviewLayer) SetConnection(value IAVCaptureConnection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConnection:"), value)
}


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/isdeferredstartenabled
func (c_ CaptureVideoPreviewLayer) IsDeferredStartEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDeferredStartEnabled"))
	return rv
}


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/isdeferredstartenabled
func (c_ CaptureVideoPreviewLayer) SetIsDeferredStartEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDeferredStartEnabled:"), value)
}


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/isdeferredstartsupported
func (c_ CaptureVideoPreviewLayer) IsDeferredStartSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDeferredStartSupported"))
	return rv
}


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/isdeferredstartsupported
func (c_ CaptureVideoPreviewLayer) SetIsDeferredStartSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDeferredStartSupported:"), value)
}


// A Boolean value that indicates whether the layer is rendering video frames from its source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/ispreviewing
func (c_ CaptureVideoPreviewLayer) IsPreviewing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPreviewing"))
	return rv
}


// A Boolean value that indicates whether the layer is rendering video frames from its source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/ispreviewing
func (c_ CaptureVideoPreviewLayer) SetIsPreviewing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPreviewing:"), value)
}


// A capture session with visual output to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/session
func (c_ CaptureVideoPreviewLayer) Session() IAVCaptureSession {
	rv := objc.Send[CaptureSession](c_.ID, objc.Sel("session"))
	return rv
}


// A capture session with visual output to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/session
func (c_ CaptureVideoPreviewLayer) SetSession(value IAVCaptureSession) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSession:"), value)
}


// A value that indicates how the layer displays video content within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/videogravity
func (c_ CaptureVideoPreviewLayer) VideoGravity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoGravity"))
	return rv
}


// A value that indicates how the layer displays video content within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/videogravity
func (c_ CaptureVideoPreviewLayer) SetVideoGravity(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoGravity:"), value)
}



