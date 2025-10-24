// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ICCameraDevice */


/* debug [class_header]: Header for ICCameraDevice */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICCameraDevice */
// An interface definition for the [ICCameraDevice] class.
type IICCameraDevice interface {
	IICDevice
	
/* debug [class_interface_properties]: Properties for ICCameraDevice */
	// properties:
	IsAccessRestrictedAppleDevice() unsafe.Pointer
	SetIsAccessRestrictedAppleDevice(value unsafe.Pointer)
	MediaFiles() ICCameraItem
	SetMediaFiles(value ICCameraItem)
	TimeOffset() unsafe.Pointer
	SetTimeOffset(value unsafe.Pointer)
	MountPoint() unsafe.Pointer
	SetMountPoint(value unsafe.Pointer)
	ContentCatalogPercentCompleted() unsafe.Pointer
	SetContentCatalogPercentCompleted(value unsafe.Pointer)
	TetheredCaptureEnabled() unsafe.Pointer
	SetTetheredCaptureEnabled(value unsafe.Pointer)
	BatteryLevel() unsafe.Pointer
	SetBatteryLevel(value unsafe.Pointer)
	BatteryLevelAvailable() unsafe.Pointer
	SetBatteryLevelAvailable(value unsafe.Pointer)
	Contents() ICCameraItem
	SetContents(value ICCameraItem)
	IsEjectable() unsafe.Pointer
	SetIsEjectable(value unsafe.Pointer)
	ICloudPhotosEnabled() unsafe.Pointer
	SetICloudPhotosEnabled(value unsafe.Pointer)
	IsLocked() unsafe.Pointer
	SetIsLocked(value unsafe.Pointer)
	PtpEventHandler() unsafe.Pointer
	SetPtpEventHandler(value unsafe.Pointer)
	MediaPresentation() unsafe.Pointer
	SetMediaPresentation(value unsafe.Pointer)
	AccessRestrictedAppleDevice() bool
	Ejectable() bool
	Locked() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICCameraDevice */
	// methods:
	CancelDownload()
	RequestDownloadFile()
	RequestTakePicture()
	RequestDeleteFiles()
	CancelDelete()
	RequestSendPTPCommand()
	RequestSyncClock()
	RequestReadData()
	RequestReadDataFromFileAtOffsetLengthReadDelegateDidReadDataSelectorContextInfo(file ICCameraFile, offset unsafe.Pointer, length unsafe.Pointer, readDelegate objc.IObject, selector objc.SEL, contextInfo unsafe.Pointer)
	Files()
	FilesOfType(fileUTType string) unsafe.Pointer
	RequestDeleteFilesWithFiles(files []CCameraItem)
	RequestDeleteFilesDeleteFailedCompletion(files []CCameraItem, deleteFailed unsafe.Pointer, completion unsafe.Pointer) foundation.Progress
	RequestDownloadFileOptionsDownloadDelegateDidDownloadSelectorContextInfo(file ICCameraFile, options foundation.IDictionary, downloadDelegate unsafe.Pointer, selector objc.SEL, contextInfo unsafe.Pointer)
	RequestSendPTPCommandOutDataCompletion(ptpCommand objc.IObject /* cross-framework: NSData */, ptpData objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer)
	RequestSendPTPCommandOutDataSendCommandDelegateDidSendCommandSelectorContextInfo(command objc.IObject /* cross-framework: NSData */, data objc.IObject /* cross-framework: NSData */, sendCommandDelegate objc.IObject, selector objc.SEL, contextInfo unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICCameraDevice */
// Alloc allocates a new instance without initialization.
func (ic _ICCameraDeviceClass) Alloc() ICCameraDevice {
	rv := objc.Send[ICCameraDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICCameraDevice */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICCameraDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICCameraDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICCameraDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICCameraDevice */

// Cancels a download from the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507628-canceldownload
func (i_ ICCameraDevice) CancelDownload() {
	objc.Send[objc.ID](i_.ID, objc.Sel("cancelDownload"))
}/* debug [instance_methods/method]: CancelDownload */


// Downloads a file from the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507818-requestdownloadfile
func (i_ ICCameraDevice) RequestDownloadFile() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestDownloadFile"))
}/* debug [instance_methods/method]: RequestDownloadFile */


// Captures a new image using the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507903-requesttakepicture
func (i_ ICCameraDevice) RequestTakePicture() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestTakePicture"))
}/* debug [instance_methods/method]: RequestTakePicture */


// Deletes files from the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507921-requestdeletefiles
func (i_ ICCameraDevice) RequestDeleteFiles() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestDeleteFiles"))
}/* debug [instance_methods/method]: RequestDeleteFiles */


// Cancels the current delete operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507931-canceldelete
func (i_ ICCameraDevice) CancelDelete() {
	objc.Send[objc.ID](i_.ID, objc.Sel("cancelDelete"))
}/* debug [instance_methods/method]: CancelDelete */


// Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507967-requestsendptpcommand
func (i_ ICCameraDevice) RequestSendPTPCommand() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSendPTPCommand"))
}/* debug [instance_methods/method]: RequestSendPTPCommand */


