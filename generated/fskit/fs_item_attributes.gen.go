// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSItemAttributes */


/* debug [class_header]: Header for FSItemAttributes */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSItemAttributes */
// An interface definition for the [FSItemAttributes] class.
type IFSItemAttributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSItemAttributes */
	// properties:
	AccessTime() unsafe.Pointer
	SetAccessTime(value unsafe.Pointer)
	AddedTime() unsafe.Pointer
	SetAddedTime(value unsafe.Pointer)
	AllocSize() uint64
	SetAllocSize(value uint64)
	BackupTime() unsafe.Pointer
	SetBackupTime(value unsafe.Pointer)
	BirthTime() unsafe.Pointer
	SetBirthTime(value unsafe.Pointer)
	ChangeTime() unsafe.Pointer
	SetChangeTime(value unsafe.Pointer)
	FileID() FSItemID
	SetFileID(value FSItemID)
	Flags() uint32 /* not a class type */
	SetFlags(value uint32 /* not a class type */)
	Gid() uint32 /* not a class type */
	SetGid(value uint32 /* not a class type */)
	InhibitKernelOffloadedIO() bool
	SetInhibitKernelOffloadedIO(value bool)
	LinkCount() uint32 /* not a class type */
	SetLinkCount(value uint32 /* not a class type */)
	Mode() uint32 /* not a class type */
	SetMode(value uint32 /* not a class type */)
	ModifyTime() unsafe.Pointer
	SetModifyTime(value unsafe.Pointer)
	ParentID() FSItemID
	SetParentID(value FSItemID)
	Size() uint64
	SetSize(value uint64)
	SupportsLimitedXAttrs() bool
	SetSupportsLimitedXAttrs(value bool)
	Type() FSItemType
	SetType(value FSItemType)
	Uid() uint32 /* not a class type */
	SetUid(value uint32 /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSItemAttributes */
	// methods:
	InvalidateAllProperties()
	IsValid(attribute FSItemAttribute) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSItemAttributes */
// Alloc allocates a new instance without initialization.
func (fc _FSItemAttributesClass) Alloc() FSItemAttributes {
	rv := objc.Send[FSItemAttributes](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSItemAttributes */
// Attributes of an item, such as size, creation and modification times, and user and group identifiers.


// Attributes of an item, such as size, creation and modification times, and user and group identifiers.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSItemAttributes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSItemAttributes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSItemAttributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSItemAttributes */

// Marks all attributes inactive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/invalidateAllProperties()
func (f_ FSItemAttributes) InvalidateAllProperties() {
	objc.Send[objc.ID](f_.ID, objc.Sel("invalidateAllProperties"))
}/* debug [instance_methods/method]: InvalidateAllProperties */


// Returns a Boolean value that indicates whether the attribute is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/isValid(_:)
func (f_ FSItemAttributes) IsValid(attribute FSItemAttribute) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isValid:"), attribute)
	return rv
}/* debug [instance_methods/method]: IsValid */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSItemAttributes */

// The item’s last-accessed time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/accessTime
func (f_ FSItemAttributes) AccessTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("accessTime"))
	return rv
}/* debug [instance_properties/getter]: accessTime */


// The item’s last-accessed time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/accessTime
func (f_ FSItemAttributes) SetAccessTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAccessTime:"), value)
}/* debug [instance_properties/setter]: accessTime */


// The item’s added time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/addedTime
func (f_ FSItemAttributes) AddedTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("addedTime"))
	return rv
}/* debug [instance_properties/getter]: addedTime */


// The item’s added time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/addedTime
func (f_ FSItemAttributes) SetAddedTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAddedTime:"), value)
}/* debug [instance_properties/setter]: addedTime */


// The item’s allocated size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/allocSize
func (f_ FSItemAttributes) AllocSize() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("allocSize"))
	return rv
}/* debug [instance_properties/getter]: allocSize */


// The item’s allocated size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/allocSize
func (f_ FSItemAttributes) SetAllocSize(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAllocSize:"), value)
}/* debug [instance_properties/setter]: allocSize */


// The item’s last-backup time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/backupTime
func (f_ FSItemAttributes) BackupTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("backupTime"))
	return rv
}/* debug [instance_properties/getter]: backupTime */


// The item’s last-backup time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/backupTime
func (f_ FSItemAttributes) SetBackupTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBackupTime:"), value)
}/* debug [instance_properties/setter]: backupTime */


// The item’s creation time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/birthTime
func (f_ FSItemAttributes) BirthTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("birthTime"))
	return rv
}/* debug [instance_properties/getter]: birthTime */


// The item’s creation time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/birthTime
func (f_ FSItemAttributes) SetBirthTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBirthTime:"), value)
}/* debug [instance_properties/setter]: birthTime */


// The item’s last-changed time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/changeTime
func (f_ FSItemAttributes) ChangeTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("changeTime"))
	return rv
}/* debug [instance_properties/getter]: changeTime */


// The item’s last-changed time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/changeTime
func (f_ FSItemAttributes) SetChangeTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setChangeTime:"), value)
}/* debug [instance_properties/setter]: changeTime */


