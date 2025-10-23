// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mReserved2] class.
var (
	MReserved2Class     _mReserved2Class
	MReserved2ClassOnce sync.Once
)

func getmReserved2Class() _mReserved2Class {
	MReserved2ClassOnce.Do(func() {
		MReserved2Class = _mReserved2Class{objc.GetClass("mReserved2")}
	})
	return MReserved2Class
}

type _mReserved2Class struct {
	class objc.Class
}

// An interface definition for the [mReserved2] class.
type ImReserved2 interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mReserved2
type mReserved2 struct {
	objectivec.Object
}

// mReserved2From constructs a [mReserved2] from an unsafe.Pointer.
func mReserved2From(ptr unsafe.Pointer) mReserved2 {
	return mReserved2{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mReserved2Class) Alloc() mReserved2 {
	rv := objc.Send[mReserved2](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mReserved2Class) New() mReserved2 {
	rv := objc.Send[mReserved2](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReserved2) Init() mReserved2 {
	rv := objc.Send[mReserved2](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReserved2) Autorelease() mReserved2 {
	rv := objc.Send[mReserved2](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReserved2 creates a new mReserved2 instance.
func NewmReserved2() mReserved2 {
	return getmReserved2Class().New()
}