// Synchronizes the camera’s clock with the computer’s clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508062-requestsyncclock
func (i_ ICCameraDevice) RequestSyncClock() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSyncClock"))
}/* debug [instance_methods/method]: RequestSyncClock */


// Asynchronously reads data of a specified length from a specified offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508079-requestreaddata
func (i_ ICCameraDevice) RequestReadData() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestReadData"))
}/* debug [instance_methods/method]: RequestReadData */


// Asynchronously reads data of a specified length from a specified offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508079-requestreaddatafromfile
func (i_ ICCameraDevice) RequestReadDataFromFileAtOffsetLengthReadDelegateDidReadDataSelectorContextInfo(file ICCameraFile, offset unsafe.Pointer, length unsafe.Pointer, readDelegate objc.IObject, selector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestReadDataFromFile:atOffset:length:readDelegate:didReadDataSelector:contextInfo:"), file, offset, length, readDelegate, selector, contextInfo)
}/* debug [instance_methods/method]: RequestReadDataFromFileAtOffsetLengthReadDelegateDidReadDataSelectorContextInfo */


// Returns an array of files of the selected type on the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508121-files
func (i_ ICCameraDevice) Files() {
	objc.Send[objc.ID](i_.ID, objc.Sel("files"))
}/* debug [instance_methods/method]: Files */


// Returns an array of files of the selected type on the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508121-filesoftype
func (i_ ICCameraDevice) FilesOfType(fileUTType string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("filesOfType:"), objc.String(fileUTType))
	return rv
}/* debug [instance_methods/method]: FilesOfType */


// Deletes files from the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestDeleteFiles(_:)
func (i_ ICCameraDevice) RequestDeleteFilesWithFiles(files []CCameraItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestDeleteFiles:"), files)
}/* debug [instance_methods/method]: RequestDeleteFilesWithFiles */


// Deletes files from the camera, with the ability to catch failures and execute a completion block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestDeleteFiles(_:deleteFailed:completion:)
func (i_ ICCameraDevice) RequestDeleteFilesDeleteFailedCompletion(files []CCameraItem, deleteFailed unsafe.Pointer, completion unsafe.Pointer) foundation.Progress {
	rv := objc.Send[foundation.Progress](i_.ID, objc.Sel("requestDeleteFiles:deleteFailed:completion:"), files, deleteFailed, completion)
	return rv
}/* debug [instance_methods/method]: RequestDeleteFilesDeleteFailedCompletion */


// Downloads a file from the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestDownloadFile(_:options:downloadDelegate:didDownloadSelector:contextInfo:)
func (i_ ICCameraDevice) RequestDownloadFileOptionsDownloadDelegateDidDownloadSelectorContextInfo(file ICCameraFile, options foundation.IDictionary, downloadDelegate unsafe.Pointer, selector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestDownloadFile:options:downloadDelegate:didDownloadSelector:contextInfo:"), file, options, downloadDelegate, selector, contextInfo)
}/* debug [instance_methods/method]: RequestDownloadFileOptionsDownloadDelegateDidDownloadSelectorContextInfo */


// Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestSendPTPCommand(_:outData:completion:)
func (i_ ICCameraDevice) RequestSendPTPCommandOutDataCompletion(ptpCommand objc.IObject /* cross-framework: NSData */, ptpData objc.IObject /* cross-framework: NSData */, completion unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSendPTPCommand:outData:completion:"), ptpCommand, ptpData, completion)
}/* debug [instance_methods/method]: RequestSendPTPCommandOutDataCompletion */


