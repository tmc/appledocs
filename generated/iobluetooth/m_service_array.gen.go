// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mServiceArray] class.
var (
	MServiceArrayClass     _mServiceArrayClass
	MServiceArrayClassOnce sync.Once
)

func getmServiceArrayClass() _mServiceArrayClass {
	MServiceArrayClassOnce.Do(func() {
		MServiceArrayClass = _mServiceArrayClass{objc.GetClass("mServiceArray")}
	})
	return MServiceArrayClass
}

type _mServiceArrayClass struct {
	class objc.Class
}

// An interface definition for the [mServiceArray] class.
type ImServiceArray interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mServiceArray
type mServiceArray struct {
	objectivec.Object
}

// mServiceArrayFrom constructs a [mServiceArray] from an unsafe.Pointer.
func mServiceArrayFrom(ptr unsafe.Pointer) mServiceArray {
	return mServiceArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mServiceArrayClass) Alloc() mServiceArray {
	rv := objc.Send[mServiceArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mServiceArrayClass) New() mServiceArray {
	rv := objc.Send[mServiceArray](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mServiceArray) Init() mServiceArray {
	rv := objc.Send[mServiceArray](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mServiceArray) Autorelease() mServiceArray {
	rv := objc.Send[mServiceArray](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmServiceArray creates a new mServiceArray instance.
func NewmServiceArray() mServiceArray {
	return getmServiceArrayClass().New()
}




