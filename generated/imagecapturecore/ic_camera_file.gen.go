// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	BurstUUID() string /* primitive/slice/pointer. */
	GpsString() string /* primitive/slice/pointer. */
	BurstFavorite() bool /* primitive/slice/pointer. */
	SetBurstFavorite(value bool /* primitive/slice/pointer. */)
	BurstPicked() bool /* primitive/slice/pointer. */
	SetBurstPicked(value bool /* primitive/slice/pointer. */)
	CreatedFilename() string /* primitive/slice/pointer. */
	SetCreatedFilename(value string /* primitive/slice/pointer. */)
	Duration() float64 /* primitive/slice/pointer. */
	SetDuration(value float64 /* primitive/slice/pointer. */)
	ExifCreationDate() foundation.objc.IObject /* cross-framework: Date */
	SetExifCreationDate(value foundation.objc.IObject /* cross-framework: Date */)
	ExifModificationDate() foundation.objc.IObject /* cross-framework: Date */
	SetExifModificationDate(value foundation.objc.IObject /* cross-framework: Date */)
	FileCreationDate() foundation.objc.IObject /* cross-framework: Date */
	SetFileCreationDate(value foundation.objc.IObject /* cross-framework: Date */)
	FileModificationDate() foundation.objc.IObject /* cross-framework: Date */
	SetFileModificationDate(value foundation.objc.IObject /* cross-framework: Date */)
	FileSize() unsafe.Pointer
	SetFileSize(value unsafe.Pointer)
	Fingerprint() string /* primitive/slice/pointer. */
	SetFingerprint(value string /* primitive/slice/pointer. */)
	FirstPicked() bool /* primitive/slice/pointer. */
	SetFirstPicked(value bool /* primitive/slice/pointer. */)
	GroupUUID() string /* primitive/slice/pointer. */
	SetGroupUUID(value string /* primitive/slice/pointer. */)
	Height() int /* primitive/slice/pointer. */
	SetHeight(value int /* primitive/slice/pointer. */)
	HighFramerate() bool /* primitive/slice/pointer. */
	SetHighFramerate(value bool /* primitive/slice/pointer. */)
	Orientation() unsafe.Pointer
	SetOrientation(value unsafe.Pointer)
	OriginalFilename() string /* primitive/slice/pointer. */
	SetOriginalFilename(value string /* primitive/slice/pointer. */)
	OriginatingAssetID() string /* primitive/slice/pointer. */
	SetOriginatingAssetID(value string /* primitive/slice/pointer. */)
	PairedRawImage() ICCameraFile /* already interface */
	SetPairedRawImage(value ICCameraFile /* already interface */)
	RelatedUUID() string /* primitive/slice/pointer. */
	SetRelatedUUID(value string /* primitive/slice/pointer. */)
	SidecarFiles() ICCameraItem /* already interface */
	SetSidecarFiles(value ICCameraItem /* already interface */)
	TimeLapse() bool /* primitive/slice/pointer. */
	SetTimeLapse(value bool /* primitive/slice/pointer. */)
	Width() int /* primitive/slice/pointer. */
	SetWidth(value int /* primitive/slice/pointer. */)
	// methods:
}

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



// The burst UUID of the file if it is in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/burstUUID
func (i_ ICCameraFile) BurstUUID() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("burstUUID"))
	return rv
}


// The GPS String of the file in standard format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/gpsString
func (i_ ICCameraFile) GpsString() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("gpsString"))
	return rv
}


// A Boolean value that indicates this file is the burst favorite in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/burstfavorite
func (i_ ICCameraFile) BurstFavorite() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("burstFavorite"))
	return rv
}


// A Boolean value that indicates this file is the burst favorite in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/burstfavorite
func (i_ ICCameraFile) SetBurstFavorite(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBurstFavorite:"), value)
}


// A Boolean value that indicates whether this file is user picked in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/burstpicked
func (i_ ICCameraFile) BurstPicked() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("burstPicked"))
	return rv
}


// A Boolean value that indicates whether this file is user picked in a burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/burstpicked
func (i_ ICCameraFile) SetBurstPicked(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBurstPicked:"), value)
}


// The created name of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/createdfilename
func (i_ ICCameraFile) CreatedFilename() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("createdFilename"))
	return rv
}


// The created name of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/createdfilename
func (i_ ICCameraFile) SetCreatedFilename(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCreatedFilename:"), objc.String(value))
}


// The duration, in seconds, of an audio or video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/duration
func (i_ ICCameraFile) Duration() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](i_.ID, objc.Sel("duration"))
	return rv
}


// The duration, in seconds, of an audio or video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/duration
func (i_ ICCameraFile) SetDuration(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDuration:"), value)
}


// The
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/exifcreationdate
func (i_ ICCameraFile) ExifCreationDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("exifCreationDate"))
	return rv
}


// The
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/exifcreationdate
func (i_ ICCameraFile) SetExifCreationDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExifCreationDate:"), value)
}


// The
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/exifmodificationdate
func (i_ ICCameraFile) ExifModificationDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("exifModificationDate"))
	return rv
}


