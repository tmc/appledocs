// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mGETProgress] class.
var (
	MGETProgressClass     _mGETProgressClass
	MGETProgressClassOnce sync.Once
)

func getmGETProgressClass() _mGETProgressClass {
	MGETProgressClassOnce.Do(func() {
		MGETProgressClass = _mGETProgressClass{objc.GetClass("mGETProgress")}
	})
	return MGETProgressClass
}

type _mGETProgressClass struct {
	class objc.Class
}

// An interface definition for the [mGETProgress] class.
type ImGETProgress interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mGETProgress
type mGETProgress struct {
	objectivec.Object
}

// mGETProgressFrom constructs a [mGETProgress] from an unsafe.Pointer.
func mGETProgressFrom(ptr unsafe.Pointer) mGETProgress {
	return mGETProgress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mGETProgressClass) Alloc() mGETProgress {
	rv := objc.Send[mGETProgress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mGETProgressClass) New() mGETProgress {
	rv := objc.Send[mGETProgress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mGETProgress) Init() mGETProgress {
	rv := objc.Send[mGETProgress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mGETProgress) Autorelease() mGETProgress {
	rv := objc.Send[mGETProgress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmGETProgress creates a new mGETProgress instance.
func NewmGETProgress() mGETProgress {
	return getmGETProgressClass().New()
}




