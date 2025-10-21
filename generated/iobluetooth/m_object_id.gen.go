// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mObjectID] class.
var (
	MObjectIDClass     _mObjectIDClass
	MObjectIDClassOnce sync.Once
)

func getmObjectIDClass() _mObjectIDClass {
	MObjectIDClassOnce.Do(func() {
		MObjectIDClass = _mObjectIDClass{objc.GetClass("mObjectID")}
	})
	return MObjectIDClass
}

type _mObjectIDClass struct {
	class objc.Class
}

// An interface definition for the [mObjectID] class.
type ImObjectID interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mObjectID
type mObjectID struct {
	objectivec.Object
}

// mObjectIDFrom constructs a [mObjectID] from an unsafe.Pointer.
func mObjectIDFrom(ptr unsafe.Pointer) mObjectID {
	return mObjectID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mObjectIDClass) Alloc() mObjectID {
	rv := objc.Send[mObjectID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mObjectIDClass) New() mObjectID {
	rv := objc.Send[mObjectID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mObjectID) Init() mObjectID {
	rv := objc.Send[mObjectID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mObjectID) Autorelease() mObjectID {
	rv := objc.Send[mObjectID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmObjectID creates a new mObjectID instance.
func NewmObjectID() mObjectID {
	return getmObjectIDClass().New()
}




