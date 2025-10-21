// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mFolderListingFileRef] class.
var (
	MFolderListingFileRefClass     _mFolderListingFileRefClass
	MFolderListingFileRefClassOnce sync.Once
)

func getmFolderListingFileRefClass() _mFolderListingFileRefClass {
	MFolderListingFileRefClassOnce.Do(func() {
		MFolderListingFileRefClass = _mFolderListingFileRefClass{objc.GetClass("mFolderListingFileRef")}
	})
	return MFolderListingFileRefClass
}

type _mFolderListingFileRefClass struct {
	class objc.Class
}

// An interface definition for the [mFolderListingFileRef] class.
type ImFolderListingFileRef interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mFolderListingFileRef
type mFolderListingFileRef struct {
	objectivec.Object
}

// mFolderListingFileRefFrom constructs a [mFolderListingFileRef] from an unsafe.Pointer.
func mFolderListingFileRefFrom(ptr unsafe.Pointer) mFolderListingFileRef {
	return mFolderListingFileRef{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mFolderListingFileRefClass) Alloc() mFolderListingFileRef {
	rv := objc.Send[mFolderListingFileRef](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mFolderListingFileRefClass) New() mFolderListingFileRef {
	rv := objc.Send[mFolderListingFileRef](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mFolderListingFileRef) Init() mFolderListingFileRef {
	rv := objc.Send[mFolderListingFileRef](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mFolderListingFileRef) Autorelease() mFolderListingFileRef {
	rv := objc.Send[mFolderListingFileRef](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmFolderListingFileRef creates a new mFolderListingFileRef instance.
func NewmFolderListingFileRef() mFolderListingFileRef {
	return getmFolderListingFileRefClass().New()
}