// The item’s file identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/fileID
func (f_ FSItemAttributes) FileID() FSItemID {
	rv := objc.Send[FSItemID](f_.ID, objc.Sel("fileID"))
	return rv
}/* debug [instance_properties/getter]: fileID */


// The item’s file identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/fileID
func (f_ FSItemAttributes) SetFileID(value FSItemID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileID:"), value)
}/* debug [instance_properties/setter]: fileID */


// The item’s behavior flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/flags
func (f_ FSItemAttributes) Flags() uint32 /* not a class type */ {
	rv := objc.Send[uint32](f_.ID, objc.Sel("flags"))
	return rv
}/* debug [instance_properties/getter]: flags */


// The item’s behavior flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/flags
func (f_ FSItemAttributes) SetFlags(value uint32 /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFlags:"), value)
}/* debug [instance_properties/setter]: flags */


// The group identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/gid
func (f_ FSItemAttributes) Gid() uint32 /* not a class type */ {
	rv := objc.Send[uint32](f_.ID, objc.Sel("gid"))
	return rv
}/* debug [instance_properties/getter]: gid */


// The group identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/gid
func (f_ FSItemAttributes) SetGid(value uint32 /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setGid:"), value)
}/* debug [instance_properties/setter]: gid */


// A Boolean value that indicates whether the file system overrides the per-volume settings for kernel offloaded I/O for a specific file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/inhibitKernelOffloadedIO
func (f_ FSItemAttributes) InhibitKernelOffloadedIO() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("inhibitKernelOffloadedIO"))
	return rv
}/* debug [instance_properties/getter]: inhibitKernelOffloadedIO */


// A Boolean value that indicates whether the file system overrides the per-volume settings for kernel offloaded I/O for a specific file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/inhibitKernelOffloadedIO
func (f_ FSItemAttributes) SetInhibitKernelOffloadedIO(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInhibitKernelOffloadedIO:"), value)
}/* debug [instance_properties/setter]: inhibitKernelOffloadedIO */


// The number of hard links to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/linkCount
func (f_ FSItemAttributes) LinkCount() uint32 /* not a class type */ {
	rv := objc.Send[uint32](f_.ID, objc.Sel("linkCount"))
	return rv
}/* debug [instance_properties/getter]: linkCount */


// The number of hard links to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/linkCount
func (f_ FSItemAttributes) SetLinkCount(value uint32 /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLinkCount:"), value)
}/* debug [instance_properties/setter]: linkCount */


// The mode of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/mode
func (f_ FSItemAttributes) Mode() uint32 /* not a class type */ {
	rv := objc.Send[uint32](f_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// The mode of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/mode
func (f_ FSItemAttributes) SetMode(value uint32 /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */


// The item’s last-modified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/modifyTime
func (f_ FSItemAttributes) ModifyTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("modifyTime"))
	return rv
}/* debug [instance_properties/getter]: modifyTime */


// The item’s last-modified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/modifyTime
func (f_ FSItemAttributes) SetModifyTime(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setModifyTime:"), value)
}/* debug [instance_properties/setter]: modifyTime */


// The identifier of the item’s parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/parentID
func (f_ FSItemAttributes) ParentID() FSItemID {
	rv := objc.Send[FSItemID](f_.ID, objc.Sel("parentID"))
	return rv
}/* debug [instance_properties/getter]: parentID */


// The identifier of the item’s parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/parentID
func (f_ FSItemAttributes) SetParentID(value FSItemID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setParentID:"), value)
}/* debug [instance_properties/setter]: parentID */


// The item’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/size
func (f_ FSItemAttributes) Size() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The item’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/size
func (f_ FSItemAttributes) SetSize(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */


// A Boolean value that indicates whether the item supports a limited set of extended attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/supportsLimitedXAttrs
func (f_ FSItemAttributes) SupportsLimitedXAttrs() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsLimitedXAttrs"))
	return rv
}/* debug [instance_properties/getter]: supportsLimitedXAttrs */


// A Boolean value that indicates whether the item supports a limited set of extended attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/supportsLimitedXAttrs
func (f_ FSItemAttributes) SetSupportsLimitedXAttrs(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsLimitedXAttrs:"), value)
}/* debug [instance_properties/setter]: supportsLimitedXAttrs */


// The item type, such as a regular file, directory, or symbolic link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/type
func (f_ FSItemAttributes) Type() FSItemType {
	rv := objc.Send[FSItemType](f_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The item type, such as a regular file, directory, or symbolic link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/type
func (f_ FSItemAttributes) SetType(value FSItemType) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// The user identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/uid
func (f_ FSItemAttributes) Uid() uint32 /* not a class type */ {
	rv := objc.Send[uint32](f_.ID, objc.Sel("uid"))
	return rv
}/* debug [instance_properties/getter]: uid */


// The user identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/uid
func (f_ FSItemAttributes) SetUid(value uint32 /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUid:"), value)
}/* debug [instance_properties/setter]: uid */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSItemAttributes */



