// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ICCameraItem] class.
var (
	ICCameraItemClass     _ICCameraItemClass
	ICCameraItemClassOnce sync.Once
)

func getICCameraItemClass() _ICCameraItemClass {
	ICCameraItemClassOnce.Do(func() {
		ICCameraItemClass = _ICCameraItemClass{objc.GetClass("ICCameraItem")}
	})
	return ICCameraItemClass
}

type _ICCameraItemClass struct {
	class objc.Class
}

// An interface definition for the [ICCameraItem] class.
type IICCameraItem interface {
	objectivec.IObject
	// properties:
	Name() string
	UTI() string
	CreationDate() foundation.Date
	SetCreationDate(value foundation.Date)
	Device() ICCameraDevice
	SetDevice(value ICCameraDevice)
	FileSystemPath() string
	SetFileSystemPath(value string)
	IsInTemporaryStore() bool
	SetIsInTemporaryStore(value bool)
	IsLocked() bool
	SetIsLocked(value bool)
	IsRaw() bool
	SetIsRaw(value bool)
	LargeThumbnailIfAvailable() appkit.Image
	SetLargeThumbnailIfAvailable(value appkit.Image)
	Metadata() unsafe.Pointer
	SetMetadata(value unsafe.Pointer)
	MetadataIfAvailable() string
	SetMetadataIfAvailable(value string)
	ModificationDate() foundation.Date
	SetModificationDate(value foundation.Date)
	ParentFolder() ICCameraFolder
	SetParentFolder(value ICCameraFolder)
	PtpObjectHandle() unsafe.Pointer
	SetPtpObjectHandle(value unsafe.Pointer)
	Thumbnail() appkit.Image
	SetThumbnail(value appkit.Image)
	ThumbnailIfAvailable() appkit.Image
	SetThumbnailIfAvailable(value appkit.Image)
	UserData() foundation.MutableDictionary
	SetUserData(value foundation.MutableDictionary)
	WasAddedAfterContentCatalogCompleted() bool
	SetWasAddedAfterContentCatalogCompleted(value bool)
	// methods:
}

// An abstract class that represents a camera item.
//
// The ImageCaptureCore framework defines two concrete subclasses of camera items: and .


// An abstract class that represents a camera item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem
type ICCameraItem struct {
	objectivec.Object
}

// ICCameraItemFrom constructs a [ICCameraItem] from an unsafe.Pointer.
//
// An abstract class that represents a camera item.
func ICCameraItemFrom(ptr unsafe.Pointer) ICCameraItem {
	return ICCameraItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ICCameraItemClass) Alloc() ICCameraItem {
	rv := objc.Send[ICCameraItem](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ICCameraItemClass) New() ICCameraItem {
	rv := objc.Send[ICCameraItem](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICCameraItem) Init() ICCameraItem {
	rv := objc.Send[ICCameraItem](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICCameraItem) Autorelease() ICCameraItem {
	rv := objc.Send[ICCameraItem](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraItem creates a new ICCameraItem instance.
func NewICCameraItem() ICCameraItem {
	return getICCameraItemClass().New()
}



// The item’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/name
func (i_ ICCameraItem) Name() string {
	rv := objc.Send[string](i_.ID, objc.Sel("name"))
	return rv
}


// The item’s uniform type identifier (UTI) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/uti
func (i_ ICCameraItem) UTI() string {
	rv := objc.Send[string](i_.ID, objc.Sel("UTI"))
	return rv
}


// The item’s creation date, usually the same as its
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/creationdate
func (i_ ICCameraItem) CreationDate() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("creationDate"))
	return rv
}


// The item’s creation date, usually the same as its
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/creationdate
func (i_ ICCameraItem) SetCreationDate(value foundation.Date) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCreationDate:"), value)
}


// The item’s parent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/device
func (i_ ICCameraItem) Device() ICCameraDevice {
	rv := objc.Send[ICCameraDevice](i_.ID, objc.Sel("device"))
	return rv
}


// The item’s parent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/device
func (i_ ICCameraItem) SetDevice(value ICCameraDevice) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDevice:"), value)
}


// The item’s file system path on a camera using the mass storage transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/filesystempath
func (i_ ICCameraItem) FileSystemPath() string {
	rv := objc.Send[string](i_.ID, objc.Sel("fileSystemPath"))
	return rv
}


// The item’s file system path on a camera using the mass storage transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/filesystempath
func (i_ ICCameraItem) SetFileSystemPath(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileSystemPath:"), objc.String(value))
}


// A Boolean value that indicates whether this item is in a temporary store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/isintemporarystore
func (i_ ICCameraItem) IsInTemporaryStore() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isInTemporaryStore"))
	return rv
}


