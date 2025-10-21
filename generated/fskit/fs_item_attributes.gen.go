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
	IsValid(attribute unsafe.Pointer) bool
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
func (f_ FSItemAttributes) IsValid(attribute unsafe.Pointer) bool {
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
func (f_ FSItemAttributes) FileID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileID"))
	return rv
}


// SetFileID sets the value of the fileID property.
// The item’s file identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/fileID
func (f_ FSItemAttributes) SetFileID(value unsafe.Pointer) {
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



