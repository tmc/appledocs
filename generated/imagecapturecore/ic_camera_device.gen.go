// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	IICDevice
	// properties:
	BatteryLevel() int
	SetBatteryLevel(value int)
	BatteryLevelAvailable() bool
	SetBatteryLevelAvailable(value bool)
	ContentCatalogPercentCompleted() int
	SetContentCatalogPercentCompleted(value int)
	Contents() ICCameraItem
	SetContents(value ICCameraItem)
	ICloudPhotosEnabled() bool
	SetICloudPhotosEnabled(value bool)
	IsAccessRestrictedAppleDevice() bool
	SetIsAccessRestrictedAppleDevice(value bool)
	IsEjectable() bool
	SetIsEjectable(value bool)
	IsLocked() bool
	SetIsLocked(value bool)
	MediaFiles() ICCameraItem
	SetMediaFiles(value ICCameraItem)
	MediaPresentation() ICMediaPresentation
	SetMediaPresentation(value ICMediaPresentation)
	MountPoint() objc.IObject /* cross-framework: NSString */
	SetMountPoint(value objc.IObject /* cross-framework: NSString */)
	PtpEventHandler() unsafe.Pointer
	SetPtpEventHandler(value unsafe.Pointer)
	TetheredCaptureEnabled() bool
	SetTetheredCaptureEnabled(value bool)
	TimeOffset() float64
	SetTimeOffset(value float64)
	// methods:
	CancelDelete()
	RequestDeleteFiles(files []ICCameraItem)
}

// An object that represents a camera.


// An object that represents a camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice
type ICCameraDevice struct {
	ICDevice
}

// ICCameraDeviceFrom constructs a [ICCameraDevice] from an unsafe.Pointer.
//
// An object that represents a camera.
func ICCameraDeviceFrom(ptr unsafe.Pointer) ICCameraDevice {
	return ICCameraDevice{
		ICDevice: ICDeviceFrom(ptr),
	}
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/cancelDelete()
func (i_ ICCameraDevice) CancelDelete() {
	objc.Send[objc.ID](i_.ID, objc.Sel("cancelDelete"))
}


// Deletes files from the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestDeleteFiles(_:)
func (i_ ICCameraDevice) RequestDeleteFiles(files []ICCameraItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestDeleteFiles:"), files)
}


// The battery charge level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/batterylevel
func (i_ ICCameraDevice) BatteryLevel() int {
	rv := objc.Send[int](i_.ID, objc.Sel("batteryLevel"))
	return rv
}


// The battery charge level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/batterylevel
func (i_ ICCameraDevice) SetBatteryLevel(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBatteryLevel:"), value)
}


// A Boolean value that indicates whether the battery charge level is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/batterylevelavailable
func (i_ ICCameraDevice) BatteryLevelAvailable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("batteryLevelAvailable"))
	return rv
}


// A Boolean value that indicates whether the battery charge level is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/batterylevelavailable
func (i_ ICCameraDevice) SetBatteryLevelAvailable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBatteryLevelAvailable:"), value)
}


// The percentage of the camera’s content that has been catalogued.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/contentcatalogpercentcompleted
func (i_ ICCameraDevice) ContentCatalogPercentCompleted() int {
	rv := objc.Send[int](i_.ID, objc.Sel("contentCatalogPercentCompleted"))
	return rv
}


// The percentage of the camera’s content that has been catalogued.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/contentcatalogpercentcompleted
func (i_ ICCameraDevice) SetContentCatalogPercentCompleted(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentCatalogPercentCompleted:"), value)
}


// All image, movie, and audio files stored on the camera, in an order that reflects the camera’s storage folder structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/contents
func (i_ ICCameraDevice) Contents() ICCameraItem {
	rv := objc.Send[ICCameraItem](i_.ID, objc.Sel("contents"))
	return rv
}


// All image, movie, and audio files stored on the camera, in an order that reflects the camera’s storage folder structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/contents
func (i_ ICCameraDevice) SetContents(value ICCameraItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContents:"), value)
}


// A Boolean value indicating whether the iCloud Photo Library is enabled on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/icloudphotosenabled
func (i_ ICCameraDevice) ICloudPhotosEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("iCloudPhotosEnabled"))
	return rv
}


// A Boolean value indicating whether the iCloud Photo Library is enabled on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/icloudphotosenabled
func (i_ ICCameraDevice) SetICloudPhotosEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setICloudPhotosEnabled:"), value)
}


// A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/isaccessrestrictedappledevice
func (i_ ICCameraDevice) IsAccessRestrictedAppleDevice() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isAccessRestrictedAppleDevice"))
	return rv
}


// A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/isaccessrestrictedappledevice
func (i_ ICCameraDevice) SetIsAccessRestrictedAppleDevice(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsAccessRestrictedAppleDevice:"), value)
}


// A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/isejectable
func (i_ ICCameraDevice) IsEjectable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEjectable"))
	return rv
}


// A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/isejectable
func (i_ ICCameraDevice) SetIsEjectable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEjectable:"), value)
}


// A Boolean value indicating whether the device is locked, preventing deletion of any asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/islocked
func (i_ ICCameraDevice) IsLocked() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isLocked"))
	return rv
}


// A Boolean value indicating whether the device is locked, preventing deletion of any asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/islocked
func (i_ ICCameraDevice) SetIsLocked(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsLocked:"), value)
}


// All image, movie and audio files stored on the camera, without regard to the camera’s storage folder structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mediafiles
func (i_ ICCameraDevice) MediaFiles() ICCameraItem {
	rv := objc.Send[ICCameraItem](i_.ID, objc.Sel("mediaFiles"))
	return rv
}


// All image, movie and audio files stored on the camera, without regard to the camera’s storage folder structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mediafiles
func (i_ ICCameraDevice) SetMediaFiles(value ICCameraItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaFiles:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mediapresentation
func (i_ ICCameraDevice) MediaPresentation() ICMediaPresentation {
	rv := objc.Send[ICMediaPresentation](i_.ID, objc.Sel("mediaPresentation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mediapresentation
func (i_ ICCameraDevice) SetMediaPresentation(value ICMediaPresentation) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaPresentation:"), value)
}


// The file system mount point for a camera using the mass storage transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mountpoint
func (i_ ICCameraDevice) MountPoint() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("mountPoint"))
	return rv
}


// The file system mount point for a camera using the mass storage transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/mountpoint
func (i_ ICCameraDevice) SetMountPoint(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMountPoint:"), value)
}


// A closure for handling PTP event packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/ptpeventhandler
func (i_ ICCameraDevice) PtpEventHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("ptpEventHandler"))
	return rv
}


// A closure for handling PTP event packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/ptpeventhandler
func (i_ ICCameraDevice) SetPtpEventHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPtpEventHandler:"), value)
}


// A Boolean value indicating whether tethered capture is enabled on the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/tetheredcaptureenabled
func (i_ ICCameraDevice) TetheredCaptureEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("tetheredCaptureEnabled"))
	return rv
}


// A Boolean value indicating whether tethered capture is enabled on the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/tetheredcaptureenabled
func (i_ ICCameraDevice) SetTetheredCaptureEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTetheredCaptureEnabled:"), value)
}


// The time offset, in seconds, between the camera’s clock and the computer’s clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/timeoffset
func (i_ ICCameraDevice) TimeOffset() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("timeOffset"))
	return rv
}


// The time offset, in seconds, between the camera’s clock and the computer’s clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/timeoffset
func (i_ ICCameraDevice) SetTimeOffset(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTimeOffset:"), value)
}



