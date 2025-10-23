// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mName] class.
var (
	MNameClass     _mNameClass
	MNameClassOnce sync.Once
)

func getmNameClass() _mNameClass {
	MNameClassOnce.Do(func() {
		MNameClass = _mNameClass{objc.GetClass("mName")}
	})
	return MNameClass
}

type _mNameClass struct {
	class objc.Class
}

// An interface definition for the [mName] class.
type ImName interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mName
type mName struct {
	objectivec.Object
}

// mNameFrom constructs a [mName] from an unsafe.Pointer.
func mNameFrom(ptr unsafe.Pointer) mName {
	return mName{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mNameClass) Alloc() mName {
	rv := objc.Send[mName](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mNameClass) New() mName {
	rv := objc.Send[mName](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mName) Init() mName {
	rv := objc.Send[mName](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mName) Autorelease() mName {
	rv := objc.Send[mName](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmName creates a new mName instance.
func NewmName() mName {
	return getmNameClass().New()
}