// The
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/exifmodificationdate
func (i_ ICCameraFile) SetExifModificationDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExifModificationDate:"), value)
}


// The creation date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filecreationdate
func (i_ ICCameraFile) FileCreationDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("fileCreationDate"))
	return rv
}


// The creation date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filecreationdate
func (i_ ICCameraFile) SetFileCreationDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileCreationDate:"), value)
}


// The modification date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filemodificationdate
func (i_ ICCameraFile) FileModificationDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("fileModificationDate"))
	return rv
}


// The modification date of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filemodificationdate
func (i_ ICCameraFile) SetFileModificationDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileModificationDate:"), value)
}


// The size of the file, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filesize
func (i_ ICCameraFile) FileSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fileSize"))
	return rv
}


// The size of the file, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/filesize
func (i_ ICCameraFile) SetFileSize(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/fingerprint
func (i_ ICCameraFile) Fingerprint() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("fingerprint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/fingerprint
func (i_ ICCameraFile) SetFingerprint(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFingerprint:"), objc.String(value))
}


// A Boolean value that indicates whether a file is autopicked by Photos to represent the burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/firstpicked
func (i_ ICCameraFile) FirstPicked() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("firstPicked"))
	return rv
}


// A Boolean value that indicates whether a file is autopicked by Photos to represent the burst.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/firstpicked
func (i_ ICCameraFile) SetFirstPicked(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFirstPicked:"), value)
}


// The group
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/groupuuid
func (i_ ICCameraFile) GroupUUID() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("groupUUID"))
	return rv
}


// The group
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/groupuuid
func (i_ ICCameraFile) SetGroupUUID(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupUUID:"), objc.String(value))
}


// The height of an image or movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/height
func (i_ ICCameraFile) Height() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](i_.ID, objc.Sel("height"))
	return rv
}


// The height of an image or movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/height
func (i_ ICCameraFile) SetHeight(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHeight:"), value)
}


// A Boolean value that indicates whether the file is a slow motion or high-frame-rate video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/highframerate
func (i_ ICCameraFile) HighFramerate() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("highFramerate"))
	return rv
}


// A Boolean value that indicates whether the file is a slow motion or high-frame-rate video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/highframerate
func (i_ ICCameraFile) SetHighFramerate(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHighFramerate:"), value)
}


// The orientation to use when downloading the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/orientation
func (i_ ICCameraFile) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("orientation"))
	return rv
}


// The orientation to use when downloading the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/orientation
func (i_ ICCameraFile) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOrientation:"), value)
}


// The original name of the file on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/originalfilename
func (i_ ICCameraFile) OriginalFilename() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("originalFilename"))
	return rv
}


// The original name of the file on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/originalfilename
func (i_ ICCameraFile) SetOriginalFilename(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOriginalFilename:"), objc.String(value))
}


// The originating asset ID of an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/originatingassetid
func (i_ ICCameraFile) OriginatingAssetID() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("originatingAssetID"))
	return rv
}


// The originating asset ID of an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/originatingassetid
func (i_ ICCameraFile) SetOriginatingAssetID(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOriginatingAssetID:"), objc.String(value))
}


// A sidecar file containing the logical
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/pairedrawimage
func (i_ ICCameraFile) PairedRawImage() ICCameraFile /* already interface */ {
	rv := objc.Send[ICCameraFile](i_.ID, objc.Sel("pairedRawImage"))
	return rv
}


// A sidecar file containing the logical
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/pairedrawimage
func (i_ ICCameraFile) SetPairedRawImage(value ICCameraFile /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPairedRawImage:"), value)
}


// A related UUID correlating several images from an Apple device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/relateduuid
func (i_ ICCameraFile) RelatedUUID() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("relatedUUID"))
	return rv
}


// A related UUID correlating several images from an Apple device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/relateduuid
func (i_ ICCameraFile) SetRelatedUUID(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRelatedUUID:"), objc.String(value))
}


// An array of two camera files associated with this file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/sidecarfiles
func (i_ ICCameraFile) SidecarFiles() ICCameraItem /* already interface */ {
	rv := objc.Send[ICCameraItem](i_.ID, objc.Sel("sidecarFiles"))
	return rv
}


// An array of two camera files associated with this file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/sidecarfiles
func (i_ ICCameraFile) SetSidecarFiles(value ICCameraItem /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSidecarFiles:"), value)
}


// A Boolean value that indicates whether the file is a time-lapse video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/timelapse
func (i_ ICCameraFile) TimeLapse() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("timeLapse"))
	return rv
}


// A Boolean value that indicates whether the file is a time-lapse video file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/timelapse
func (i_ ICCameraFile) SetTimeLapse(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTimeLapse:"), value)
}


// The width of an image or movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/width
func (i_ ICCameraFile) Width() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](i_.ID, objc.Sel("width"))
	return rv
}


// The width of an image or movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafile/width
func (i_ ICCameraFile) SetWidth(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWidth:"), value)
}



