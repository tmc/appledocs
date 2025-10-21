// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHObjectChangeDetails] class.
var (
	PHObjectChangeDetailsClass     _PHObjectChangeDetailsClass
	PHObjectChangeDetailsClassOnce sync.Once
)

func getPHObjectChangeDetailsClass() _PHObjectChangeDetailsClass {
	PHObjectChangeDetailsClassOnce.Do(func() {
		PHObjectChangeDetailsClass = _PHObjectChangeDetailsClass{objc.GetClass("PHObjectChangeDetails")}
	})
	return PHObjectChangeDetailsClass
}

type _PHObjectChangeDetailsClass struct {
	class objc.Class
}

// An interface definition for the [PHObjectChangeDetails] class.
type IPHObjectChangeDetails interface {
	objectivec.IObject
}

// A description of changes that occurred in an asset or collection object.
//
// A object provides detailed information about differences between two states of an asset or collection object—one that you previously obtained and an updated state that would result if you fetched that entity again. You observe changes by adopting the protocol and registering your observer with the shared object. When Photos notifies your observer of a change, you get change details by passing the object you’re interested in to the method. For an asset collection or collection list, a object describe changes only to the collection’s properties. If you’re instead interested in changes to the collection’s membership, fetch the collection’s contents and use the method to track changes to the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHObjectChangeDetails
type PHObjectChangeDetails struct {
	objectivec.Object
}

// PHObjectChangeDetailsFrom constructs a [PHObjectChangeDetails] from an unsafe.Pointer.
//
// A description of changes that occurred in an asset or collection object.
func PHObjectChangeDetailsFrom(ptr unsafe.Pointer) PHObjectChangeDetails {
	return PHObjectChangeDetails{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHObjectChangeDetailsClass) Alloc() PHObjectChangeDetails {
	rv := objc.Send[PHObjectChangeDetails](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHObjectChangeDetailsClass) New() PHObjectChangeDetails {
	rv := objc.Send[PHObjectChangeDetails](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHObjectChangeDetails) Init() PHObjectChangeDetails {
	rv := objc.Send[PHObjectChangeDetails](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHObjectChangeDetails) Autorelease() PHObjectChangeDetails {
	rv := objc.Send[PHObjectChangeDetails](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHObjectChangeDetails creates a new PHObjectChangeDetails instance.
func NewPHObjectChangeDetails() PHObjectChangeDetails {
	return getPHObjectChangeDetailsClass().New()
}


// The indexes of objects in the fetch result whose content or metadata have been updated.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/changedindexes
func (p_ PHObjectChangeDetails) ChangedIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("changedIndexes"))
	return rv
}


// SetChangedIndexes sets the value of the changedIndexes property.
// The indexes of objects in the fetch result whose content or metadata have been updated.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phfetchresultchangedetails/changedindexes
func (p_ PHObjectChangeDetails) SetChangedIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setChangedIndexes:"), value)
}

// An object that reflects the original state of the asset or collection it represents.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobjectchangedetails/objectbeforechanges
func (p_ PHObjectChangeDetails) ObjectBeforeChanges() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("objectBeforeChanges"))
	return rv
}


// SetObjectBeforeChanges sets the value of the objectBeforeChanges property.
// An object that reflects the original state of the asset or collection it represents.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobjectchangedetails/objectbeforechanges
func (p_ PHObjectChangeDetails) SetObjectBeforeChanges(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObjectBeforeChanges:"), value)
}

// A Boolean value that indicates whether the asset’s photo or video content has changed.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobjectchangedetails/assetcontentchanged
func (p_ PHObjectChangeDetails) AssetContentChanged() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("assetContentChanged"))
	return rv
}


// SetAssetContentChanged sets the value of the assetContentChanged property.
// A Boolean value that indicates whether the asset’s photo or video content has changed.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobjectchangedetails/assetcontentchanged
func (p_ PHObjectChangeDetails) SetAssetContentChanged(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAssetContentChanged:"), value)
}

// An object that reflects the current state of the asset or collection it represents.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobjectchangedetails/objectafterchanges
func (p_ PHObjectChangeDetails) ObjectAfterChanges() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("objectAfterChanges"))
	return rv
}


// SetObjectAfterChanges sets the value of the objectAfterChanges property.
// An object that reflects the current state of the asset or collection it represents.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobjectchangedetails/objectafterchanges
func (p_ PHObjectChangeDetails) SetObjectAfterChanges(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObjectAfterChanges:"), value)
}

// A Boolean value that indicates whether the object has been deleted from the Photos library.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobjectchangedetails/objectwasdeleted
func (p_ PHObjectChangeDetails) ObjectWasDeleted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("objectWasDeleted"))
	return rv
}


// SetObjectWasDeleted sets the value of the objectWasDeleted property.
// A Boolean value that indicates whether the object has been deleted from the Photos library.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobjectchangedetails/objectwasdeleted
func (p_ PHObjectChangeDetails) SetObjectWasDeleted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObjectWasDeleted:"), value)
}



