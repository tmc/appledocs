// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ICCameraDevice] class.
var (
	ICCameraDeviceClass     _ICCameraDeviceClass
	ICCameraDeviceClassOnce sync.Once
)

func getICCameraDeviceClass() _ICCameraDeviceClass {
	ICCameraDeviceClassOnce.Do(func() {
		ICCameraDeviceClass = _ICCameraDeviceClass{objc.GetClass("ICCameraDevice")}
	})
	return ICCameraDeviceClass
}

type _ICCameraDeviceClass struct {
	class objc.Class
}

// An interface definition for the [ICCameraDevice] class.
type IICCameraDevice interface {
	objectivec.IObject
	CancelDelete()
	RequestSyncClock()
}

// An object that represents a camera.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice
type ICCameraDevice struct {
	objectivec.Object
}

// ICCameraDeviceFrom constructs a [ICCameraDevice] from an unsafe.Pointer.
//
// An object that represents a camera.
func ICCameraDeviceFrom(ptr unsafe.Pointer) ICCameraDevice {
	return ICCameraDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ICCameraDeviceClass) Alloc() ICCameraDevice {
	rv := objc.Send[ICCameraDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ICCameraDeviceClass) New() ICCameraDevice {
	rv := objc.Send[ICCameraDevice](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICCameraDevice) Init() ICCameraDevice {
	rv := objc.Send[ICCameraDevice](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICCameraDevice) Autorelease() ICCameraDevice {
	rv := objc.Send[ICCameraDevice](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraDevice creates a new ICCameraDevice instance.
func NewICCameraDevice() ICCameraDevice {
	return getICCameraDeviceClass().New()
}


// Cancels the current delete operation.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/cancelDelete()
func (i_ ICCameraDevice) CancelDelete() {
	objc.Send[objc.ID](i_.ID, objc.Sel("cancelDelete"))
}

// Synchronizes the camera’s clock with the computer’s clock.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestSyncClock()
func (i_ ICCameraDevice) RequestSyncClock() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSyncClock"))
}



