// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ICCameraFile */


/* debug [class_header]: Header for ICCameraFile */
// The class instance for the [ICCameraFile] class.
var (
	ICCameraFileClass     _ICCameraFileClass
	ICCameraFileClassOnce sync.Once
)

func getICCameraFileClass() _ICCameraFileClass {
	ICCameraFileClassOnce.Do(func() {
		ICCameraFileClass = _ICCameraFileClass{objc.GetClass("ICCameraFile")}
	})
	return ICCameraFileClass
}

type _ICCameraFileClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICCameraFile */
// An interface definition for the [ICCameraFile] class.
type IICCameraFile interface {
	IICCameraItem
	
/* debug [class_interface_properties]: Properties for ICCameraFile */
	// properties:
	Duration() unsafe.Pointer
	SetDuration(value unsafe.Pointer)
	Orientation() unsafe.Pointer
	SetOrientation(value unsafe.Pointer)
	SidecarFiles() ICCameraItem
	SetSidecarFiles(value ICCameraItem)
	FileSize() unsafe.Pointer
	SetFileSize(value unsafe.Pointer)
	BurstFavorite() unsafe.Pointer
	SetBurstFavorite(value unsafe.Pointer)
	BurstPicked() unsafe.Pointer
	SetBurstPicked(value unsafe.Pointer)
	BurstUUID() unsafe.Pointer
	SetBurstUUID(value unsafe.Pointer)
	CreatedFilename() unsafe.Pointer
	SetCreatedFilename(value unsafe.Pointer)
	ExifCreationDate() unsafe.Pointer
	SetExifCreationDate(value unsafe.Pointer)
	ExifModificationDate() unsafe.Pointer
	SetExifModificationDate(value unsafe.Pointer)
	FileCreationDate() unsafe.Pointer
	SetFileCreationDate(value unsafe.Pointer)
	FileModificationDate() unsafe.Pointer
	SetFileModificationDate(value unsafe.Pointer)
	FirstPicked() unsafe.Pointer
	SetFirstPicked(value unsafe.Pointer)
	GpsString() unsafe.Pointer
	SetGpsString(value unsafe.Pointer)
	GroupUUID() unsafe.Pointer
	SetGroupUUID(value unsafe.Pointer)
	Height() unsafe.Pointer
	SetHeight(value unsafe.Pointer)
	HighFramerate() unsafe.Pointer
	SetHighFramerate(value unsafe.Pointer)
	OriginalFilename() unsafe.Pointer
	SetOriginalFilename(value unsafe.Pointer)
	OriginatingAssetID() unsafe.Pointer
	SetOriginatingAssetID(value unsafe.Pointer)
	PairedRawImage() ICCameraFile
	SetPairedRawImage(value ICCameraFile)
	RelatedUUID() unsafe.Pointer
	SetRelatedUUID(value unsafe.Pointer)
	TimeLapse() unsafe.Pointer
	SetTimeLapse(value unsafe.Pointer)
	Width() unsafe.Pointer
	SetWidth(value unsafe.Pointer)
	Fingerprint() unsafe.Pointer
	SetFingerprint(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICCameraFile */
	// methods:
	RequestMetadataDictionaryWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer, unsafe.Pointer))
	RequestReadDataAtOffsetLengthCompletion(offset unsafe.Pointer, length unsafe.Pointer, completion func(unsafe.Pointer, unsafe.Pointer))
	RequestThumbnailDataWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer, unsafe.Pointer))
	RequestDownload()
	RequestDownloadWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer, unsafe.Pointer)) foundation.Progress
	RequestSecurityScopedURLWithCompletion(completion func(unsafe.Pointer, unsafe.Pointer))
	RequestFingerprintWithCompletion(completion func(unsafe.Pointer, unsafe.Pointer))
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICCameraFile */
// Alloc allocates a new instance without initialization.
func (ic _ICCameraFileClass) Alloc() ICCameraFile {
	rv := objc.Send[ICCameraFile](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICCameraFileClass) New() ICCameraFile {
	rv := objc.Send[ICCameraFile](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICCameraFile) Init() ICCameraFile {
	rv := objc.Send[ICCameraFile](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICCameraFile) Autorelease() ICCameraFile {
	rv := objc.Send[ICCameraFile](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraFile creates a new ICCameraFile instance.
func NewICCameraFile() ICCameraFile {
	return getICCameraFileClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICCameraFile */
// An object that represents a file on a camera.


// An object that represents a file on a camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile
type ICCameraFile struct {
	ICCameraItem
}

// ICCameraFileFrom constructs a [ICCameraFile] from an unsafe.Pointer.
//
// An object that represents a file on a camera.
func ICCameraFileFrom(ptr unsafe.Pointer) ICCameraFile {
	return ICCameraFile{
		ICCameraItem: ICCameraItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICCameraFile *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICCameraFile */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/4359698-fingerprintforfile
func (ic _ICCameraFileClass) FingerprintForFile() {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("fingerprintForFile"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FingerprintForFile) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/4359698-fingerprintforfileaturl
func (ic _ICCameraFileClass) FingerprintForFileAtURL(url foundation.URL) string {
	rv := objc.Send[string](objc.ID(ic.class), objc.Sel("fingerprintForFileAtURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FingerprintForFileAtURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICCameraFile */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICCameraFile */

// Requests metadata and executes the completion block in place of the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131468-requestmetadatadictionarywithopt
func (i_ ICCameraFile) RequestMetadataDictionaryWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestMetadataDictionaryWithOptions:completion:"), options, completion)
}/* debug [instance_methods/method]: RequestMetadataDictionaryWithOptionsCompletion */


// Requests to asynchronously read data of a specified length from a specified offset, then executes the completion block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131470-requestreaddataatoffset
func (i_ ICCameraFile) RequestReadDataAtOffsetLengthCompletion(offset unsafe.Pointer, length unsafe.Pointer, completion func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestReadDataAtOffset:length:completion:"), offset, length, completion)
}/* debug [instance_methods/method]: RequestReadDataAtOffsetLengthCompletion */


// Requests a thumbnail and executes the completion block in place of the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131471-requestthumbnaildatawithoptions
func (i_ ICCameraFile) RequestThumbnailDataWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestThumbnailDataWithOptions:completion:"), options, completion)
}/* debug [instance_methods/method]: RequestThumbnailDataWithOptionsCompletion */


// Requests a download and executes the completion block in place of the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3142906-requestdownload
func (i_ ICCameraFile) RequestDownload() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestDownload"))
}/* debug [instance_methods/method]: RequestDownload */


// Requests a download and executes the completion block in place of the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3142906-requestdownloadwithoptions
func (i_ ICCameraFile) RequestDownloadWithOptionsCompletion(options objc.IObject, completion func(unsafe.Pointer, unsafe.Pointer)) foundation.Progress {
	rv := objc.Send[foundation.Progress](i_.ID, objc.Sel("requestDownloadWithOptions:completion:"), options, completion)
	return rv
}/* debug [instance_methods/method]: RequestDownloadWithOptionsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/4149740-requestsecurityscopedurlwithcomp
func (i_ ICCameraFile) RequestSecurityScopedURLWithCompletion(completion func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSecurityScopedURLWithCompletion:"), completion)
}/* debug [instance_methods/method]: RequestSecurityScopedURLWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/4359699-requestfingerprintwithcompletion
func (i_ ICCameraFile) RequestFingerprintWithCompletion(completion func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestFingerprintWithCompletion:"), completion)
}/* debug [instance_methods/method]: RequestFingerprintWithCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICCameraFile */

// The duration, in seconds, of an audio or video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/1388981-duration
func (i_ ICCameraFile) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The duration, in seconds, of an audio or video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/1388981-duration
func (i_ ICCameraFile) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// The orientation to use when downloading the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/1388983-orientation
func (i_ ICCameraFile) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("orientation"))
	return rv
}/* debug [instance_properties/getter]: orientation */


// The orientation to use when downloading the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/1388983-orientation
func (i_ ICCameraFile) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOrientation:"), value)
}/* debug [instance_properties/setter]: orientation */


// An array of two camera files associated with this file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/1389003-sidecarfiles
func (i_ ICCameraFile) SidecarFiles() ICCameraItem {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("sidecarFiles"))
	return rv
}/* debug [instance_properties/getter]: sidecarFiles */


// An array of two camera files associated with this file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/1389003-sidecarfiles
func (i_ ICCameraFile) SetSidecarFiles(value ICCameraItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSidecarFiles:"), value)
}/* debug [instance_properties/setter]: sidecarFiles */


// The size of the file, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/1389013-filesize
func (i_ ICCameraFile) FileSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fileSize"))
	return rv
}/* debug [instance_properties/getter]: fileSize */


// The size of the file, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/1389013-filesize
func (i_ ICCameraFile) SetFileSize(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileSize:"), value)
}/* debug [instance_properties/setter]: fileSize */


// A Boolean value that indicates this file is the burst favorite in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131450-burstfavorite
func (i_ ICCameraFile) BurstFavorite() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("burstFavorite"))
	return rv
}/* debug [instance_properties/getter]: burstFavorite */


// A Boolean value that indicates this file is the burst favorite in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131450-burstfavorite
func (i_ ICCameraFile) SetBurstFavorite(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBurstFavorite:"), value)
}/* debug [instance_properties/setter]: burstFavorite */


// A Boolean value that indicates whether this file is user picked in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131451-burstpicked
func (i_ ICCameraFile) BurstPicked() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("burstPicked"))
	return rv
}/* debug [instance_properties/getter]: burstPicked */


