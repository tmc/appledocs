// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mAttributeID] class.
var (
	MAttributeIDClass     _mAttributeIDClass
	MAttributeIDClassOnce sync.Once
)

func getmAttributeIDClass() _mAttributeIDClass {
	MAttributeIDClassOnce.Do(func() {
		MAttributeIDClass = _mAttributeIDClass{objc.GetClass("mAttributeID")}
	})
	return MAttributeIDClass
}

type _mAttributeIDClass struct {
	class objc.Class
}

// An interface definition for the [mAttributeID] class.
type ImAttributeID interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/mAttributeID
type mAttributeID struct {
	objectivec.Object
}

// mAttributeIDFrom constructs a [mAttributeID] from an unsafe.Pointer.
func mAttributeIDFrom(ptr unsafe.Pointer) mAttributeID {
	return mAttributeID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mAttributeIDClass) Alloc() mAttributeID {
	rv := objc.Send[mAttributeID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mAttributeIDClass) New() mAttributeID {
	rv := objc.Send[mAttributeID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAttributeID) Init() mAttributeID {
	rv := objc.Send[mAttributeID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAttributeID) Autorelease() mAttributeID {
	rv := objc.Send[mAttributeID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAttributeID creates a new mAttributeID instance.
func NewmAttributeID() mAttributeID {
	return getmAttributeIDClass().New()
}