// Sends a Picture Transfer Protocol (PTP) command to a camera asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/requestSendPTPCommand(_:outData:sendCommandDelegate:didSendCommand:contextInfo:)
func (i_ ICCameraDevice) RequestSendPTPCommandOutDataSendCommandDelegateDidSendCommandSelectorContextInfo(command objc.IObject /* cross-framework: NSData */, data objc.IObject /* cross-framework: NSData */, sendCommandDelegate objc.IObject, selector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSendPTPCommand:outData:sendCommandDelegate:didSendCommandSelector:contextInfo:"), command, data, sendCommandDelegate, selector, contextInfo)
}/* debug [instance_methods/method]: RequestSendPTPCommandOutDataSendCommandDelegateDidSendCommandSelectorContextInfo */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICCameraDevice */

// A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507742-isaccessrestrictedappledevice
func (i_ ICCameraDevice) IsAccessRestrictedAppleDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isAccessRestrictedAppleDevice"))
	return rv
}/* debug [instance_properties/getter]: isAccessRestrictedAppleDevice */


// A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507742-isaccessrestrictedappledevice
func (i_ ICCameraDevice) SetIsAccessRestrictedAppleDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsAccessRestrictedAppleDevice:"), value)
}/* debug [instance_properties/setter]: isAccessRestrictedAppleDevice */


// All image, movie and audio files stored on the camera, without regard to the camera’s storage folder structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507811-mediafiles
func (i_ ICCameraDevice) MediaFiles() ICCameraItem {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("mediaFiles"))
	return rv
}/* debug [instance_properties/getter]: mediaFiles */


// All image, movie and audio files stored on the camera, without regard to the camera’s storage folder structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507811-mediafiles
func (i_ ICCameraDevice) SetMediaFiles(value ICCameraItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaFiles:"), value)
}/* debug [instance_properties/setter]: mediaFiles */


// The time offset, in seconds, between the camera’s clock and the computer’s clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507878-timeoffset
func (i_ ICCameraDevice) TimeOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("timeOffset"))
	return rv
}/* debug [instance_properties/getter]: timeOffset */


// The time offset, in seconds, between the camera’s clock and the computer’s clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507878-timeoffset
func (i_ ICCameraDevice) SetTimeOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTimeOffset:"), value)
}/* debug [instance_properties/setter]: timeOffset */


// The file system mount point for a camera using the mass storage transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507897-mountpoint
func (i_ ICCameraDevice) MountPoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mountPoint"))
	return rv
}/* debug [instance_properties/getter]: mountPoint */


// The file system mount point for a camera using the mass storage transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507897-mountpoint
func (i_ ICCameraDevice) SetMountPoint(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMountPoint:"), value)
}/* debug [instance_properties/setter]: mountPoint */


// The percentage of the camera’s content that has been catalogued.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507899-contentcatalogpercentcompleted
func (i_ ICCameraDevice) ContentCatalogPercentCompleted() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("contentCatalogPercentCompleted"))
	return rv
}/* debug [instance_properties/getter]: contentCatalogPercentCompleted */


// The percentage of the camera’s content that has been catalogued.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1507899-contentcatalogpercentcompleted
func (i_ ICCameraDevice) SetContentCatalogPercentCompleted(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentCatalogPercentCompleted:"), value)
}/* debug [instance_properties/setter]: contentCatalogPercentCompleted */


// A Boolean value indicating whether tethered capture is enabled on the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508004-tetheredcaptureenabled
func (i_ ICCameraDevice) TetheredCaptureEnabled() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("tetheredCaptureEnabled"))
	return rv
}/* debug [instance_properties/getter]: tetheredCaptureEnabled */


// A Boolean value indicating whether tethered capture is enabled on the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508004-tetheredcaptureenabled
func (i_ ICCameraDevice) SetTetheredCaptureEnabled(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTetheredCaptureEnabled:"), value)
}/* debug [instance_properties/setter]: tetheredCaptureEnabled */


// The battery charge level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508043-batterylevel
func (i_ ICCameraDevice) BatteryLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("batteryLevel"))
	return rv
}/* debug [instance_properties/getter]: batteryLevel */


