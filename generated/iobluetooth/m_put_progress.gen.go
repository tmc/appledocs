// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mPUTProgress] class.
var (
	MPUTProgressClass     _mPUTProgressClass
	MPUTProgressClassOnce sync.Once
)

func getmPUTProgressClass() _mPUTProgressClass {
	MPUTProgressClassOnce.Do(func() {
		MPUTProgressClass = _mPUTProgressClass{objc.GetClass("mPUTProgress")}
	})
	return MPUTProgressClass
}

type _mPUTProgressClass struct {
	class objc.Class
}

// An interface definition for the [mPUTProgress] class.
type ImPUTProgress interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mPUTProgress
type mPUTProgress struct {
	objectivec.Object
}

// mPUTProgressFrom constructs a [mPUTProgress] from an unsafe.Pointer.
func mPUTProgressFrom(ptr unsafe.Pointer) mPUTProgress {
	return mPUTProgress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mPUTProgressClass) Alloc() mPUTProgress {
	rv := objc.Send[mPUTProgress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mPUTProgressClass) New() mPUTProgress {
	rv := objc.Send[mPUTProgress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPUTProgress) Init() mPUTProgress {
	rv := objc.Send[mPUTProgress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPUTProgress) Autorelease() mPUTProgress {
	rv := objc.Send[mPUTProgress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPUTProgress creates a new mPUTProgress instance.
func NewmPUTProgress() mPUTProgress {
	return getmPUTProgressClass().New()
}