// A Boolean value that indicates whether this item is in a temporary store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/isintemporarystore
func (i_ ICCameraItem) SetIsInTemporaryStore(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsInTemporaryStore:"), value)
}


// A Boolean value that indicates whether the storage card in the camera is locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/islocked
func (i_ ICCameraItem) IsLocked() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isLocked"))
	return rv
}


// A Boolean value that indicates whether the storage card in the camera is locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/islocked
func (i_ ICCameraItem) SetIsLocked(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsLocked:"), value)
}


// A Boolean value indicating whether the item is a raw image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/israw
func (i_ ICCameraItem) IsRaw() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isRaw"))
	return rv
}


// A Boolean value indicating whether the item is a raw image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/israw
func (i_ ICCameraItem) SetIsRaw(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsRaw:"), value)
}


// A large thumbnail for the item if one is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/largethumbnailifavailable
func (i_ ICCameraItem) LargeThumbnailIfAvailable() appkit.Image {
	rv := objc.Send[appkit.Image](i_.ID, objc.Sel("largeThumbnailIfAvailable"))
	return rv
}


// A large thumbnail for the item if one is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/largethumbnailifavailable
func (i_ ICCameraItem) SetLargeThumbnailIfAvailable(value appkit.Image) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLargeThumbnailIfAvailable:"), value)
}


// The item’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/metadata
func (i_ ICCameraItem) Metadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("metadata"))
	return rv
}


// The item’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/metadata
func (i_ ICCameraItem) SetMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMetadata:"), value)
}


// The item’s metadata if it is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/metadataifavailable
func (i_ ICCameraItem) MetadataIfAvailable() string {
	rv := objc.Send[string](i_.ID, objc.Sel("metadataIfAvailable"))
	return rv
}


// The item’s metadata if it is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/metadataifavailable
func (i_ ICCameraItem) SetMetadataIfAvailable(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMetadataIfAvailable:"), objc.String(value))
}


// The item’s modification date, usually the same as its
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/modificationdate
func (i_ ICCameraItem) ModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("modificationDate"))
	return rv
}


// The item’s modification date, usually the same as its
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/modificationdate
func (i_ ICCameraItem) SetModificationDate(value foundation.Date) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setModificationDate:"), value)
}


// This item’s parent folder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/parentfolder
func (i_ ICCameraItem) ParentFolder() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](i_.ID, objc.Sel("parentFolder"))
	return rv
}


// This item’s parent folder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/parentfolder
func (i_ ICCameraItem) SetParentFolder(value ICCameraFolder) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setParentFolder:"), value)
}


// The item’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/ptpobjecthandle
func (i_ ICCameraItem) PtpObjectHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("ptpObjectHandle"))
	return rv
}


// The item’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/ptpobjecthandle
func (i_ ICCameraItem) SetPtpObjectHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPtpObjectHandle:"), value)
}


// The item’s thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/thumbnail
func (i_ ICCameraItem) Thumbnail() appkit.Image {
	rv := objc.Send[appkit.Image](i_.ID, objc.Sel("thumbnail"))
	return rv
}


// The item’s thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/thumbnail
func (i_ ICCameraItem) SetThumbnail(value appkit.Image) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThumbnail:"), value)
}


// The item’s thumbnail if it is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/thumbnailifavailable
func (i_ ICCameraItem) ThumbnailIfAvailable() appkit.Image {
	rv := objc.Send[appkit.Image](i_.ID, objc.Sel("thumbnailIfAvailable"))
	return rv
}


// The item’s thumbnail if it is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/thumbnailifavailable
func (i_ ICCameraItem) SetThumbnailIfAvailable(value appkit.Image) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThumbnailIfAvailable:"), value)
}


// A mutable dictionary to store arbitrary key-value pairs associated with a camera item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/userdata
func (i_ ICCameraItem) UserData() foundation.MutableDictionary {
	rv := objc.Send[foundation.MutableDictionary](i_.ID, objc.Sel("userData"))
	return rv
}


// A mutable dictionary to store arbitrary key-value pairs associated with a camera item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/userdata
func (i_ ICCameraItem) SetUserData(value foundation.MutableDictionary) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUserData:"), value)
}


// A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/wasaddedaftercontentcatalogcompleted
func (i_ ICCameraItem) WasAddedAfterContentCatalogCompleted() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("wasAddedAfterContentCatalogCompleted"))
	return rv
}


// A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/wasaddedaftercontentcatalogcompleted
func (i_ ICCameraItem) SetWasAddedAfterContentCatalogCompleted(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWasAddedAfterContentCatalogCompleted:"), value)
}