// The battery charge level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508043-batterylevel
func (i_ ICCameraDevice) SetBatteryLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBatteryLevel:"), value)
}/* debug [instance_properties/setter]: batteryLevel */


// A Boolean value that indicates whether the battery charge level is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508053-batterylevelavailable
func (i_ ICCameraDevice) BatteryLevelAvailable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("batteryLevelAvailable"))
	return rv
}/* debug [instance_properties/getter]: batteryLevelAvailable */


// A Boolean value that indicates whether the battery charge level is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508053-batterylevelavailable
func (i_ ICCameraDevice) SetBatteryLevelAvailable(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBatteryLevelAvailable:"), value)
}/* debug [instance_properties/setter]: batteryLevelAvailable */


// All image, movie, and audio files stored on the camera, in an order that reflects the camera’s storage folder structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508088-contents
func (i_ ICCameraDevice) Contents() ICCameraItem {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// All image, movie, and audio files stored on the camera, in an order that reflects the camera’s storage folder structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/1508088-contents
func (i_ ICCameraDevice) SetContents(value ICCameraItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */


// A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3142890-isejectable
func (i_ ICCameraDevice) IsEjectable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isEjectable"))
	return rv
}/* debug [instance_properties/getter]: isEjectable */


// A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3142890-isejectable
func (i_ ICCameraDevice) SetIsEjectable(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEjectable:"), value)
}/* debug [instance_properties/setter]: isEjectable */


// A Boolean value indicating whether the iCloud Photo Library is enabled on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3142891-icloudphotosenabled
func (i_ ICCameraDevice) ICloudPhotosEnabled() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("iCloudPhotosEnabled"))
	return rv
}/* debug [instance_properties/getter]: iCloudPhotosEnabled */


// A Boolean value indicating whether the iCloud Photo Library is enabled on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3142891-icloudphotosenabled
func (i_ ICCameraDevice) SetICloudPhotosEnabled(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setICloudPhotosEnabled:"), value)
}/* debug [instance_properties/setter]: iCloudPhotosEnabled */


// A Boolean value indicating whether the device is locked, preventing deletion of any asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3142892-islocked
func (i_ ICCameraDevice) IsLocked() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isLocked"))
	return rv
}/* debug [instance_properties/getter]: isLocked */


// A Boolean value indicating whether the device is locked, preventing deletion of any asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3142892-islocked
func (i_ ICCameraDevice) SetIsLocked(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsLocked:"), value)
}/* debug [instance_properties/setter]: isLocked */


// A closure for handling PTP event packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3393297-ptpeventhandler
func (i_ ICCameraDevice) PtpEventHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("ptpEventHandler"))
	return rv
}/* debug [instance_properties/getter]: ptpEventHandler */


// A closure for handling PTP event packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3393297-ptpeventhandler
func (i_ ICCameraDevice) SetPtpEventHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPtpEventHandler:"), value)
}/* debug [instance_properties/setter]: ptpEventHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3601135-mediapresentation
func (i_ ICCameraDevice) MediaPresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mediaPresentation"))
	return rv
}/* debug [instance_properties/getter]: mediaPresentation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameradevice/3601135-mediapresentation
func (i_ ICCameraDevice) SetMediaPresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaPresentation:"), value)
}/* debug [instance_properties/setter]: mediaPresentation */


// A Boolean value indicating whether the device is an Apple device, passcode-locked, and connected to an untrusted host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/isAccessRestrictedAppleDevice
func (i_ ICCameraDevice) AccessRestrictedAppleDevice() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("accessRestrictedAppleDevice"))
	return rv
}/* debug [instance_properties/getter]: accessRestrictedAppleDevice */


// A Boolean value indicating whether the device can be ‘soft’ removed or disconnected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/isEjectable
func (i_ ICCameraDevice) Ejectable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("ejectable"))
	return rv
}/* debug [instance_properties/getter]: ejectable */


// A Boolean value indicating whether the device is locked, preventing deletion of any asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraDevice/isLocked
func (i_ ICCameraDevice) Locked() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("locked"))
	return rv
}/* debug [instance_properties/getter]: locked */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICCameraDevice */



