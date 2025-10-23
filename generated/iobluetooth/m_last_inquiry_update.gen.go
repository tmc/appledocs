// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mLastInquiryUpdate] class.
var (
	MLastInquiryUpdateClass     _mLastInquiryUpdateClass
	MLastInquiryUpdateClassOnce sync.Once
)

func getmLastInquiryUpdateClass() _mLastInquiryUpdateClass {
	MLastInquiryUpdateClassOnce.Do(func() {
		MLastInquiryUpdateClass = _mLastInquiryUpdateClass{objc.GetClass("mLastInquiryUpdate")}
	})
	return MLastInquiryUpdateClass
}

type _mLastInquiryUpdateClass struct {
	class objc.Class
}

// An interface definition for the [mLastInquiryUpdate] class.
type ImLastInquiryUpdate interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mLastInquiryUpdate
type mLastInquiryUpdate struct {
	objectivec.Object
}

// mLastInquiryUpdateFrom constructs a [mLastInquiryUpdate] from an unsafe.Pointer.
func mLastInquiryUpdateFrom(ptr unsafe.Pointer) mLastInquiryUpdate {
	return mLastInquiryUpdate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mLastInquiryUpdateClass) Alloc() mLastInquiryUpdate {
	rv := objc.Send[mLastInquiryUpdate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mLastInquiryUpdateClass) New() mLastInquiryUpdate {
	rv := objc.Send[mLastInquiryUpdate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLastInquiryUpdate) Init() mLastInquiryUpdate {
	rv := objc.Send[mLastInquiryUpdate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLastInquiryUpdate) Autorelease() mLastInquiryUpdate {
	rv := objc.Send[mLastInquiryUpdate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLastInquiryUpdate creates a new mLastInquiryUpdate instance.
func NewmLastInquiryUpdate() mLastInquiryUpdate {
	return getmLastInquiryUpdateClass().New()
}




