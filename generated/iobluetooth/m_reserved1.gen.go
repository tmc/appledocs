// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mReserved1] class.
var (
	MReserved1Class     _mReserved1Class
	MReserved1ClassOnce sync.Once
)

func getmReserved1Class() _mReserved1Class {
	MReserved1ClassOnce.Do(func() {
		MReserved1Class = _mReserved1Class{objc.GetClass("mReserved1")}
	})
	return MReserved1Class
}

type _mReserved1Class struct {
	class objc.Class
}

// An interface definition for the [mReserved1] class.
type ImReserved1 interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mReserved1
type mReserved1 struct {
	objectivec.Object
}

// mReserved1From constructs a [mReserved1] from an unsafe.Pointer.
func mReserved1From(ptr unsafe.Pointer) mReserved1 {
	return mReserved1{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mReserved1Class) Alloc() mReserved1 {
	rv := objc.Send[mReserved1](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mReserved1Class) New() mReserved1 {
	rv := objc.Send[mReserved1](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReserved1) Init() mReserved1 {
	rv := objc.Send[mReserved1](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReserved1) Autorelease() mReserved1 {
	rv := objc.Send[mReserved1](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReserved1 creates a new mReserved1 instance.
func NewmReserved1() mReserved1 {
	return getmReserved1Class().New()
}




