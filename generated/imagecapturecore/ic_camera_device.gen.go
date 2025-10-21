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

// A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/isaccessrestrictedappledevice
func (i_ ICCameraDevice) IsAccessRestrictedAppleDevice() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isAccessRestrictedAppleDevice"))
	return rv
}


// SetIsAccessRestrictedAppleDevice sets the value of the isAccessRestrictedAppleDevice property.
// A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/isaccessrestrictedappledevice
func (i_ ICCameraDevice) SetIsAccessRestrictedAppleDevice(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsAccessRestrictedAppleDevice:"), value)
}

// A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/isejectable
func (i_ ICCameraDevice) IsEjectable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEjectable"))
	return rv
}


// SetIsEjectable sets the value of the isEjectable property.
// A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/isejectable
func (i_ ICCameraDevice) SetIsEjectable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEjectable:"), value)
}

// All image, movie, and audio files stored on the camera, in an order that reflects the camera’s storage folder structure.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/contents
func (i_ ICCameraDevice) Contents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("contents"))
	return rv
}


// SetContents sets the value of the contents property.
// All image, movie, and audio files stored on the camera, in an order that reflects the camera’s storage folder structure.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/contents
func (i_ ICCameraDevice) SetContents(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContents:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mediapresentation
func (i_ ICCameraDevice) MediaPresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mediaPresentation"))
	return rv
}


// SetMediaPresentation sets the value of the mediaPresentation property.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mediapresentation
func (i_ ICCameraDevice) SetMediaPresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaPresentation:"), value)
}

// All image, movie and audio files stored on the camera, without regard to the camera’s storage folder structure.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mediafiles
func (i_ ICCameraDevice) MediaFiles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mediaFiles"))
	return rv
}


// SetMediaFiles sets the value of the mediaFiles property.
// All image, movie and audio files stored on the camera, without regard to the camera’s storage folder structure.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mediafiles
func (i_ ICCameraDevice) SetMediaFiles(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaFiles:"), value)
}

// The battery charge level.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/batterylevel
func (i_ ICCameraDevice) BatteryLevel() int {
	rv := objc.Send[int](i_.ID, objc.Sel("batteryLevel"))
	return rv
}


// SetBatteryLevel sets the value of the batteryLevel property.
// The battery charge level.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/batterylevel
func (i_ ICCameraDevice) SetBatteryLevel(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBatteryLevel:"), value)
}

// The time offset, in seconds, between the camera’s clock and the computer’s clock.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/timeoffset
func (i_ ICCameraDevice) TimeOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("timeOffset"))
	return rv
}


// SetTimeOffset sets the value of the timeOffset property.
// The time offset, in seconds, between the camera’s clock and the computer’s clock.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/timeoffset
func (i_ ICCameraDevice) SetTimeOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTimeOffset:"), value)
}

// A closure for handling PTP event packets.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/ptpeventhandler
func (i_ ICCameraDevice) PtpEventHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("ptpEventHandler"))
	return rv
}


// SetPtpEventHandler sets the value of the ptpEventHandler property.
// A closure for handling PTP event packets.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/ptpeventhandler
func (i_ ICCameraDevice) SetPtpEventHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPtpEventHandler:"), value)
}

// A Boolean value indicating whether the iCloud Photo Library is enabled on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/icloudphotosenabled
func (i_ ICCameraDevice) ICloudPhotosEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("iCloudPhotosEnabled"))
	return rv
}


// SetICloudPhotosEnabled sets the value of the iCloudPhotosEnabled property.
// A Boolean value indicating whether the iCloud Photo Library is enabled on the device.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/icloudphotosenabled
func (i_ ICCameraDevice) SetICloudPhotosEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setICloudPhotosEnabled:"), value)
}

// A Boolean value indicating whether the device is locked, preventing deletion of any asset.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/islocked
func (i_ ICCameraDevice) IsLocked() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isLocked"))
	return rv
}


// SetIsLocked sets the value of the isLocked property.
// A Boolean value indicating whether the device is locked, preventing deletion of any asset.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/islocked
func (i_ ICCameraDevice) SetIsLocked(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsLocked:"), value)
}

// The percentage of the camera’s content that has been catalogued.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/contentcatalogpercentcompleted
func (i_ ICCameraDevice) ContentCatalogPercentCompleted() int {
	rv := objc.Send[int](i_.ID, objc.Sel("contentCatalogPercentCompleted"))
	return rv
}


// SetContentCatalogPercentCompleted sets the value of the contentCatalogPercentCompleted property.
// The percentage of the camera’s content that has been catalogued.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/contentcatalogpercentcompleted
func (i_ ICCameraDevice) SetContentCatalogPercentCompleted(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentCatalogPercentCompleted:"), value)
}

// A Boolean value that indicates whether the battery charge level is available.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/batterylevelavailable
func (i_ ICCameraDevice) BatteryLevelAvailable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("batteryLevelAvailable"))
	return rv
}


// SetBatteryLevelAvailable sets the value of the batteryLevelAvailable property.
// A Boolean value that indicates whether the battery charge level is available.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/batterylevelavailable
func (i_ ICCameraDevice) SetBatteryLevelAvailable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBatteryLevelAvailable:"), value)
}

// A Boolean value indicating whether tethered capture is enabled on the camera.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/tetheredcaptureenabled
func (i_ ICCameraDevice) TetheredCaptureEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("tetheredCaptureEnabled"))
	return rv
}


// SetTetheredCaptureEnabled sets the value of the tetheredCaptureEnabled property.
// A Boolean value indicating whether tethered capture is enabled on the camera.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/tetheredcaptureenabled
func (i_ ICCameraDevice) SetTetheredCaptureEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTetheredCaptureEnabled:"), value)
}

// The file system mount point for a camera using the mass storage transport type.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mountpoint
func (i_ ICCameraDevice) MountPoint() string {
	rv := objc.Send[string](i_.ID, objc.Sel("mountPoint"))
	return rv
}


// SetMountPoint sets the value of the mountPoint property.
// The file system mount point for a camera using the mass storage transport type.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mountpoint
func (i_ ICCameraDevice) SetMountPoint(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMountPoint:"), objc.String(value))
}



