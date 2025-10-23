// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mLastNameUpdate] class.
var (
	MLastNameUpdateClass     _mLastNameUpdateClass
	MLastNameUpdateClassOnce sync.Once
)

func getmLastNameUpdateClass() _mLastNameUpdateClass {
	MLastNameUpdateClassOnce.Do(func() {
		MLastNameUpdateClass = _mLastNameUpdateClass{objc.GetClass("mLastNameUpdate")}
	})
	return MLastNameUpdateClass
}

type _mLastNameUpdateClass struct {
	class objc.Class
}

// An interface definition for the [mLastNameUpdate] class.
type ImLastNameUpdate interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mLastNameUpdate
type mLastNameUpdate struct {
	objectivec.Object
}

// mLastNameUpdateFrom constructs a [mLastNameUpdate] from an unsafe.Pointer.
func mLastNameUpdateFrom(ptr unsafe.Pointer) mLastNameUpdate {
	return mLastNameUpdate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mLastNameUpdateClass) Alloc() mLastNameUpdate {
	rv := objc.Send[mLastNameUpdate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mLastNameUpdateClass) New() mLastNameUpdate {
	rv := objc.Send[mLastNameUpdate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLastNameUpdate) Init() mLastNameUpdate {
	rv := objc.Send[mLastNameUpdate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLastNameUpdate) Autorelease() mLastNameUpdate {
	rv := objc.Send[mLastNameUpdate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLastNameUpdate creates a new mLastNameUpdate instance.
func NewmLastNameUpdate() mLastNameUpdate {
	return getmLastNameUpdateClass().New()
}




