// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mMTU] class.
var (
	MMTUClass     _mMTUClass
	MMTUClassOnce sync.Once
)

func getmMTUClass() _mMTUClass {
	MMTUClassOnce.Do(func() {
		MMTUClass = _mMTUClass{objc.GetClass("mMTU")}
	})
	return MMTUClass
}

type _mMTUClass struct {
	class objc.Class
}

// An interface definition for the [mMTU] class.
type ImMTU interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mMTU
type mMTU struct {
	objectivec.Object
}

// mMTUFrom constructs a [mMTU] from an unsafe.Pointer.
func mMTUFrom(ptr unsafe.Pointer) mMTU {
	return mMTU{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mMTUClass) Alloc() mMTU {
	rv := objc.Send[mMTU](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mMTUClass) New() mMTU {
	rv := objc.Send[mMTU](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mMTU) Init() mMTU {
	rv := objc.Send[mMTU](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mMTU) Autorelease() mMTU {
	rv := objc.Send[mMTU](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmMTU creates a new mMTU instance.
func NewmMTU() mMTU {
	return getmMTUClass().New()
}




