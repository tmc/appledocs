// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
}

// A class that monitors the physical orientation of a capture device and provides adjustment angles to keep images level, relative to gravity.
//
// Correctly rotate the photos and movies your app captures, and optionally, a live camera preview, by applying a coordinator’s and properties, respectively. Each rotation coordinator instance updates its properties so that your app can observe them and immediately apply them to the relevant components.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/RotationCoordinator/init(device:previewLayer:)
func NewCaptureDeviceRotationCoordinatorWithDevicePreviewLayer(device unsafe.Pointer, previewLayer unsafe.Pointer) CaptureDeviceRotationCoordinator {
	instance := getCaptureDeviceRotationCoordinatorClass().Alloc()
	rv := objc.Send[CaptureDeviceRotationCoordinator](instance.ID, objc.Sel("initWithDevice:previewLayer:"), device, previewLayer)
	rv.Autorelease()
	return rv
}



