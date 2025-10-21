// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIsClosed] class.
var (
	MIsClosedClass     _mIsClosedClass
	MIsClosedClassOnce sync.Once
)

func getmIsClosedClass() _mIsClosedClass {
	MIsClosedClassOnce.Do(func() {
		MIsClosedClass = _mIsClosedClass{objc.GetClass("mIsClosed")}
	})
	return MIsClosedClass
}

type _mIsClosedClass struct {
	class objc.Class
}

// An interface definition for the [mIsClosed] class.
type ImIsClosed interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mIsClosed
type mIsClosed struct {
	objectivec.Object
}

// mIsClosedFrom constructs a [mIsClosed] from an unsafe.Pointer.
func mIsClosedFrom(ptr unsafe.Pointer) mIsClosed {
	return mIsClosed{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIsClosedClass) Alloc() mIsClosed {
	rv := objc.Send[mIsClosed](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIsClosedClass) New() mIsClosed {
	rv := objc.Send[mIsClosed](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIsClosed) Init() mIsClosed {
	rv := objc.Send[mIsClosed](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIsClosed) Autorelease() mIsClosed {
	rv := objc.Send[mIsClosed](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIsClosed creates a new mIsClosed instance.
func NewmIsClosed() mIsClosed {
	return getmIsClosedClass().New()
}




