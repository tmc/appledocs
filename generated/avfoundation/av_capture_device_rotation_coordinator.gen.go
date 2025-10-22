// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [CaptureDeviceRotationCoordinator] class.
var (
	CaptureDeviceRotationCoordinatorClass     _CaptureDeviceRotationCoordinatorClass
	CaptureDeviceRotationCoordinatorClassOnce sync.Once
)

func getCaptureDeviceRotationCoordinatorClass() _CaptureDeviceRotationCoordinatorClass {
	CaptureDeviceRotationCoordinatorClassOnce.Do(func() {
		CaptureDeviceRotationCoordinatorClass = _CaptureDeviceRotationCoordinatorClass{objc.GetClass("AVCaptureDeviceRotationCoordinator")}
	})
	return CaptureDeviceRotationCoordinatorClass
}

type _CaptureDeviceRotationCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [CaptureDeviceRotationCoordinator] class.
type ICaptureDeviceRotationCoordinator interface {
	objectivec.IObject
	Device() AVCaptureDevice
	SetDevice(value IAVCaptureDevice)
	PreviewLayer() quartzcore.Layer
	SetPreviewLayer(value quartzcore.ILayer)
	VideoRotationAngleForHorizonLevelCapture() float64
	SetVideoRotationAngleForHorizonLevelCapture(value float64)
	VideoRotationAngleForHorizonLevelPreview() float64
	SetVideoRotationAngleForHorizonLevelPreview(value float64)
}

// A class that monitors the physical orientation of a capture device and provides adjustment angles to keep images level, relative to gravity.
//
// Correctly rotate the photos and movies your app captures, and optionally, a live camera preview, by applying a coordinator’s and properties, respectively. Each rotation coordinator instance updates its properties so that your app can observe them and immediately apply them to the relevant components.


// A class that monitors the physical orientation of a capture device and provides adjustment angles to keep images level, relative to gravity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator

type CaptureDeviceRotationCoordinator struct {
	objectivec.Object
}

// CaptureDeviceRotationCoordinatorFrom constructs a [CaptureDeviceRotationCoordinator] from an unsafe.Pointer.
//
// A class that monitors the physical orientation of a capture device and provides adjustment angles to keep images level, relative to gravity.
func CaptureDeviceRotationCoordinatorFrom(ptr unsafe.Pointer) CaptureDeviceRotationCoordinator {
	return CaptureDeviceRotationCoordinator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceRotationCoordinatorClass) Alloc() CaptureDeviceRotationCoordinator {
	rv := objc.Send[CaptureDeviceRotationCoordinator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureDeviceRotationCoordinatorClass) New() CaptureDeviceRotationCoordinator {
	rv := objc.Send[CaptureDeviceRotationCoordinator](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDeviceRotationCoordinator) Init() CaptureDeviceRotationCoordinator {
	rv := objc.Send[CaptureDeviceRotationCoordinator](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDeviceRotationCoordinator) Autorelease() CaptureDeviceRotationCoordinator {
	rv := objc.Send[CaptureDeviceRotationCoordinator](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDeviceRotationCoordinator creates a new CaptureDeviceRotationCoordinator instance.
func NewCaptureDeviceRotationCoordinator() CaptureDeviceRotationCoordinator {
	return getCaptureDeviceRotationCoordinatorClass().New()
}




// Creates a coordinator that provides separate compensation angles for content your app takes with a capture device, and for your app’s camera preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/init(device:previewLayer:)

func NewCaptureDeviceRotationCoordinatorWithDevicePreviewLayer(device IAVCaptureDevice, previewLayer quartzcore.ILayer) CaptureDeviceRotationCoordinator {
	instance := getCaptureDeviceRotationCoordinatorClass().Alloc()
	rv := objc.Send[CaptureDeviceRotationCoordinator](instance.ID, objc.Sel("initWithDevice:previewLayer:"), device, previewLayer)
	rv.Autorelease()
	return rv
}



// The capture device the coordinator monitors to track its physical rotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/rotationcoordinator/device

func (c_ CaptureDeviceRotationCoordinator) Device() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("device"))
	return rv
}


// The capture device the coordinator monitors to track its physical rotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/rotationcoordinator/device

func (c_ CaptureDeviceRotationCoordinator) SetDevice(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDevice:"), value)
}


// The layer that displays a camera preview the coordinator calculates a video rotation angle for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/rotationcoordinator/previewlayer

func (c_ CaptureDeviceRotationCoordinator) PreviewLayer() quartzcore.Layer {
	rv := objc.Send[quartzcore.Layer](c_.ID, objc.Sel("previewLayer"))
	return rv
}


// The layer that displays a camera preview the coordinator calculates a video rotation angle for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/rotationcoordinator/previewlayer

func (c_ CaptureDeviceRotationCoordinator) SetPreviewLayer(value quartzcore.ILayer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewLayer:"), value)
}


// An angle the coordinator provides your app to apply to photos or videos it captures with the device so that they’re level relative to gravity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/rotationcoordinator/videorotationangleforhorizonlevelcapture

func (c_ CaptureDeviceRotationCoordinator) VideoRotationAngleForHorizonLevelCapture() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoRotationAngleForHorizonLevelCapture"))
	return rv
}


// An angle the coordinator provides your app to apply to photos or videos it captures with the device so that they’re level relative to gravity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/rotationcoordinator/videorotationangleforhorizonlevelcapture

func (c_ CaptureDeviceRotationCoordinator) SetVideoRotationAngleForHorizonLevelCapture(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoRotationAngleForHorizonLevelCapture:"), value)
}


// An angle the coordinator provides your app to apply to the preview layer so that it’s level relative to gravity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/rotationcoordinator/videorotationangleforhorizonlevelpreview

func (c_ CaptureDeviceRotationCoordinator) VideoRotationAngleForHorizonLevelPreview() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoRotationAngleForHorizonLevelPreview"))
	return rv
}


// An angle the coordinator provides your app to apply to the preview layer so that it’s level relative to gravity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/rotationcoordinator/videorotationangleforhorizonlevelpreview

func (c_ CaptureDeviceRotationCoordinator) SetVideoRotationAngleForHorizonLevelPreview(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoRotationAngleForHorizonLevelPreview:"), value)
}


