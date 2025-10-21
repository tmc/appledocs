// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ICCameraFile] class.
type IICCameraFile interface {
	IICCameraItem
	RequestSecurityScopedURLWithCompletion(completion unsafe.Pointer)
}

// An object that represents a file on a camera.
//
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

// Alloc allocates a new instance without initialization.
func (ic _ICCameraFileClass) Alloc() ICCameraFile {
	rv := objc.Send[ICCameraFile](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/requestSecurityScopedURL(completion:)
func (i_ ICCameraFile) RequestSecurityScopedURLWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSecurityScopedURLWithCompletion:"), completion)
}

// The burst UUID of the file if it is in a burst.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/burstUUID
func (i_ ICCameraFile) BurstUUID() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("burstUUID"))
	return rv
}

// The duration, in seconds, of an audio or video file.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/duration
func (i_ ICCameraFile) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("duration"))
	return rv
}

// The width of an image or movie frame.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/width
func (i_ ICCameraFile) Width() int {
	rv := objc.Send[int](i_.ID, objc.Sel("width"))
	return rv
}

// A Boolean value that indicates this file is the burst favorite in a burst.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/burstfavorite
func (i_ ICCameraFile) BurstFavorite() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("burstFavorite"))
	return rv
}


// SetBurstFavorite sets the value of the burstFavorite property.
// A Boolean value that indicates this file is the burst favorite in a burst.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/burstfavorite
func (i_ ICCameraFile) SetBurstFavorite(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBurstFavorite:"), value)
}

// A Boolean value that indicates whether this file is user picked in a burst.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/burstpicked
func (i_ ICCameraFile) BurstPicked() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("burstPicked"))
	return rv
}


// SetBurstPicked sets the value of the burstPicked property.
// A Boolean value that indicates whether this file is user picked in a burst.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/burstpicked
func (i_ ICCameraFile) SetBurstPicked(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBurstPicked:"), value)
}

// The created name of the file.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/createdfilename
func (i_ ICCameraFile) CreatedFilename() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("createdFilename"))
	return rv
}


// SetCreatedFilename sets the value of the createdFilename property.
// The created name of the file.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/createdfilename
func (i_ ICCameraFile) SetCreatedFilename(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCreatedFilename:"), value)
}

// The
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/exifcreationdate
func (i_ ICCameraFile) ExifCreationDate() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("exifCreationDate"))
	return rv
}


// SetExifCreationDate sets the value of the exifCreationDate property.
// The

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/exifcreationdate
func (i_ ICCameraFile) SetExifCreationDate(value foundation.IDate) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExifCreationDate:"), value)
}

// The
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/exifmodificationdate
func (i_ ICCameraFile) ExifModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("exifModificationDate"))
	return rv
}


// SetExifModificationDate sets the value of the exifModificationDate property.
// The

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/exifmodificationdate
func (i_ ICCameraFile) SetExifModificationDate(value foundation.IDate) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExifModificationDate:"), value)
}

// The creation date of the file.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filecreationdate
func (i_ ICCameraFile) FileCreationDate() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("fileCreationDate"))
	return rv
}


// SetFileCreationDate sets the value of the fileCreationDate property.
// The creation date of the file.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filecreationdate
func (i_ ICCameraFile) SetFileCreationDate(value foundation.IDate) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileCreationDate:"), value)
}

// The modification date of the file.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filemodificationdate
func (i_ ICCameraFile) FileModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("fileModificationDate"))
	return rv
}


// SetFileModificationDate sets the value of the fileModificationDate property.
// The modification date of the file.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filemodificationdate
func (i_ ICCameraFile) SetFileModificationDate(value foundation.IDate) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileModificationDate:"), value)
}

// The size of the file, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filesize
func (i_ ICCameraFile) FileSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fileSize"))
	return rv
}


// SetFileSize sets the value of the fileSize property.
// The size of the file, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filesize
func (i_ ICCameraFile) SetFileSize(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/fingerprint
func (i_ ICCameraFile) Fingerprint() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("fingerprint"))
	return rv
}


// SetFingerprint sets the value of the fingerprint property.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/fingerprint
func (i_ ICCameraFile) SetFingerprint(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFingerprint:"), value)
}

// A Boolean value that indicates whether a file is autopicked by Photos to represent the burst.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/firstpicked
func (i_ ICCameraFile) FirstPicked() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("firstPicked"))
	return rv
}


