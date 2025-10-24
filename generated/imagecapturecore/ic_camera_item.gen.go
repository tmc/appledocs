// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ICCameraItem */


/* debug [class_header]: Header for ICCameraItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICCameraItem */
// An interface definition for the [ICCameraItem] class.
type IICCameraItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ICCameraItem */
	// properties:
	Uti() unsafe.Pointer
	SetUti(value unsafe.Pointer)
	ParentFolder() ICCameraFolder
	SetParentFolder(value ICCameraFolder)
	ModificationDate() unsafe.Pointer
	SetModificationDate(value unsafe.Pointer)
	MetadataIfAvailable() unsafe.Pointer
	SetMetadataIfAvailable(value unsafe.Pointer)
	IsInTemporaryStore() unsafe.Pointer
	SetIsInTemporaryStore(value unsafe.Pointer)
	CreationDate() unsafe.Pointer
	SetCreationDate(value unsafe.Pointer)
	LargeThumbnailIfAvailable() Image get /* not a class type */
	SetLargeThumbnailIfAvailable(value Image get /* not a class type */)
	PtpObjectHandle() unsafe.Pointer
	SetPtpObjectHandle(value unsafe.Pointer)
	FileSystemPath() unsafe.Pointer
	SetFileSystemPath(value unsafe.Pointer)
	ThumbnailIfAvailable() Image get /* not a class type */
	SetThumbnailIfAvailable(value Image get /* not a class type */)
	UserData() foundation.MutableDictionary
	SetUserData(value foundation.MutableDictionary)
	Name() unsafe.Pointer
	SetName(value unsafe.Pointer)
	WasAddedAfterContentCatalogCompleted() unsafe.Pointer
	SetWasAddedAfterContentCatalogCompleted(value unsafe.Pointer)
	IsLocked() unsafe.Pointer
	SetIsLocked(value unsafe.Pointer)
	Device() ICCameraDevice
	SetDevice(value ICCameraDevice)
	IsRaw() unsafe.Pointer
	SetIsRaw(value unsafe.Pointer)
	Metadata() unsafe.Pointer
	SetMetadata(value unsafe.Pointer)
	Thumbnail() Image get /* not a class type */
	SetThumbnail(value Image get /* not a class type */)
	InTemporaryStore() bool
	Locked() bool
	Raw() bool
	UTI() objc.IObject /* cross-framework: NSString */
	AddedAfterContentCatalogCompleted() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICCameraItem */
	// methods:
	FlushMetadataCache()
	FlushThumbnailCache()
	RequestMetadata()
	RequestThumbnail()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICCameraItem */
// Alloc allocates a new instance without initialization.
func (ic _ICCameraItemClass) Alloc() ICCameraItem {
	rv := objc.Send[ICCameraItem](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICCameraItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICCameraItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICCameraItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICCameraItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICCameraItem */

// Deletes the item’s cached metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/3131477-flushmetadatacache
func (i_ ICCameraItem) FlushMetadataCache() {
	objc.Send[objc.ID](i_.ID, objc.Sel("flushMetadataCache"))
}/* debug [instance_methods/method]: FlushMetadataCache */


// Deletes the item’s cached thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/3131478-flushthumbnailcache
func (i_ ICCameraItem) FlushThumbnailCache() {
	objc.Send[objc.ID](i_.ID, objc.Sel("flushThumbnailCache"))
}/* debug [instance_methods/method]: FlushThumbnailCache */


// Requests metadata for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/3131480-requestmetadata
func (i_ ICCameraItem) RequestMetadata() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestMetadata"))
}/* debug [instance_methods/method]: RequestMetadata */


// Requests a thumbnail for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/3131481-requestthumbnail
func (i_ ICCameraItem) RequestThumbnail() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestThumbnail"))
}/* debug [instance_methods/method]: RequestThumbnail */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICCameraItem */

// The item’s uniform type identifier (UTI) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388977-uti
func (i_ ICCameraItem) Uti() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("uti"))
	return rv
}/* debug [instance_properties/getter]: uti */


// The item’s uniform type identifier (UTI) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388977-uti
func (i_ ICCameraItem) SetUti(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUti:"), value)
}/* debug [instance_properties/setter]: uti */


// This item’s parent folder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388978-parentfolder
func (i_ ICCameraItem) ParentFolder() ICCameraFolder {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("parentFolder"))
	return rv
}/* debug [instance_properties/getter]: parentFolder */


// This item’s parent folder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388978-parentfolder
func (i_ ICCameraItem) SetParentFolder(value ICCameraFolder) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setParentFolder:"), value)
}/* debug [instance_properties/setter]: parentFolder */


// The item’s modification date, usually the same as its modification date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388979-modificationdate
func (i_ ICCameraItem) ModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("modificationDate"))
	return rv
}/* debug [instance_properties/getter]: modificationDate */


// The item’s modification date, usually the same as its modification date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388979-modificationdate
func (i_ ICCameraItem) SetModificationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setModificationDate:"), value)
}/* debug [instance_properties/setter]: modificationDate */


// The item’s metadata if it is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388985-metadataifavailable
func (i_ ICCameraItem) MetadataIfAvailable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("metadataIfAvailable"))
	return rv
}/* debug [instance_properties/getter]: metadataIfAvailable */


// The item’s metadata if it is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388985-metadataifavailable
func (i_ ICCameraItem) SetMetadataIfAvailable(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMetadataIfAvailable:"), value)
}/* debug [instance_properties/setter]: metadataIfAvailable */


// A Boolean value that indicates whether this item is in a temporary store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388987-isintemporarystore
func (i_ ICCameraItem) IsInTemporaryStore() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isInTemporaryStore"))
	return rv
}/* debug [instance_properties/getter]: isInTemporaryStore */


// A Boolean value that indicates whether this item is in a temporary store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388987-isintemporarystore
func (i_ ICCameraItem) SetIsInTemporaryStore(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsInTemporaryStore:"), value)
}/* debug [instance_properties/setter]: isInTemporaryStore */


// The item’s creation date, usually the same as its creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388989-creationdate
func (i_ ICCameraItem) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("creationDate"))
	return rv
}/* debug [instance_properties/getter]: creationDate */


// The item’s creation date, usually the same as its creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388989-creationdate
func (i_ ICCameraItem) SetCreationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCreationDate:"), value)
}/* debug [instance_properties/setter]: creationDate */


// A large thumbnail for the item if one is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388995-largethumbnailifavailable
func (i_ ICCameraItem) LargeThumbnailIfAvailable() Image get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("largeThumbnailIfAvailable"))
	return rv
}/* debug [instance_properties/getter]: largeThumbnailIfAvailable */


// A large thumbnail for the item if one is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388995-largethumbnailifavailable
func (i_ ICCameraItem) SetLargeThumbnailIfAvailable(value Image get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLargeThumbnailIfAvailable:"), value)
}/* debug [instance_properties/setter]: largeThumbnailIfAvailable */


// The item’s object handle value, if the camera uses the protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388997-ptpobjecthandle
func (i_ ICCameraItem) PtpObjectHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("ptpObjectHandle"))
	return rv
}/* debug [instance_properties/getter]: ptpObjectHandle */


// The item’s object handle value, if the camera uses the protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388997-ptpobjecthandle
func (i_ ICCameraItem) SetPtpObjectHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPtpObjectHandle:"), value)
}/* debug [instance_properties/setter]: ptpObjectHandle */


// The item’s file system path on a camera using the mass storage transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388999-filesystempath
func (i_ ICCameraItem) FileSystemPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fileSystemPath"))
	return rv
}/* debug [instance_properties/getter]: fileSystemPath */


// The item’s file system path on a camera using the mass storage transport type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1388999-filesystempath
func (i_ ICCameraItem) SetFileSystemPath(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFileSystemPath:"), value)
}/* debug [instance_properties/setter]: fileSystemPath */


// The item’s thumbnail if it is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389001-thumbnailifavailable
func (i_ ICCameraItem) ThumbnailIfAvailable() Image get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("thumbnailIfAvailable"))
	return rv
}/* debug [instance_properties/getter]: thumbnailIfAvailable */


// The item’s thumbnail if it is readily available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389001-thumbnailifavailable
func (i_ ICCameraItem) SetThumbnailIfAvailable(value Image get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThumbnailIfAvailable:"), value)
}/* debug [instance_properties/setter]: thumbnailIfAvailable */


// A mutable dictionary to store arbitrary key-value pairs associated with a camera item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389009-userdata
func (i_ ICCameraItem) UserData() foundation.MutableDictionary {
	rv := objc.Send[foundation.MutableDictionary](i_.ID, objc.Sel("userData"))
	return rv
}/* debug [instance_properties/getter]: userData */


// A mutable dictionary to store arbitrary key-value pairs associated with a camera item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389009-userdata
func (i_ ICCameraItem) SetUserData(value foundation.MutableDictionary) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUserData:"), value)
}/* debug [instance_properties/setter]: userData */


// The item’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389011-name
func (i_ ICCameraItem) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The item’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389011-name
func (i_ ICCameraItem) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389015-wasaddedaftercontentcatalogcompl
func (i_ ICCameraItem) WasAddedAfterContentCatalogCompleted() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("wasAddedAfterContentCatalogCompleted"))
	return rv
}/* debug [instance_properties/getter]: wasAddedAfterContentCatalogCompleted */


// A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389015-wasaddedaftercontentcatalogcompl
func (i_ ICCameraItem) SetWasAddedAfterContentCatalogCompleted(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setWasAddedAfterContentCatalogCompleted:"), value)
}/* debug [instance_properties/setter]: wasAddedAfterContentCatalogCompleted */


// A Boolean value that indicates whether the storage card in the camera is locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389017-islocked
func (i_ ICCameraItem) IsLocked() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isLocked"))
	return rv
}/* debug [instance_properties/getter]: isLocked */


// A Boolean value that indicates whether the storage card in the camera is locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389017-islocked
func (i_ ICCameraItem) SetIsLocked(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsLocked:"), value)
}/* debug [instance_properties/setter]: isLocked */


// The item’s parent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389019-device
func (i_ ICCameraItem) Device() ICCameraDevice {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The item’s parent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389019-device
func (i_ ICCameraItem) SetDevice(value ICCameraDevice) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// A Boolean value indicating whether the item is a raw image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389021-israw
func (i_ ICCameraItem) IsRaw() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("isRaw"))
	return rv
}/* debug [instance_properties/getter]: isRaw */


// A Boolean value indicating whether the item is a raw image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/1389021-israw
func (i_ ICCameraItem) SetIsRaw(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsRaw:"), value)
}/* debug [instance_properties/setter]: isRaw */


// The item’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/3131479-metadata
func (i_ ICCameraItem) Metadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// The item’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/3131479-metadata
func (i_ ICCameraItem) SetMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMetadata:"), value)
}/* debug [instance_properties/setter]: metadata */


// The item’s thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/3131482-thumbnail
func (i_ ICCameraItem) Thumbnail() Image get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("thumbnail"))
	return rv
}/* debug [instance_properties/getter]: thumbnail */


// The item’s thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccameraitem/3131482-thumbnail
func (i_ ICCameraItem) SetThumbnail(value Image get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThumbnail:"), value)
}/* debug [instance_properties/setter]: thumbnail */


// A Boolean value that indicates whether this item is in a temporary store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/isInTemporaryStore
func (i_ ICCameraItem) InTemporaryStore() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("inTemporaryStore"))
	return rv
}/* debug [instance_properties/getter]: inTemporaryStore */


// A Boolean value that indicates whether the storage card in the camera is locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/isLocked
func (i_ ICCameraItem) Locked() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("locked"))
	return rv
}/* debug [instance_properties/getter]: locked */


// A Boolean value indicating whether the item is a raw image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/isRaw
func (i_ ICCameraItem) Raw() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("raw"))
	return rv
}/* debug [instance_properties/getter]: raw */


// The item’s uniform type identifier (UTI) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/uti
func (i_ ICCameraItem) UTI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("UTI"))
	return rv
}/* debug [instance_properties/getter]: UTI */


// A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/wasAddedAfterContentCatalogCompleted
func (i_ ICCameraItem) AddedAfterContentCatalogCompleted() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("addedAfterContentCatalogCompleted"))
	return rv
}/* debug [instance_properties/getter]: addedAfterContentCatalogCompleted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICCameraItem */



