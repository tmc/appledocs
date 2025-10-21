// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mActionInProgress] class.
var (
	MActionInProgressClass     _mActionInProgressClass
	MActionInProgressClassOnce sync.Once
)

func getmActionInProgressClass() _mActionInProgressClass {
	MActionInProgressClassOnce.Do(func() {
		MActionInProgressClass = _mActionInProgressClass{objc.GetClass("mActionInProgress")}
	})
	return MActionInProgressClass
}

type _mActionInProgressClass struct {
	class objc.Class
}

// An interface definition for the [mActionInProgress] class.
type ImActionInProgress interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mActionInProgress
type mActionInProgress struct {
	objectivec.Object
}

// mActionInProgressFrom constructs a [mActionInProgress] from an unsafe.Pointer.
func mActionInProgressFrom(ptr unsafe.Pointer) mActionInProgress {
	return mActionInProgress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mActionInProgressClass) Alloc() mActionInProgress {
	rv := objc.Send[mActionInProgress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mActionInProgressClass) New() mActionInProgress {
	rv := objc.Send[mActionInProgress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mActionInProgress) Init() mActionInProgress {
	rv := objc.Send[mActionInProgress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mActionInProgress) Autorelease() mActionInProgress {
	rv := objc.Send[mActionInProgress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmActionInProgress creates a new mActionInProgress instance.
func NewmActionInProgress() mActionInProgress {
	return getmActionInProgressClass().New()
}