// A Boolean value that indicates whether this file is user picked in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131451-burstpicked
func (i_ ICCameraFile) SetBurstPicked(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBurstPicked:"), value)
}/* debug [instance_properties/setter]: burstPicked */


// The burst UUID of the file if it is in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131452-burstuuid
func (i_ ICCameraFile) BurstUUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("burstUUID"))
	return rv
}/* debug [instance_properties/getter]: burstUUID */


// The burst UUID of the file if it is in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131452-burstuuid
func (i_ ICCameraFile) SetBurstUUID(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBurstUUID:"), value)
}/* debug [instance_properties/setter]: burstUUID */


// The created name of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131453-createdfilename
func (i_ ICCameraFile) CreatedFilename() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("createdFilename"))
	return rv
}/* debug [instance_properties/getter]: createdFilename */


// The created name of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131453-createdfilename
func (i_ ICCameraFile) SetCreatedFilename(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCreatedFilename:"), value)
}/* debug [instance_properties/setter]: createdFilename */


// The creation date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131454-exifcreationdate
func (i_ ICCameraFile) ExifCreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("exifCreationDate"))
	return rv
}/* debug [instance_properties/getter]: exifCreationDate */


// The creation date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131454-exifcreationdate
func (i_ ICCameraFile) SetExifCreationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExifCreationDate:"), value)
}/* debug [instance_properties/setter]: exifCreationDate */


// The modification date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131455-exifmodificationdate
func (i_ ICCameraFile) ExifModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("exifModificationDate"))
	return rv
}/* debug [instance_properties/getter]: exifModificationDate */


// The modification date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131455-exifmodificationdate
func (i_ ICCameraFile) SetExifModificationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExifModificationDate:"), value)
}/* debug [instance_properties/setter]: exifModificationDate */


// The creation date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131456-filecreationdate
func (i_ ICCameraFile) FileCreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fileCreationDate"))
	return rv
}/* debug [instance_properties/getter]: fileCreationDate */


// The creation date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131456-filecreationdate
func (i_ ICCameraFile) SetFileCreationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileCreationDate:"), value)
}/* debug [instance_properties/setter]: fileCreationDate */


// The modification date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131457-filemodificationdate
func (i_ ICCameraFile) FileModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fileModificationDate"))
	return rv
}/* debug [instance_properties/getter]: fileModificationDate */


// The modification date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131457-filemodificationdate
func (i_ ICCameraFile) SetFileModificationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileModificationDate:"), value)
}/* debug [instance_properties/setter]: fileModificationDate */


// A Boolean value that indicates whether a file is autopicked by Photos to represent the burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131458-firstpicked
func (i_ ICCameraFile) FirstPicked() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("firstPicked"))
	return rv
}/* debug [instance_properties/getter]: firstPicked */


// A Boolean value that indicates whether a file is autopicked by Photos to represent the burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131458-firstpicked
func (i_ ICCameraFile) SetFirstPicked(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFirstPicked:"), value)
}/* debug [instance_properties/setter]: firstPicked */


// The GPS String of the file in standard format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131459-gpsstring
func (i_ ICCameraFile) GpsString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("gpsString"))
	return rv
}/* debug [instance_properties/getter]: gpsString */


// The GPS String of the file in standard format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131459-gpsstring
func (i_ ICCameraFile) SetGpsString(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGpsString:"), value)
}/* debug [instance_properties/setter]: gpsString */


// The group of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131460-groupuuid
func (i_ ICCameraFile) GroupUUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("groupUUID"))
	return rv
}/* debug [instance_properties/getter]: groupUUID */


// The group of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131460-groupuuid
func (i_ ICCameraFile) SetGroupUUID(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupUUID:"), value)
}/* debug [instance_properties/setter]: groupUUID */


// The height of an image or movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131461-height
func (i_ ICCameraFile) Height() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// The height of an image or movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131461-height
func (i_ ICCameraFile) SetHeight(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// A Boolean value that indicates whether the file is a slow motion or high-frame-rate video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131462-highframerate
func (i_ ICCameraFile) HighFramerate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("highFramerate"))
	return rv
}/* debug [instance_properties/getter]: highFramerate */


// A Boolean value that indicates whether the file is a slow motion or high-frame-rate video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131462-highframerate
func (i_ ICCameraFile) SetHighFramerate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHighFramerate:"), value)
}/* debug [instance_properties/setter]: highFramerate */


// The original name of the file on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131463-originalfilename
func (i_ ICCameraFile) OriginalFilename() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("originalFilename"))
	return rv
}/* debug [instance_properties/getter]: originalFilename */


// The original name of the file on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131463-originalfilename
func (i_ ICCameraFile) SetOriginalFilename(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOriginalFilename:"), value)
}/* debug [instance_properties/setter]: originalFilename */


// The originating asset ID of an or file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131464-originatingassetid
func (i_ ICCameraFile) OriginatingAssetID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("originatingAssetID"))
	return rv
}/* debug [instance_properties/getter]: originatingAssetID */


// The originating asset ID of an or file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131464-originatingassetid
func (i_ ICCameraFile) SetOriginatingAssetID(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOriginatingAssetID:"), value)
}/* debug [instance_properties/setter]: originatingAssetID */


// A sidecar file containing the logical compliment of a or other two-format image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131465-pairedrawimage
func (i_ ICCameraFile) PairedRawImage() ICCameraFile {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("pairedRawImage"))
	return rv
}/* debug [instance_properties/getter]: pairedRawImage */


// A sidecar file containing the logical compliment of a or other two-format image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131465-pairedrawimage
func (i_ ICCameraFile) SetPairedRawImage(value ICCameraFile) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPairedRawImage:"), value)
}/* debug [instance_properties/setter]: pairedRawImage */


// A related UUID correlating several images from an Apple device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131466-relateduuid
func (i_ ICCameraFile) RelatedUUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("relatedUUID"))
	return rv
}/* debug [instance_properties/getter]: relatedUUID */


// A related UUID correlating several images from an Apple device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131466-relateduuid
func (i_ ICCameraFile) SetRelatedUUID(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRelatedUUID:"), value)
}/* debug [instance_properties/setter]: relatedUUID */


// A Boolean value that indicates whether the file is a time-lapse video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131473-timelapse
func (i_ ICCameraFile) TimeLapse() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("timeLapse"))
	return rv
}/* debug [instance_properties/getter]: timeLapse */


// A Boolean value that indicates whether the file is a time-lapse video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131473-timelapse
func (i_ ICCameraFile) SetTimeLapse(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTimeLapse:"), value)
}/* debug [instance_properties/setter]: timeLapse */


// The width of an image or movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131474-width
func (i_ ICCameraFile) Width() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// The width of an image or movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/3131474-width
func (i_ ICCameraFile) SetWidth(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/4359697-fingerprint
func (i_ ICCameraFile) Fingerprint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fingerprint"))
	return rv
}/* debug [instance_properties/getter]: fingerprint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/4359697-fingerprint
func (i_ ICCameraFile) SetFingerprint(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFingerprint:"), value)
}/* debug [instance_properties/setter]: fingerprint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICCameraFile */



