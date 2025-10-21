// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSItemAttributes] class.
var (
	FSItemAttributesClass     _FSItemAttributesClass
	FSItemAttributesClassOnce sync.Once
)

func getFSItemAttributesClass() _FSItemAttributesClass {
	FSItemAttributesClassOnce.Do(func() {
		FSItemAttributesClass = _FSItemAttributesClass{objc.GetClass("FSItemAttributes")}
	})
	return FSItemAttributesClass
}

type _FSItemAttributesClass struct {
	class objc.Class
}

// An interface definition for the [FSItemAttributes] class.
type IFSItemAttributes interface {
	objectivec.IObject
	InvalidateAllProperties()
	IsValid(attribute FSItemAttribute) bool
}

// Attributes of an item, such as size, creation and modification times, and user and group identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes
type FSItemAttributes struct {
	objectivec.Object
}

// FSItemAttributesFrom constructs a [FSItemAttributes] from an unsafe.Pointer.
//
// Attributes of an item, such as size, creation and modification times, and user and group identifiers.
func FSItemAttributesFrom(ptr unsafe.Pointer) FSItemAttributes {
	return FSItemAttributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSItemAttributesClass) Alloc() FSItemAttributes {
	rv := objc.Send[FSItemAttributes](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSItemAttributesClass) New() FSItemAttributes {
	rv := objc.Send[FSItemAttributes](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSItemAttributes) Init() FSItemAttributes {
	rv := objc.Send[FSItemAttributes](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSItemAttributes) Autorelease() FSItemAttributes {
	rv := objc.Send[FSItemAttributes](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItemAttributes creates a new FSItemAttributes instance.
func NewFSItemAttributes() FSItemAttributes {
	return getFSItemAttributesClass().New()
}


// Marks all attributes inactive.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/invalidateAllProperties()
func (f_ FSItemAttributes) InvalidateAllProperties() {
	objc.Send[objc.ID](f_.ID, objc.Sel("invalidateAllProperties"))
}

// Returns a Boolean value that indicates whether the attribute is valid.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/isValid(_:)
func (f_ FSItemAttributes) IsValid(attribute FSItemAttribute) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isValid:"), attribute)
	return rv
}

// The item’s added time.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/addedTime
func (f_ FSItemAttributes) AddedTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("addedTime"))
	return rv
}


// SetAddedTime sets the value of the addedTime property.
// The item’s added time.

//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/addedTime
func (f_ FSItemAttributes) SetAddedTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAddedTime:"), value)
}

// The item’s file identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/fileID
func (f_ FSItemAttributes) FileID() FSItemID {
	rv := objc.Send[FSItemID](f_.ID, objc.Sel("fileID"))
	return rv
}


// SetFileID sets the value of the fileID property.
// The item’s file identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/fileID
func (f_ FSItemAttributes) SetFileID(value IFSItemID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileID:"), value)
}

// The mode of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/mode
func (f_ FSItemAttributes) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
// The mode of the item.

//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/mode
func (f_ FSItemAttributes) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMode:"), value)
}

// A Boolean value that indicates whether the item supports a limited set of extended attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/supportsLimitedXAttrs
func (f_ FSItemAttributes) SupportsLimitedXAttrs() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsLimitedXAttrs"))
	return rv
}


// SetSupportsLimitedXAttrs sets the value of the supportsLimitedXAttrs property.
// A Boolean value that indicates whether the item supports a limited set of extended attributes.

//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/supportsLimitedXAttrs
func (f_ FSItemAttributes) SetSupportsLimitedXAttrs(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsLimitedXAttrs:"), value)
}

// The item’s last-accessed time.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/accesstime
func (f_ FSItemAttributes) AccessTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("accessTime"))
	return rv
}


// SetAccessTime sets the value of the accessTime property.
// The item’s last-accessed time.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/accesstime
func (f_ FSItemAttributes) SetAccessTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAccessTime:"), value)
}

// The item’s allocated size.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/allocsize
func (f_ FSItemAttributes) AllocSize() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("allocSize"))
	return rv
}


// SetAllocSize sets the value of the allocSize property.
// The item’s allocated size.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/allocsize
func (f_ FSItemAttributes) SetAllocSize(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAllocSize:"), value)
}

// The item’s last-backup time.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/backuptime
func (f_ FSItemAttributes) BackupTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("backupTime"))
	return rv
}


// SetBackupTime sets the value of the backupTime property.
// The item’s last-backup time.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/backuptime
func (f_ FSItemAttributes) SetBackupTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBackupTime:"), value)
}

// The item’s creation time.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/birthtime
func (f_ FSItemAttributes) BirthTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("birthTime"))
	return rv
}


// SetBirthTime sets the value of the birthTime property.
// The item’s creation time.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/birthtime
func (f_ FSItemAttributes) SetBirthTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBirthTime:"), value)
}

// The item’s last-changed time.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/changetime
func (f_ FSItemAttributes) ChangeTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("changeTime"))
	return rv
}


// SetChangeTime sets the value of the changeTime property.
// The item’s last-changed time.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/changetime
func (f_ FSItemAttributes) SetChangeTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setChangeTime:"), value)
}

// The item’s behavior flags.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/flags
func (f_ FSItemAttributes) Flags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("flags"))
	return rv
}


// SetFlags sets the value of the flags property.
// The item’s behavior flags.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/flags
func (f_ FSItemAttributes) SetFlags(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFlags:"), value)
}

// The group identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/gid
func (f_ FSItemAttributes) Gid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("gid"))
	return rv
}


// SetGid sets the value of the gid property.
// The group identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/gid
func (f_ FSItemAttributes) SetGid(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setGid:"), value)
}

// A Boolean value that indicates whether the file system overrides the per-volume settings for kernel offloaded I/O for a specific file.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/inhibitkerneloffloadedio
func (f_ FSItemAttributes) InhibitKernelOffloadedIO() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("inhibitKernelOffloadedIO"))
	return rv
}


// SetInhibitKernelOffloadedIO sets the value of the inhibitKernelOffloadedIO property.
// A Boolean value that indicates whether the file system overrides the per-volume settings for kernel offloaded I/O for a specific file.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/inhibitkerneloffloadedio
func (f_ FSItemAttributes) SetInhibitKernelOffloadedIO(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInhibitKernelOffloadedIO:"), value)
}

// The number of hard links to the item.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/linkcount
func (f_ FSItemAttributes) LinkCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("linkCount"))
	return rv
}


// SetLinkCount sets the value of the linkCount property.
// The number of hard links to the item.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/linkcount
func (f_ FSItemAttributes) SetLinkCount(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLinkCount:"), value)
}

// The item’s last-modified time.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/modifytime
func (f_ FSItemAttributes) ModifyTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("modifyTime"))
	return rv
}


// SetModifyTime sets the value of the modifyTime property.
// The item’s last-modified time.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/modifytime
func (f_ FSItemAttributes) SetModifyTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setModifyTime:"), value)
}

// The identifier of the item’s parent.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/parentid
func (f_ FSItemAttributes) ParentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("parentID"))
	return rv
}


// SetParentID sets the value of the parentID property.
// The identifier of the item’s parent.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/parentid
func (f_ FSItemAttributes) SetParentID(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setParentID:"), value)
}

// The item’s size.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/size
func (f_ FSItemAttributes) Size() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("size"))
	return rv
}


// SetSize sets the value of the size property.
// The item’s size.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/size
func (f_ FSItemAttributes) SetSize(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSize:"), value)
}

// The item type, such as a regular file, directory, or symbolic link.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/type
func (f_ FSItemAttributes) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The item type, such as a regular file, directory, or symbolic link.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/type
func (f_ FSItemAttributes) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setType:"), value)
}

// The user identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/uid
func (f_ FSItemAttributes) Uid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("uid"))
	return rv
}


// SetUid sets the value of the uid property.
// The user identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/attributes/uid
func (f_ FSItemAttributes) SetUid(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUid:"), value)
}



