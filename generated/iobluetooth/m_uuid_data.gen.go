// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mUUIDData] class.
var (
	MUUIDDataClass     _mUUIDDataClass
	MUUIDDataClassOnce sync.Once
)

func getmUUIDDataClass() _mUUIDDataClass {
	MUUIDDataClassOnce.Do(func() {
		MUUIDDataClass = _mUUIDDataClass{objc.GetClass("mUUIDData")}
	})
	return MUUIDDataClass
}

type _mUUIDDataClass struct {
	class objc.Class
}

// An interface definition for the [mUUIDData] class.
type ImUUIDData interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/mUUIDData
type mUUIDData struct {
	objectivec.Object
}

// mUUIDDataFrom constructs a [mUUIDData] from an unsafe.Pointer.
func mUUIDDataFrom(ptr unsafe.Pointer) mUUIDData {
	return mUUIDData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mUUIDDataClass) Alloc() mUUIDData {
	rv := objc.Send[mUUIDData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mUUIDDataClass) New() mUUIDData {
	rv := objc.Send[mUUIDData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mUUIDData) Init() mUUIDData {
	rv := objc.Send[mUUIDData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mUUIDData) Autorelease() mUUIDData {
	rv := objc.Send[mUUIDData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmUUIDData creates a new mUUIDData instance.
func NewmUUIDData() mUUIDData {
	return getmUUIDDataClass().New()
}




