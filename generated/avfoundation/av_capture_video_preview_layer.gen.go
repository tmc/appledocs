// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	ILayer
	

	// properties:
	Connection() IAVCaptureConnection
	DeferredStartEnabled() bool
	SetDeferredStartEnabled(value bool)
	DeferredStartSupported() bool
	Session() IAVCaptureSession
	SetSession(value IAVCaptureSession)
	VideoGravity() LayerVideoGravity /* typedef */
	SetVideoGravity(value LayerVideoGravity /* typedef */)
	IsDeferredStartEnabled() bool
	SetIsDeferredStartEnabled(value bool)
	IsDeferredStartSupported() bool
	SetIsDeferredStartSupported(value bool)
	IsPreviewing() bool
	SetIsPreviewing(value bool)


	

	// methods:
	CaptureDevicePointOfInterestForPoint(pointInLayer corefoundation.CGPoint) corefoundation.CGPoint
	PointForCaptureDevicePointOfInterest(captureDevicePointOfInterest corefoundation.CGPoint) corefoundation.CGPoint
	RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates corefoundation.CGRect) corefoundation.CGRect
	MetadataOutputRectOfInterestForRect(rectInLayerCoordinates corefoundation.CGRect) corefoundation.CGRect
	SetSessionWithNoConnection(session IAVCaptureSession)
	TransformedMetadataObjectForMetadataObject(metadataObject IAVMetadataObject) IMetadataObject


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureVideoPreviewLayerClass) Alloc() CaptureVideoPreviewLayer {
	rv := objc.Send[CaptureVideoPreviewLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A Core Animation layer that displays video from a camera device.
//
// Use this layer to provide a preview of the content the camera captures. A convenient way to use this class in iOS is to set it as the backing layer for a view as shown below.


// A Core Animation layer that displays video from a camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer
type CaptureVideoPreviewLayer struct {
	Layer
}

// CaptureVideoPreviewLayerFrom constructs a [CaptureVideoPreviewLayer] from an unsafe.Pointer.
//
// A Core Animation layer that displays video from a camera device.
func CaptureVideoPreviewLayerFrom(ptr unsafe.Pointer) CaptureVideoPreviewLayer {
	return CaptureVideoPreviewLayer{
		Layer: LayerFrom(ptr),
	}
}






// Creates a layer to preview the visual output of a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/init(session:)
func NewCaptureVideoPreviewLayerWithSession(session IAVCaptureSession) CaptureVideoPreviewLayer {
	instance := getCaptureVideoPreviewLayerClass().Alloc()
	rv := objc.Send[CaptureVideoPreviewLayer](instance.ID, objc.Sel("initWithSession:"), session)
	rv.Autorelease()
	return rv
}


// Creates a layer to preview the visual output of a capture session, without making connections to eligible video inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/init(sessionWithNoConnection:)
func NewCaptureVideoPreviewLayerWithSessionWithNoConnection(session IAVCaptureSession) CaptureVideoPreviewLayer {
	instance := getCaptureVideoPreviewLayerClass().Alloc()
	rv := objc.Send[CaptureVideoPreviewLayer](instance.ID, objc.Sel("initWithSessionWithNoConnection:"), session)
	rv.Autorelease()
	return rv
}







// Returns a new layer to preview the visual output of a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/layerWithSession:
func (cc _CaptureVideoPreviewLayerClass) LayerWithSession(session IAVCaptureSession) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layerWithSession:"), session)
	return rv
}


// Returns a new layer to preview the visual output of a capture session, without making connections to eligible video inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/layerWithSessionWithNoConnection:
func (cc _CaptureVideoPreviewLayerClass) LayerWithSessionWithNoConnection(session IAVCaptureSession) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layerWithSessionWithNoConnection:"), session)
	return rv
}












// Converts a point from layer coordinates to the coordinate space of the capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/captureDevicePointConverted(fromLayerPoint:)
func (c_ CaptureVideoPreviewLayer) CaptureDevicePointOfInterestForPoint(pointInLayer corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("captureDevicePointOfInterestForPoint:"), pointInLayer)
	return rv
}


// Converts a point from the coordinate space of the capture device to the coordinate space of the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/layerPointConverted(fromCaptureDevicePoint:)
func (c_ CaptureVideoPreviewLayer) PointForCaptureDevicePointOfInterest(captureDevicePointOfInterest corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("pointForCaptureDevicePointOfInterest:"), captureDevicePointOfInterest)
	return rv
}


// Converts a rectangle from metadata output coordinates to the coordinate space of the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/layerRectConverted(fromMetadataOutputRect:)
func (c_ CaptureVideoPreviewLayer) RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("rectForMetadataOutputRectOfInterest:"), rectInMetadataOutputCoordinates)
	return rv
}


// Converts a rectangle from layer coordinates to the coordinate space of the metadata output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/metadataOutputRectConverted(fromLayerRect:)
func (c_ CaptureVideoPreviewLayer) MetadataOutputRectOfInterestForRect(rectInLayerCoordinates corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("metadataOutputRectOfInterestForRect:"), rectInLayerCoordinates)
	return rv
}


// Associates a session with the layer without automatically forming a connection to an eligible input port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/setSessionWithNoConnection(_:)
func (c_ CaptureVideoPreviewLayer) SetSessionWithNoConnection(session IAVCaptureSession) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSessionWithNoConnection:"), session)
}


// Converts a metadata object’s visual properties to layer coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/transformedMetadataObject(for:)
func (c_ CaptureVideoPreviewLayer) TransformedMetadataObjectForMetadataObject(metadataObject IAVMetadataObject) IMetadataObject {
	rv := objc.Send[MetadataObject](c_.ID, objc.Sel("transformedMetadataObjectForMetadataObject:"), metadataObject)
	return rv
}







// An object that describes the connection from the layer to a particular input port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/connection
func (c_ CaptureVideoPreviewLayer) Connection() IAVCaptureConnection {
	rv := objc.Send[CaptureConnection](c_.ID, objc.Sel("connection"))
	return rv
}


// A value that indicates whether to defer starting this preview layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/isDeferredStartEnabled
func (c_ CaptureVideoPreviewLayer) DeferredStartEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("deferredStartEnabled"))
	return rv
}


// A value that indicates whether to defer starting this preview layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/isDeferredStartEnabled
func (c_ CaptureVideoPreviewLayer) SetDeferredStartEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeferredStartEnabled:"), value)
}


// A value that indicates whether the preview layer supports deferred start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/isDeferredStartSupported
func (c_ CaptureVideoPreviewLayer) DeferredStartSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("deferredStartSupported"))
	return rv
}


// A capture session with visual output to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/session
func (c_ CaptureVideoPreviewLayer) Session() IAVCaptureSession {
	rv := objc.Send[CaptureSession](c_.ID, objc.Sel("session"))
	return rv
}


// A capture session with visual output to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/session
func (c_ CaptureVideoPreviewLayer) SetSession(value IAVCaptureSession) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSession:"), value)
}


// A value that indicates how the layer displays video content within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/videoGravity
func (c_ CaptureVideoPreviewLayer) VideoGravity() LayerVideoGravity /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("videoGravity"))
	return rv
}


// A value that indicates how the layer displays video content within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/videoGravity
func (c_ CaptureVideoPreviewLayer) SetVideoGravity(value LayerVideoGravity /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoGravity:"), value)
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