// SetFirstPicked sets the value of the firstPicked property.
// A Boolean value that indicates whether a file is autopicked by Photos to represent the burst.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/firstpicked
func (i_ ICCameraFile) SetFirstPicked(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFirstPicked:"), value)
}

// The GPS String of the file in standard format.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/gpsstring
func (i_ ICCameraFile) GpsString() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("gpsString"))
	return rv
}


// SetGpsString sets the value of the gpsString property.
// The GPS String of the file in standard format.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/gpsstring
func (i_ ICCameraFile) SetGpsString(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGpsString:"), value)
}

// The group
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/groupuuid
func (i_ ICCameraFile) GroupUUID() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("groupUUID"))
	return rv
}


// SetGroupUUID sets the value of the groupUUID property.
// The group

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/groupuuid
func (i_ ICCameraFile) SetGroupUUID(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupUUID:"), value)
}

// The height of an image or movie frame.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/height
func (i_ ICCameraFile) Height() int {
	rv := objc.Send[int](i_.ID, objc.Sel("height"))
	return rv
}


// SetHeight sets the value of the height property.
// The height of an image or movie frame.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/height
func (i_ ICCameraFile) SetHeight(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHeight:"), value)
}

// A Boolean value that indicates whether the file is a slow motion or high-frame-rate video file.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/highframerate
func (i_ ICCameraFile) HighFramerate() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("highFramerate"))
	return rv
}


// SetHighFramerate sets the value of the highFramerate property.
// A Boolean value that indicates whether the file is a slow motion or high-frame-rate video file.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/highframerate
func (i_ ICCameraFile) SetHighFramerate(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHighFramerate:"), value)
}

// The orientation to use when downloading the image.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/orientation
func (i_ ICCameraFile) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("orientation"))
	return rv
}


// SetOrientation sets the value of the orientation property.
// The orientation to use when downloading the image.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/orientation
func (i_ ICCameraFile) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOrientation:"), value)
}

// The original name of the file on disk.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/originalfilename
func (i_ ICCameraFile) OriginalFilename() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("originalFilename"))
	return rv
}


// SetOriginalFilename sets the value of the originalFilename property.
// The original name of the file on disk.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/originalfilename
func (i_ ICCameraFile) SetOriginalFilename(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOriginalFilename:"), value)
}

// The originating asset ID of an
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/originatingassetid
func (i_ ICCameraFile) OriginatingAssetID() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("originatingAssetID"))
	return rv
}


// SetOriginatingAssetID sets the value of the originatingAssetID property.
// The originating asset ID of an

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/originatingassetid
func (i_ ICCameraFile) SetOriginatingAssetID(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOriginatingAssetID:"), value)
}

// A sidecar file containing the logical
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/pairedrawimage
func (i_ ICCameraFile) PairedRawImage() ICCameraFile {
	rv := objc.Send[ICCameraFile](i_.ID, objc.Sel("pairedRawImage"))
	return rv
}


// SetPairedRawImage sets the value of the pairedRawImage property.
// A sidecar file containing the logical

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/pairedrawimage
func (i_ ICCameraFile) SetPairedRawImage(value ICCameraFile) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPairedRawImage:"), value)
}

// A related UUID correlating several images from an Apple device.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/relateduuid
func (i_ ICCameraFile) RelatedUUID() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("relatedUUID"))
	return rv
}


// SetRelatedUUID sets the value of the relatedUUID property.
// A related UUID correlating several images from an Apple device.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/relateduuid
func (i_ ICCameraFile) SetRelatedUUID(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRelatedUUID:"), value)
}

// An array of two camera files associated with this file.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/sidecarfiles
func (i_ ICCameraFile) SidecarFiles() ICCameraItem {
	rv := objc.Send[ICCameraItem](i_.ID, objc.Sel("sidecarFiles"))
	return rv
}


// SetSidecarFiles sets the value of the sidecarFiles property.
// An array of two camera files associated with this file.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/sidecarfiles
func (i_ ICCameraFile) SetSidecarFiles(value ICCameraItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSidecarFiles:"), value)
}

// A Boolean value that indicates whether the file is a time-lapse video file.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/timelapse
func (i_ ICCameraFile) TimeLapse() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("timeLapse"))
	return rv
}


// SetTimeLapse sets the value of the timeLapse property.
// A Boolean value that indicates whether the file is a time-lapse video file.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/timelapse
func (i_ ICCameraFile) SetTimeLapse(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTimeLapse:"), value)
}



