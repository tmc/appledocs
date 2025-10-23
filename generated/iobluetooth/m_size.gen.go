// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mSize] class.
var (
	MSizeClass     _mSizeClass
	MSizeClassOnce sync.Once
)

func getmSizeClass() _mSizeClass {
	MSizeClassOnce.Do(func() {
		MSizeClass = _mSizeClass{objc.GetClass("mSize")}
	})
	return MSizeClass
}

type _mSizeClass struct {
	class objc.Class
}

// An interface definition for the [mSize] class.
type ImSize interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/mSize
type mSize struct {
	objectivec.Object
}

// mSizeFrom constructs a [mSize] from an unsafe.Pointer.
func mSizeFrom(ptr unsafe.Pointer) mSize {
	return mSize{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mSizeClass) Alloc() mSize {
	rv := objc.Send[mSize](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mSizeClass) New() mSize {
	rv := objc.Send[mSize](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mSize) Init() mSize {
	rv := objc.Send[mSize](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mSize) Autorelease() mSize {
	rv := objc.Send[mSize](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmSize creates a new mSize instance.
func NewmSize() mSize {
	return getmSizeClass().New()
}




