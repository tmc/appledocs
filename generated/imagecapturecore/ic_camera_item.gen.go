// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// An abstract class that represents a camera item.
//
// The ImageCaptureCore framework defines two concrete subclasses of camera items: and .
//
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


// A mutable dictionary to store arbitrary key-value pairs associated with a camera item.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/userdata
func (i_ ICCameraItem) UserData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("userData"))
	return rv
}


// SetUserData sets the value of the userData property.
// A mutable dictionary to store arbitrary key-value pairs associated with a camera item.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/userdata
func (i_ ICCameraItem) SetUserData(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUserData:"), value)
}

// The item’s thumbnail if it is readily available.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/thumbnailifavailable
func (i_ ICCameraItem) ThumbnailIfAvailable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("thumbnailIfAvailable"))
	return rv
}


// SetThumbnailIfAvailable sets the value of the thumbnailIfAvailable property.
// The item’s thumbnail if it is readily available.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/thumbnailifavailable
func (i_ ICCameraItem) SetThumbnailIfAvailable(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThumbnailIfAvailable:"), value)
}

// The item’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/name
func (i_ ICCameraItem) Name() string {
	rv := objc.Send[string](i_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The item’s name.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/name
func (i_ ICCameraItem) SetName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setName:"), objc.String(value))
}

// The item’s metadata if it is readily available.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/metadataifavailable
func (i_ ICCameraItem) MetadataIfAvailable() string {
	rv := objc.Send[string](i_.ID, objc.Sel("metadataIfAvailable"))
	return rv
}


// SetMetadataIfAvailable sets the value of the metadataIfAvailable property.
// The item’s metadata if it is readily available.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/metadataifavailable
func (i_ ICCameraItem) SetMetadataIfAvailable(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMetadataIfAvailable:"), objc.String(value))
}

// A Boolean value indicating whether the item is a raw image file.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/israw
func (i_ ICCameraItem) IsRaw() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isRaw"))
	return rv
}


// SetIsRaw sets the value of the isRaw property.
// A Boolean value indicating whether the item is a raw image file.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/israw
func (i_ ICCameraItem) SetIsRaw(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsRaw:"), value)
}

// A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/wasaddedaftercontentcatalogcompleted
func (i_ ICCameraItem) WasAddedAfterContentCatalogCompleted() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("wasAddedAfterContentCatalogCompleted"))
	return rv
}


// SetWasAddedAfterContentCatalogCompleted sets the value of the wasAddedAfterContentCatalogCompleted property.
// A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/wasaddedaftercontentcatalogcompleted
func (i_ ICCameraItem) SetWasAddedAfterContentCatalogCompleted(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWasAddedAfterContentCatalogCompleted:"), value)
}

// The item’s metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/metadata
func (i_ ICCameraItem) Metadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// The item’s metadata.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/metadata
func (i_ ICCameraItem) SetMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMetadata:"), value)
}

// This item’s parent folder.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/parentfolder
func (i_ ICCameraItem) ParentFolder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("parentFolder"))
	return rv
}


// SetParentFolder sets the value of the parentFolder property.
// This item’s parent folder.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/parentfolder
func (i_ ICCameraItem) SetParentFolder(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setParentFolder:"), value)
}

// A Boolean value that indicates whether this item is in a temporary store.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/isintemporarystore
func (i_ ICCameraItem) IsInTemporaryStore() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isInTemporaryStore"))
	return rv
}


// SetIsInTemporaryStore sets the value of the isInTemporaryStore property.
// A Boolean value that indicates whether this item is in a temporary store.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/isintemporarystore
func (i_ ICCameraItem) SetIsInTemporaryStore(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsInTemporaryStore:"), value)
}

// The item’s file system path on a camera using the mass storage transport type.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/filesystempath
func (i_ ICCameraItem) FileSystemPath() string {
	rv := objc.Send[string](i_.ID, objc.Sel("fileSystemPath"))
	return rv
}


// SetFileSystemPath sets the value of the fileSystemPath property.
// The item’s file system path on a camera using the mass storage transport type.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/filesystempath
func (i_ ICCameraItem) SetFileSystemPath(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileSystemPath:"), objc.String(value))
}

// A large thumbnail for the item if one is readily available.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/largethumbnailifavailable
func (i_ ICCameraItem) LargeThumbnailIfAvailable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("largeThumbnailIfAvailable"))
	return rv
}


// SetLargeThumbnailIfAvailable sets the value of the largeThumbnailIfAvailable property.
// A large thumbnail for the item if one is readily available.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/largethumbnailifavailable
func (i_ ICCameraItem) SetLargeThumbnailIfAvailable(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLargeThumbnailIfAvailable:"), value)
}

// The item’s parent device.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/device
func (i_ ICCameraItem) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("device"))
	return rv
}


// SetDevice sets the value of the device property.
// The item’s parent device.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/device
func (i_ ICCameraItem) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDevice:"), value)
}

// A Boolean value that indicates whether the storage card in the camera is locked.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/islocked
func (i_ ICCameraItem) IsLocked() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isLocked"))
	return rv
}


// SetIsLocked sets the value of the isLocked property.
// A Boolean value that indicates whether the storage card in the camera is locked.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/islocked
func (i_ ICCameraItem) SetIsLocked(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsLocked:"), value)
}

// The item’s creation date, usually the same as its
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/creationdate
func (i_ ICCameraItem) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("creationDate"))
	return rv
}


// SetCreationDate sets the value of the creationDate property.
// The item’s creation date, usually the same as its

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/creationdate
func (i_ ICCameraItem) SetCreationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCreationDate:"), value)
}

// The item’s
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/ptpobjecthandle
func (i_ ICCameraItem) PtpObjectHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("ptpObjectHandle"))
	return rv
}


// SetPtpObjectHandle sets the value of the ptpObjectHandle property.
// The item’s

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/ptpobjecthandle
func (i_ ICCameraItem) SetPtpObjectHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPtpObjectHandle:"), value)
}

// A Boolean value indicating whether the item is a raw image file.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/isRaw
func (i_ ICCameraItem) Raw() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("raw"))
	return rv
}

// The item’s modification date, usually the same as its modification date.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/modificationDate
func (i_ ICCameraItem) ModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("modificationDate"))
	return rv
}

// The item’s thumbnail.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/thumbnail
func (i_ ICCameraItem) Thumbnail() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](i_.ID, objc.Sel("thumbnail"))
	return rv
}

// The item’s uniform type identifier (UTI) string.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/uti
func (i_ ICCameraItem) UTI() string {
	rv := objc.Send[string](i_.ID, objc.Sel("UTI"))
	return rv
}



